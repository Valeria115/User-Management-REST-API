package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateUser(t *testing.T) {
	setupTestData()
	tests := map[string]struct {
		body       string
		wantStatus int
		wantInBody string
	}{
		"success": {
			body:       `{"name": "Charlie", "email": "charlie@example.com"}`,
			wantStatus: http.StatusCreated,
			wantInBody: "charlie@example.com",
		},
		"no_email": {
			body:       `{"name": "Bob"}`,
			wantStatus: http.StatusBadRequest,
			wantInBody: "Name and Email required",
		},
		"invalid_json": {
			body:       `{invalid}`,
			wantStatus: http.StatusBadRequest,
			wantInBody: "Invalid JSON",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			body := bytes.NewBufferString(tt.body)
			req, _ := http.NewRequest("POST", "/users", body)

			w := httptest.NewRecorder()

			createUser(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("status = %d, want %d", w.Code, tt.wantStatus)
			}

			if tt.wantInBody != "" && !strings.Contains(w.Body.String(), tt.wantInBody) {
				t.Errorf("body doesnt contain %q, got %s", tt.wantInBody, w.Body.String())
			}
		})
	}
}
