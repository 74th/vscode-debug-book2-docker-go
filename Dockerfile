FROM golang:1.16-buster as build
RUN go install github.com/go-delve/delve/cmd/dlv@v1.6.1
WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .
RUN go build -o app main.go

FROM debian:buster as production
WORKDIR /app
COPY --from=build /app/app .

CMD ["/app/app"]