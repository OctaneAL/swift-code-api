package handlers

import (
	"context"
	"net/http"

	"github.com/OctaneAL/swift-code-api/internal/data"
	"gitlab.com/distributed_lab/logan/v3"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	swiftCodesQCtxKey
	configCtxKey
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}

func CtxSwiftCodesQ(q data.SwiftCodesQ) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, swiftCodesQCtxKey, q)
	}
}

func SwiftCodesQ(r *http.Request) data.SwiftCodesQ {
	return r.Context().Value(swiftCodesQCtxKey).(data.SwiftCodesQ).New()
}
