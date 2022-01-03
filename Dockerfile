FROM golang:1.17-alpine AS build
WORKDIR /app
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .
RUN go build -o app

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/app /app
COPY --from=build /app/apps/configs/config.json apps/configs/config.json
EXPOSE 8080
CMD ["/app"]