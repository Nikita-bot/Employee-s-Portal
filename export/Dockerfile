FROM golang:alpine as builder

WORKDIR /build

RUN apk update --no-cache add bash git gcc gettext musl-dev

COPY . . 

RUN go mod download

RUN go build -ldflags="-s -w" -o /build/export ./export.go

FROM alpine

RUN apk update --no-cache

WORKDIR /app

COPY --from=builder /build/export /app/export

CMD ["/app/export"]