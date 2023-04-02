export CGO_ENABLED=1
export GO111MODULE=on

.PHONY: build

GPP_BOT_VERSION := latests
OUTPUT_DIR=./build/_output

build: # @HELP build the Go binaries and run all validations (default)
build:
	go build ${BUILD_FLAGS} -o ${OUTPUT_DIR}/gppbot ./cmd/gpp-bot

debug: BUILD_FLAGS += -gcflags=all="-N -l"
debug: build # @HELP build the Go binaries with debug symbols

test: # @HELP run the unit tests for both large scale and full calibration
test: build
	go test -race github.com/joshuazhu78/xg-bot/...

all: build

clean:: # @HELP remove all the build artifacts
	rm -rf ${OUTPUT_DIR}
	go clean -testcache
