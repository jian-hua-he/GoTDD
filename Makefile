VOLUME_PATH = /go/src/go-tdd
IMAGE = golang:1.14
BASE_COMMAND = docker run --rm -v $(PWD):$(VOLUME_PATH) $(IMAGE)

check-target:
ifndef TARGET
	$(error TARGET is undefined)
endif

test-all:
	$(BASE_COMMAND) bash -c "cd $(VOLUME_PATH) && go test --race -cover ./..."

test: check-target
	$(BASE_COMMAND) bash -c "cd $(VOLUME_PATH) && go test --race -cover ./$(TARGET)"

bench: check-target
	$(BASE_COMMAND) bash -c "cd $(VOLUME_PATH) && go test ./$(TARGET) -bench=./concurrency"
