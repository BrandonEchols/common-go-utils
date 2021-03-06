package common

import (
	"fmt"
	"go.uber.org/zap"
)

/*
	This is initially a dummy interface (for testing purposes). Once InitializeLogger is called, this becomes a
	SugaredLogger from the https://github.com/uber-go/zap package.
*/
var Logger ILeveledLogger = LeveledLogger{}

/*
	InitializeLogger is used to initialize the logger to a Zap logger that is either Production, or Development,
	based on the is_production_logger boolean parameter. This will determine if Debug* logs will be output or not.
	@params
		is_production_logger bool True if the Logger should be made as a Zap-Production logger.
*/
func InitializeLogger(is_production_logger bool) {
	if is_production_logger {
		zapLogger, _ := zap.NewProduction(zap.AddStacktrace(zap.PanicLevel))
		Logger = zapLogger.Sugar()
	} else {
		zapLogger, _ := zap.NewDevelopment(zap.AddStacktrace(zap.DPanicLevel))
		Logger = zapLogger.Sugar()
	}
}

/*
	The following dummy functions, struct, and interface, allow for easy testing. Once InitializeLogger has been
	called, the Logger will be a SugaredLogger. See https://github.com/uber-go/zap for more information
*/
type LeveledLogger struct{}

type ILeveledLogger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Panic(args ...interface{})
	Fatal(args ...interface{})
	Debugf(template string, args ...interface{})
	Infof(template string, args ...interface{})
	Warnf(template string, args ...interface{})
	Errorf(template string, args ...interface{})
	DPanicf(template string, args ...interface{})
	Panicf(template string, args ...interface{})
	Fatalf(template string, args ...interface{})
}

func (l LeveledLogger) Debug(args ...interface{}) {
	fmt.Println(args...)
}

func (l LeveledLogger) Info(args ...interface{}) {
	fmt.Println(args...)
}

func (l LeveledLogger) Warn(args ...interface{}) {
	fmt.Println(args...)
}

func (l LeveledLogger) Error(args ...interface{}) {
	fmt.Println(args...)
}

func (l LeveledLogger) Panic(args ...interface{}) {
	fmt.Println(args...)
}

func (l LeveledLogger) Fatal(args ...interface{}) {
	fmt.Println(args...)
}

func (l LeveledLogger) Debugf(template string, args ...interface{}) {
	fmt.Printf(template+"\n", args...)
}

func (l LeveledLogger) Infof(template string, args ...interface{}) {
	fmt.Printf(template+"\n", args...)
}

func (l LeveledLogger) Warnf(template string, args ...interface{}) {
	fmt.Printf(template+"\n", args...)
}

func (l LeveledLogger) Errorf(template string, args ...interface{}) {
	fmt.Printf(template+"\n", args...)
}

func (l LeveledLogger) DPanicf(template string, args ...interface{}) {
	fmt.Printf(template+"\n", args...)
}

func (l LeveledLogger) Panicf(template string, args ...interface{}) {
	fmt.Printf(template+"\n", args...)
}

func (l LeveledLogger) Fatalf(template string, args ...interface{}) {
	fmt.Printf(template+"\n", args...)
}
