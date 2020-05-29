GO=`which go`

test:
	$(GO) test .
	$(GO) test -benchmem -bench .