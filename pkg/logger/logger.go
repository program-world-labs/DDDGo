package logger

import (
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel/trace"
)

// Interface -.
type Interface interface {
	Trace() *Event
	Debug() *Event
	Info() *Event
	Warn() *Event
	Error() *Event
	Fatal() *Event
	Panic() *Event
	Err(error) *Event
	Log() *Event
	Span(trace.Span) *Event
	ErrorSpan(trace.Span) *Event
	Print(...interface{})
	Printf(string, ...interface{})
	Write(p []byte) (n int, err error)
	Output(io.Writer) Logger

}

var _ Interface = (*Logger)(nil)

type Logger struct {
	*zerolog.Logger
	projectID string
}

type Event struct {
	*zerolog.Event
}

// See: https://cloud.google.com/logging/docs/reference/v2/rest/v2/LogEntry#LogSeverity
func logLevelSeverity(l zerolog.Level) string {
	return map[zerolog.Level]string{
		zerolog.DebugLevel: "DEBUG",
		zerolog.InfoLevel:  "INFO",
		zerolog.WarnLevel:  "WARNING",
		zerolog.ErrorLevel: "ERROR",
		zerolog.PanicLevel: "CRITICAL",
		zerolog.FatalLevel: "CRITICAL",
	}[l]
}

// NewProductionLogger returns a configured logger for production.
// It outputs info level and above logs with sampling.
const severityFieldName = "severity"
const timeFieldName = "time"

func NewProductionLogger(projectID string) *Logger {
	logLevel := zerolog.InfoLevel
	zerolog.SetGlobalLevel(logLevel)

	zerolog.LevelFieldName = severityFieldName
	zerolog.LevelFieldMarshalFunc = func(l zerolog.Level) string {
		return logLevelSeverity(l)
	}
	zerolog.TimestampFieldName = timeFieldName
	zerolog.TimeFieldFormat = time.RFC3339Nano

	// default sampler
	sampler := &zerolog.BasicSampler{N: 1}

	logger := zerolog.New(os.Stderr).Sample(sampler).With().Timestamp().Logger()

	return &Logger{&logger, projectID}
}

// NewDevelopmentLogger returns a configured logger for development.
// It outputs debug level and above logs, and sampling is disabled.
func NewDevelopmentLogger(projectID string) *Logger {
	logLevel := zerolog.DebugLevel
	zerolog.SetGlobalLevel(logLevel)

	zerolog.LevelFieldName = severityFieldName
	zerolog.LevelFieldMarshalFunc = func(l zerolog.Level) string {
		return logLevelSeverity(l)
	}
	zerolog.TimestampFieldName = timeFieldName
	zerolog.TimeFieldFormat = time.RFC3339Nano

	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()

	return &Logger{&logger, projectID}
}

// To use method chain we need followings

func (l *Logger) Trace() *Event {
	e := l.Logger.Trace()

	return &Event{e}
}

func (l *Logger) Debug() *Event {
	e := l.Logger.Debug()

	return &Event{e}
}

func (l *Logger) Info() *Event {
	e := l.Logger.Info()

	return &Event{e}
}

func (l *Logger) Warn() *Event {
	e := l.Logger.Warn()

	return &Event{e}
}

func (l *Logger) Error() *Event {
	e := l.Logger.Error()

	return &Event{e}
}

func (l *Logger) Err(err error) *Event {
	e := l.Logger.Err(err)

	return &Event{e}
}

func (l *Logger) Fatal() *Event {
	e := l.Logger.Fatal()

	return &Event{e}
}

func (l *Logger) Panic() *Event {
	e := l.Logger.Panic()

	return &Event{e}
}

func (l *Logger) WithLevel(level zerolog.Level) *Event {
	e := l.Logger.WithLevel(level)

	return &Event{e}
}

func (l *Logger) Log() *Event {
	e := l.Logger.Log()

	return &Event{e}
}

func (l *Logger) Print(v ...interface{}) {
	l.Logger.Print(v...)
}

func (l *Logger) Printf(format string, v ...interface{}) {
	l.Logger.Printf(format, v...)
}

func (l Logger) Write(p []byte) (n int, err error) {
	n, err = l.Logger.Write(p)

	return n, err
}

func (l Logger) Output(w io.Writer) Logger {
	logger := l.Logger.Output(w)

	return Logger{
		&logger,
		l.projectID,
	}
}

func (l *Logger) Span(span trace.Span) *Event {
	e := l.Logger.Info().Str("trace", "projects/"+l.projectID+"/traces/"+span.SpanContext().TraceID().String()).Str("span", span.SpanContext().SpanID().String())

	return &Event{e}
}

func (l *Logger) ErrorSpan(span trace.Span) *Event {
	e := l.Logger.Error().Str("trace", "projects/"+l.projectID+"/traces/"+span.SpanContext().TraceID().String()).Str("span", span.SpanContext().SpanID().String())

	return &Event{e}
}
