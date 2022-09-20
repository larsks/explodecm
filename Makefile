PKG=github.com/larsks/explodecm

GO=go
LINT=golangci-lint run

GOARCH=$(shell go env GOARCH)
GOOS=$(shell go env GOOS)
PROG = build/$(GOARCH)-$(GOOS)-explodecm

COMMIT = $(shell git rev-parse --short=10 HEAD)
DATE = $(shell date -u +"%Y-%m-%dT%H:%M:%S")

GOLDFLAGS = \
	    -X '$(PKG)/version.BuildRef=$(COMMIT)' \
	    -X '$(PKG)/version.BuildDate=$(DATE)'

$(PROG): build main.go
	$(GO) build -o $(PROG) -ldflags "$(GOLDFLAGS)"

lint: .last_linted

.last_linted: main.go
	$(LINT) && touch $@ || rm -f $@

build:
	mkdir build

clean:
	go clean
	rm -f $(PROG) .last_linted
