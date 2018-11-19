package mlog

type NullLogger struct {
}

func NewNullLogger() *NullLogger {
	return &NullLogger{}
}

func (l *NullLogger) Debug(msg string, v ...interface{}) {

}

func (l *NullLogger) Debugf(format string, v ...interface{}) {

}

func (l *NullLogger) Info(msg string, v ...interface{}) {

}
func (l *NullLogger) Infof(format string, v ...interface{}) {

}

func (l *NullLogger) Notice(msg string, v ...interface{}) {

}

func (l *NullLogger) Noticef(format string, v ...interface{}) {

}

func (l *NullLogger) Warning(msg string, v ...interface{}) {

}

func (l *NullLogger) Warningf(format string, v ...interface{}) {

}

func (l *NullLogger) Error(msg string, v ...interface{}) {

}

func (l *NullLogger) Errorf(format string, v ...interface{}) {

}

func (l *NullLogger) Critical(msg string, v ...interface{}) {

}

func (l *NullLogger) Criticalf(format string, v ...interface{}) {

}

func (l *NullLogger) Alert(msg string, v ...interface{}) {

}

func (l *NullLogger) Alertf(msg string, v ...interface{}) {

}

func (l *NullLogger) Emergency(msg string, v ...interface{}) {

}

func (l *NullLogger) Emergencyf(format string, v ...interface{}) {

}

func (l *NullLogger) Print(v ...interface{}) {

}

func (l *NullLogger) Println(v ...interface{}) {

}

func (l *NullLogger) Printf(format string, v ...interface{}) {

}

func (l *NullLogger) Fatal(v ...interface{}) {

}

func (l *NullLogger) Fatalln(v ...interface{}) {

}

func (l *NullLogger) Fatalf(format string, v ...interface{}) {

}

func (l *NullLogger) Panic(v ...interface{}) {

}

func (l *NullLogger) Panicln(v ...interface{}) {

}

func (l *NullLogger) Panicf(format string, v ...interface{}) {

}
