package lib_test

import (
	"context"
	"errors"
	"time"

	"core-gin/lib"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	gormlogger "gorm.io/gorm/logger"
)

var _ = Describe("Logger", func() {
	var logger lib.Logger

	BeforeEach(func() {
		logger = lib.GetLogger()
	})

	Context("when calling the GetGormLogger method", func() {
		It("should return a GormLogger struct", func() {
			Expect(logger.GetGormLogger()).To(BeAssignableToTypeOf(&lib.GormLogger{}))
		})
	})

	Context("when calling the GetFxLogger method", func() {
		It("should return a FxLogger struct", func() {
			Expect(logger.GetFxLogger()).To(BeAssignableToTypeOf(&lib.FxLogger{}))
		})
	})
})

var _ = Describe("FxLogger", func() {
	var fxLogger *lib.FxLogger

	BeforeEach(func() {
		// Create a test logger
		zapLogger, err := zap.NewDevelopment()
		Expect(err).NotTo(HaveOccurred())
		otelzapLogger := otelzap.New(zapLogger)
		logger := lib.NewSugaredLogger(otelzapLogger)
		fxLogger = &lib.FxLogger{Logger: logger}
	})

	It("should log OnStartExecuting events", func() {
		onStartExecuting := &fxevent.OnStartExecuting{
			FunctionName: "function",
			CallerName:   "caller",
		}
		fxLogger.LogEvent(onStartExecuting)
	})

	It("should log OnStartExecuted events", func() {
		onStartExecuted := &fxevent.OnStartExecuted{
			FunctionName: "function",
			CallerName:   "caller",
			Method:       "jhon",
			Runtime:      time.Duration(10 * time.Millisecond),
		}
		fxLogger.LogEvent(onStartExecuted)
	})

	It("should log OnStartExecuted events with errors", func() {
		onStartExecutedWithError := &fxevent.OnStartExecuted{
			FunctionName: "function",
			CallerName:   "caller",
			Method:       "jhon",
			Err:          errors.New("test Error"),
			Runtime:      time.Duration(10 * time.Millisecond),
		}
		fxLogger.LogEvent(onStartExecutedWithError)
	})

	It("should log OnStopExecuting events", func() {
		onStopExecuting := &fxevent.OnStopExecuting{
			FunctionName: "function",
			CallerName:   "caller",
		}
		fxLogger.LogEvent(onStopExecuting)
	})

	It("should log OnStopExecuted events", func() {
		onStopExecuted := &fxevent.OnStopExecuted{
			FunctionName: "function",
			CallerName:   "caller",
			Runtime:      time.Duration(10 * time.Millisecond),
		}
		fxLogger.LogEvent(onStopExecuted)
	})

	It("should log OnStopExecuted events with erros", func() {
		onStopExecuted := &fxevent.OnStopExecuted{
			FunctionName: "function",
			CallerName:   "caller",
			Runtime:      time.Duration(10 * time.Millisecond),
			Err:          errors.New("test Error"),
		}
		fxLogger.LogEvent(onStopExecuted)
	})

	It("should log Supplied events with erros", func() {
		onSupplied := &fxevent.Supplied{
			TypeName:   "type",
			ModuleName: "module",
			Err:        errors.New("test Error"),
		}
		fxLogger.LogEvent(onSupplied)
	})

	It("should log Provided events with erros", func() {
		onProvided := &fxevent.Provided{
			ConstructorName: "constructor",
			OutputTypeNames: []string{"type"},
			ModuleName:      "module",
			Err:             errors.New("test Error"),
		}
		fxLogger.LogEvent(onProvided)
	})

	It("should log Decorated events with erros", func() {
		onDecorated := &fxevent.Decorated{
			DecoratorName:   "decorator",
			ModuleName:      "module",
			OutputTypeNames: []string{"type"},
			Err:             errors.New("test Error"),
		}
		fxLogger.LogEvent(onDecorated)
	})

	It("should log Invoking events with erros", func() {
		onInvoking := &fxevent.Invoking{
			FunctionName: "function",
			ModuleName:   "module",
		}
		fxLogger.LogEvent(onInvoking)
	})

	It("should log Started events", func() {
		onStarted := &fxevent.Started{
			Err: nil,
		}
		fxLogger.LogEvent(onStarted)
	})

	It("should log LoggerInitialized events", func() {
		onLoggerInitialized := &fxevent.LoggerInitialized{
			Err: nil,
		}
		fxLogger.LogEvent(onLoggerInitialized)
	})
})

