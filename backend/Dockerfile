#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /go/src/app
COPY . .
RUN go build -o /go/bin/app .

#final stage
FROM golang:alpine
RUN apk --no-cache add ca-certificates
RUN apk add --no-cache git
COPY --from=builder /go/bin/app /app
EXPOSE 8080
CMD ["/app"]