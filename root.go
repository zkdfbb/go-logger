package logger

import (
	"fmt"
	"github.com/mattn/go-isatty"
	"os"
)

var (
	log *Logger
)

func init() {
	log = NewLogger()
	if isatty.IsTerminal(os.Stdout.Fd()) {
		log.Detach("console")
		config := &ConsoleConfig{Color: true}
		log.Attach("console", LOGGER_LEVEL_DEBUG, config)
	}
}

func Detach(adapterName string) error {
	return log.Detach(adapterName)
}

func Attach(adapterName string, level int, config Config) error {
	return log.Attach(adapterName, level, config)
}

func SetLevel(level string) {
	outputs := []*outputLogger{}
	for _, output := range log.outputs {
		output.Level = log.LoggerLevel(level)
	}
	log.outputs = outputs
}

func LoggerLevel(levelStr string) int {
	return log.LoggerLevel(levelStr)
}

func Emergency(msg string, a ...interface{}) {
	if len(a) > 0 {
		msg := fmt.Sprintf(msg, a...)
		log.Writer(LOGGER_LEVEL_EMERGENCY, msg)
	} else {
		log.Writer(LOGGER_LEVEL_EMERGENCY, msg)
	}
}

func Emergencyf(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	log.Writer(LOGGER_LEVEL_EMERGENCY, msg)
}

func Alert(msg string, a ...interface{}) {
	if len(a) > 0 {
		msg := fmt.Sprintf(msg, a...)
		log.Writer(LOGGER_LEVEL_ALERT, msg)
	} else {
		log.Writer(LOGGER_LEVEL_ALERT, msg)
	}
}

func Alertf(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	log.Writer(LOGGER_LEVEL_ALERT, msg)
}

func Critical(msg string, a ...interface{}) {
	if len(a) > 0 {
		msg := fmt.Sprintf(msg, a...)
		log.Writer(LOGGER_LEVEL_CRITICAL, msg)
	} else {
		log.Writer(LOGGER_LEVEL_CRITICAL, msg)
	}
}

func Criticalf(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	log.Writer(LOGGER_LEVEL_CRITICAL, msg)
}

func Error(msg string, a ...interface{}) {
	if len(a) > 0 {
		msg := fmt.Sprintf(msg, a...)
		log.Writer(LOGGER_LEVEL_ERROR, msg)
	} else {
		log.Writer(LOGGER_LEVEL_ERROR, msg)
	}
}

func Errorf(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	log.Writer(LOGGER_LEVEL_ERROR, msg)
}

func Warning(msg string, a ...interface{}) {
	if len(a) > 0 {
		msg := fmt.Sprintf(msg, a...)
		log.Writer(LOGGER_LEVEL_WARNING, msg)
	} else {
		log.Writer(LOGGER_LEVEL_WARNING, msg)
	}
}

func Warningf(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	log.Writer(LOGGER_LEVEL_WARNING, msg)
}

func Warn(msg string, a ...interface{}) {
	if len(a) > 0 {
		msg := fmt.Sprintf(msg, a...)
		log.Writer(LOGGER_LEVEL_WARNING, msg)
	} else {
		log.Writer(LOGGER_LEVEL_WARNING, msg)
	}
}

func Warnf(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	log.Writer(LOGGER_LEVEL_WARNING, msg)
}

func Notice(msg string, a ...interface{}) {
	if len(a) > 0 {
		msg := fmt.Sprintf(msg, a...)
		log.Writer(LOGGER_LEVEL_NOTICE, msg)
	} else {
		log.Writer(LOGGER_LEVEL_NOTICE, msg)
	}
}

func Noticef(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	log.Writer(LOGGER_LEVEL_NOTICE, msg)
}

func Info(msg string, a ...interface{}) {
	if len(a) > 0 {
		msg := fmt.Sprintf(msg, a...)
		log.Writer(LOGGER_LEVEL_INFO, msg)
	} else {
		log.Writer(LOGGER_LEVEL_INFO, msg)
	}
}

func Infof(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	log.Writer(LOGGER_LEVEL_INFO, msg)
}

func Debug(msg string, a ...interface{}) {
	if len(a) > 0 {
		msg := fmt.Sprintf(msg, a...)
		log.Writer(LOGGER_LEVEL_DEBUG, msg)
	} else {
		log.Writer(LOGGER_LEVEL_DEBUG, msg)
	}
}

func Debugf(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	log.Writer(LOGGER_LEVEL_DEBUG, msg)
}
