FROM alpine:3.14
RUN apk add --no-cache ca-certificates

WORKDIR /go
COPY ./backend /go/backend
EXPOSE 8080

CMD ["/go/backend"]
