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

#### type Printer

```go
type Printer interface {
	Print(v ...interface{})
	Printf(format string, v ...interface{})
	Println(v ...interface{})
}
```

Printer is anything that can print via the 3 standard methods.
