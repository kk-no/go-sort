.PHONY: test
test:
	go test ./... -v

.PHONY: bench
bench:
	go test -bench . ./... -benchmem