FROM golang:latest

WORKDIR /go/src/golamapi

COPY . .

CMD ["aah", "r", "-i", "golamapi"]