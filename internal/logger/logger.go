package logger

import (
	"fmt"

	"github.com/go-logr/logr"
)

type customLogger struct {
	name      string
	verbosity int
}

func (l *customLogger) Enabled(level int) bool {
	return level <= l.verbosity
}

func (l *customLogger) Info(level int, msg string, keysAndValues ...interface{}) {
	l.log(level, msg, keysAndValues...)
}

func (l *customLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	keysAndValues = append([]interface{}{"error", err}, keysAndValues...)
	l.log(0, msg, keysAndValues...)
}

func (l *customLogger) WithValues(keysAndValues ...interface{}) logr.LogSink {
	return l
}

func (l *customLogger) WithName(name string) logr.LogSink {
	return &customLogger{
		name:      l.name + "." + name,
		verbosity: l.verbosity,
	}
}

func (l *customLogger) Init(info logr.RuntimeInfo) {
	// No initialization needed
}

func (l *customLogger) log(level int, msg string, keysAndValues ...interface{}) {
	if !l.Enabled(level) {
		return
	}
	fmt.Printf("%s: %s", l.name, msg)
	if len(keysAndValues) > 0 {
		fmt.Printf(" |")
		for i := 0; i < len(keysAndValues); i += 2 {
			fmt.Printf(" %v=%v", keysAndValues[i], keysAndValues[i+1])
		}
	}
	fmt.Println()
}

func NewConsoleLogger(verbosity int, jsonFormat bool) logr.Logger {
	return logr.New(&customLogger{
		name:      "ROOT",
		verbosity: verbosity,
	})
}
