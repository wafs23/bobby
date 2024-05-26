FROM golang:1.22-alpine3.20 AS builder

RUN go version

WORKDIR /go/src/app
COPY . .
RUN go build -o main main.go

FROM alpine:3.20
WORKDIR /go/src/app

COPY --from=builder /go/src/app/main .
# COPY app.env .
COPY ./templates ./
ADD templates ./templates
COPY ./assets ./
ADD assets ./assets

EXPOSE 9033

RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser

CMD ["/go/src/app/main"]