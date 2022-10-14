GO111MODULE := on
DOCKER_TAG := $(or ${GIT_TAG_NAME}, latest)

all: s3-event-logger

.PHONY: s3-event-logger
s3-event-logger:
	go build -o bin/s3-event-logger
	strip bin/s3-event-logger

.PHONY: dockerimages
dockerimages:
	docker build -t mwennrich/s3-event-logger:${DOCKER_TAG} .

.PHONY: dockerpush
dockerpush:
	docker push mwennrich/s3-event-logger:${DOCKER_TAG}

.PHONY: clean
clean:
	rm -f bin/*
