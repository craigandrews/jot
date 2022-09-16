package jot // import "github.com/craigandrews/jot"

Package jot contains a simple logger for developers making notes during
development.

By default, and when enabled, it forwards all requests to the standard
logger via the `log.Print*` functions.

FUNCTIONS

func Disable()
    Disable output from the standard Jotter.

func Enable()
    Enable output from the standard Jotter.

func Enabled() bool
    Enabled returns true if the standard Jotter is enabled.

func Print(v ...interface{})
    Print via the standard Jotter. Arguments are handled in the manner of
    fmt.Print.

func Printf(format string, v ...interface{})
    Printf via the standard Jotter. Arguments are handled in the manner of
    fmt.Printf.

func Println(v ...interface{})
    Println via the standard Jotter. Arguments are handled in the manner of
    fmt.Println.

func SetPrinter(printer Printer)
    SetPrinter changes the Printer instance used by the standard Jotter.


TYPES

type Jotter struct {
	// Has unexported fields.
}
    Jotter is a simple logger that must be manually enabled at runtime.

func New(printer Printer) *Jotter
    New Jotter instance.

func (j *Jotter) Disable()
    Disable printing from this Jotter.

func (j *Jotter) Enable()
    Enable printing from this Jotter.

func (j *Jotter) Enabled() bool
    Enabled returns true if this Jotter is enabled.

func (j *Jotter) Print(v ...interface{})
    Print via the wrapped Printer if enabled. Arguments are handled in the
    manner of fmt.Print.

func (j *Jotter) Printf(format string, v ...interface{})
    Printf via the wrapped Printer if enabled. Arguments are handled in the
    manner of fmt.Printf.

func (j *Jotter) Println(v ...interface{})
    Println via the wrapped Printer if enabled. Arguments are handled in the
    manner of fmt.Println.

type Printer interface {
	Print(v ...interface{})
	Printf(format string, v ...interface{})
	Println(v ...interface{})
}
    Printer is anything that can print via the 3 standard methods.

