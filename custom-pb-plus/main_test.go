package main

import (
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/pocketbase/pocketbase/core"
    pbrouter "github.com/pocketbase/pocketbase/tools/router"
)

// eventFactory creates a minimal RequestEvent for testing routes.
func eventFactory(w http.ResponseWriter, r *http.Request) (*core.RequestEvent, pbrouter.EventCleanupFunc) {
    return &core.RequestEvent{Event: pbrouter.Event{Response: w, Request: r}}, nil
}

func buildTestMux() http.Handler {
    r := pbrouter.NewRouter[*core.RequestEvent](eventFactory)
    r.GET("/hello", handleHello)
    r.GET("/healthz", handleHealthz)
    r.GET("/version", handleVersion)
    mux, _ := r.BuildMux()
    return mux
}

func TestHello(t *testing.T) {
    mux := buildTestMux()
    req := httptest.NewRequest(http.MethodGet, "/hello", nil)
    rec := httptest.NewRecorder()
    mux.ServeHTTP(rec, req)

    if rec.Code != http.StatusOK {
        t.Fatalf("expected 200, got %d", rec.Code)
    }
    if got := rec.Body.String(); got != "Hello world!" {
        t.Fatalf("unexpected body: %q", got)
    }
}

func TestHealthz(t *testing.T) {
    mux := buildTestMux()
    req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
    rec := httptest.NewRecorder()
    mux.ServeHTTP(rec, req)

    if rec.Code != http.StatusOK {
        t.Fatalf("expected 200, got %d", rec.Code)
    }
    var data map[string]string
    if err := json.Unmarshal(rec.Body.Bytes(), &data); err != nil {
        t.Fatalf("invalid json: %v", err)
    }
    if data["status"] != "ok" {
        t.Fatalf("expected status=ok, got %v", data)
    }
}

func TestVersion(t *testing.T) {
    mux := buildTestMux()
    req := httptest.NewRequest(http.MethodGet, "/version", nil)
    rec := httptest.NewRecorder()
    mux.ServeHTTP(rec, req)

    if rec.Code != http.StatusOK {
        t.Fatalf("expected 200, got %d", rec.Code)
    }
    var data map[string]string
    if err := json.Unmarshal(rec.Body.Bytes(), &data); err != nil {
        t.Fatalf("invalid json: %v", err)
    }
    if _, ok := data["version"]; !ok {
        t.Fatalf("missing version key: %v", data)
    }
    if _, ok := data["commit"]; !ok {
        t.Fatalf("missing commit key: %v", data)
    }
    if _, ok := data["builtAt"]; !ok {
        t.Fatalf("missing builtAt key: %v", data)
    }
}

