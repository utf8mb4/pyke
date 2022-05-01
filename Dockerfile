FROM golang:1.17 as build

ENV GOPROXY=https://goproxy.cn,direct

WORKDIR /
ADD go.sum go.sum
ADD go.mod go.mod
ADD main.go main.go

RUN go build -o app main.go

FROM centos:7

WORKDIR /
COPY --from=build /app  /app

CMD ["./app"]
