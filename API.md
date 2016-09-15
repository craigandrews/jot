# Jot

```go
import "github.com/doozr/jot"
```

Package jot contains a simple logger for developers making notes during
development.

By default, and when enabled, it forwards all requests to the standard logger
via the `log.Print*` functions.

## Usage

#### func  Disable

```go
func Disable()
```
Disable output from the standard Jotter.

#### func  Enable

```go
func Enable()
```
Enable output from the standard Jotter.

#### func  Enabled

```go
func Enabled() bool
```
Enabled returns true if the standard Jotter is enabled.

#### func  Print

```go
func Print(v ...interface{})
```
Print via the standard Jotter. Arguments are handled in the manner of fmt.Print.

#### func  Printf

```go
func Printf(format string, v ...interface{})
```
Printf via the standard Jotter. Arguments are handled in the manner of
fmt.Printf.

#### func  Println

```go
func Println(v ...interface{})
```
Println via the standard Jotter. Arguments are handled in the manner of
fmt.Println.

#### func  SetPrinter

```go
func SetPrinter(printer Printer)
```
SetPrinter changes the Printer instance used by the standard Jotter.

#### type Jotter

```go
type Jotter struct {
}
```

Jotter is a simple logger that must be manually enabled at runtime.

#### func  New

```go
func New(printer Printer) *Jotter
```
New Jotter instance.

#### func (*Jotter) Disable

```go
func (j *Jotter) Disable()
```
Disable printing from this Jotter.

#### func (*Jotter) Enable

```go
func (j *Jotter) Enable()
```
Enable printing from this Jotter.

#### func (*Jotter) Enabled

```go
func (j *Jotter) Enabled() bool
```
Enabled returns true if this Jotter is enabled.

#### func (*Jotter) Print

```go
func (j *Jotter) Print(v ...interface{})
```
Print via the wrapped Printer if enabled. Arguments are handled in the manner of
fmt.Print.

#### func (*Jotter) Printf

```go
func (j *Jotter) Printf(format string, v ...interface{})
```
Printf via the wrapped Printer if enabled. Arguments are handled in the manner
of fmt.Printf.

#### func (*Jotter) Println

```go
func (j *Jotter) Println(v ...interface{})
```
Println via the wrapped Printer if enabled. Arguments are handled in the manner
of fmt.Println.

#### type LoggerShim

```go
type LoggerShim struct{}
```

LoggerShim is an empty object that serves as a way to instantiate the standard
logger.

#### func (LoggerShim) Print

```go
func (l LoggerShim) Print(v ...interface{})
```
Print forwards to log.Print.

#### func (LoggerShim) Printf

```go
func (l LoggerShim) Printf(format string, v ...interface{})
```
Printf forwards to log.Printf.

#### func (LoggerShim) Println

```go
func (l LoggerShim) Println(v ...interface{})
```
Println forwards to log.Println.

#### type Printer

```go
type Printer interface {
	Print(v ...interface{})
	Printf(format string, v ...interface{})
	Println(v ...interface{})
}
```

Printer is anything that can print via the 3 standard methods.
