package zako

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouting(t *testing.T) {
	app := New()

	app.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello, zako!")
	})

	tests := []struct {
		name       string
		method     string
		path       string
		wantStatus int
		wantBody   string
	}{
		{
			name:       "GET success",
			method:     http.MethodGet,
			path:       "/hello",
			wantStatus: http.StatusOK,
			wantBody:   "hello, zako!",
		},
		{
			name:       "Not found",
			method:     http.MethodGet,
			path:       "/unknown",
			wantStatus: http.StatusNotFound,
		},
		{
			name:       "Method not allowed",
			method:     http.MethodPost,
			path:       "/hello",
			wantStatus: http.StatusMethodNotAllowed,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, tt.path, nil)
			rec := httptest.NewRecorder()

			app.ServeHTTP(rec, req)

			if rec.Code != tt.wantStatus {
				t.Errorf("status = %d, want %d", rec.Code, tt.wantStatus)
			}

			if tt.wantBody != "" && rec.Body.String() != tt.wantBody {
				t.Errorf("body = %q, want %q", rec.Body.String(), tt.wantBody)
			}
		})
	}
}
