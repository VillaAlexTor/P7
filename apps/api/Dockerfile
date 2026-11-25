FROM golang:1.24-alpine AS build
WORKDIR /src
COPY go.mod ./
COPY cmd ./cmd
COPY internal ./internal
RUN CGO_ENABLED=0 go build -o /p7-api ./cmd/server

FROM alpine:3.21
RUN adduser -D -H -u 10001 p7
USER p7
COPY --from=build /p7-api /p7-api
EXPOSE 8080
ENTRYPOINT ["/p7-api"]

