FROM golang:1.16-buster AS builder

# GO ENV VARS
ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64

# Build Delve
RUN go get github.com/go-delve/delve/cmd/dlv

# COPY SRC
WORKDIR /build
COPY ./api/src .

# BUILD
RUN go build -gcflags="all=-N -l" -o main .

FROM ubuntu as prod

COPY ./models ./models
COPY --from=builder /build/main /
#COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
CMD ["/main"]

FROM builder as test
EXPOSE 8000 40000

COPY ./models ./models
COPY --from=builder /go/bin/dlv /
COPY --from=builder /build/main /
#COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
CMD ["/dlv", "--headless", "--listen=:40000", "--api-version=2", "exec", "/main"]
