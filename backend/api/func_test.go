package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Mock data for successful response
var mockResponse = `{"key": "value"}`

// Test struct to match the JSON structure
type TestData struct {
	Key string `json:"key"`
}

func TestGetData_Success(t *testing.T) {
	// Create a mock server with a successful response
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(mockResponse))
	}))
	defer server.Close()

	var result TestData
	GetData(server.URL, &result)

	// Validate the result
	if result.Key != "value" {
		t.Errorf("Expected key to be 'value', got '%s'", result.Key)
	}
}
