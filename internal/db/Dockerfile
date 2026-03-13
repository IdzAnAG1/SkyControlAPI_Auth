FROM --platform=linux/arm64 golang:latest as build

RUN GOOS="linux" GOARCH="arm64" CGO_ENABLED=0 go install github.com/pressly/goose/v3/cmd/goose@latest

FROM --platform=linux/arm64 alpine

COPY --from=build /go/bin/goose /usr/local/bin/goose

COPY ./migrations ./migrations

ENV GOOSE_DRIVER=postgres

CMD ["goose", "-dir", "./migrations", "up"]