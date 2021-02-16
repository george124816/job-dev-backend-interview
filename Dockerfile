## Builder Image
FROM golang:1.15.8-alpine as builder
RUN apk add git
##RUN mkdir /build
COPY . /build/
WORKDIR /build
RUN go build cmd/api/main.go
RUN apk add tree
RUN tree

## Clean Image
FROM alpine:3
COPY --from=builder /build/main .
RUN mkdir -p /db/migrations
COPY --from=builder /build/db/migrations/ ./db/migrations

ENTRYPOINT [ "./main"]