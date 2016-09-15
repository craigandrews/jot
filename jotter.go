package jot

// Printer is anything that can print via the 3 standard methods.
type Printer interface {
	Print(v ...interface{})
	Printf(format string, v ...interface{})
	Println(v ...interface{})
}

// Jotter is a simple logger that must be manually enabled at runtime.
type Jotter struct {
	enabled bool
	printer Printer
}

// New Jotter instance.
func New(printer Printer) *Jotter {
	return &Jotter{
		enabled: false,
		printer: printer,
	}
}

// Enabled returns true if this Jotter is enabled.
func (j *Jotter) Enabled() bool {
	return j.enabled
}

// Enable printing from this Jotter.
func (j *Jotter) Enable() {
	j.enabled = true
}

// Disable printing from this Jotter.
func (j *Jotter) Disable() {
	j.enabled = false
}

// Print via the wrapped Printer if enabled.
// Arguments are handled in the manner of fmt.Print.
func (j *Jotter) Print(v ...interface{}) {
	if j.Enabled() {
		j.printer.Print(v...)
	}
}

// Printf via the wrapped Printer if enabled.
// Arguments are handled in the manner of fmt.Printf.
func (j *Jotter) Printf(format string, v ...interface{}) {
	if j.Enabled() {
		j.printer.Printf(format, v...)
	}
}

// Println via the wrapped Printer if enabled.
// Arguments are handled in the manner of fmt.Println.
func (j *Jotter) Println(v ...interface{}) {
	if j.Enabled() {
		j.printer.Println(v...)
	}
}
