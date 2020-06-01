package llogrus

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

//Service wrapper for logrus logging
type Service struct {
	log *logrus.Logger
}

func NewService(level string) *Service {

	var service = &Service{}
	service.init(level)

	return service
}

func (s *Service) init(level string){
	s.log = &logrus.Logger{
		Out: os.Stdout,
		Formatter: &logrus.JSONFormatter{},
	}

	s.SetLogLevel(level)
}

func (s *Service) SetLogLevel(level string){
	logLevel, err := logrus.ParseLevel(level)
	if err != nil{
		logLevel = logrus.DebugLevel
	}

	s.log.Level = logLevel
}

func (s *Service) Info(msg string, tags ...string){
	s.log.WithFields(s.parseFields(tags...)).Info(msg)
}

func (s *Service) Warn(msg string, tags ...string){
	s.log.WithFields(s.parseFields(tags...)).Warn(msg)
}

func (s *Service) Debug(msg string, tags...string){
	s.log.WithFields(s.parseFields(tags...)).Debug(msg)
}

func (s *Service) Error(msg string, err error, tags...string){
	msg = fmt.Sprintf("%v ERROR - %v", msg, err)
	s.log.WithFields(s.parseFields(tags...)).Error(msg)
}

func (s *Service) parseFields(tags...string) logrus.Fields {
	result := make(logrus.Fields, len(tags))
	for _, tag := range tags{
		els := strings.Split(tag, ":")
		result[strings.TrimSpace(els[0])] = strings.TrimSpace(els[1])
	}
	return result
}
