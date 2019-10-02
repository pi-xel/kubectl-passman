package main

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_formatResponse(t *testing.T) {
	fixture := `{"apiVersion":"client.authentication.k8s.io/v1beta1","kind":"ExecCredential","status":{"token":"my-bearer-token"}}`
	require.Equal(t, fixture, formatResponse(&response{}))
}

func Test_formatResponse_is_json(t *testing.T) {
	require.True(t, json.Valid([]byte(formatResponse(&response{}))))
}

func Test_formatResponse_populate_defaults(t *testing.T) {
	require.Contains(t, formatResponse(&response{}), "apiVersion")
}
func Test_formatResponse_override_defaults(t *testing.T) {
	require.Contains(t, formatResponse(&response{Kind: "foo"}), `"kind":"foo"`)
}

func Test_keychainFetcher_NoKeychainError(t *testing.T) {
	// panicker := func() {
	// 	// TODO: MOCK keychain.QueryItem(query) returns err=1
	// 	keychainFetcher("error")
	// }
	// require.PanicsWithValue(t, "unable to connect to keychain", panicker)
}
func Test_keychainFetcher_NoItemFoundError(t *testing.T) {
	panicker := func() {
		// TODO: MOCK keychain.QueryItem(query) returns empty array
		keychainFetcher("doesn't exist")
	}
	require.PanicsWithValue(t, "item doesn't exist", panicker)
}

func Test_keychainFetcher_ItemFound(t *testing.T) {
	var expected = "RSA"
	// TODO: MOCK keychain.QueryItem(query) returns "RSARSA"
	require.Contains(t, keychainFetcher("gabriel"), expected)
}

func Test_opgetter_happy(t *testing.T) {
	var expected = "RSA"
	// TODO: MOCK exec.Command() returns {"details": {"fields":[{"name": "password", "value": "RSARSA"}]}}
	require.Contains(t, opgetter("gabriel"), expected)
}

func Test_opgetter_op_fail(t *testing.T) {
	// TODO: MOCK exec.Command() returns (ERROR)  item mykubecreds not found
}

func Test_opgetter_password_not_found(t *testing.T) {
	// TODO: MOCK exec.Command() returns {"details": {"fields":[{"name": "notpassword", "value": "RSARSA"}]}}
}
