package jot

import "log"

// logShim is an empty object that serves as a way to instantiate the standard logger
type logShim struct{}

// Print forwards to log.Print
func (l logShim) Print(v ...interface{}) {
	log.Print(v...)
}

// Printf forwards to log.Printf
func (l logShim) Printf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

// Println forwards to log.Println
func (l logShim) Println(v ...interface{}) {
	log.Println(v...)
}
