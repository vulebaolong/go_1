FROM golang:1.26.2-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o go-backend ./cmd


FROM alpine:latest

WORKDIR /app

COPY --from=build /app/go-backend .

CMD ["./go-backend"]

# 2.26gb
# dockerignore: 2.16gb
# stage: 109mb
