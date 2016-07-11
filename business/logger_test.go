package business

type fakeLog struct {
}

func (*fakeLog) Criticalf(format string, v ...interface{}) {}
func (*fakeLog) Errorf(format string, v ...interface{})    {}
func (*fakeLog) Warningf(format string, v ...interface{})  {}
func (*fakeLog) Noticef(format string, v ...interface{})   {}
func (*fakeLog) Infof(format string, v ...interface{})     {}
func (*fakeLog) Debugf(format string, v ...interface{})    {}
