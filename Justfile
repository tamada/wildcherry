@default: help

App := 'WildCherry'
Version := `grep '^const VERSION = ' cmd/main/version.go | sed "s/^VERSION = \"\(.*\)\"/\1/g"`

# show help message
@help:
    echo "Build tool for {{ App }} {{ Version }} with Just"
    echo "Usage: just <recipe>"
    echo ""
    just --list

build: test
    go build -o wildcherry cmd/main/wildcherry.go

test:
    go test -covermode=count -coverprofile=coverage.out ./...
