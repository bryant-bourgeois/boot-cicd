package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKeyHappyCase(t *testing.T) {
	headers := http.Header{}
	headers.Add("Authorization", "ApiKey 12345")
	val, err := GetAPIKey(headers)
	if err != nil {
		t.Errorf("Auth header is malformed: %s", val)
	}
}

func TestGetAPIKeyNoPrefix(t *testing.T) {
	headers := http.Header{}
	headers.Add("Authorization", "12345")
	_, err := GetAPIKey(headers)
	if err == nil {
		t.Error("This should be returning an error")
	}
}

func TestGetAPIKeyNoSpace(t *testing.T) {
	headers := http.Header{}
	headers.Add("Authorization", "ApiKey12345")
	_, err := GetAPIKey(headers)
	if err == nil {
		t.Error("This should be returning an error")
	}
}
