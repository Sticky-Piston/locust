# # Builder
FROM golang:latest as build

RUN apt-get -y update && apt-get -y upgrade

COPY . /build

WORKDIR /build

RUN go mod download
RUN make build

# Distribution
FROM debian:latest

WORKDIR /

COPY --from=build /build/locust /locust

#RUN go mod download

#RUN make build

#USER nonroot:nonroot

ENTRYPOINT ./locust node
