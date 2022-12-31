# base go image
FROM golang:1.18-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

# build executable go with option not to enable c go (only use standard go lib)
RUN CGO_ENABLED=0 go build -o brokerApp ./cmd/api

# give enough permission to run executable go
RUN chmod +x /app/brokerApp

# build a tiny docker image
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/brokerApp /app

CMD [ "/app/brokerApp" ]