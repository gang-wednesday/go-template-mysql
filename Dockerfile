FROM golang:1.19 as builder


RUN mkdir  /app
ADD . /app

WORKDIR /app
ENV ENVIRONMENT_NAME=docker
RUN GOARCH=amd64 \
    GOOS=linux \
    CGO_ENABLED=0 \
    go mod vendor

RUN go run ./cmd/seeder/main.go
RUN go build -o ./output/server ./cmd/server/main.go
RUN go build -o ./output/migrations ./cmd/migrations/main.go
RUN go build  -o ./output/seeder ./cmd/seeder/output/seed.go



FROM alpine:latest
ENV ENVIRONMENT_NAME=docker
RUN apk add --no-cache libc6-compat 
RUN apk add --no-cache --upgrade bash
RUN mkdir -p /app/
WORKDIR /app

COPY /scripts /app/scripts/
COPY --from=builder /app/output/ /app/output
COPY --from=builder /app/cmd/seeder/output/build/ /app/output/cmd/seeder/output/build/
COPY ./.env.docker /app/output/
COPY ./.env.docker /app/output/cmd/seeder/output/build/
COPY ./.env.docker /app/output/cmd/seeder/output/
COPY ./.env.docker /app/output/cmd/seeder/
COPY ./.env.docker /app/output/cmd/
COPY ./.env.docker /app/
COPY ./scripts/ /app/
COPY --from=builder /app/internal/migrations/ /app/internal/migrations/
CMD ["bash","./migrate-and-run.sh"]
EXPOSE 9000

