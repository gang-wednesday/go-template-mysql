FROM golang:1.19 as builder


RUN mkdir  /app
ADD . /app

WORKDIR /app
ENV ENVIRONMENT_NAME=docker
RUN GOARCH=amd64 \
    GOOS=linux \
    CGO_ENABLED=0 \
    go mod vendor
RUN cd /app
RUN go build -o ./output/main ./cmd/server/main.go
RUN go build -o ./output/seeder ./cmd/seeder/main.go
RUN go build -o ./output/migrations ./cmd/migrations/main.go


CMD ["bash", "./scripts/migrate-and-run.sh"]
EXPOSE 9000

