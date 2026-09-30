package main

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func TestGetUser(t *testing.T) {
	setupTestData()
	tests := map[string]struct {
		id         int
		wantStatus int
		wantInBody string
	}{
		"success": {
			id:         1,
			wantStatus: http.StatusOK,
			wantInBody: "alice@example.com",
		},
		"not_found": {
			id:         999,
			wantStatus: http.StatusNotFound,
			wantInBody: "user not found",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			req, _ := http.NewRequest(http.MethodGet, "/users/"+strconv.Itoa(tt.id), nil)
			w := httptest.NewRecorder()

			getUser(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("status = %d, want %d", w.Code, tt.wantStatus)
			}

			if tt.wantInBody != "" && !strings.Contains(w.Body.String(), tt.wantInBody) {
				t.Errorf("body doesnt contain %q, got %s", tt.wantInBody, w.Body.String())
			}
		})
	}

}
