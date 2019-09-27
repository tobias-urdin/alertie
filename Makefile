PWD := $(shell pwd)
GOBIN := $(shell which go)
GLIDEBIN := $(shell which glide)
PKG := $(shell awk  '/^package: / { print $$2 }' glide.yaml)

NAME := alertie
BIN := $(PWD)/$(NAME)

.PHONY: all
all: build

work:
	[ -d $(GOPATH) ] || exit 1

depend:
	$(GLIDEBIN) install

.PHONY: depend-update
depend-update:
	$(GLIDEBIN) update

.PHONY: test
test: work depend
	$(GOBIN) test ./...

.PHONY: bin
bin:
	mkdir ./bin 2> /dev/null >&1 || true

.PHONY: build
build: work bin
	$(GOBIN) build -o ./bin/alertie ./cmd/alertie/...
	$(GOBIN) build -o ./bin/alertie-api ./cmd/alertie-api/...
	$(GOBIN) build -o ./bin/alertie-worker ./cmd/alertie-worker/...

.PHONY: fmt
fmt:
	gofmt -l .

.PHONY: fmtfix
fmtfix: work
	$(GOBIN) fmt ./...

.PHONY: clean
clean:
	[ -f $(BIN) ] && rm -f $(BIN)

.PHONY: api
api:
	./bin/alertie-api -config ./bin/alertie.ini

.PHONY: worker
worker:
	./bin/alertie-worker -config ./bin/alertie.ini
