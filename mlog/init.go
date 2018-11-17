package mlog

type StdLogger interface {
	Print(v ...interface{})
	Println(v ...interface{})
	Printf(format string, v ...interface{})

	Fatal(v ...interface{})
	Fatalln(v ...interface{})
	Fatalf(format string, v ...interface{})

	Panic(v ...interface{})
	Panicln(v ...interface{})
	Panicf(format string, v ...interface{})
}

type LevelLogger interface {
	StdLogger

	Debug(msg string, v ...interface{})
	Debugf(format string, v ...interface{})

	Info(msg string, v ...interface{})
	Infof(format string, v ...interface{})

	Notice(msg string, v ...interface{})
	Noticef(format string, v ...interface{})

	Warning(msg string, v ...interface{})
	Warningf(format string, v ...interface{})

	Error(msg string, v ...interface{})
	Errorf(format string, v ...interface{})

	Critical(msg string, v ...interface{})
	Criticalf(format string, v ...interface{})

	Alert(msg string, v ...interface{})
	Alertf(msg string, v ...interface{})

	Emergency(msg string, v ...interface{})
	Emergencyf(format string, v ...interface{})

	Reload() error
	Close()
}
