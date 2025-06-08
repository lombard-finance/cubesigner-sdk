package client

import "net/http"

func (cli *Client) buildMfaHeaders(mfaHeaders MfaHeaders) map[string]string {
	return map[string]string{
		"x-cubist-mfa-id":           mfaHeaders.Id,
		"x-cubist-mfa-org-id":       cli.orgID,
		"x-cubist-mfa-confirmation": mfaHeaders.Confirmation,
	}
}

func (cli *Client) addExtraHeaders(req *http.Request) {
	for k, v := range cli.extraHeaders {
		req.Header.Add(k, v)
	}
}
