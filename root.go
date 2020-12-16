package logger

import (
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

func Attach(adapterName string, level int, config Config) error {
	return log.Attach(adapterName, level, config)
}

func Emergency(msg string, a ...interface{}) {
	if len(a) > 0 {
		log.Emergencyf(msg, a...)
	} else {
		log.Emergency(msg)
	}
}

func Emergencyf(format string, a ...interface{}) {
	log.Emergencyf(format, a...)
}

func Alert(msg string, a ...interface{}) {
	if len(a) > 0 {
		log.Alertf(msg, a...)
	} else {
		log.Alert(msg)
	}
}

func Alertf(format string, a ...interface{}) {
	log.Alertf(format, a...)
}

func Critical(msg string, a ...interface{}) {
	if len(a) > 0 {
		log.Criticalf(msg, a...)
	} else {
		log.Critical(msg)
	}
}

func Criticalf(format string, a ...interface{}) {
	log.Criticalf(format, a...)
}

func Error(msg string, a ...interface{}) {
	if len(a) > 0 {
		log.Errorf(msg, a...)
	} else {
		log.Error(msg)
	}
}

func Errorf(format string, a ...interface{}) {
	log.Errorf(format, a...)
}

func Warning(msg string, a ...interface{}) {
	if len(a) > 0 {
		log.Warningf(msg, a...)
	} else {
		log.Warning(msg)
	}
}

func Warningf(format string, a ...interface{}) {
	log.Warningf(format, a...)
}

func Warn(msg string, a ...interface{}) {
	if len(a) > 0 {
		log.Warningf(msg, a...)
	} else {
		log.Warning(msg)
	}
}

func Warnf(format string, a ...interface{}) {
	log.Warningf(format, a...)
}

func Notice(msg string, a ...interface{}) {
	if len(a) > 0 {
		log.Noticef(msg, a...)
	} else {
		log.Notice(msg)
	}
}

func Noticef(format string, a ...interface{}) {
	log.Noticef(format, a...)
}

func Info(msg string, a ...interface{}) {
	if len(a) > 0 {
		log.Infof(msg, a...)
	} else {
		log.Info(msg)
	}
}

func Infof(format string, a ...interface{}) {
	log.Infof(format, a...)
}

func Debug(msg string, a ...interface{}) {
	if len(a) > 0 {
		log.Debugf(msg, a...)
	} else {
		log.Debug(msg)
	}
}

func Debugf(format string, a ...interface{}) {
	log.Debugf(format, a...)
}
