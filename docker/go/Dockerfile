FROM golang:1.16-alpine as dev

ARG PROJECT_WORKING_DIR
WORKDIR ${PROJECT_WORKING_DIR}

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV GO111MODULE=on
ENV GOPATH=/go
RUN apk --no-cache add make

COPY ./src .

RUN go mod tidy && \
    go install github.com/cosmtrek/air@latest

CMD ["air"]

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder ${PROJECT_WORKING_DIR} /app
CMD ["/app"]