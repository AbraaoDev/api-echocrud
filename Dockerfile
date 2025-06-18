FROM golang:1.24-alpine as stage

RUN apk add --no-cache tzdata

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# compila o bin
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o server cmd/server.go

# bin para image
FROM scratch

# apenas o binario
COPY --from=stage /app/server /server

COPY --from=stage /usr/share/zoneinfo /usr/share/zoneinfo

EXPOSE 3333

ENTRYPOINT ["/server"]
