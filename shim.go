package jot

import "log"

// logShim is an empty object that serves as a way to instantiate the standard logger
type logShim struct{}

// Print forwards to log.Print
func (_ logShim) Print(v ...interface{}) {
	log.Print(v...)
}

// Printf forwards to log.Printf
func (_ logShim) Printf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

// Println forwards to log.Println
func (_ logShim) Println(v ...interface{}) {
	log.Println(v...)
}
