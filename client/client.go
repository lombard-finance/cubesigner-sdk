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
}

type MfaHeaders struct {
	Id           string
	Confirmation string
}

// New creates a new client, connecting with a standard HTTP.
func New(address, orgID, token string, logger *logrus.Entry, timeout time.Duration) (*Client, error) {

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

func (cli *Client) addExtraHeaders(req *http.Request) {
	for k, v := range cli.extraHeaders {
		req.Header.Add(k, v)
	}
}

func (cli *Client) get(endpoint *url.URL, overrideHeaders map[string]string, page *Page) (io.Reader, error) {
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

	opCtx, cancel := context.WithTimeout(context.Background(), cli.timeout)
	req, err := http.NewRequestWithContext(opCtx, http.MethodGet, endpoint.String(), nil)
	if err != nil {
		cancel()
		return nil, errors.Wrap(err, "failed to create GET request")
	}

	cli.addExtraHeaders(req)

	for key, value := range overrideHeaders {
		req.Header.Set(key, value)
	}

	resp, err := cli.client.Do(req)
	if err != nil {
		cancel()
		return nil, errors.Wrap(err, "failed to call GET endpoint")
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		// Nothing found.  This is not an error, so we return nil on both counts.
		cancel()
		return nil, nil
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		cancel()
		return nil, errors.Wrap(err, "failed to read GET response")
	}

	statusFamily := resp.StatusCode / 100
	if statusFamily != 2 {
		cancel()
		log.WithFields(logrus.Fields{
			"status_code": resp.StatusCode,
			"data":        string(data),
		}).Trace("GET failed")
		return nil, errors.Errorf("Method %s, StatusCode: %d, Endpoint: %s", http.MethodGet, resp.StatusCode, endpoint)
	}

	cancel()
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
	// copy body if not nil
	var buf bytes.Buffer
	var tee io.Reader
	if body != nil {
		tee = io.TeeReader(body, &buf)

		bodyBytes, err := io.ReadAll(tee)
		if err != nil {
			cli.logger.WithError(err).Warn("failed to read request body")
		} else {
			cli.logger.Tracef("request body: %s", bodyBytes)
		}
	}

	log := cli.logger.WithFields(logrus.Fields{
		"address":  cli.address,
		"endpoint": endpoint.String(),
		"method":   method,
	})

	if page != nil {
		page.Apply(endpoint)
		log = log.WithField("query", endpoint.Query().Encode())
	}

	// build request
	opCtx, cancel := context.WithTimeout(context.Background(), cli.timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(opCtx, method, endpoint.String(), &buf)
	if err != nil {
		return nil, 0, errors.Wrap(err, "create request with context")
	}

	// add headers
	cli.addExtraHeaders(req)
	req.Header.Set("Content-type", "application/json")
	//req.Header.Set("Accept", "application/json")

	for key, value := range overrideHeaders {
		req.Header.Set(key, value)
	}

	// do the request
	resp, err := cli.client.Do(req)
	if err != nil {
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

// close closes the client, freeing up resources.
// TODO
func (cli *Client) close() {
}

func (cli *Client) getMfaHeaders(mfaHeaders MfaHeaders) map[string]string {
	return map[string]string{
		"x-cubist-mfa-id":           mfaHeaders.Id,
		"x-cubist-mfa-org-id":       cli.orgID,
		"x-cubist-mfa-confirmation": mfaHeaders.Confirmation,
	}
}