var _ = Describe("GormLogger", func() {
	var gormLogger *lib.GormLogger

	BeforeEach(func() {
		// Create a test logger
		zapLogger, err := zap.NewDevelopment()
		Expect(err).NotTo(HaveOccurred())
		otelzapLogger := otelzap.New(zapLogger)
		logger := lib.NewSugaredLogger(otelzapLogger)
		gormLogger = &lib.GormLogger{Logger: logger}
	})

	It("should log info messages", func() {
		gormLogger.LogMode(gormlogger.Info)
		gormLogger.Info(context.Background(), "This is an info message")
	})

	It("should log Warn messages", func() {
		gormLogger.LogMode(gormlogger.Warn)
		gormLogger.Warn(context.Background(), "This is an Warn message")
	})

	It("should log Error messages", func() {
		gormLogger.LogMode(gormlogger.Error)
		gormLogger.Error(context.Background(), "This is an Error message")
	})

	It("should log trace messages for LogLevel=Info", func() {
		gormLogger.LogMode(gormlogger.Info)
		begin := time.Now()
		gormLogger.Trace(context.Background(), begin, func() (string, int64) {
			return "SELECT * FROM test", 1
		}, nil)
	})

	It("should log trace messages for LogLevel=Warn", func() {
		gormLogger.LogMode(gormlogger.Warn)
		begin := time.Now()
		gormLogger.Trace(context.Background(), begin, func() (string, int64) {
			return "SELECT * FROM test", 1
		}, nil)
	})

	It("should log trace messages for LogLevel=Error", func() {
		gormLogger.LogMode(gormlogger.Error)
		begin := time.Now()
		gormLogger.Trace(context.Background(), begin, func() (string, int64) {
			return "SELECT * FROM test", 1
		}, nil)
	})
})

var _ = Describe("newLogger", func() {
	It("should create a logger with Debug level", func() {
		env := lib.Env{
			LogLevel: "debug",
		}
		logger := lib.NewLogger(env)
		Expect(logger.Level().Enabled(zapcore.DebugLevel)).To(BeTrue())
	})

	It("should create a logger with Info level", func() {
		env := lib.Env{
			LogLevel: "info",
		}
		logger := lib.NewLogger(env)
		Expect(logger.Level().Enabled(zapcore.InfoLevel)).To(BeTrue())
	})

	It("should create a logger with Warn level", func() {
		env := lib.Env{
			LogLevel: "warn",
		}
		logger := lib.NewLogger(env)
		Expect(logger.Level().Enabled(zapcore.WarnLevel)).To(BeTrue())
	})

	It("should create a logger with Error level", func() {
		env := lib.Env{
			LogLevel: "error",
		}
		logger := lib.NewLogger(env)
		Expect(logger.Level().Enabled(zapcore.ErrorLevel)).To(BeTrue())
	})

	It("should create a logger with Fatal level", func() {
		env := lib.Env{
			LogLevel: "fatal",
		}
		logger := lib.NewLogger(env)
		Expect(logger.Level().Enabled(zapcore.FatalLevel)).To(BeTrue())
	})

	It("should create a logger with Panic level", func() {
		env := lib.Env{
			LogLevel: "",
		}
		logger := lib.NewLogger(env)
		Expect(logger.Level().Enabled(zapcore.PanicLevel)).To(BeTrue())
	})
})

var _ = Describe("Printf", func() {
	It("should log a message with arguments", func() {
		env := lib.Env{
			LogLevel: "debug",
		}
		logger := lib.NewLogger(env)
		fxLogger := &lib.FxLogger{Logger: logger}
		fxLogger.Printf("%s %d", "foo", 1)

		// Check that the message was logged at the Debug level
		Expect(logger.Level().Enabled(zapcore.DebugLevel)).To(BeTrue())
	})
})
