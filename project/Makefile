
export GOPATH=/root/projects/go-exercise/project
export GOBIN=${GOPATH}/bin

all:
	@echo "Usage: make <target>"

hello:
	@cd ${GOPATH}/src/$@; \
	go install; \
	cd -;

main:
	@cd ${GOPATH}/src/$@; \
	go install; \
	cd -;

clean:
	rm -rf bin/*
	rm -rf pkg/*
