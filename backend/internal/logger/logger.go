package logger

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Level string

const (
	INFO  Level = "INFO"
	WARN  Level = "WARN"
	ERROR Level = "ERROR"
	DEBUG Level = "DEBUG"
)

type Logger struct {
	Level Level
}

func New(level Level) *Logger {
	return &Logger{Level: level}
}

// convert to JSON
func (l *Logger) log(level Level, msg string, values ...interface{}) {
	if !l.shouldLog(level) {
		return
	}

	entry := map[string]interface{}{
		"time":    time.Now().Format(time.RFC3339),
		"level":   level,
		"message": msg,
	}

	if len(values) == 1 {
		entry["data"] = values[0]
	} else if len(values) > 1 {
		entry["data"] = values
	}

	b, err := json.Marshal(entry)
	if err != nil {
		fmt.Fprintf(os.Stderr, `{"time":"%s","level":"ERROR","message":"log marshall failed","error":"%v"}`+"\n",
			time.Now().Format(time.RFC3339), err)
		return
	}

	fmt.Fprintln(os.Stdout, string(b))
}

func (l *Logger) shouldLog(level Level) bool {
	order := map[Level]int{
		ERROR: 1,
		WARN:  2,
		INFO:  3,
		DEBUG: 4,
	}
	return order[level] <= order[l.Level]
}

func (l *Logger) Info(msg string, v ...interface{})  { l.log(INFO, msg, v...) }
func (l *Logger) Warn(msg string, v ...interface{})  { l.log(WARN, msg, v...) }
func (l *Logger) Error(msg string, v ...interface{}) { l.log(ERROR, msg, v...) }
func (l *Logger) Debug(msg string, v ...interface{}) { l.log(DEBUG, msg, v...) }
