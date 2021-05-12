FROM golang:1.16-buster AS builder

# GO ENV VARS
ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64

# COPY SRC
WORKDIR /build
COPY worker/src .

# BUILD
RUN go mod tidy
RUN go build -o main .

FROM ubuntu as prod
COPY --from=builder /build/main /
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
CMD ["/main"]

FROM golang:1.16-buster AS test

# GO ENV VARS
ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64

# Build Delve
RUN go get github.com/go-delve/delve/cmd/dlv
# Build Air
RUN go get -u github.com/cosmtrek/air

EXPOSE 8000 40000

RUN mkdir /worker
WORKDIR /worker

#COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
#CMD ["/dlv", "--headless", "--listen=:40000", "--api-version=2", "exec", "/main"]
CMD ["air"]
