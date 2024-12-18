FROM golang:1.21-alpine AS builder

WORKDIR /app

# Installa le dipendenze di sistema necessarie per la compilazione
RUN apk add --no-cache \
    gcc \
    musl-dev \
    sqlite-dev \
    # Aggiungi questi pacchetti per risolvere problemi di compilazione
    linux-headers \
    build-base

# Copia i file del progetto
COPY go.mod go.sum* ./
COPY *.go ./

# Genera go.sum se non esiste
RUN if [ ! -f go.sum ]; then go mod tidy; fi

# Scarica le dipendenze
RUN go mod download

# Imposta le variabili per la compilazione cross-platform
ENV CGO_ENABLED=1 \
    GOOS=linux \
    CGO_CFLAGS="-D_LARGEFILE64_SOURCE"

# Costruisce l'applicazione
RUN go build \
    -ldflags "-s -w" \
    -o qotd-server

# Imposta i permessi di esecuzione
RUN chmod +x qotd-server

# Immagine finale
FROM alpine:latest

WORKDIR /root/

# Installa le dipendenze di runtime
RUN apk add --no-cache \
    sqlite

# Copia il binario dal builder
COPY --from=builder /app/qotd-server .

# Imposta i permessi di esecuzione nel nuovo stage
RUN chmod +x qotd-server

# Espone la porta QOTD
EXPOSE 17

# Comando di avvio
CMD ["./qotd-server"]