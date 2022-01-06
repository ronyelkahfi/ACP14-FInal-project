FROM golang:1.17-alpine AS build
WORKDIR /app
COPY go.mod .
COPY go.sum .
COPY .env .

RUN go mod download

COPY . .
RUN go build -o app

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/app /acp14
COPY --from=build /app/apps/configs/config.json apps/configs/config.json
EXPOSE 8080
CMD ["/acp14"]