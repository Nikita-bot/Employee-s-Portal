FROM golang:alpine as builder

WORKDIR /build

RUN apk update --no-cache add bash git gcc gettext musl-dev

COPY . . 

RUN go mod download

RUN go build -ldflags="-s -w" -o /build/portal cmd/main.go

FROM alpine

RUN apk update --no-cache

WORKDIR /app

COPY ./log /app/log
COPY ./static /app/static
COPY ./migrations /app/migrations
COPY --from=builder /build/portal /app/portal

CMD ["/app/portal"]