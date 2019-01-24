PKGS := $(shell go list ./... | grep -v vendor)
BIN := bin
PROTO_PATH := ./proto
PROTO_FILES := ./proto/*.proto
PROTO_GO_OUT := ./src/proto
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
proto-all: proto-go proto-swift

.PHONY: proto-go
proto-go:
	protoc \
		-I${PROTO_PATH} \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/gogo/protobuf/protobuf \
		--gogoslick_out=$(PROTO_GO_OUT) \
		${PROTO_FILES}

.PHONY: proto-swift
proto-swift:
	@./hack/gen-proto-swift.sh $(PROTO_PATH) $(PROTO_SWIFT_OUT)
