package httpapi

import (
	"context"
	"net"
	"net/http"

	"github.com/google/uuid"
)

type handler struct {
	serverID string
	next     http.Handler
}

func newServerMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		h := &handler{
			serverID: uuid.New().String(), //TODO: use concrete identifier
			next:     next,
		}
		return h
	}
}

// ServeHTTP implements http.Handler.
func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := h.createRequestContext(r)
	h.next.ServeHTTP(w, r.WithContext(ctx))
}

func (h handler) createRequestContext(req *http.Request) context.Context {
	ctx := req.Context()

	var (
		xForwardedFor   = req.Header.Get("X-FORWARDED-FOR")
		xForwardedProto = req.Header.Get("X-FORWARDED-PROTO")
		xForwardedHost  = req.Header.Get("X-FORWARDED-HOST")
	)

	if xForwardedFor != "" {
		ctx = context.WithValue(ctx, contextKeyXForwardedFor, xForwardedFor)
	}
	if xForwardedProto != "" {
		ctx = context.WithValue(ctx, contextKeyXForwardedProto, xForwardedProto)
	}
	if xForwardedHost != "" {
		ctx = context.WithValue(ctx, contextKeyXForwardedHost, xForwardedHost)
	}

	ip, _, _ := net.SplitHostPort(req.RemoteAddr)
	ctx = context.WithValue(ctx, contextKeyClientIP, ip)
	ctx = context.WithValue(ctx, contextKeyEndpoint, req.URL.RequestURI())

	ctx = context.WithValue(ctx, contextKeyServerID, h.serverID)

	return ctx
}
