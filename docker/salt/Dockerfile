FROM golang:1.19-alpine AS builder

WORKDIR /src
COPY cmd/salt/*.go ./
COPY pkg/helpers/*.go ./pkg/helpers/
COPY go.mod .
COPY go.sum .
RUN go mod tidy

RUN go build -o cmd/salt .

FROM alpine
WORKDIR /src
COPY --from=builder /src .

EXPOSE 9000
