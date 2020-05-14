# 절규 ![Go testing](https://github.com/rafaelcn/jeolgyu/workflows/Go%20test/badge.svg?branch=master)

This is a simple logger to be used in a Go project to output any message sent to
it to one of its sinks. Currently there are three implemented sinks, the
`SinkFile`, the `SinkOutput` and the `SinkBoth`. The latter outputs messages to
both a file and the standard output

The aim of the project is to be a simple usage logger and not currently worrying
about performance so don't expect that.

To use it on your project a `go get -u github.com/rafaelcn/jeolgyu` will
download it. If you wish to update it just write the same code again on the
terminal.

# Usage

The logger can write to a file, the output or both. Jeolgyu has the concept of
sinks and you can use one of them as you initialize the logger. Code samples can
be seen below for a better perspective of the project.

When the logger is using the SinkFile option the formatted message will be
encoded using *JSON*. No other message encoding standards are implemented yet
on the project.

```
import (
    ...

    "github.com/rafaelcn/jeolgyu"
)

func main() {
    // SinkBoth will use a logger file and also the stdout/stderr when printing
    // messages. The second parameter to New indicates the relative path of the
    // written logger on the filesystem.
    logger := jeolgyu.New(SinkBoth, "")

    // The created file will have be named as a timestamp, according to when the
    // logger has been created.

    logger.Info("Well, this is a nice message.")

    // The Info call will sink the message to the stdout/stderr and also to the
    // filesystem with the respective contents:
    //
    // stdout:
    // "<timestamp of the message> [+] Well, this is a nice message"
    //
    // file:
    // {"level":"info","what":"Well, this is a nice message","when":"<timestamp>"}
}
```

There's a more complex sample below showing some other features of the logger.

```
import (
    ...

    "github.com/rafaelcn/jeolgyu"
)

func main() {
    logger := jeolgyu.New(SinkBoth, "")

    _, err := ioutil.ReadFile(..., ...)

    if err != nil {
        logger.Error("Something wen't wrong, sorry. Reason %v", err)
    }
}
```

If you have any questions regarding the use of the software you can use the
Github [issue system](https://github.com/rafaelcn/jeolgyu/issues/new).

# Performance

The project was tested as the bencharks written on the
[jeolgyu_bench_test.go](https://github.com/rafaelcn/jeolgyu/blob/master/jeolgyu_bench_test.go).

| Package |  Sink  |    Time       | Objects allocated |
|---------|--------|---------------|-------------------|
| jeolgyu | File   | 7641 ns/op    | 18 allocs/op      |
| jeolgyu | Output | 599605 ns/op  | 8 allocs/op       |
| jeolgyu | Both   |  - ns/op      | - allocs/op       |
