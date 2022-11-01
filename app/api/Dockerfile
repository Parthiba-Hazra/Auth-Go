FROM golang:alpine as builder 

WORKDIR /go/src/app

# This repo for live reloading while building the application
RUN go install github.com/cespare/reflex@latest

COPY go.mod .
COPY go.sum .

RUN go mod tidy

COPY . .

RUN go build -o ./run .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the executable image from builder
COPY --from=builder /go/src/app/run .

EXPOSE 8080

CMD ["./run"]