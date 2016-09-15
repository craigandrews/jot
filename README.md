# Jot

Install:

```
go get github.com/doozr/jot
```

Use:

```go
import "github.com/doozr/jot"
```

Jot is a simple logger for developers making notes during development. It is
similar in concept and use to the `debug` log level that many loggers provide. A
way of making notes and annotations to the code that appear at runtime.

See the [API documentatin](API.md) for more detail.

### Getting A Jotter Instance

By default a standard `Jotter` instance is created that forwards everything to
the standard logger. This means that changes to the standard logger
configuration are immediately reflected in Jot output.

Alternatively a customer `Jotter` instance can be created using `jot.New`. `Jotter` wraps an instance of `Printer`, which can be any object that implements
`Print`, `Printf` and `Println` in the same way as the `fmt` package.
Coincidentally the `log.Logger` type conforms to this interface.

```go
logger := log.New(os.Stderr, "", 0)
jotter := jot.New(logger)
```

### Using Jotter

Use Jotter just as you would a logger. The only difference is that the `jot`
lines may not appear unless the Jotter instance is enabled.

```go
jot.Print("A thing has happened: ", thing)
jot.Printf("Got %d requests from user %s", len(requests), user.ID)
jot.Println("Just like Print but with a newline if the Printer supports it")
```

### Enabling Jotter

The `Jotter` instance can be enabled by called `Enable()` on it, or the on the
package to enable the standard `Jotter`. It can be turned off again by calling
`Disable()`. This is useful for being able to turn on `Jotter` at runtime via
some secret API call.

```go
jot.Enable()
jot.Print("This is printed")
jot.Disable()
jot.Print("This is not")
```

A useful way to enable Jotter could be to use an environment variable. This is
not enabled by default to prevent a generic way of enabling detailed output for
any program that uses Jotter, but it is easy to add.

```go
if os.Getenv("JOTTER_ENABLE") == "true" {
	jot.Enable()
}
```

Then, to use it:

```
export JOTTER_ENABLE=true
my_program
```

### Example

```go
jot.Print("Calling connectToThing", someParam)
jot.Printf("User: %s ACL: %s", user, acl)
result, err := connectToThing(someParam)

if err != nil {
	log.Printf("Error connecting to thing with %s: %v", someParam, err)
}

if (result == "specific value of interest to developer") {
	jot.Printf("TRACER: weird result occurred")
}
```

In this example the log line will be printed in case of error as is proper, but
the jot lines are only printed if the standard `Jotter` instance is enabled.
This ensures that noisy log lines that only help developers can be avoided
unless absolutely required.
