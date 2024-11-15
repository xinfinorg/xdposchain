package log

import (
	"os"
)

var (
	root          = &logger{[]interface{}{}, new(swapHandler)}
	StdoutHandler = StreamHandler(os.Stdout, LogfmtFormat())
	StderrHandler = StreamHandler(os.Stderr, LogfmtFormat())
)

func init() {
	root.SetHandler(DiscardHandler())
}

// New returns a new logger with the given context.
// New is a convenient alias for Root().New
func New(ctx ...interface{}) Logger {
	return root.New(ctx...)
}

// Root returns the root logger
func Root() Logger {
	return root
}

// The following functions bypass the exported logger methods (logger.Debug,
// etc.) to keep the call depth the same for all paths to logger.write so
// runtime.Caller(2) always refers to the call site in client code.

// Trace is a convenient alias for Root().Trace
//
// Log a message at the trace level with context key/value pairs
//
// # Usage
//
//	log.Trace("msg")
//	log.Trace("msg", "key1", val1)
//	log.Trace("msg", "key1", val1, "key2", val2)
func Trace(msg string, ctx ...interface{}) {
	root.write(msg, LvlTrace, ctx)
}

// Debug is a convenient alias for Root().Debug
//
// Log a message at the debug level with context key/value pairs
//
// # Usage Examples
//
//	log.Debug("msg")
//	log.Debug("msg", "key1", val1)
//	log.Debug("msg", "key1", val1, "key2", val2)
func Debug(msg string, ctx ...interface{}) {
	root.write(msg, LvlDebug, ctx)
}

// Info is a convenient alias for Root().Info
//
// Log a message at the info level with context key/value pairs
//
// # Usage Examples
//
//	log.Info("msg")
//	log.Info("msg", "key1", val1)
//	log.Info("msg", "key1", val1, "key2", val2)
func Info(msg string, ctx ...interface{}) {
	root.write(msg, LvlInfo, ctx)
}

// Warn is a convenient alias for Root().Warn
//
// Log a message at the warn level with context key/value pairs
//
// # Usage Examples
//
//	log.Warn("msg")
//	log.Warn("msg", "key1", val1)
//	log.Warn("msg", "key1", val1, "key2", val2)
func Warn(msg string, ctx ...interface{}) {
	root.write(msg, LvlWarn, ctx)
}

// Error is a convenient alias for Root().Error
//
// Log a message at the error level with context key/value pairs
//
// # Usage Examples
//
//	log.Error("msg")
//	log.Error("msg", "key1", val1)
//	log.Error("msg", "key1", val1, "key2", val2)
func Error(msg string, ctx ...interface{}) {
	root.write(msg, LvlError, ctx)
}

// Crit is a convenient alias for Root().Crit
//
// Log a message at the crit level with context key/value pairs, and then exit.
//
// # Usage Examples
//
//	log.Crit("msg")
//	log.Crit("msg", "key1", val1)
//	log.Crit("msg", "key1", val1, "key2", val2)
func Crit(msg string, ctx ...interface{}) {
	root.write(msg, LvlCrit, ctx)
	os.Exit(1)
}
