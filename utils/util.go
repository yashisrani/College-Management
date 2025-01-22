package utils

import (
	"flag"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/yashisrani/Go-Backend/model"
)

// In Go, the init() function is a special function that is automatically executed before the main() function. It's used to initialize the package-level variables or perform any necessary setup tasks.

var Logger logrus.Logger

func init() {
	Logger = *logrus.New()
	Logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	Logger.Out = os.Stdout
}

func SetLogger() {
	logLevel := flag.String(model.LogLevel, model.LogLevelInfo, "log-level(debug,error,warn,info)")
	flag.Parse()

	switch logLevel {
	case &model.LogLevelDebug:
		Logger.SetLevel(logrus.DebugLevel)
	case &model.LogLevelError:
		Logger.SetLevel(logrus.ErrorLevel)
	case &model.LogLevelWarn:
		Logger.SetLevel(logrus.WarnLevel)
	default:
		Logger.SetLevel(logrus.InfoLevel)
	}
}

func Log(logLevel, packageLevel, functionName string, message, parameter interface{}) {
	switch logLevel {
	case model.LogLevelDebug:
		if parameter != nil {
			Logger.Debugf("packageLevel: %s, functionName: %s, message: %v, parameter: %v", packageLevel, functionName, message, parameter)
		} else {
			Logger.Debugf("packageLevel: %s, functionName: %s, message: %v", packageLevel, functionName, message)
		}
	case model.LogLevelError:
		if parameter != nil {
			Logger.Errorf("packageLevel: %s, functionName: %s, message: %v, parameter: %v", packageLevel, functionName, message, parameter)
		} else {
			Logger.Errorf("packageLevel: %s, functionName: %s, message: %v", packageLevel, functionName, message)
		}
	case model.LogLevelWarn:
		if parameter != nil {
			Logger.Warnf("packageLevel: %s, functionName: %s, message: %v, parameter: %v", packageLevel, functionName, message, parameter)
		} else {
			Logger.Warnf("packageLevel: %s, functionName: %s, message: %v", packageLevel, functionName, message)
		}
	case model.LogLevelFatal:
		if parameter != nil {
			Logger.Fatalf("packageLevel: %s, functionName: %s, message: %v, parameter: %v", packageLevel, functionName, message, parameter)
		} else {
			Logger.Fatalf("packageLevel: %s, functionName: %s, message: %v", packageLevel, functionName, message)
		}
	default :
		if parameter != nil {
			Logger.Infof("packageLevel: %s, functionName: %s, message: %v, parameter: %v", packageLevel, functionName, message, parameter)
		} else {
			Logger.Infof("packageLevel: %s, functionName: %s, message: %v", packageLevel, functionName, message)
		}
	}
}
