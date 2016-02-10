package goa

import (
	"fmt"
	"os"

	"golang.org/x/net/context"
	"gopkg.in/inconshreveable/log15.v2"
)

// Logger provides a logging interface can be implemented to
// consume logs from the generated application.  For maximum
// compatibility with external logging packages, the variadic
// `args` argument should contain an odd number of arguments.
// The first of these can be used as the log "message", and the
// remaning will be paired in order as key/value pairs.
// This applies to log15 and logrus.  Other adapters
// may treat this differently.  Goa's logger interface
// does not concern itself with log levels.  Each adapter is free
// to implement them as written here, or apply leveled filters.
type Logger interface {
	Debug(ctx context.Context, args ...interface{})
	Info(ctx context.Context, args ...interface{})
	Warning(ctx context.Context, args ...interface{})
	Error(ctx context.Context, args ...interface{})
	Critical(ctx context.Context, args ...interface{})
	Fatal(ctx context.Context, args ...interface{})
}

// Ensure that our NoopLogger satisfies the interface
var _ Logger = (*NoopLogger)(nil)

// Ensure that our NoopLogger satisfies the interface
var _ Logger = (*Log15Adapter)(nil)

// NoopLogger is a logger that does nothing
type NoopLogger struct{}

// NewNoopLogger returns a new NoopLogger
func NewNoopLogger() *NoopLogger {
	return &NoopLogger{}
}

// Debug consumes a Debug message
func (n *NoopLogger) Debug(ctx context.Context, args ...interface{}) {}

// Info consumes an Info message
func (n *NoopLogger) Info(ctx context.Context, args ...interface{}) {}

// Warning consumes a Warning message
func (n *NoopLogger) Warning(ctx context.Context, args ...interface{}) {}

// Error consumes an Error message
func (n *NoopLogger) Error(ctx context.Context, args ...interface{}) {}

// Critical consumes a Critical message
func (n *NoopLogger) Critical(ctx context.Context, args ...interface{}) {}

// Fatal consumes a fatal message and exits.
func (n *NoopLogger) Fatal(ctx context.Context, args ...interface{}) {
	fmt.Println(args...)
	os.Exit(1)
}

type Log15Adapter struct {
	logger log15.Logger
}

func NewLog15Adapter(logger log15.Logger) *Log15Adapter {
	return &Log15Adapter{logger: logger}
}

// Debug consumes a Debug message
func (l *Log15Adapter) Debug(ctx context.Context, args ...interface{}) {
	if len(args) > 0 {
		first := args[0]
		l.logger.Debug(first.(string), args[1:])
	}
}

// Info consumes an Info message
func (l *Log15Adapter) Info(ctx context.Context, args ...interface{}) {
	if len(args) > 0 {
		first := args[0]
		l.logger.Info(first.(string), args[1:])
	}
}

// Warning consumes a Warning message
func (l *Log15Adapter) Warning(ctx context.Context, args ...interface{}) {
	if len(args) > 0 {
		first := args[0]
		l.logger.Warn(first.(string), args[1:])
	}
}

// Error consumes an Error message
func (l *Log15Adapter) Error(ctx context.Context, args ...interface{}) {
	if len(args) > 0 {
		first := args[0]
		l.logger.Error(first.(string), args[1:])
	}
}

// Critical consumes a Critical message
func (l *Log15Adapter) Critical(ctx context.Context, args ...interface{}) {
	if len(args) > 0 {
		first := args[0]
		l.logger.Crit(first.(string), args[1:])
	}
}

// Fatal consumes a fatal message and exits.
func (l *Log15Adapter) Fatal(ctx context.Context, args ...interface{}) {
	if len(args) > 0 {
		first := args[0]
		l.logger.Crit(first.(string), args[1:])
	}
	os.Exit(1)
}
