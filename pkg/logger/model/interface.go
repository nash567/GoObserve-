package model

// A Logger provides methods for logging messages.
type Logger interface {
	// Debug emits a "DEBUG" level log message.
	Debug(msg string)

	// Info emits an "INFO" level log message.
	Info(msg string)

	// Warn emits a "WARN" level log message.
	Warn(msg string)

	// Error emits an "ERROR" level log message.
	Error(msg string)

	// WithField adds a field to the logger and returns a new Logger.
	WithField(key string, value interface{}) Logger

	// WithFields adds multiple fields to the logger and returns a new Logger.
	WithFields(fields Fields) Logger

	// WithError adds a field called "error" to the logger and returns a new Logger.
	WithError(err error) Logger

	// Fatal emits a "FATAL" level log message.
	Fatal(msg string)
}

//go:generate mockery --name=Logger --outpkg mocks

type LevelSetter interface {
	SetLevel(Level)
	GetLevel() Level
}

//go:generate mockery --name=LevelSetter --outpkg mocks
