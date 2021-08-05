FROM golang:alpine AS builder
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go env -w GOPROXY=https://goproxy.cn && CGO_ENABLED=0 go build -ldflags "-w -s" -o main

FROM alpine
RUN apk add gettext
RUN rm -rf /app
WORKDIR /app
COPY --from=builder /app/main /app/main
COPY run.sh /app
COPY static /app/static
COPY config/config.tmpl config/config.tmpl
CMD ["sh", "./run.sh"]
