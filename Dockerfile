FROM golang:1.11 AS builder
WORKDIR /go/src
COPY ./ .
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
RUN go get -d -v ./...
RUN go build -a -installsuffix cgo -ldflags="-s -w" -o main

FROM scratch AS runtime
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/src/.env ./
COPY --from=builder /go/src/main ./
EXPOSE 8080/tcp
ENTRYPOINT ["./main"]