package main_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"suggestion-engine/cmd/api"
)

// Test the endpoint /suggest responds correctly
func TestSuggestEndpoint(t *testing.T) {
	server := main.NewServer()

	req := httptest.NewRequest(http.MethodGet, "/suggest?q=python", nil)
	rec := httptest.NewRecorder()

	// call handler
	server.HandleSuggest(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status code 200, but got %d", rec.Code)
	}

	var data map[string]interface{}
	if err := json.Unmarshal(rec.Body.Bytes(), &data); err != nil {
		t.Fatalf("error on decoded JSON: %v", err)
	}

	// check response struct
	if _, ok := data["query"]; !ok {
		t.Errorf("response JSON should be contains the field 'query'")
	}

	if _, ok := data["suggestions"]; !ok {
		t.Errorf("response JSON should be contains the field 'suggestions'")
	}

	suggestions := data["suggestions"].([]interface{})
	if len(suggestions) == 0 {
		t.Errorf("expected at least a suggestions, but got 0")
	}
}
