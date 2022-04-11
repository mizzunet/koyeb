FROM golang:latest-alpine AS builder
RUN mkdir /build
ADD go.mod go.sum hello.go /build/
WORKDIR /build
RUN go build

FROM alpine
RUN echo hellow
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/helloworld /app/
COPY views/ /app/views
WORKDIR /app
CMD ["./helloworld"]
