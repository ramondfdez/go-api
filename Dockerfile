FROM golang:latest as builder

WORKDIR /go/src/app
COPY . .

RUN go mod download

RUN CGO_ENABLED=0 go build -o /go/bin/app

FROM alpine:latest

COPY --from=builder /go/bin/app /
CMD ["/app"]