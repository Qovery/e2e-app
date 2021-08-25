FROM golang:1.16.3-buster as build

WORKDIR /app
COPY . .
RUN go get && go build -o /app.bin main.go

FROM debian:buster-slim as run

RUN apt-get update && apt-get install -y ca-certificates && apt-get clean
COPY --from=build /app.bin /usr/bin/app
RUN chmod +x /usr/bin/app
CMD app