package logger

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

const DefaultLogLevel = zerolog.DebugLevel

type ILogger interface {
	Debugf(f string, v ...interface{})
	Infof(f string, v ...interface{})
	Warningf(f string, v ...interface{})
	Errorf(f string, v ...interface{})
	Error(e error)
	Fatalf(f string, v ...interface{})
	Fatal(e error)
}

type LogConfig struct {
	Level    string `env:"LOG_LEVEL" envDefault:"debug"`
	IsPretty bool   `env:"LOG_PRETTY" envDefault:"false"`
}

type _Logger struct {
	logger zerolog.Logger
}

func NewLogger(appName string, c LogConfig) ILogger {
	log.Logger = log.With().Str("module", appName).Logger()

	if c.IsPretty {
		log.Logger = log.Logger.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	}

	// Set correct log level
	if l := c.Level; l != "" {
		level, err := zerolog.ParseLevel(l)
		if err != nil {
			fmt.Printf("invalid log level '%s': %s\n", l, err.Error())
			os.Exit(1)
		}
		log.Info().Msgf("Set global log level to %s", level)
		zerolog.SetGlobalLevel(level)
	} else {
		log.Info().Msgf("Global log level set to default: %s", DefaultLogLevel)
		zerolog.SetGlobalLevel(DefaultLogLevel)
	}

	return &_Logger{
		logger: log.Logger,
	}
}

func (l *_Logger) Debugf(f string, v ...interface{}) {
	l.logger.Debug().Msgf(f, v...)
}

func (l *_Logger) Infof(f string, v ...interface{}) {
	l.logger.Info().Msgf(f, v...)
}

func (l *_Logger) Warningf(f string, v ...interface{}) {
	l.logger.Warn().Msgf(f, v...)
}

func (l *_Logger) Errorf(f string, v ...interface{}) {
	l.logger.Error().Msgf(f, v...)
}

func (l *_Logger) Error(e error) {
	l.logger.Error().Msg(e.Error())
}

func (l *_Logger) Fatalf(f string, v ...interface{}) {
	l.logger.Fatal().Msgf(f, v...)
}

func (l *_Logger) Fatal(e error) {
	l.logger.Fatal().Msg(e.Error())
}
