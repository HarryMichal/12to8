TEST ?= unit
VERSION ?= locked

test: build
ifeq (acceptance,${TEST})
ifeq (latest,${VERSION})
	cd 925r && git fetch && git reset --hard origin/master
endif
	cp 925r.yml 925r
	cd 925r && docker build -t 925r:upstream -f scripts/docker/Dockerfile .
endif
	./tests/${TEST}.sh

build:
ifeq (latest,${VERSION})
	go get -u
	go mod tidy
	git diff
endif
	go build
