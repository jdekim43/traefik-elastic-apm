package traefik_elastic_apm

import (
	"context"
	"go.elastic.co/apm/module/apmhttp/v2"
	"go.elastic.co/apm/v2"
	"net/http"
)

type Config struct {
	ServiceName string
	Environment string
}

func CreateConfig() *Config {
	return &Config{}
}

type ServiceTracing struct {
	next   http.Handler
	tracer *apm.Tracer
	name   string
}

func NewServiceTracing(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	var serviceName = config.ServiceName
	if len(serviceName) == 0 {
		serviceName = name
	}

	tracer, err := apm.NewTracerOptions(apm.TracerOptions{
		ServiceName:        serviceName,
		ServiceEnvironment: config.Environment,
	})

	if err != nil {
		return nil, err
	}

	return &ServiceTracing{
		tracer: tracer,
		next:   next,
		name:   name,
	}, nil
}

func (a *ServiceTracing) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	wrappedHandler := apmhttp.Wrap(a.next, apmhttp.WithTracer(a.tracer))
	wrappedHandler.ServeHTTP(rw, req)
}
