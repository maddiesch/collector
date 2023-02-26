# The root directory is the location of this Makefile
ROOT_DIR := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

GOLANG := go
GOLANG_TEST_FLAGS ?= -v
GOLANG_TEST_RUN ?= .
GOLANG_TEST_PATH ?= ./...
GOLANG_TEST_TIMEOUT ?= 30s

.PHONY: test
test:
	${GOLANG} test ${GOLANG_TEST_FLAGS} -run ${GOLANG_TEST_RUN} ${GOLANG_TEST_PATH} -timeout ${GOLANG_TEST_TIMEOUT}

GO_FILES := $(shell find ${ROOT_DIR} -type f -name '*.go' -not -name '*_test.go')
GO_TEST_FILES := $(shell find ${ROOT_DIR} -type f -name '*_test.go')

SQL_FILES := $(shell find ${ROOT_DIR} -type f -name '*.sql')

tmp/coverage.out: ${GO_FILES} ${GO_TEST_FILES} ${ROOT_DIR}/go.mod ${ROOT_DIR}/go.sum
	@mkdir -p tmp
	${GOLANG} test -race -covermode=atomic -coverprofile=$@ ./...

bin/collector-ctl: ${GO_FILES} ${SQL_FILES} ${ROOT_DIR}/go.mod ${ROOT_DIR}/go.sum
	${GOLANG} build -o ${ROOT_DIR}/$@ ./cmd/collector-ctl

.PHONY: build
build: bin/collector-ctl
