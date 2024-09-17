package logger

import (
	"context"
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
)

func New(conf Configuration) *Logger {
	var skipFrameCount = 3 + conf.SkipFrameCount
	var l *Logger

	if conf.LogFormat == LogFormatJSON {
		l = &Logger{
			stdoutLogger: createLogger(os.Stdout, skipFrameCount),
			stderrLogger: createLogger(os.Stderr, skipFrameCount),
		}

		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	} else {
		l = &Logger{
			stdoutLogger: createLogger(createConsoleWriter(os.Stdout), skipFrameCount),
			stderrLogger: createLogger(createConsoleWriter(os.Stderr), skipFrameCount),
		}
	}

	zerolog.SetGlobalLevel(parseLogLevel(conf.LogLevel))

	return l
}

func (l Logger) Error(ctx context.Context, err error, logs ...KV) {
	event := l.stderrLogger.Error()

	for _, log := range logs {
		event = event.Any(log.Key, log.Value)
	}

	event.Err(err).Send()
}

func (l Logger) Warn(ctx context.Context, msg string, logs ...KV) {
	event := l.stdoutLogger.Warn()

	for _, log := range logs {
		event = event.Any(log.Key, log.Value)
	}

	event.Msg(msg)
}

func (l Logger) Info(ctx context.Context, msg string, logs ...KV) {
	event := l.stdoutLogger.Info()

	for _, log := range logs {
		event = event.Any(log.Key, log.Value)
	}

	event.Msg(msg)
}

func (l Logger) Debug(ctx context.Context, msg string, logs ...KV) {
	event := l.stdoutLogger.Debug()

	for _, log := range logs {
		event = event.Any(log.Key, log.Value)
	}

	event.Msg(msg)
}

func (l *Logger) Close(ctx context.Context) error {
	return nil
}

func parseLogLevel(logLevel LogLevel) zerolog.Level {
	switch logLevel {
	case LogLevelInfo:
		return zerolog.InfoLevel
	case LogLevelWarn:
		return zerolog.WarnLevel
	default:
		return zerolog.DebugLevel
	}
}

func createLogger(w io.Writer, skipFrameCount int) zerolog.Logger {
	return zerolog.New(w).
		With().
		Timestamp().
		CallerWithSkipFrameCount(skipFrameCount).
		Logger()
}

func createConsoleWriter(w io.Writer) zerolog.ConsoleWriter {
	return zerolog.ConsoleWriter{
		Out:        w,
		TimeFormat: time.DateTime,
	}
}

func ParseLogLevel(logLevel string) LogLevel {
	switch logLevel {
	case "WARN":
		return LogLevelWarn
	case "INFO":
		return LogLevelInfo
	default:
		return LogLevelDebug
	}
}

func ParseLogFormat(logFormat string) LogFormat {
	switch logFormat {
	case "JSON":
		return LogFormatJSON
	default:
		return LogFormatPlain
	}
}
