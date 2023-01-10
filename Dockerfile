FROM golang:1.19 as builder


RUN mkdir  /app
ADD . /app

WORKDIR /app

RUN GOARCH=amd64 \
    GOOS=linux \
    CGO_ENABLED=0 \
    go mod vendor
RUN cd /app
RUN go build -o ./output/main ./cmd/server/main.go
RUN go build -o ./output/seeder ./cmd/seeder/main.go
RUN go build -o ./output/migrations ./cmd/migrations/main.go


FROM alpine:latest

RUN mkdir /app/
RUN apk add --no-cache libc6-compat 

WORKDIR /app

COPY --from=builder /app/.env.docker /app/
COPY --from=builder /app/output/ /app/

ENV ENVIRONMENT_NAME=docker
EXPOSE 9000
CMD ["./main"]
