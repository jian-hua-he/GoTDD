VOLUME_PATH = "/go/src/go-tdd"
IMAGE = golang:1.14
BASE_COMMAND = docker run --rm -v $(PWD):$(VOLUME_PATH) $(IMAGE)

test-all:
	$(BASE_COMMAND) bash -c "cd $(VOLUME_PATH) && go test --race -cover ./..."
