package log

import "go.uber.org/zap/zapcore"

// Defines common log fields.
const (
	KeyRequestID   string = "requestID"
	KeyUsername    string = "username"
	KeyWatcherName string = "watcher"
)

// Field is an alias for the field structure in the underlying log frame.
type Field = zapcore.Field
