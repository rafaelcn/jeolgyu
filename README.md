# 절규 ![Go testing](https://github.com/rafaelcn/jeolgyu/workflows/Go%20test/badge.svg?branch=master)

This is a simple logger to be used in a Go project to output any message sent to
it to one of its sinks. Currently there are three implemented sinks, the
`SinkFile`, the `SinkOutput` and the `SinkBoth`. The latter outputs messages to
both a file and the standard output

The aim of the project is to be a simple usage logger and not currently worrying
about performance so don't expect that.