package api

import (
	"net/http/httptest"
	"path/filepath"
	"soundspace/internal/store"
	"strings"
	"testing"
)

func TestHealthEndpoint(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "api.db"))
	defer s.Close()
	r := httptest.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	New(s).Handler().ServeHTTP(w, r)
	if w.Code != 204 {
		t.Fatalf("%d", w.Code)
	}
}
func TestCreateEndpoint(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "api2.db"))
	defer s.Close()
	r := httptest.NewRequest("POST", "/records", strings.NewReader(`{"id":"x","title":"X","location":"L","status":"draft"}`))
	w := httptest.NewRecorder()
	New(s).Handler().ServeHTTP(w, r)
	if w.Code != 201 {
		t.Fatalf("%d %s", w.Code, w.Body.String())
	}
}
