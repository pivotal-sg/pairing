GIT_VERSION := $(shell git describe --abbrev=4 --dirty --always --tags)

all: pairing.linux pairing.darwin
.PHONY: all

pairing.linux: main.go
	GOOS=linux go build -ldflags "-X main.version=$(GIT_VERSION)" -o $@ $< 

pairing.darwin: main.go
	GOOS=darwin go build -ldflags "-X main.version=$(GIT_VERSION)" -o $@ $<

.PHONY: clean
clean:
	rm pairing.linux pairing.darwin
