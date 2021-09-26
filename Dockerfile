FROM golang:1.17-alpine AS build
WORKDIR /app

COPY . .
RUN go mod tidy
RUN go build -o nayatechid *.go

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/nayatechid .
COPY --from=build /app/public ./public/
EXPOSE 8080
RUN ls /app
CMD ["/app/nayatechid"]
