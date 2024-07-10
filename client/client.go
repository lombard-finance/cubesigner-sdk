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
	"reflect"
	"strings"
	"time"

	"github.com/lombard-finance/cubesigner-sdk/client/pagination"
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

func (cli *Client) get(endpoint string, overrideHeaders map[string]string, page *pagination.Page) (io.Reader, error) {
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

	if page != nil {
		page.Apply(requestEndpoint)
		log.WithField("query", requestEndpoint.Query().Encode())
	}

	opCtx, cancel := context.WithTimeout(context.Background(), cli.timeout)
	req, err := http.NewRequestWithContext(opCtx, http.MethodGet, requestEndpoint.String(), nil)
	if err != nil {
		cancel()
		return nil, errors.Wrap(err, "failed to create GET request")
	}

	cli.addExtraHeaders(req)
	if overrideHeaders != nil {
		for key, value := range overrideHeaders {
			req.Header.Set(key, value)
		}
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

func (cli *Client) post(endpoint string, body io.Reader, overrideHeaders map[string]string, page *pagination.Page) (io.Reader, int, error) {
	return cli.requestWithBody(endpoint, http.MethodPost, body, overrideHeaders, page)
}

func (cli *Client) put(endpoint string, body io.Reader, overrideHeaders map[string]string, page *pagination.Page) (io.Reader, int, error) {
	return cli.requestWithBody(endpoint, http.MethodPut, body, overrideHeaders, page)
}

func (cli *Client) patch(endpoint string, body io.Reader, overrideHeaders map[string]string, page *pagination.Page) (io.Reader, int, error) {
	return cli.requestWithBody(endpoint, http.MethodPatch, body, overrideHeaders, page)
}

func (cli *Client) requestWithBody(endpoint string, method string, body io.Reader, overrideHeaders map[string]string, page *pagination.Page) (io.Reader, int, error) {
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

	// replace path variables
	endpoint = strings.Replace(endpoint, ":org_id", url.PathEscape(cli.orgID), -1)

	log := cli.logger.WithFields(map[string]interface{}{
		"address":  cli.address,
		"endpoint": endpoint,
		"method":   method,
	})

	// build url
	requestEndpoint, err := url.Parse(fmt.Sprintf("%s%s", strings.TrimSuffix(cli.base.String(), "/"), endpoint))
	if err != nil {
		return nil, 0, errors.Wrap(err, "invalid endpoint")
	}

	if page != nil {
		page.Apply(requestEndpoint)
		log.WithField("query", requestEndpoint.Query().Encode())
	}

	// build request
	opCtx, cancel := context.WithTimeout(context.Background(), cli.timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(opCtx, method, requestEndpoint.String(), tee)
	if err != nil {
		return nil, 0, errors.Wrap(err, "create request with context")
	}

	// add headers
	cli.addExtraHeaders(req)
	req.Header.Set("Content-type", "application/json")
	//req.Header.Set("Accept", "application/json")

	if overrideHeaders != nil {
		for key, value := range overrideHeaders {
			req.Header.Set(key, value)
		}
	}

	// do the request
	resp, err := cli.client.Do(req)
	if err != nil {
		return nil, 0, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, errors.Wrap(err, "read response")
	}

	log.WithField("status_code", resp.StatusCode)

	statusFamily := resp.StatusCode / 100
	if statusFamily != 2 {
		return nil, 0, errors.Errorf("%s request with status code: %d, message: %s", method, resp.StatusCode, string(responseData))
	}
	return bytes.NewReader(responseData), resp.StatusCode, nil
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
