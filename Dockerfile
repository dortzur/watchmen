FROM golang:1.12.1

ENV PORT 3000
ENV GO_ENV production
ENV GIN_MODE release
WORKDIR /usr/local/go/src/watchmen
ADD . /usr/local/go/src/watchmen
RUN go build -o main .
CMD ["./main"]
