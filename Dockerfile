FROM golang:1.12.1

ENV PORT 3000
WORKDIR /usr/local/go/src/watchmen
ADD . /usr/local/go/src/watchmen
RUN go build -o main .
CMD ["./main"]
