# Builder
FROM golang:latest

RUN apt-get -y update && apt-get -y upgrade

COPY . /build

WORKDIR /build
RUN go build

# Distribution
FROM debian:latest

RUN apt-get -y update && apt-get -y upgrade

WORKDIR /app 

EXPOSE 9090

COPY --from=builder /build/locust ./locust

CMD ./locust