package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCalculateHandler(t *testing.T) {
	tests := []struct {
		payload      string
		expectedCode int
		expectedBody string
	}{
		{`{"expression": "2 + 2"}`, http.StatusOK, `{"result":"4.000000"}`},
		{`{"expression": "10 / 0"}`, http.StatusUnprocessableEntity, `{"error":"division by zero"}`},
		{`{"expression": "invalid"}`, http.StatusUnprocessableEntity, `{"error":"invalid expression"}`},
	}

	for _, test := range tests {
		req := httptest.NewRequest(http.MethodPost, "/api/v1/calculate", bytes.NewBuffer([]byte(test.payload)))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		calculateHandler(w, req)

		res := w.Result()
		if res.StatusCode != test.expectedCode {
			t.Errorf("expected status %d, got %d", test.expectedCode, res.StatusCode)
		}

		body := w.Body.String()
		if body != test.expectedBody {
			t.Errorf("expected body %s, got %s", test.expectedBody, body)
		}
	}
}
