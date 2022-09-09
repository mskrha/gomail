BUILD		= gomail
VERSION		?= 0.0
REGISTRY	= 'registry.docker.srv.sobotiste.info'

all: clean format build

clean:
	rm -f $(BUILD)

format:
	go fmt

build:
	CGO_ENABLED=0 go build -ldflags "-X main.version=$(VERSION)" -o $(BUILD) *.go

docker: clean format build
	docker build --compress --tag $(REGISTRY)/$(BUILD):$(VERSION) --tag $(REGISTRY)/$(BUILD):latest .
	docker push $(REGISTRY)/$(BUILD):$(VERSION)
	docker push $(REGISTRY)/$(BUILD):latest
	rm -f $(BUILD)
