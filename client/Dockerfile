# Build app in a stock Go builder container
FROM golang:1.11.4-alpine3.8 as builder

RUN apk add --no-cache make gcc musl-dev linux-headers git

COPY . /go/src/testms
WORKDIR /go/src/testms

RUN go get github.com/tools/godep && godep restore
RUN go test ./...
RUN cd cmd && go build -o testapp

## Pull binaries into a second stage deploy alpine container
FROM alpine:latest
RUN apk add --no-cache ca-certificates

COPY --from=builder /go/src/testms/cmd/testapp /testms/

WORKDIR /testms

RUN chmod +x testapp

CMD ["./testapp"]