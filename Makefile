PKGS := $(shell go list ./... | grep -v vendor)
BIN := bin
PROTO_PATH := ./proto
PROTO_FILES := ./proto/*.proto
PROTO_GO_OUT := ./src
PROTO_SWIFT_OUT := ./swift
PROTO_DOC := ./doc

.PHONY: run
run:
	dev_appserver.py appengine/app_dev.yaml

.PHONY: test
test:
	go test -v -parallel=4 $(PKGS)

.PHONY: dep
dep:
	dep ensure -v

.PHONY: vet
vet:
	go vet $(PKGS)

.PHONY: lint
lint:
	golint $(PKGS)

.PHONY: clean
clean:
	rm -rf $(BIN)

.PHONY: coverage
coverage:
	go test -v -race -covermode=atomic -coverpkg=./... -coverprofile=coverage.txt ./...

.PHONY: reviewdog
reviewdog:
	reviewdog -reporter=github-pr-review

.PHONY: deploy-dev
deploy-dev:
	gcloud app deploy appengine/app_dev.yaml --project $(PROJECT_ID)

.PHONY: deploy-prod
deploy-prod:
	gcloud app deploy appengine/app_prod.yaml --project $(PROJECT_ID)

.PHONY: deploy-cron
deploy-cron:
	gcloud app deploy appengine/cron.yaml --project $(PROJECT_ID)

.PHONY: proto-all
proto-all: proto-go proto-swift proto-doc

.PHONY: proto-go
proto-go:
	protoc \
		-I=. \
		-I=${GOPATH}/src \
		-I=${GOPATH}/src/github.com/gogo/protobuf/protobuf \
		--gogoslick_out=$(PROTO_GO_OUT) \
		./proto/api.proto

.PHONY: proto-swift
proto-swift:
	cp -f ./proto/api.proto ./proto/swift.proto
	sed -i -e '/import .*gogo\/protobuf.*/d' proto/swift.proto
	sed -i -e 's/ \[.*\]//g' proto/swift.proto
	@mkdir -p $(PROTO_SWIFT_OUT)
	protoc \
		--swift_out=$(PROTO_SWIFT_OUT) \
		--proto_path=$(PROTO_PATH) \
		--swift_opt=Visibility=Public \
		./proto/swift.proto
	mv ./swift/swift.pb.swift ./swift/api.pb.swift

.PHONY: proto-doc
proto-doc:
	@mkdir -p $(PROTO_DOC)
	protoc \
		-I/usr/local/include \
		-I. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--doc_out=./doc \
		--doc_opt=markdown,api.md \
		./proto/api.proto
