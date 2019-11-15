package mlog

type StdLogger interface {
	Print(v ...interface{})
	Println(v ...interface{})
	Printf(format string, v ...interface{})
}

type LevelLogger interface {
	StdLogger

	Debug(v ...interface{})
	Debugf(format string, v ...interface{})

	Info(v ...interface{})
	Infof(format string, v ...interface{})

	Notice(v ...interface{})
	Noticef(format string, v ...interface{})

	Warning(v ...interface{})
	Warningf(format string, v ...interface{})

	Error(v ...interface{})
	Errorf(format string, v ...interface{})

	Critical(v ...interface{})
	Criticalf(format string, v ...interface{})

	Alert(v ...interface{})
	Alertf(format string, v ...interface{})

	Emergency(v ...interface{})
	Emergencyf(format string, v ...interface{})

	Close()
}

const (
	LDEBUG     = 100
	LINFO      = 200
	LNOTICE    = 250
	LWARNING   = 300
	LERROR     = 400
	LCRITICAL  = 500
	LALERT     = 550
	LEMERGENCY = 600
	NONE       = 1000
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
	NONE:       "[NONE]",
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
