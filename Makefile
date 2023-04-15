export CGO_ENABLED=1
export GO111MODULE=on

.PHONY: build

GPP_ASN1_VERSION     := latests
ONOS_PROTOC_VERSION  := v1.0.2
OUTPUT_DIR            =./build/_output
SPEC_TXT             ?=38331-gb0

build: # @HELP build the Go binaries and run all validations (default)
build:
	go build ${BUILD_FLAGS} -o ${OUTPUT_DIR}/extractor ./cmd/extractor

debug: BUILD_FLAGS += -gcflags=all="-N -l"
debug: build # @HELP build the Go binaries with debug symbols

test: # @HELP run the unit tests for both large scale and full calibration
test: build
	go test -race github.com/joshuazhu78/gpp-asn/...

extract: # @HELP extract the asn1 syntax and saved to asn folder
extract: build
	${OUTPUT_DIR}/extractor --specTxt=spec/$(SPEC_TXT).txt

protos-gen: # @HELP generates the proto file from asn1 syntax
protos-gen: extract
	asn1c -B asn/CSI-ReportConfig.asn1 > protos/CSI-ReportConfig.proto

protos-go: # @HELP compile the protobuf files (using protoc-go Docker)
protos-go:
	docker run -it \
		-v `pwd`:/go/src/github.com/joshuazhu78/gpp-asn \
		-v `pwd`/../onosproject/onos-lib-go:/go/src/github.com/onosproject/onos-lib-go \
		-w /go/src/github.com/joshuazhu78/gpp-asn \
		--entrypoint /go/src/github.com/joshuazhu78/gpp-asn/build/bin/compile-protos.sh \
		onosproject/protoc-go:${ONOS_PROTOC_VERSION}

all: build

clean:: # @HELP remove all the build artifacts
	rm -rf ${OUTPUT_DIR}
	go clean -testcache
