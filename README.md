# Jot

Install:<img align="right" width="180" style="margin: 12px" src="https://cdn.rawgit.com/doozr/jot/master/jot.svg">


```
go get github.com/doozr/jot
```

Use:

```go
import "github.com/doozr/jot"
```

Jot is a simple logger for developers making notes during development. Like
writing notes in the margin or keeping a log of execution, it allows a more
detailed record of what is actually going on that the normal logger may provide.

It is similar in concept and use to the `debug` log level that many loggers
provide. A way of making annotations in the code that appear in the logs at
runtime.

See the [API documentation](API.md) for more detail.

### Getting A Jot Instance

By default a standard jotter is created that forwards everything to the standard
logger. This means that changes to the standard logger configuration are
immediately reflected in Jot output.

Alternatively a custom `Jotter` instance can be created using `jot.New`.
`Jotter` wraps an instance of `Printer`, which can be any object that implements
`Print`, `Printf` and `Println` in the same way as the `fmt` package.
Coincidentally the `log.Logger` type conforms to this interface.

```go
logger := log.New(os.Stderr, "", 0)
jotter := jot.New(logger)
```

### Using Jot

Use Jot just as you would a logger. The only difference is that the `jot`
lines will not appear unless the `Jotter` instance is enabled.

```go
jot.Print("A thing has happened: ", thing)
jot.Printf("Got %d requests from user %s", len(requests), user.ID)
jot.Println("Just like Print but with a newline if the Printer supports it")
```

### Enabling Jot

A `Jotter` instance can be enabled by called `Enable()` on it. Call `Enable()`
on the package to enable the standard jotter. It can be turned off again by
calling `Disable()`. This is useful for being able to turn Jot on and off at
runtime via an API call or other signal.

```go
jot.Enable()
jot.Print("This is printed")
jot.Disable()
jot.Print("This is not")
```

A useful way to enable Jot could be to use an environment variable. This is
not enabled by default to prevent a generic way of enabling detailed output for
any program that uses Jot, but it is easy to add.

```go
if os.Getenv("JOTTER_ENABLE") == "true" {
	jot.Enable()
}
```

Then, to use it:

```
export JOT_ENABLE=true
my_program
```

### Conditional Jotting

Sometimes it may be desirable to defer jotting a particular message unless Jot
is definitely enabled. This could be because it takes a significant time to
calculate the thing being jotted that is unnecessary in normal execution. This
is handled by checking the Jot status with the `Enabled` method.

```go
if jot.Enabled() {
	jot.Print("Result of claculation: ", longRunningCalculation())
}
```

### Example

```go
func listen(client Client, ch chan Message,
			done chan struct{}, wg *sync.WaitGroup) {
	jot.Print("Started forwardMessage")
	defer func() {
		jot.Print("Finished forwardMessage")
		wg.Done()
	}()

	for {
		select {
		case <-done:
			jot.Print("forwardMessage detected done channel close")
			return
		default:
			message, err := client.Read()
			if err != nil {
				log.Print("Error reading from client: ", err)
				return
			}
			log.Printf(
				"Received message %d from %s to %s",
				message.ID, message.From, message.To)
			jot.Printf(
				"Message id: %d body: %s route: %s",
				message.ID, message.Text, message.Route)

			ch <- message
		}
	}
}
```

In this example the `log` lines will be printed on message arrival and in case
of error as is proper. The `jot` lines only appear if the standard jotter is
enabled, and can be used to help debug message or synchronisation issues. If the
standard jotter is disabled they get out of the way to ensure the normal
operational logs are clear and concise.
