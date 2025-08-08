FROM golang:1.22.1

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /usr/local/bin/shift-app .

CMD ["/usr/local/bin/shift-app"]
