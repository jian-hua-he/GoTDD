VOLUME_PATH = /go/src/go-tdd
IMAGE = golang:1.14
BASE_COMMAND = docker run --rm -v $(PWD):$(VOLUME_PATH) $(IMAGE)

TARGET ?= concurrency

test-all:
	$(BASE_COMMAND) bash -c "cd $(VOLUME_PATH) && go test --race -cover ./..."

bench:
	$(BASE_COMMAND) bash -c "cd $(VOLUME_PATH) && go test ./$(TARGET) -bench=./concurrency"
