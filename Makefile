BINARY_NAME=${GOPATH}/bin/roa-webapp
DIR_WITH_MAIN=cmd/roawa
TESTS        ?= ./...

build:
	cd ${DIR_WITH_MAIN} ;go build -o ${BINARY_NAME}

run:
	${BINARY_NAME}

test:
	set -a; go test $(TESTS)