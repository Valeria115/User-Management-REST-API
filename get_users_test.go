package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUsers(t *testing.T) {
	setupTestData()
	tests := map[string]struct {
		wantStatus int
	}{
		"success": {
			wantStatus: http.StatusOK,
		},
		"no_success": {
			wantStatus: http.StatusOK,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			req, _ := http.NewRequest(http.MethodGet, "/user", nil)

			w := httptest.NewRecorder()

			getUsers(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("status=%d, want %d", w.Code, tt.wantStatus)
			}

		})
	}
}
