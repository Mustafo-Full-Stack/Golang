# Этап 1: сборка программы.
FROM golang:1.27-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /bin/api ./cmd/api

# Этап 2: компактный образ приложения.
FROM alpine:3.20
WORKDIR /app
COPY --from=build /bin/api /bin/api
EXPOSE 8080
CMD ["/bin/api"]
