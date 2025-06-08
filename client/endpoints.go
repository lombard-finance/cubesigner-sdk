package client

import (
	"fmt"
	"net/url"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/pkg/errors"
)

// Parameter names
const (
	ParamOrgID     = "org_id"
	ParamPubkey    = "pubkey"
	ParamKeyID     = "key_id"
	ParamMfaID     = "mfa_id"
	ParamRoleID    = "role_id"
	ParamAddress   = "address"
	QueryParamVote = "mfa_vote"
)

// Endpoint patterns
var (
	// Babylon endpoints
	SignBabylonStaking      = fmt.Sprintf("/v0/org/{%s}/babylon/staking/{%s}", ParamOrgID, ParamPubkey)
	SignBabylonRegistration = fmt.Sprintf("/v0/org/{%s}/babylon/registration/{%s}", ParamOrgID, ParamPubkey)

	// Blob endpoints
	SignBlob = fmt.Sprintf("/v1/org/{%s}/blob/sign/{%s}", ParamOrgID, ParamKeyID)

	// BTC endpoints
	SignBtcTaproot = fmt.Sprintf("/v0/org/{%s}/btc/taproot/sign/{%s}", ParamOrgID, ParamPubkey)
	SignBtcSegWit  = fmt.Sprintf("/v0/org/{%s}/btc/sign/{%s}", ParamOrgID, ParamPubkey)
	SignBtcPsbt    = fmt.Sprintf("/v0/org/{%s}/btc/psbt/sign/{%s}", ParamOrgID, ParamPubkey)

	// EVM endpoints
	SignEvmEip712 = fmt.Sprintf("/v0/org/{%s}/evm/eip712/sign/{%s}", ParamOrgID, ParamPubkey)

	// Key endpoints
	CreateKey   = fmt.Sprintf("/v0/org/{%s}/keys", ParamOrgID)
	GetKeyInOrg = fmt.Sprintf("/v0/org/{%s}/keys/{%s}", ParamOrgID, ParamKeyID)

	// MFA endpoints
	ListMfaRequests = fmt.Sprintf("/v0/org/{%s}/mfa", ParamOrgID)
	MfaRequest      = fmt.Sprintf("/v0/org/{%s}/mfa/{%s}", ParamOrgID, ParamMfaID)

	// Role endpoints
	CreateRoleToken = fmt.Sprintf("/v0/org/{%s}/roles/{%s}/tokens", ParamOrgID, ParamRoleID)
	AddKeysToRole   = fmt.Sprintf("/v0/org/{%s}/roles/{%s}/add_keys", ParamOrgID, ParamRoleID)
	GetKeysInRole   = fmt.Sprintf("/v0/org/{%s}/roles/{%s}/keys", ParamOrgID, ParamRoleID)

	// Signer session endpoints
	RefreshToken = fmt.Sprintf("/v1/org/{%s}/token/refresh", ParamOrgID)

	// User endpoints
	AboutMeLegacy = "/v0/about_me"
	AboutMe       = fmt.Sprintf("/v0/org/{%s}/user/me", ParamOrgID)
)

// buildEndpoint replaces path parameters in endpoint patterns
func buildEndpoint(pattern string, params map[string]interface{}) string {
	if len(params) == 0 {
		return pattern
	}

	result := pattern
	for key, value := range params {
		placeholder := "{" + key + "}"
		stringValue := parameterToString(value, "")
		escapedValue := url.PathEscape(stringValue)
		result = strings.ReplaceAll(result, placeholder, escapedValue)
	}
	return result
}

// buildEndpoint replaces path parameters in endpoint patterns and adds query parameters if any
func buildEndpointWithQuery(pattern string, pathParams map[string]interface{}, queryParams map[string]interface{}) string {
	basePath := buildEndpoint(pattern, pathParams)

	if len(queryParams) == 0 {
		return basePath
	}

	u, _ := url.Parse(basePath)
	q := u.Query()
	for key, value := range queryParams {
		q.Set(key, parameterToString(value, ""))
	}
	u.RawQuery = q.Encode()

	return u.String()
}

// BuildFullEndpoint builds the endpoint, auto-replaces the org_id, and parses as a URL
func (c *Client) BuildFullEndpoint(pattern string, pathParams map[string]interface{}, queryParams map[string]interface{}) (*url.URL, error) {
	if pathParams == nil {
		pathParams = make(map[string]interface{})
	}
	if _, exists := pathParams[ParamOrgID]; !exists {
		pathParams[ParamOrgID] = c.orgID
	}

	endpoint := buildEndpointWithQuery(pattern, pathParams, queryParams)

	// Parse and combine with base URL
	requestEndpoint, err := url.Parse(fmt.Sprintf("%s%s", strings.TrimSuffix(c.base.String(), "/"), endpoint))
	if err != nil {
		return nil, errors.Wrap(err, "invalid endpoint")
	}

	return requestEndpoint, nil
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

// MatchesEndpoint checks if an actual path matches an endpoint pattern
func MatchesEndpoint(actualPath, endpointPattern string) bool {
	// Convert endpoint pattern to regex pattern
	// /v0/org/{org_id}/btc/sign/{pubkey} -> ^/v0/org/[^/]+/btc/sign/[^/]+$
	pattern := regexp.QuoteMeta(endpointPattern)
	pattern = regexp.MustCompile(`\{[^}]+\}`).ReplaceAllString(pattern, `[^/]+`)
	pattern = "^" + pattern + "$"

	matched, _ := regexp.MatchString(pattern, actualPath)
	return matched
}
