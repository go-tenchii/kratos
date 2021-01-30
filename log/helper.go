package log

import (
	"fmt"
	"os"
	"strconv"
)

var nop Logger = new(nopLogger)

// Option .
type Option func(*options)

type options struct {
	level   Level
	verbose int
}

// AllowLevel .
func AllowLevel(l Level) Option {
	return func(o *options) {
		o.level = l
	}
}

// AllowVerbose .
func AllowVerbose(v int) Option {
	return func(o *options) {
		o.verbose = v
	}
}

// Helper is a logger helper.
type Helper struct {
	opts  options
	debug Logger
	info  Logger
	warn  Logger
	err   Logger
}

// NewHelper new a logger helper.
func NewHelper(name string, logger Logger, opts ...Option) *Helper {
	v, _ := strconv.Atoi(os.Getenv("KRATOS_LOG_VERBOSE"))
	options := options{
		level:   GetLevel(os.Getenv("KRATOS_LOG_LEVEL")),
		verbose: v,
	}
	for _, o := range opts {
		o(&options)
	}
	log := With(logger, "module", name)
	return &Helper{
		opts:  options,
		debug: Debug(log),
		info:  Info(log),
		warn:  Warn(log),
		err:   Error(log),
	}
}

// V logs a message at verbose level.
func (h *Helper) V(v int) Verbose {
	if h.opts.verbose >= v {
		return Verbose{h.info}
	}
	return Verbose{}
}

// Debug logs a message at debug level.
func (h *Helper) Debug(a ...interface{}) {
	if h.opts.level.Enabled(LevelDebug) {
		h.debug.Print("log", fmt.Sprint(a...))
	}
}

// Debugf logs a message at debug level.
func (h *Helper) Debugf(format string, a ...interface{}) {
	if h.opts.level.Enabled(LevelDebug) {
		h.debug.Print("log", fmt.Sprintf(format, a...))
	}
}

// Debugw logs a message at debug level.
func (h *Helper) Debugw(kvpair ...interface{}) {
	if h.opts.level.Enabled(LevelDebug) {
		h.debug.Print(kvpair...)
	}
}

// Info logs a message at info level.
func (h *Helper) Info(a ...interface{}) {
	if h.opts.level.Enabled(LevelInfo) {
		h.info.Print("log", fmt.Sprint(a...))
	}
}

// Infof logs a message at info level.
func (h *Helper) Infof(format string, a ...interface{}) {
	if h.opts.level.Enabled(LevelInfo) {
		h.info.Print("log", fmt.Sprintf(format, a...))
	}
}

// Infow logs a message at info level.
func (h *Helper) Infow(kvpair ...interface{}) {
	if h.opts.level.Enabled(LevelInfo) {
		h.info.Print(kvpair...)
	}
}

// Warn logs a message at warn level.
func (h *Helper) Warn(a ...interface{}) {
	if h.opts.level.Enabled(LevelWarn) {
		h.warn.Print("log", fmt.Sprint(a...))
	}
}

// Warnf logs a message at warnf level.
func (h *Helper) Warnf(format string, a ...interface{}) {
	if h.opts.level.Enabled(LevelWarn) {
		h.warn.Print("log", fmt.Sprintf(format, a...))
	}
}

// Warnw logs a message at warnf level.
func (h *Helper) Warnw(kvpair ...interface{}) {
	if h.opts.level.Enabled(LevelWarn) {
		h.warn.Print(kvpair...)
	}
}

// Error logs a message at error level.
func (h *Helper) Error(a ...interface{}) {
	if h.opts.level.Enabled(LevelError) {
		h.err.Print("log", fmt.Sprint(a...))
	}
}

// Errorf logs a message at error level.
func (h *Helper) Errorf(format string, a ...interface{}) {
	if h.opts.level.Enabled(LevelError) {
		h.err.Print("log", fmt.Sprintf(format, a...))
	}
}

// Errorw logs a message at error level.
func (h *Helper) Errorw(kvpair ...interface{}) {
	if h.opts.level.Enabled(LevelError) {
		h.err.Print(kvpair...)
	}
}
