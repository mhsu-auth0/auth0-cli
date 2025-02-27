package validators

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNoKey(t *testing.T) {
	err := APIKey("")
	require.EqualError(t, err, "you have not configured API keys yet. To do so, run `auth0 login` which will configure your API keys from Auth0")
}

func TestKeyTooShort(t *testing.T) {
	err := APIKey("123")
	require.EqualError(t, err, "the API key provided is too short, it must be at least 12 characters long")
}

func TestHTTPMethod(t *testing.T) {
	err := HTTPMethod("GET")
	require.NoError(t, err)
}

func TestHTTPMethodInvalid(t *testing.T) {
	err := HTTPMethod("invalid")
	require.Equal(t, "invalid is not an acceptable HTTP method (GET, POST, DELETE)", fmt.Sprintf("%s", err))
}

func TestHTTPMethodLowercase(t *testing.T) {
	err := HTTPMethod("post")
	require.NoError(t, err)
}

func TestRequestSourceAPI(t *testing.T) {
	err := RequestSource("API")
	require.NoError(t, err)
}

func TestRequestSourceDashboard(t *testing.T) {
	err := RequestSource("dashboard")
	require.NoError(t, err)
}

func TestRequestStatusSucceeded(t *testing.T) {
	err := RequestStatus("succeeded")
	require.NoError(t, err)
}

func TestRequestStatusFailed(t *testing.T) {
	err := RequestStatus("failed")
	require.NoError(t, err)
}

func TestRequestStatusInvalid(t *testing.T) {
	err := RequestStatus("invalid")
	require.Equal(t, "invalid is not an acceptable request status (SUCCEEDED, FAILED)", fmt.Sprintf("%s", err))
}

func TestRequestSourceInvalid(t *testing.T) {
	err := RequestSource("invalid")
	require.Equal(t, "invalid is not an acceptable source (API, DASHBOARD)", fmt.Sprintf("%s", err))
}

func TestStatusCode(t *testing.T) {
	err := StatusCode("200")
	require.NoError(t, err)
}

func TestStatusCodeUnusedInStripe(t *testing.T) {
	err := StatusCode("300")
	require.Equal(t, "Provided status code 300 is not in the range of acceptable status codes (200's, 400's, 500's)", fmt.Sprintf("%s", err))
}

func TestStatusCodeType(t *testing.T) {
	err := StatusCodeType("2Xx")
	require.NoError(t, err)
}

func TestStatusCodeTypeUnusedInStripe(t *testing.T) {
	err := StatusCodeType("3XX")
	require.Equal(t, "Provided status code type 3XX is not a valid type (2XX, 4XX, 5XX)", fmt.Sprintf("%s", err))
}

func TestStatusCodeNotXs(t *testing.T) {
	err := StatusCodeType("201")
	require.Equal(t, "Provided status code type 201 is not a valid type (2XX, 4XX, 5XX)", fmt.Sprintf("%s", err))
}
