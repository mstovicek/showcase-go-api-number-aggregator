DEPEND=github.com/Masterminds/glide

all: ;

depend:
	go get -u -v $(DEPEND)
	$(GOPATH)/bin/glide install

run-app:
	go run app/*.go

run-mock:
	go run mock/*.go

fmt:
	go fmt $(shell glide novendor)

vet:
	go vet $(shell glide novendor)

lint:
	go get -u -v github.com/golang/lint/golint
	for file in $(shell find . -name '*.go' -not -path './vendor/*'); do golint $${file}; done

test:
	go test ./...