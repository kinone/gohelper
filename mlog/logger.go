package mlog

import (
	"log"
	"os"
)

const (
	LDEBUG     = 100
	LINFO      = 200
	LNOTICE    = 250
	LWARNING   = 300
	LERROR     = 400
	LCRITICAL  = 500
	LALERT     = 550
	LEMERGENCY = 600
)

var levelPrefix = map[int]string{
	LDEBUG:     "[DEBUG]",
	LINFO:      "[INFO]",
	LNOTICE:    "[NOTICE]",
	LWARNING:   "[WARNING]",
	LERROR:     "[ERROR]",
	LCRITICAL:  "[CRITICAL]",
	LALERT:     "[ALERT]",
	LEMERGENCY: "[EMERGENCY]",
}

var levelString = map[string]int{
	"debug":     LDEBUG,
	"info":      LINFO,
	"notice":    LNOTICE,
	"warning":   LWARNING,
	"error":     LERROR,
	"critical":  LCRITICAL,
	"alert":     LALERT,
	"emergency": LEMERGENCY,
}

func ConvertLogLevel(level string) int {
	l, e := levelString[level]
	if !e {
		return LDEBUG
	}

	return l
}

type Logger struct {
	*log.Logger
	level      int
	strLogfile string
	logFile    *os.File
}

func NewLogger(file string, strLevel string) *Logger {
	level := ConvertLogLevel(strLevel)

	if len(file) == 0 { // 没有配置log file
		return &Logger{log.New(os.Stderr, "", log.LstdFlags), level, file, nil}
	}

	var err error
	logFile, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if nil != err {
		log.Fatal("can not open log file", err)
	}

	logger := log.New(logFile, "", log.LstdFlags)

	return &Logger{logger, level, file, logFile}
}

func (l *Logger) Log(level int, msg string, context ...interface{}) {
	lp, exist := levelPrefix[level]
	if !exist { // 不存在的level
		return
	}

	if l.level > level { // 不需要处理的level
		return
	}

	l.Println(lp, msg, context)
}

func (l *Logger) Debug(msg string, context ...interface{}) {
	l.Log(LDEBUG, msg, context)
}

func (l *Logger) Info(msg string, context ...interface{}) {
	l.Log(LINFO, msg, context)
}

func (l *Logger) Notice(msg string, context ...interface{}) {
	l.Log(LNOTICE, msg, context)
}

func (l *Logger) Warning(msg string, context ...interface{}) {
	l.Log(LWARNING, msg, context)
}

func (l *Logger) Error(msg string, context ...interface{}) {
	l.Log(LERROR, msg, context)
}

func (l *Logger) Critical(msg string, context ...interface{}) {
	l.Log(LCRITICAL, msg, context)
}

func (l *Logger) Alert(msg string, context ...interface{}) {
	l.Log(LALERT, msg, context)
}

func (l *Logger) Emergency(msg string, context ...interface{}) {
	l.Log(LEMERGENCY, msg, context)
}

func (l *Logger) Reload() error {
	lf, err := os.OpenFile(l.strLogfile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if nil != err {
		return err
	}
	l.SetOutput(lf)
	l.logFile.Close()
	l.logFile = lf

	return nil
}

func (l *Logger) Close() {
	l.logFile.Close()
}
