package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	expected := "test-api-key"
	testHeader := http.Header{}
	testHeader.Add("Authorization", "ApiKey "+expected)

	result, err := GetAPIKey(testHeader)
	if err != nil {
		t.Errorf("Test failed: Unexpected error: %v", err)
		return
	}

	if result != expected {
		t.Errorf("Test failed: expected %s, got %s", expected, result)
	}
}
