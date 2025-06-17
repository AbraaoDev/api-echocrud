FROM golang:1.24-alpine as stage1

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# compila o bin
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o server

# bin para image
FROM scratch

# apenas o binario
COPY --from=stage1 /app/server / 

ENTRYPOINT ["/server"]
