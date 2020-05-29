GO=`which go`

test:
tests:
	$(GO) test .
	$(GO) test -benchmem -bench .