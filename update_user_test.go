package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func TestUpdateUser(t *testing.T) {
	setupTestData()
	tests := map[string]struct {
		id         int
		body       string
		wantStatus int
		wantInBody string
	}{
		"success": {
			id:         1,
			body:       `{"name": "Bob Updated", "email": "bob_new_unique@example.com"}`,
			wantStatus: http.StatusOK,
			wantInBody: "bob_new_unique@example.com",
		},
		"not_found": {
			id:         999,
			body:       `{"name": "no_name", "email": "no_name@example.com"}`,
			wantStatus: http.StatusNotFound,
			wantInBody: "user not found",
		},
		"invalid_json": {
			id:         2,
			body:       `{invalid}`,
			wantStatus: http.StatusBadRequest,
			wantInBody: "Invalid JSON",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			body := bytes.NewBufferString(tt.body)
			req, _ := http.NewRequest(http.MethodPut, "/users/"+strconv.Itoa(tt.id), body)

			w := httptest.NewRecorder()

			updateUser(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("status = %d, want %d", w.Code, tt.wantStatus)
			}

			if tt.wantInBody != "" && !strings.Contains(w.Body.String(), tt.wantInBody) {
				t.Errorf("body doesnt contain %q, got %s", tt.wantInBody, w.Body.String())
			}
		})
	}
}
