FROM golang:1.17-alpine AS build
RUN mkdir /app
ADD . /app
WORKDIR /app
COPY go.mod .
COPY go.sum .
COPY json /apps
RUN go mod download

COPY . .
RUN go build -o app

FROM alpine:latest
COPY --from=build /app/app /app

EXPOSE 8080
CMD ["/app"]