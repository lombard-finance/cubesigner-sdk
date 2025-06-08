package client

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Client struct {
	logger *logrus.Entry

	base         *url.URL
	client       *http.Client
	timeout      time.Duration
	address      string
	extraHeaders map[string]string

	orgID string
	token string

	// Context-based cancellation
	ctx       context.Context
	cancel    context.CancelFunc
	closeOnce sync.Once
}

type MfaHeaders struct {
	Id           string
	Confirmation string
}

// New creates a new client, connecting with a standard HTTP.
func New(address, orgID, token string, logger *logrus.Entry, timeout time.Duration) (*Client, error) {
	ctx, cancel := context.WithCancel(context.Background())

	client := &http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout: timeout,
				//KeepAlive:     0,
			}).DialContext,
			MaxIdleConns:        32,
			MaxConnsPerHost:     32,
			MaxIdleConnsPerHost: 32,
			IdleConnTimeout:     600 * time.Second,
		},
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       timeout,
	}

	if !strings.HasPrefix(address, "https") {
		address = fmt.Sprintf("https://%s", address)
	}
	if !strings.HasSuffix(address, "/") {
		address = fmt.Sprintf("%s/", address)
	}
	base, err := url.Parse(address)
	if err != nil {
		cancel()
		return nil, errors.Wrap(err, "invalid URL")
	}

	return &Client{
		logger:  logger.WithField("client", "cubist"),
		base:    base,
		client:  client,
		timeout: timeout,
		address: address,
		orgID:   orgID,
		token:   token,
		extraHeaders: map[string]string{
			"Authorization": token,
		},
		ctx:    ctx,
		cancel: cancel,
	}, nil
}

// Name provides the name of the service.
func (cli *Client) Name() string {
	return "Standard (HTTP)"
}

// Address provides the address for the connection.
func (cli *Client) Address() string {
	return cli.address
}

func (cli *Client) get(endpoint *url.URL, overrideHeaders map[string]string, page *Page) (io.Reader, error) {
	// Check if client is closed
	if cli.isClosed() {
		return nil, errors.New("client is closed")
	}

	log := cli.logger.WithFields(logrus.Fields{
		"id":       fmt.Sprintf("%02x", rand.Int31()),
		"address":  cli.address,
		"endpoint": endpoint.String(),
		"method":   http.MethodGet,
	})

	log.Trace("request")

	if page != nil {
		page.Apply(endpoint)
		log = log.WithField("query", endpoint.Query().Encode())
	}

	// Use client context as parent - auto-cancels if client is closed
	opCtx, cancel := context.WithTimeout(cli.ctx, cli.timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(opCtx, http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return nil, errors.Wrap(err, "create GET request")
	}

	// Add headers
	cli.addExtraHeaders(req)
	for key, value := range overrideHeaders {
		req.Header.Set(key, value)
	}

	// Do the request
	resp, err := cli.client.Do(req)
	if err != nil {
		if cli.isClosed() {
			return nil, errors.New("client was closed during request")
		}
		return nil, errors.Wrap(err, "call GET endpoint")
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		// Nothing found.  This is not an error, so we return nil on both counts.
		return nil, nil
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "read GET response")
	}

	statusFamily := resp.StatusCode / 100
	if statusFamily != 2 {
		log.WithFields(logrus.Fields{
			"status_code": resp.StatusCode,
			"data":        string(data),
		}).Trace("GET failed")
		return nil, errors.Errorf("Method %s, StatusCode: %d, Endpoint: %s", http.MethodGet, resp.StatusCode, endpoint)
	}

	log.WithField("response", string(data)).Trace("response")

	return bytes.NewReader(data), nil
}

func (cli *Client) post(endpoint *url.URL, body io.Reader, overrideHeaders map[string]string, page *Page) (io.Reader, int, error) {
	return cli.requestWithBody(endpoint, http.MethodPost, body, overrideHeaders, page)
}

func (cli *Client) put(endpoint *url.URL, body io.Reader, overrideHeaders map[string]string, page *Page) (io.Reader, int, error) {
	return cli.requestWithBody(endpoint, http.MethodPut, body, overrideHeaders, page)
}

func (cli *Client) patch(endpoint *url.URL, body io.Reader, overrideHeaders map[string]string, page *Page) (io.Reader, int, error) {
	return cli.requestWithBody(endpoint, http.MethodPatch, body, overrideHeaders, page)
}

func (cli *Client) requestWithBody(endpoint *url.URL, method string, body io.Reader, overrideHeaders map[string]string, page *Page) (io.Reader, int, error) {
	// Check if client is closed
	if cli.isClosed() {
		return nil, 0, errors.New("client is closed")
	}

	log := cli.logger.WithFields(logrus.Fields{
		"address":  cli.address,
		"endpoint": endpoint.String(),
		"method":   method,
	})

	// Copy body if not nil
	var buf bytes.Buffer
	var tee io.Reader
	if body != nil {
		tee = io.TeeReader(body, &buf)

		bodyBytes, err := io.ReadAll(tee)
		if err != nil {
			log.WithError(err).Warn("read request body")
		} else {
			log.Tracef("request body: %s", bodyBytes)
		}
	}

	if page != nil {
		page.Apply(endpoint)
		log = log.WithField("query", endpoint.Query().Encode())
	}

	// Use client context as parent - auto-cancels if client is closed
	opCtx, cancel := context.WithTimeout(cli.ctx, cli.timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(opCtx, method, endpoint.String(), &buf)
	if err != nil {
		return nil, 0, errors.Wrap(err, "create request with context")
	}

	// Add headers
	cli.addExtraHeaders(req)
	req.Header.Set("Content-type", "application/json")
	//req.Header.Set("Accept", "application/json")
	for key, value := range overrideHeaders {
		req.Header.Set(key, value)
	}

	// Do the request
	resp, err := cli.client.Do(req)
	if err != nil {
		if cli.isClosed() {
			return nil, 0, errors.New("client was closed during request")
		}
		return nil, 0, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, errors.Wrap(err, "read response")
	}

	log.WithField("status_code", resp.StatusCode)

	statusFamily := resp.StatusCode / 100
	if statusFamily != 2 {
		log.Trace("failed")
		return nil, 0, errors.Errorf("Method: %s, StatusCode: %d, Endpoint: %s", method, resp.StatusCode, endpoint)
	}

	return bytes.NewReader(data), resp.StatusCode, nil
}

// Close closes the client, freeing up resources and cancelling all operations
func (cli *Client) Close() {
	cli.closeOnce.Do(func() {
		cli.logger.Debug("closing client")

		// Cancel all ongoing and future operations
		cli.cancel()

		// Close idle connections
		if cli.client != nil && cli.client.Transport != nil {
			if transport, ok := cli.client.Transport.(*http.Transport); ok {
				transport.CloseIdleConnections()
			}
		}

		// Clear sensitive data
		cli.token = ""
		if cli.extraHeaders != nil {
			delete(cli.extraHeaders, "Authorization")
		}

		cli.logger.Debug("client closed")
	})
}

// isClosed returns true if the client has been closed
func (cli *Client) isClosed() bool {
	select {
	case <-cli.ctx.Done():
		return true
	default:
		return false
	}
}
