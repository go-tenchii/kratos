package log

import "strconv"

// Verbose is logging verbose.
type Verbose int

func (v Verbose) String() string {
	return strconv.Itoa(int(v))
}

// Enabled compare whether the logging verbose is enabled.
func (v Verbose) Enabled(n Verbose) bool {
	return v > n
}
