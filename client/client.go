package client

import (
	"bytes"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"io"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"
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

type Params struct {
	Logger  *logrus.Entry
	Timeout time.Duration
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

func (cli *Client) get(endpoint string) (io.Reader, error) {
	log := cli.logger.WithField("id", fmt.Sprintf("%02x", rand.Int31())).
		WithField("address", cli.address).
		WithField("endpoint", endpoint)
	log.Trace("GET request")

	// replace known path variables
	endpoint = strings.Replace(endpoint, ":org_id", url.PathEscape(cli.orgID), -1)

	requestEndpoint, err := url.Parse(fmt.Sprintf("%s%s", strings.TrimSuffix(cli.base.String(), "/"), endpoint))
	if err != nil {
		return nil, errors.Wrap(err, "invalid endpoint")
	}

	opCtx, cancel := context.WithTimeout(context.Background(), cli.timeout)
	req, err := http.NewRequestWithContext(opCtx, http.MethodGet, requestEndpoint.String(), nil)
	if err != nil {
		cancel()
		return nil, errors.Wrap(err, "failed to create GET request")
	}
	cli.addExtraHeaders(req)

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
		log.Trace("status_code", resp.StatusCode)
		log.Trace("data", string(data))
		log.Trace("GET failed")
		return nil, errors.Errorf("Method %s, StatusCode: %d, Endpoint: %s, Data: %s", http.MethodGet, resp.StatusCode, endpoint, data)
	}
	cancel()

	log.Trace("response", string(data))
	log.Trace("GET response")

	return bytes.NewReader(data), nil
}

func (cli *Client) post(endpoint string, body io.Reader) (io.Reader, error) {
	return cli.requestWithBody(endpoint, http.MethodPost, body)
}

func (cli *Client) put(endpoint string, body io.Reader) (io.Reader, error) {
	return cli.requestWithBody(endpoint, http.MethodPut, body)
}

func (cli *Client) patch(endpoint string, body io.Reader) (io.Reader, error) {
	return cli.requestWithBody(endpoint, http.MethodPatch, body)
}

func (cli *Client) requestWithBody(endpoint string, method string, body io.Reader) (io.Reader, error) {
	log := cli.logger.WithField("id", fmt.Sprintf("%02x", rand.Int31())).
		WithField("address", cli.address).
		WithField("endpoint", endpoint).
		WithField("method", method)

	// replace path variables
	endpoint = strings.Replace(endpoint, ":org_id", url.PathEscape(cli.orgID), -1)

	// build url
	requestEndpoint, err := url.Parse(fmt.Sprintf("%s%s", strings.TrimSuffix(cli.base.String(), "/"), endpoint))
	log.Info(requestEndpoint.String()) // TODO: remove
	if err != nil {
		return nil, errors.Wrap(err, "invalid endpoint")
	}

	// build request
	opCtx, cancel := context.WithTimeout(context.Background(), cli.timeout)
	req, err := http.NewRequestWithContext(opCtx, method, requestEndpoint.String(), body)
	if err != nil {
		cancel()
		return nil, errors.Wrap(err, "create request with context")
	}

	// add headers
	cli.addExtraHeaders(req)
	req.Header.Set("Content-type", "application/json")
	//req.Header.Set("Accept", "application/json")

	// do the request
	resp, err := cli.client.Do(req)
	if err != nil {
		cancel()
		return nil, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		cancel()
		return nil, errors.Wrap(err, "read response")
	}

	statusFamily := resp.StatusCode / 100
	if statusFamily != 2 {
		log.Trace("status_code", resp.StatusCode)
		log.Trace("data", string(data))
		log.Tracef("%s failed", method)
		cancel()
		return nil, errors.Errorf("Method %s, StatusCode: %d, Endpoint: %s, Data: %s", method, resp.StatusCode, endpoint, data)
	}
	cancel()

	log.Trace("response", string(data))
	log.Tracef("%s response", method)

	return bytes.NewReader(data), nil
}

// close closes the client, freeing up resources.
func (cli *Client) close() {
}

// parameterToString convert interface{} parameters to string, using a delimiter if format is provided.
func parameterToString(obj interface{}, collectionFormat string) string {
	var delimiter string

	switch collectionFormat {
	case "pipes":
		delimiter = "|"
	case "ssv":
		delimiter = " "
	case "tsv":
		delimiter = "\t"
	case "csv":
		delimiter = ","
	}

	if reflect.TypeOf(obj).Kind() == reflect.Slice {
		return strings.Trim(strings.Replace(fmt.Sprint(obj), " ", delimiter, -1), "[]")
	} else if t, ok := obj.(time.Time); ok {
		return t.Format(time.RFC3339)
	}

	return fmt.Sprintf("%v", obj)
}
