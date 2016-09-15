// Package jot contains a simple logger for developers making notes during development.
//
// By default, and when enabled, it forwards all requests to the standard logger
// via the `log.Print*` functions.
package jot

var jotter = Jotter{
	enabled: false,
	printer: logShim{},
}

// SetPrinter changes the Printer instance used by the standard Jotter.
func SetPrinter(printer Printer) {
	jotter.printer = printer
}

// Enabled returns true if the standard Jotter is enabled.
func Enabled() bool {
	return jotter.Enabled()
}

// Enable output from the standard Jotter.
func Enable() {
	jotter.Enable()
}

// Disable output from the standard Jotter.
func Disable() {
	jotter.Disable()
}

// Print via the standard Jotter.
// Arguments are handled in the manner of fmt.Print.
func Print(v ...interface{}) {
	jotter.Print(v...)
}

// Printf via the standard Jotter.
// Arguments are handled in the manner of fmt.Printf.
func Printf(format string, v ...interface{}) {
	jotter.Printf(format, v...)
}

// Println via the standard Jotter.
// Arguments are handled in the manner of fmt.Println.
func Println(v ...interface{}) {
	jotter.Println(v...)
}
