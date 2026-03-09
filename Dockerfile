FROM golang:latest as build

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN GOOS="linux" GOARCH="arm64" CGO_ENABLED=0 go build -o scAuth ./cmd/main/main.go

FROM --platform=linux/arm64 alpine

COPY --from=build /app/scAuth /app/scAuth

WORKDIR /app

EXPOSE 8080

CMD ["/app/scAuth"]
