package jot

import "log"

// LoggerShim is an empty object that serves as a way to instantiate the standard logger.
type LoggerShim struct{}

// Print forwards to log.Print.
func (l LoggerShim) Print(v ...interface{}) {
	log.Print(v...)
}

// Printf forwards to log.Printf.
func (l LoggerShim) Printf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

// Println forwards to log.Println.
func (l LoggerShim) Println(v ...interface{}) {
	log.Println(v...)
}
