FROM golang:1.8

WORKDIR /go/src/golamapi

COPY . .

CMD ["aah", "r", "-i", "golamapi"]