package contract

type Logger interface {
	Debug(v ...interface{})
	Debugf(format string, v ...interface{})
	Debugw(format string, v ...interface{})

	Info(v ...interface{})
	Infof(format string, v ...interface{})
	Infow(format string, v ...interface{})

	Error(v ...interface{})
	Errorf(format string, v ...interface{})
	Errorw(format string, v ...interface{})

}
