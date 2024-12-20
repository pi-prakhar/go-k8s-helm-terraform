FROM golang:1.23.1 as builder
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY cmd ./cmd
COPY internal ./internal
COPY metrics ./metrics

RUN go mod download

ENV GOOS=linux
ENV CGO_ENABLED=0
RUN go build -o main ./cmd/server

FROM gcr.io/distroless/base
# FROM debian:bullseye-slim

WORKDIR /app
USER 1000:1000
COPY --from=builder --chown=1000:1000  /app/main /app/main
CMD [ "./main" ]
