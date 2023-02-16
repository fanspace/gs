package core

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"os"
	"strings"
)

const (
	LevelDebug = "debug"
	LevelInfo  = "info"
	LevelWarn  = "warn"
	LevelError = "error"
)

type Field = zapcore.Field

var Int64 = zap.Int64
var Int = zap.Int
var String = zap.String
var Any = zap.Any
var Err = zap.Error

type LoggerConfiguration struct {
	EnableConsole bool
	ConsoleJson   bool
	ConsoleLevel  string
	EnableFile    bool
	FileJson      bool
	FileLevel     string
	FileLocation  string
}

type Logger struct {
	zap          *zap.Logger
	consoleLevel zap.AtomicLevel
	fileLevel    zap.AtomicLevel
}

func getZapLevel(level string) zapcore.Level {
	switch level {
	case LevelInfo:
		return zapcore.InfoLevel
	case LevelWarn:
		return zapcore.WarnLevel
	case LevelDebug:
		return zapcore.DebugLevel
	case LevelError:
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}

func makeEncoder(json bool) zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	if json {
		return zapcore.NewJSONEncoder(encoderConfig)
	}
	return zapcore.NewConsoleEncoder(encoderConfig)
}

type LoggerGrpcProduce struct {
	Info   string
	Module string
}

/*
func (ep *LoggerGrpcProduce)Write(p []byte) (n int, err error)  {
	go SendLogger("", "", "error", string(p))
	return 1, nil
}
*/

func NewLogger(config *LoggerConfiguration) *Logger {
	cores := []zapcore.Core{}
	logger := &Logger{
		consoleLevel: zap.NewAtomicLevelAt(getZapLevel(config.ConsoleLevel)),
		fileLevel:    zap.NewAtomicLevelAt(getZapLevel(config.FileLevel)),
	}

	if config.EnableConsole {
		writer := zapcore.Lock(os.Stdout)
		core := zapcore.NewCore(makeEncoder(config.ConsoleJson), writer, logger.consoleLevel)
		cores = append(cores, core)
	}

	if config.EnableFile {
		writer := zapcore.AddSync(&lumberjack.Logger{
			Filename: config.FileLocation,
			MaxSize:  100,
			Compress: true,
		})
		core := zapcore.NewCore(makeEncoder(config.FileJson), writer, logger.fileLevel)
		cores = append(cores, core)
	}
	/*

		ep := new(LoggerGrpcProduce)
		writer := zapcore.AddSync(ep)
		core := zapcore.NewCore(makeEncoder(config.FileJson), writer, zapcore.ErrorLevel)
		cores = append(cores, core)
	*/

	combinedCore := zapcore.NewTee(cores...)

	logger.zap = zap.New(combinedCore,
		zap.AddCallerSkip(2),
		zap.AddCaller(),
	)

	return logger
}

func (l *Logger) ChangeLevels(config *LoggerConfiguration) {
	l.consoleLevel.SetLevel(getZapLevel(config.ConsoleLevel))
	l.fileLevel.SetLevel(getZapLevel(config.FileLevel))
}

func (l *Logger) SetConsoleLevel(level string) {
	l.consoleLevel.SetLevel(getZapLevel(level))
}

func (l *Logger) With(fields ...Field) *Logger {
	newlogger := *l
	newlogger.zap = newlogger.zap.With(fields...)
	return &newlogger
}

func (l *Logger) StdLog(fields ...Field) *log.Logger {
	return zap.NewStdLog(l.With(fields...).zap.WithOptions(getStdLogOption()))
}

func (l *Logger) Debug(message string, fields ...Field) {
	l.zap.Debug(message, fields...)
}

func (l *Logger) Info(message string, fields ...Field) {
	l.zap.Info(message, fields...)
}

func (l *Logger) Warn(message string, fields ...Field) {
	l.zap.Warn(message, fields...)
}

func (l *Logger) Error(message string, fields ...Field) {

	l.zap.Error(message, fields...)
}

func (l *Logger) Critical(message string, fields ...Field) {
	l.zap.Error(message, fields...)
}

type stdLogLevelInterpreterCore struct {
	wrappedCore zapcore.Core
}

