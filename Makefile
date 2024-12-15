ARGS="help"

.PHONY: all
all: cmd/stpw

cmd/stpw:
	go build -o bin/stpw main.go

.PHONY: test
test:
	go test ./...

.PHONY: run
run:
	@go run main.go --file tmp/stpw.txt $(ARGS)

.PHONY: clean
clean:
	rm -f bin/stpw
