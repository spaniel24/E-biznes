FROM golang:1.18.4-alpine3.16

RUN apk add --no-cache git
RUN apk add build-base

WORKDIR /app/e-biznes_go

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o ./out/e-biznes_go .

EXPOSE 8080

CMD ["./out/e-biznes_go"]