func stdLogInterpretZapEntry(entry zapcore.Entry) zapcore.Entry {
	message := entry.Message
	if strings.Index(message, "[DEBUG]") == 0 {
		entry.Level = zapcore.DebugLevel
		entry.Message = message[7:]
	} else if strings.Index(message, "[DEBG]") == 0 {
		entry.Level = zapcore.DebugLevel
		entry.Message = message[6:]
	} else if strings.Index(message, "[WARN]") == 0 {
		entry.Level = zapcore.WarnLevel
		entry.Message = message[6:]
	} else if strings.Index(message, "[ERROR]") == 0 {
		entry.Level = zapcore.ErrorLevel
		entry.Message = message[7:]
	} else if strings.Index(message, "[EROR]") == 0 {
		entry.Level = zapcore.ErrorLevel
		entry.Message = message[6:]
	} else if strings.Index(message, "[ERR]") == 0 {
		entry.Level = zapcore.ErrorLevel
		entry.Message = message[5:]
	} else if strings.Index(message, "[INFO]") == 0 {
		entry.Level = zapcore.InfoLevel
		entry.Message = message[6:]
	}

	return entry
}

func (s *stdLogLevelInterpreterCore) Enabled(lvl zapcore.Level) bool {
	return s.wrappedCore.Enabled(lvl)
}

func (s *stdLogLevelInterpreterCore) With(fields []zapcore.Field) zapcore.Core {
	return s.wrappedCore.With(fields)
}

func (s *stdLogLevelInterpreterCore) Check(entry zapcore.Entry, checkedEntry *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	entry = stdLogInterpretZapEntry(entry)
	return s.wrappedCore.Check(entry, checkedEntry)
}

func (s *stdLogLevelInterpreterCore) Write(entry zapcore.Entry, fields []zapcore.Field) error {
	entry = stdLogInterpretZapEntry(entry)
	return s.wrappedCore.Write(entry, fields)
}

func (s *stdLogLevelInterpreterCore) Sync() error {
	return s.wrappedCore.Sync()
}

func getStdLogOption() zap.Option {
	return zap.WrapCore(
		func(core zapcore.Core) zapcore.Core {
			return &stdLogLevelInterpreterCore{core}
		},
	)
}

func defaultLog(level, msg string, fields ...Field) {
	log := struct {
		Level   string  `json:"level"`
		Message string  `json:"msg"`
		Fields  []Field `json:"fields,omitempty"`
	}{
		level,
		msg,
		fields,
	}

	if b, err := json.Marshal(log); err != nil {
		fmt.Printf(`{"level":"error","msg":"failed to encode log message"}%s`, "\n")
	} else {
		fmt.Printf("%s\n", b)

	}
}

func defaultDebugLog(msg string, fields ...Field) {
	defaultLog("debug", msg, fields...)
}

func defaultInfoLog(msg string, fields ...Field) {
	defaultLog("info", msg, fields...)
}

func defaultWarnLog(msg string, fields ...Field) {
	defaultLog("warn", msg, fields...)
}

func defaultErrorLog(msg string, fields ...Field) {
	defaultLog("error", msg, fields...)
}

func defaultCriticalLog(msg string, fields ...Field) {
	defaultLog("error", msg, fields...)
}

var globalLogger *Logger

func InitGlobalLogger(logger *Logger) {
	globalLogger = logger
	Debug = globalLogger.Debug
	Info = globalLogger.Info
	Warn = globalLogger.Warn
	Error = globalLogger.Error
	Critical = globalLogger.Critical
}

func RedirectStdLog(logger *Logger) {
	zap.RedirectStdLogAt(logger.zap.With(zap.String("source", "stdlog")), zapcore.InfoLevel)
}

type LogFunc func(string, ...Field)

func GloballyDisableDebugLogForTest() {
	globalLogger.consoleLevel.SetLevel(zapcore.ErrorLevel)
}

func GloballyEnableDebugLogForTest() {
	globalLogger.consoleLevel.SetLevel(zapcore.DebugLevel)
}

var Debug LogFunc = defaultDebugLog
var Info LogFunc = defaultInfoLog
var Warn LogFunc = defaultWarnLog
var Error LogFunc = defaultErrorLog
var Critical LogFunc = defaultCriticalLog
