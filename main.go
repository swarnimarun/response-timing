package reponsetiming

import (
	"context"
	"net/http"
	"time"
)

// add config to select header name
type Config struct {
	TimingHeaderName string `json:"TimingHeaderName,omitempty"`
}

func CreateConfig() *Config {
	return &Config{TimingHeaderName: "X-TIME-TAKEN"}
}

type ResponseTiming struct {
	next   http.Handler
	name   string
	config *Config
}

func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &ResponseTiming{
		next:   next,
		name:   name,
		config: config,
	}, nil
}

func (a *ResponseTiming) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	t := time.Now()
	a.next.ServeHTTP(resp, req)
	resp.Header().Set(a.config.TimingHeaderName, time.Since(t).String())
}
