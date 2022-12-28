package lib

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestGetLogger(t *testing.T) {
	// Reset globalLog to nil
	globalLog = nil

	// Get the logger
	logger := GetLogger()

	// Assert that the returned logger is not nil
	assert.NotNil(t, logger)
}

func TestFxLogger_LogEvent(t *testing.T) {
	// Create a test logger
	zapLogger, err := zap.NewDevelopment()
	if err != nil {
		t.Fatal(err)
	}
	otelzapLogger := otelzap.New(zapLogger)
	logger := newSugaredLogger(otelzapLogger)
	fxLogger := &FxLogger{Logger: logger}

	onStartExecuting := &fxevent.OnStartExecuting{
		FunctionName: "function",
		CallerName:   "caller",
	}
	fxLogger.LogEvent(onStartExecuting)

	onStartExecuted := &fxevent.OnStartExecuted{
		FunctionName: "function",
		CallerName:   "caller",
		Method:       "jhon",
		Runtime:      time.Duration(10 * time.Millisecond),
	}
	fxLogger.LogEvent(onStartExecuted)

	onStartExecutedWithError := &fxevent.OnStartExecuted{
		FunctionName: "function",
		CallerName:   "caller",
		Method:       "jhon",
		Err:          errors.New("test Error"),
		Runtime:      time.Duration(10 * time.Millisecond),
	}
	fxLogger.LogEvent(onStartExecutedWithError)

	onStopExecuting := &fxevent.OnStopExecuting{
		FunctionName: "function",
		CallerName:   "caller",
	}
	fxLogger.LogEvent(onStopExecuting)

	onStopExecuted := &fxevent.OnStopExecuted{
		FunctionName: "function",
		CallerName:   "caller",
		Runtime:      time.Duration(10 * time.Millisecond),
	}
	fxLogger.LogEvent(onStopExecuted)

	onStopExecutedWithError := &fxevent.OnStopExecuted{
		FunctionName: "function",
		CallerName:   "caller",
		Err:          errors.New("test Error"),
		Runtime:      time.Duration(10 * time.Millisecond),
	}
	fxLogger.LogEvent(onStopExecutedWithError)

	supplied := &fxevent.Supplied{
		TypeName:   "foo",
		ModuleName: "barr",
		Err:        errors.New("test Error"),
	}
	fxLogger.LogEvent(supplied)

	provided := &fxevent.Provided{
		ConstructorName: "foo",
		OutputTypeNames: []string{},
		ModuleName:      "barr",
		Err:             errors.New("test Error"),
	}
	fxLogger.LogEvent(provided)

	decorated := &fxevent.Decorated{
		DecoratorName:   "foo",
		ModuleName:      "barr",
		OutputTypeNames: []string{},
		Err:             errors.New("test Error"),
	}
	fxLogger.LogEvent(decorated)

	invoking := &fxevent.Invoking{
		FunctionName: "foo",
		ModuleName:   "barr",
	}
	fxLogger.LogEvent(invoking)

	started := &fxevent.Started{
		Err: errors.New("test Error"),
	}
	fxLogger.LogEvent(started)

	loggerInitialized := &fxevent.LoggerInitialized{
		Err: errors.New("test Error"),
	}
	fxLogger.LogEvent(loggerInitialized)
}

func TestNewLogger(t *testing.T) {
	t.Run("creates a logger with debug level", func(t *testing.T) {
		env := Env{LogLevel: "debug"}
		logger := newLogger(env)
		if logger.SugaredLogger == nil {
			t.Errorf("newLogger() = %v, want %v", logger, &Logger{SugaredLogger: nil})
		}
		if logger.Level() != zapcore.DebugLevel {
			t.Errorf("newLogger() log level = %v, want %v", logger.Level(), zapcore.DebugLevel)
		}
	})
	t.Run("creates a logger with info level", func(t *testing.T) {
		env := Env{LogLevel: "info"}
		logger := newLogger(env)
		if logger.SugaredLogger == nil {
			t.Errorf("newLogger() = %v, want %v", logger, &Logger{SugaredLogger: nil})
		}
		if logger.Level() != zapcore.InfoLevel {
			t.Errorf("newLogger() log level = %v, want %v", logger.Level(), zapcore.DebugLevel)
		}
	})

	t.Run("creates a logger with warn level", func(t *testing.T) {
		env := Env{LogLevel: "warn"}
		logger := newLogger(env)
		if logger.SugaredLogger == nil {
			t.Errorf("newLogger() = %v, want %v", logger, &Logger{SugaredLogger: nil})
		}
		if logger.Level() != zapcore.WarnLevel {
			t.Errorf("newLogger() log level = %v, want %v", logger.Level(), zapcore.DebugLevel)
		}
	})

	t.Run("creates a logger with error level", func(t *testing.T) {
		env := Env{LogLevel: "error"}
		logger := newLogger(env)
		if logger.SugaredLogger == nil {
			t.Errorf("newLogger() = %v, want %v", logger, &Logger{SugaredLogger: nil})
		}
		if logger.Level() != zapcore.ErrorLevel {
			t.Errorf("newLogger() log level = %v, want %v", logger.Level(), zapcore.DebugLevel)
		}
	})

	t.Run("creates a logger with fatal level", func(t *testing.T) {
		env := Env{LogLevel: "fatal"}
		logger := newLogger(env)
		if logger.SugaredLogger == nil {
			t.Errorf("newLogger() = %v, want %v", logger, &Logger{SugaredLogger: nil})
		}
		if logger.Level() != zapcore.FatalLevel {
			t.Errorf("newLogger() log level = %v, want %v", logger.Level(), zapcore.DebugLevel)
		}
	})

	t.Run("creates a logger with panic level", func(t *testing.T) {
		env := Env{LogLevel: ""}
		logger := newLogger(env)
		if logger.SugaredLogger == nil {
			t.Errorf("newLogger() = %v, want %v", logger, &Logger{SugaredLogger: nil})
		}
		if logger.Level() != zapcore.PanicLevel {
			t.Errorf("newLogger() log level = %v, want %v", logger.Level(), zapcore.DebugLevel)
		}
	})
}

func TestGetGormLogger(t *testing.T) {
	logger := newSugaredLogger(otelzapLogger)
	gormLogger := logger.GetGormLogger()

	// Ensure that the returned value is a GormLogger
	_, ok := gormLogger.(*GormLogger)
	if !ok {
		t.Errorf("GetGormLogger() = %T, want *GormLogger", gormLogger)
	}
}

func TestGetFxLogger(t *testing.T) {
	logger := &Logger{}

	fxLogger := logger.GetFxLogger()
	if reflect.TypeOf(fxLogger) != reflect.TypeOf(&FxLogger{}) {
		t.Errorf("GetFxLogger() = %v, want %v", reflect.TypeOf(fxLogger), reflect.TypeOf(&FxLogger{}))
	}
	if fxLogger == nil {
		t.Errorf("GetFxLogger().Logger = nil, want non-nil")
	}
}
