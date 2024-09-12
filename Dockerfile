FROM golang:1.23.1 as builder
ARG version=dev
WORKDIR /go/src/app
COPY . .
WORKDIR /go/src/app/cmd/taskmanager
RUN go build -o /go/bin/main .

FROM debian:buster-slim
COPY --from=builder /go/bin/main /go/bin/main
ENV PATH="/go/bin:${PATH}"
CMD ["main"]