REVISION = "$(REVISION_NUM)\#$(shell date -u +%Y-%m-%d-%H:%M:%S)"

dbuild: 
	docker build -t gcomp . --build-arg  REVISION_NUM=$(REVISION)

build:
	go build -o gcomp -ldflags "-X 'github.com/msolimans/gitcomp/pkg/build.Revision=$(REVISION)' -X 'github.com/msolimans/gitcomp/pkg/build.Time=$(shell date)'" cmd/main.go

test: 
	go test -v -count=1 ./...
