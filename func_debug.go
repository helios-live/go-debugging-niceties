package debugging

import (
	"log"
)

// DebugLevel is the level of debug verbosity
var DebugLevel int = 99

// Debug is an alias for Debugln, for backwards compatibility
func Debug(level int, p ...interface{}) {
	Debugln(level, p...)
}

// Debugln just like log.Println but writes the p params if level <= utils.DebugLevel
func Debugln(level int, p ...interface{}) {
	if level <= DebugLevel {
		log.Println(p...)
	}
}

// Debugf writes the formated format with p params if level <= utils.DebugLevel
func Debugf(level int, format string, p ...interface{}) {
	if level <= DebugLevel {
		log.Printf(format, p...)
	}
}
