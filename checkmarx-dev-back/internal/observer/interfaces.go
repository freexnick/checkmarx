package observer

import "context"

type Logger interface {
	Error(context.Context, error, ...KV)
	Warn(context.Context, string, ...KV)
	Info(context.Context, string, ...KV)
	Debug(context.Context, string, ...KV)
	Close(context.Context) error
}
