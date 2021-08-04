FROM golang:alpine AS builder
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go env -w GOPROXY=https://goproxy.cn && CGO_ENABLED=0 go build -ldflags "-w -s" -o main

FROM alpine
RUN rm -rf /app
WORKDIR /app
COPY --from=builder /app/main /app/main
COPY run.sh /app
COPY static /app/static
ADD config/config.yml config/config.yml
CMD ["sh", "./run.sh"]
