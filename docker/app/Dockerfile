FROM golang:1.19-alpine AS builder

WORKDIR /src
COPY cmd/app/*.go ./
COPY internal/app/user/*.go ./internal/app/user/
COPY pkg/mongodb/*.go ./pkg/mongodb/
COPY pkg/helpers/*.go ./pkg/helpers/
COPY pkg/validator/*.go ./pkg/validator/
COPY go.mod .
COPY go.sum .
RUN go mod tidy

RUN go build -o cmd/app .

FROM alpine
WORKDIR /src
COPY --from=builder /src .

EXPOSE 3000