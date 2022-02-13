FROM golang:1.16.3

WORKDIR /cmd/hello-go

COPY bin/hello-go .

ENTRYPOINT ["./hello-go"]