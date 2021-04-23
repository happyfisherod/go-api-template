FROM golang:1.16-buster AS builder

# GO ENV VARS
ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64

# COPY SRC
WORKDIR /build
COPY models .
COPY worker/src .

# BUILD
RUN go build -o main .

FROM ubuntu as prod
COPY --from=builder /build/main /
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
CMD ["/main"]

FROM builder as test
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
