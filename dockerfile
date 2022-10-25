FROM golang:1.19.2-alpine3.16 AS build
RUN apk --no-cache add gcc g++ make ca-certificates
WORKDIR /go/src/github.com/lamhai1401/simplebank-ex
COPY ./ ./
RUN go mod tidy
RUN GO111MODULE=on go build -o /go/bin/app

FROM alpine:3.16
WORKDIR /usr/bin
COPY --from=build /go/bin .
EXPOSE 8080
CMD ["app"]