# List available recipes
default:
    @just --list

# Alter tailor swatches
alter:
    @tailor alter

# Run linters
lint:
    actionlint
    golangci-lint run

fmt:
    go fmt .
    taplo fmt
    templ fmt .

# Check what tailor would change and measure
measure:
    @tailor baste
    @tailor measure

# Run code generation
codegen:
    go generate ./...

# Build a binary for current OS
build:
    goreleaser build --clean --snapshot --single-target -o .

# Migrate a local database
migrate:
    dbmate -u sqlite:db.sqlite3 up 

# Run locally
[arg('ARGS', help='arguments passed to pocket-expo')]
run *ARGS: build migrate
    ./pocket-expo {{ARGS}}
