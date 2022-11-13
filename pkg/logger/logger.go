package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"runtime"
	"time"
)

type Level int8

type Fields map[string]interface{}

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic
)

func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "debug"
	case LevelInfo:
		return "info"
	case LevelWarn:
		return "warn"
	case LevelError:
		return "error"
	case LevelFatal:
		return "fatal"
	case LevelPanic:
		return "panic"
	}
	return ""
}

type Logger struct {
	logger  *log.Logger
	ctx     context.Context
	fields  Fields
	callers []string
}

func NewLogger(w io.Writer, prefix string, flag int) *Logger {
	l := log.New(w, prefix, flag)
	return &Logger{
		logger: l,
	}
}

func (l *Logger) clone() *Logger {
	nl := *l
	return &nl
}

// WithFields 设置日志公共字段。
func (l *Logger) WithFields(f Fields) *Logger {
	nl := l.clone()
	if nl.fields == nil {
		nl.fields = make(Fields)
	}
	for k, v := range f {
		nl.fields[k] = v
	}
	return nl
}

// WithContext 设置日志上下文属性。
func (l *Logger) WithContext(ctx context.Context) *Logger {
	nl := l.clone()
	nl.ctx = ctx
	return nl
}

// WithCaller 设置当前某一层调用栈的信息（程序计数器、文件信息、行号）。
func (l *Logger) WithCaller(skip int) *Logger {
	nl := l.clone()
	pc, file, line, ok := runtime.Caller(skip)
	if ok {
		f := runtime.FuncForPC(pc)
		nl.callers = []string{fmt.Sprintf("%s: %d %s", file, line, f.Name())}
	}

	return nl
}

// WithCallersFrames 设置当前的整个调用栈信息。
func (l *Logger) WithCallersFrames() *Logger {
	maxCallerDepth := 25
	minCallerDepth := 1
	var callers []string
	pcs := make([]uintptr, maxCallerDepth)
	depth := runtime.Callers(minCallerDepth, pcs)
	frames := runtime.CallersFrames(pcs[:depth])
	for frame, more := frames.Next(); more; frame, more = frames.Next() {
		callers = append(callers, fmt.Sprintf("%s: %d %s", frame.File, frame.Line, frame.Function))
		if !more {
			break
		}
	}
	nl := l.clone()
	nl.callers = callers
	return nl
}

func (l *Logger) LogFormat(level Level, message string) map[string]interface{} {
	data := make(Fields, len(l.fields)+4)
	data["level"] = level.String()
	data["time"] = time.Now().Local().UnixNano()
	data["message"] = message
	data["callers"] = l.callers
	if len(l.fields) > 0 {
		for k, v := range l.fields {
			if _, ok := data[k]; !ok {
				data[k] = v
			}
		}
	}
	return data
}

func (l *Logger) WriteLog(level Level, message string) {
	body, _ := json.Marshal(l.WithCaller(3).LogFormat(level, message))
	content := string(body)
	switch level {
	case LevelDebug:
		l.logger.Print(content)
	case LevelInfo:
		l.logger.Print(content)
	case LevelWarn:
		l.logger.Print(content)
	case LevelError:
		l.logger.Print(content)
	case LevelFatal:
		l.logger.Fatal(content)
	case LevelPanic:
		l.logger.Panic(content)
	}
}

func (l *Logger) DeBug(v ...interface{}) {
	l.WriteLog(LevelDebug, fmt.Sprint(v...))
}

func (l *Logger) DeBugf(format string, v ...interface{}) {
	l.WriteLog(LevelDebug, fmt.Sprintf(format, v...))
}

func (l *Logger) Info(v ...interface{}) {
	l.WriteLog(LevelInfo, fmt.Sprint(v...))
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.WriteLog(LevelInfo, fmt.Sprintf(format, v...))
}

func (l *Logger) Warn(v ...interface{}) {
	l.WriteLog(LevelWarn, fmt.Sprint(v...))
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.WriteLog(LevelWarn, fmt.Sprintf(format, v...))
}

func (l *Logger) Error(v ...interface{}) {
	l.WriteLog(LevelError, fmt.Sprint(v...))
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.WriteLog(LevelError, fmt.Sprintf(format, v...))
}

func (l *Logger) Fatal(v ...interface{}) {
	l.WriteLog(LevelFatal, fmt.Sprint(v...))
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.WriteLog(LevelFatal, fmt.Sprintf(format, v...))
}

func (l *Logger) Panic(v ...interface{}) {
	l.WriteLog(LevelPanic, fmt.Sprint(v...))
}

func (l *Logger) Panicf(format string, v ...interface{}) {
	l.WriteLog(LevelPanic, fmt.Sprintf(format, v...))
}
