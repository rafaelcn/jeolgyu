GO=`which go`

test:
tests:
	$(GO) test .
	$(GO) test -benchmem -bench .

test-verbose:
	$(GO) test -v .
	$(GO) test -benchmem -bench -v .