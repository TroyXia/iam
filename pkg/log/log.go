package log

import (
	"github.com/TroyXia/iam/pkg/log/klog"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	std = New(NewOptions())
)

// NB: right now, we always use the equivalent of sugared logging.
// This is necessary, since logr doesn't define non-suggared types,
// and using zap-specific non-suggared types would make uses tied
// directly to Zap.

// infoLogger is a logr.InfoLogger that uses Zap to log at a particular
// level.  The level has already been converted to a Zap level, which
// is to say that `logrLevel = -1*zapLevel`.
type infoLogger struct {
	level zapcore.Level
	log   *zap.Logger
}

// zapLogger is a logr.Logger that uses Zap to log.
type zapLogger struct {
	// NB: this looks very similar to zap.SugaredLogger, but
	// deals with our desire to have multiple verbosity levels.
	zapLogger *zap.Logger
	infoLogger
}

// New create logger by opts which can custmoized by command arguments.
func New(opts *Options) *zapLogger {
	if opts == nil {
		opts = NewOptions()
	}

	var zapLevel zapcore.Level
	if err := zapLevel.UnmarshalText([]byte(opts.Level)); err != nil {
		zapLevel = zapcore.InfoLevel
	}
	encodeLevel := zapcore.CapitalLevelEncoder
	// when output to local path, with color is forbidden
	if opts.Format == consoleFormat && opts.EnableColor {
		encodeLevel = zapcore.CapitalColorLevelEncoder
	}

	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "timestamp",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    encodeLevel,
		EncodeTime:     timeEncoder,
		EncodeDuration: milliSecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	loggerConfig := &zap.Config{
		Level:             zap.NewAtomicLevelAt(zapLevel),
		Development:       opts.Development,
		DisableCaller:     opts.DisableCaller,
		DisableStacktrace: opts.DisableStacktrace,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         opts.Format,
		EncoderConfig:    encoderConfig,
		OutputPaths:      opts.OutputPaths,
		ErrorOutputPaths: opts.ErrorOutputPaths,
	}

	var err error
	l, err := loggerConfig.Build(zap.AddStacktrace(zapcore.PanicLevel), zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
	logger := &zapLogger{
		zapLogger: l.Named(opts.Name),
		infoLogger: infoLogger{
			log:   l,
			level: zap.InfoLevel,
		},
	}
	klog.InitLogger(l)
	zap.RedirectStdLog(l)

	return logger
}

// Infof method output info level log.
func Infof(format string, v ...interface{}) {
	std.zapLogger.Sugar().Infof(format, v...)
}

func (l *zapLogger) Infof(format string, v ...interface{}) {
	l.zapLogger.Sugar().Infof(format, v...)
}

// Warnf method output warning level log.
func Warnf(format string, v ...interface{}) {
	std.zapLogger.Sugar().Warnf(format, v...)
}
