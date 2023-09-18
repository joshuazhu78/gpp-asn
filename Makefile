export CGO_ENABLED=1
export GO111MODULE=on

.PHONY: build

GPP_ASN1_VERSION     := latests
ONOS_PROTOC_VERSION  := v1.0.2
OUTPUT_DIR            =./build/_output
SPEC_TXT             ?=38331-h50

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
	asn1c -B asn/nr-rrc-definitions.asn1 > protos/nr-rrc-definitions.proto
	@sed -i 's/asn\/nr_rrc_definitions.v1/nr_rrc_definitions.v1/g' protos/nr-rrc-definitions.proto
	@sed -i 's/nr_rrc_definitions.v1\/nr-rrc-definitions;nrrrcdefinitions/.\/nr-rrc-definitions;nrrrcdefinitions/g' protos/nr-rrc-definitions.proto
	asn1c -B asn/nr-ue-variables.asn1 > protos/nr-ue-variables.proto
	@sed -i 's/asn\/nr_ue_variables.v1/nr_ue_variables.v1/g' protos/nr-ue-variables.proto
	@sed -i 's/nr_ue_variables.v1\/nr-ue-variables;nruevariables/.\/nr-ue-variables;nruevariables/g' protos/nr-ue-variables.proto

protos-go: # @HELP compile the protobuf files (using protoc-go Docker)
protos-go:
	@sudo rm -rf pkg/nr/*
	docker run -it \
		-v `pwd`:/go/src/github.com/joshuazhu78/gpp-asn \
		-v `pwd`/../onosproject/onos-lib-go:/go/src/github.com/onosproject/onos-lib-go \
		-w /go/src/github.com/joshuazhu78/gpp-asn \
		--entrypoint /go/src/github.com/joshuazhu78/gpp-asn/build/bin/compile-protos.sh \
		onosproject/protoc-go:${ONOS_PROTOC_VERSION}
	@sudo chown -R $(USER):$(USER) pkg/nr
	@sed -i 's/aper:"choiceIdx:1,sizeLB:16,sizeUB:16,"`/json:"two-one,omitempty" aper:"choiceIdx:1,sizeLB:16,sizeUB:16,"`/g' pkg/nr/nr-rrc-definitions/nr-rrc-definitions.pb.go
	@sed -i 's/aper:"choiceIdx:2,sizeLB:43,sizeUB:43,"`/json:"two-two,omitempty" aper:"choiceIdx:2,sizeLB:43,sizeUB:43,"`/g' pkg/nr/nr-rrc-definitions/nr-rrc-definitions.pb.go
	@sed -i 's/aper:"choiceIdx:3,sizeLB:32,sizeUB:32,"`/json:"four-one,omitempty" aper:"choiceIdx:3,sizeLB:32,sizeUB:32,"`/g' pkg/nr/nr-rrc-definitions/nr-rrc-definitions.pb.go
	@sed -i 's/aper:"choiceIdx:4,sizeLB:59,sizeUB:59,"`/json:"three-two,omitempty" aper:"choiceIdx:4,sizeLB:59,sizeUB:59,"`/g' pkg/nr/nr-rrc-definitions/nr-rrc-definitions.pb.go
	@sed -i 's/aper:"choiceIdx:5,sizeLB:48,sizeUB:48,"`/json:"six-one,omitempty" aper:"choiceIdx:5,sizeLB:48,sizeUB:48,"`/g' pkg/nr/nr-rrc-definitions/nr-rrc-definitions.pb.go
	@sed -i 's/aper:"choiceIdx:6,sizeLB:75,sizeUB:75,"`/json:"four-two,omitempty" aper:"choiceIdx:6,sizeLB:75,sizeUB:75,"`/g' pkg/nr/nr-rrc-definitions/nr-rrc-definitions.pb.go
	@sed -i 's/aper:"choiceIdx:7,sizeLB:64,sizeUB:64,"`/json:"eight-one,omitempty" aper:"choiceIdx:7,sizeLB:64,sizeUB:64,"`/g' pkg/nr/nr-rrc-definitions/nr-rrc-definitions.pb.go
	@sed -i 's/aper:"choiceIdx:8,sizeLB:107,sizeUB:107,"`/json:"four-three,omitempty" aper:"choiceIdx:8,sizeLB:107,sizeUB:107,"`/g' pkg/nr/nr-rrc-definitions/nr-rrc-definitions.pb.go
	@sed -i 's/aper:"choiceIdx:9,sizeLB:107,sizeUB:107,"`/json:"six-two,omitempty" aper:"choiceIdx:9,sizeLB:107,sizeUB:107,"`/g' pkg/nr/nr-rrc-definitions/nr-rrc-definitions.pb.go
	@sed -i 's/aper:"choiceIdx:10,sizeLB:96,sizeUB:96,"`/json:"twelve-one,omitempty" aper:"choiceIdx:10,sizeLB:96,sizeUB:96,"`/g' pkg/nr/nr-rrc-definitions/nr-rrc-definitions.pb.go
	@sed -i 's/aper:"choiceIdx:11,sizeLB:139,sizeUB:139,"`/json:"four-four,omitempty" aper:"choiceIdx:11,sizeLB:139,sizeUB:139,"`/g' pkg/nr/nr-rrc-definitions/nr-rrc-definitions.pb.go
	@sed -i 's/aper:"choiceIdx:12,sizeLB:139,sizeUB:139,"`/json:"eight-two,omitempty" aper:"choiceIdx:12,sizeLB:139,sizeUB:139,"`/g' pkg/nr/nr-rrc-definitions/nr-rrc-definitions.pb.go
	@sed -i 's/aper:"choiceIdx:13,sizeLB:128,sizeUB:128,"`/json:"sixteen-one,omitempty" aper:"choiceIdx:13,sizeLB:128,sizeUB:128,"`/g' pkg/nr/nr-rrc-definitions/nr-rrc-definitions.pb.go

protos: protos-gen protos-go

all: build

clean:: # @HELP remove all the build artifacts
	rm -rf ${OUTPUT_DIR}
	go clean -testcache
