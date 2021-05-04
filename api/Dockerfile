FROM golang:1.16-buster AS builder

# GO ENV VARS
ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64

# COPY SRC
WORKDIR /build
COPY ./api/src .

# DOWNLOAND DEPS
RUN go mod tidy
RUN go mod download

# CREATE SWAGGER DOCS
RUN go get -u github.com/swaggo/swag/cmd/swag
RUN swag init -g server/server.go

# BUILD
RUN go build -o main .

FROM ubuntu as prod
COPY ./models ./models
COPY --from=builder /build/main /
#COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
CMD ["/main"]

FROM builder as test
#COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
