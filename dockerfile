FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o penggalangan-dana

EXPOSE 8080

CMD ["/penggalangan-dana"]
