package log

import (
	"github.com/sirupsen/logrus"
	"io"
)

type LogService interface {
	Debug(format string, fields ...interface{})
	Info(format string, fields ...interface{})
	Warn(format string, fields ...interface{})
	Error(format string, fields ...interface{})
	Fatal(format string, fields ...interface{})
}

type logServiceV1 struct {
	logger *logrus.Logger
}

func NewLogServiceV1(output io.Writer) LogService {
	logger := logrus.New()
	logger.SetOutput(output)
	return logServiceV1{
		logger: logger,
	}
}

func (service logServiceV1) Debug(format string, fields ...interface{}) {
	service.logger.Debugf(format, fields...)
}

func (service logServiceV1) Info(format string, fields ...interface{}) {
	service.logger.Infof(format, fields...)
}

func (service logServiceV1) Warn(format string, fields ...interface{}) {
	service.logger.Warnf(format, fields...)
}

func (service logServiceV1) Error(format string, fields ...interface{}) {
	service.logger.Errorf(format, fields...)
}

func (service logServiceV1) Fatal(format string, fields ...interface{}) {
	service.logger.Fatalf(format, fields...)
}
