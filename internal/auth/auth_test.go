package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKeyNoAuthHeader(t *testing.T) {
	header := http.Header{}
	expectedError := "no authorization header included"

	_, err := GetAPIKey(header)
	if err.Error() != expectedError {
		t.Fatalf("Expected error to be %s, got %s", expectedError, err.Error())
	}
}

func TestGetAPIKeyMalformedAuthHeader(t *testing.T) {
	header := http.Header{}
	header.Add("Authorization", "Bearer xxxxx")
	expectedError := "malformed authorization header"

	_, err := GetAPIKey(header)
	if err.Error() != expectedError {
		t.Fatalf("Expected error to be %s, got %s", expectedError, err.Error())
	}
}

func TestGetAPIKeySuccess(t *testing.T) {
	header := http.Header{}
	header.Add("Authorization", "ApiKey xxxxx")
	expectedKey := "xxxxx"

	key, err := GetAPIKey(header)
	if err != nil {
		t.Fatalf("Expected no error, got %s", err.Error())
	}
	if key != expectedKey {
		t.Fatalf("Expected key to be %s, got %s", expectedKey, key)
	}
}
