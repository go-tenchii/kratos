package log

import "fmt"

// Verbose is logging verbose.
type Verbose struct {
	log Logger
}

// Info is equivalent to the global Info function, guarded by the value of v.
// See the documentation of V for usage.
func (v Verbose) Info(a ...interface{}) {
	if v.log != nil {
		v.log.Print("log", fmt.Sprint(a...))
	}
}

// Infof is equivalent to the global Infof function, guarded by the value of v.
// See the documentation of V for usage.
func (v Verbose) Infof(format string, a ...interface{}) {
	if v.log != nil {
		v.log.Print("log", fmt.Sprintf(format, a...))
	}
}

// Infow is equivalent to the global Infoln function, guarded by the value of v.
// See the documentation of V for usage.
func (v Verbose) Infow(kvpair ...interface{}) {
	if v.log != nil {
		v.log.Print(kvpair...)
	}
}
