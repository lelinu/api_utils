package lzap

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/url"
	"os"
	"runtime"
	"strings"
)

var (
	logConfig zap.Config
)

const(
	Windows = "windows"
	LogFileName = "log.txt"
)

//Service wrapper for zap logger
type Service struct {
	log *zap.Logger
}

type MyField zap.Field

func NewService(level string, logPath string) *Service {
	var service = &Service{}
	service.init(level, logPath)

	return service
}

//init function
func (s *Service) init(level string, logPath string){

	atomicLevel := zap.NewAtomicLevel()
	err := atomicLevel.UnmarshalText([]byte(level))
	if err != nil{
		atomicLevel = zap.NewAtomicLevelAt(zap.InfoLevel)
	}

	logConfig = zap.Config{
		Encoding:    "json",
		OutputPaths: s.getOutputPaths(logPath),
		Level: atomicLevel,
		EncoderConfig: zapcore.EncoderConfig {
			TimeKey:        "timestamp",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.RFC3339TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
	}

	//build configuration
	log, err := logConfig.Build()
	if err != nil{
		panic(err)
	}
	s.log = log

	defer s.log.Sync()
}

func (s *Service) SetLogLevel(level string){
	logConfig.Level.SetLevel(s.mapToZapLevels(level))
	s.log.Sync()
}

func (s *Service) Info(msg string, tags ...string){
	s.log.Info(msg, s.parseFields(nil, tags...)...)
	s.log.Sync()
}

func (s *Service) Warn(msg string, tags ...string){
	s.log.Warn(msg, s.parseFields(nil, tags...)...)
	s.log.Sync()
}

func (s *Service) Debug(msg string, tags...string){
	s.log.Debug(msg, s.parseFields(nil, tags...)...)
	s.log.Sync()
}

func (s *Service) Error(msg string, err error, tags...string){
	s.log.Error(msg, s.parseFields(err, tags...)...)
	s.log.Sync()
}

func (s *Service) mapToZapLevels(level string) zapcore.Level {
	switch level{
		case "warn": return zapcore.WarnLevel
		case "debug": return zapcore.DebugLevel
		case "error": return zapcore.ErrorLevel
		default: return zapcore.InfoLevel
	}
}

func (s *Service) parseFields(err error, tags...string) []zap.Field {

	var result []zap.Field
	var length = len(tags)

	if err != nil{
		length +=1
	}

	result = make([]zap.Field, length)
	for i, tag := range tags{
		els := strings.Split(tag, ":")
		result[i] = zap.Field{
			Key:       els[0],
			String:    els[1],
			Type:      zapcore.StringType,
		}
	}

	if err != nil{
		result[length-1] = zap.Field{
			Key:       "error",
			Type:      zapcore.ErrorType,
			Interface: err,
		}
	}

	return result
}

func (s *Service) getOutputPaths(logPath string) []string{
	var outputPaths = []string{"stdout"}

	if strings.TrimSpace(logPath) != ""{

		if runtime.GOOS == Windows {
			zap.RegisterSink("winfile", s.newWinFileSink)
			outputPaths = append(outputPaths, "winfile:///" + logPath)

		}else {
			outputPaths = append(outputPaths, logPath)
		}
	}

	return outputPaths
}

func (s *Service) newWinFileSink(logPath *url.URL) (zap.Sink, error) {
	return os.OpenFile(logPath.Path[1:], os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
}
