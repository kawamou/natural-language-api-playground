# 開発環境
FROM golang:1.18-alpine as dev
WORKDIR /go/src/app
ENV GO111MODULE=auto
RUN apk add --no-cache \
        alpine-sdk \
        git \
    && go get -u github.com/cosmtrek/air
CMD ["air"]

# ビルド用
FROM golang:1.18-alpine as builder
WORKDIR /go/src/app
COPY myapp /go/src/app
ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -v -o myapp

# 本番環境
FROM gcr.io/distroless/base as prod
COPY --from=builder /go/src/app/myapp /myapp
CMD ["/myapp"]