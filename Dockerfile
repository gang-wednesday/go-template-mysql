FROM golang


RUN mkdir -p /go/src/github.com/wednesday-solutions/go-template


ADD . /go/src/github.com/wednesday-solutions/go-template

RUN /go/src/github.com/wednesday-solutions/go-template/scripts/install-tooling.sh

WORKDIR /go/src/github.com/wednesday-solutions/go-template

RUN GOARCH=amd64 \
    GOOS=linux \
    CGO_ENABLED=0 \
    go mod vendor
RUN go build -o ./ ./cmd/server/main.go
CMD ["bash", "./scripts/migrate-and-run.sh"]
EXPOSE 9000