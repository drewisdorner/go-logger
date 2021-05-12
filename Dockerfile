FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/github.ibm.com/Drewis-Dorner/go-logger/
COPY . .
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/logger


FROM scratch
COPY --from=builder /go/bin/logger /logger

EXPOSE 8080
ENTRYPOINT [ "/logger" ]