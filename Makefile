.PHONY: build test clean prepare update docker

GO=CGO_ENABLED=0 go

MICROSERVICES=cmd/device-serial

.PHONY: $(MICROSERVICES)

DOCKERS=docker_device_serial_go
.PHONY: $(DOCKERS)

VERSION=$(shell cat ./VERSION)
GIT_SHA=$(shell git rev-parse HEAD)
GOFLAGS=-ldflags "-X github.com/edgexfoundry/device-serial.Version=$(VERSION)"

build: $(MICROSERVICES)
	go build ./...

cmd/device-serial:
	$(GO) build $(GOFLAGS) -o $@ ./cmd

test:
	go test ./... -cover

clean:
	rm -f $(MICROSERVICES)

prepare:
	glide install

update:
	glide update

docker: $(DOCKERS)

docker_device_serial_go:
	docker build \
		--label "git_sha=$(GIT_SHA)" \
		-t edgexfoundry/docker-device-serial-go:$(GIT_SHA) \
		-t edgexfoundry/docker-device-serial-go:$(VERSION)-dev \
		.
