#builder container
FROM golang:1.12-alpine as builder
LABEL maintainer="Khan Sadirac <khan.sadirac42@gmail.com"
WORKDIR /app
COPY . .
RUN go build -o sfcc_exporter

# main container
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app .
EXPOSE 9240
CMD ["./sfcc_exporter"]