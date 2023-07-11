GO_FILES=app.js $(shell find ./pkg -name \*.go -print)

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## tidy: format code and tidy modfile
.PHONY: tidy
tidy:
	go fmt ./...
	go mod tidy -v

## audit: run quality control checks
.PHONY: audit
audit:
	go mod verify
	go vet ./...
	#go run honnef.co/go/tools/cmd/staticcheck@2022.1.3 -checks=all,-ST1000,-U1000 ./...
	#go run golang.org/x/vuln/cmd/govulncheck@latest ./...
	go test -race -buildvcs -vet=off ./...

## static-analysis: lint the code
.PHONY: static-analysis
static-analysis:
	go mod verify
	go vet ./...
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@v1.47.2 run -E gosec	

.PHONY: unit-test
unit-test:
	go test -v ./pkg/...