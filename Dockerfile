FROM golang:1.21-alpine AS builder

WORKDIR /app

# Installa le dipendenze di sistema
RUN apk add --no-cache \
    gcc \
    musl-dev \
    sqlite-dev

# Copia i file del progetto
COPY go.mod ./
COPY *.go ./

# Scarica le dipendenze
RUN go mod download

# Costruisce l'applicazione
RUN CGO_ENABLED=1 GOOS=linux go build -o quoted

# Immagine finale più leggera
FROM alpine:latest

WORKDIR /root/

# Installa le dipendenze di runtime
RUN apk add --no-cache \
    sqlite

# Copia il binario dal builder
COPY --from=builder /app/quoted .

# Espone la porta QOTD
EXPOSE 17

# Comando di avvio
CMD ["./quoted"]