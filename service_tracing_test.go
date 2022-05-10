package traefik_elastic_apm_test

import (
	"context"
	"github.com/jdekim43/traefik-elastic-apm"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServiceTracing(t *testing.T) {
	cfg := traefik_elastic_apm.CreateConfig()
	cfg.ServiceName = "test-service"
	cfg.Environment = "development"

	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	handler, err := traefik_elastic_apm.New(ctx, next, cfg, "service-tracing")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(recorder, req)
}
