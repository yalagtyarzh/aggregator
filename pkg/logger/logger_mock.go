package logger

type MockLogger struct {
}

func NewMockLogger() ILogger {
	return &MockLogger{}
}

func (l *MockLogger) Debugf(f string, v ...interface{}) {

}

func (l *MockLogger) Infof(f string, v ...interface{}) {

}

func (l *MockLogger) Warningf(f string, v ...interface{}) {

}

func (l *MockLogger) Errorf(f string, v ...interface{}) {

}

func (l *MockLogger) Error(e error) {

}

func (l *MockLogger) Fatalf(f string, v ...interface{}) {

}

func (l *MockLogger) Fatal(e error) {

}
