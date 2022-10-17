
FROM golang:1.19-alpine3.16 AS buildstg
#can be git commit sha 
ARG REVISION_NUM=""

WORKDIR /app
# Download necessary Go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . /app
# we can also use -mod=vendor instead if we don't want to download modules everytime we build 
RUN go build -o gitcomp -ldflags "-X 'github.com/msolimans/gitcomp/pkg/build.Revision=$REVISION_NUM#$(date -u +%Y-%m-%d)' -X 'github.com/msolimans/gitcomp/pkg/build.Time=$(date)'" cmd/main.go
# we can use even smaller image like `scratch` but it won't give access to shell 
FROM alpine
COPY --from=buildstg /app/gitcomp /gitcomp
ENTRYPOINT ["/gitcomp"]