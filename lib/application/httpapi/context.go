package httpapi

import (
	"context"
)

type contextKey string

var (
	contextKeyServerID        = contextKey("id")
	contextKeyXForwardedHost  = contextKey("xForwardedHost")
	contextKeyXForwardedFor   = contextKey("xForwardedFor")
	contextKeyXForwardedProto = contextKey("xForwardedProto")
	contextKeyEndpoint        = contextKey("endpoint")
	contextKeyClientIP        = contextKey("clientIP")
)

func (c contextKey) String() string {
	return "server" + string(c)
}

// ID gets the name server from context
func ID(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(contextKeyServerID).(string)
	return id, ok
}

// XForwardedFor gets the http address server from context
func XForwardedFor(ctx context.Context) (string, bool) {
	xForwardedFor, ok := ctx.Value(contextKeyXForwardedFor).(string)
	return xForwardedFor, ok
}

// XForwardedProto for identifying the protocol (HTTP or HTTPS) that a client used to connect to your proxy or load balancer. http address server from context
func XForwardedProto(ctx context.Context) (string, bool) {
	xForwardedProto, ok := ctx.Value(contextKeyXForwardedProto).(string)
	return xForwardedProto, ok
}

// Endpoint gets the http address server from context
func Endpoint(ctx context.Context) (string, bool) {
	endpoint, ok := ctx.Value(contextKeyEndpoint).(string)
	return endpoint, ok
}

// ClientIP gets the http address server from context
func ClientIP(ctx context.Context) (string, bool) {
	clientIP, ok := ctx.Value(contextKeyClientIP).(string)
	return clientIP, ok
}
