# Fase di build
FROM golang:1.21-alpine AS build

WORKDIR /app

# Copia solo il file go.mod e genera go.sum
COPY go.mod ./
RUN go mod tidy

# Copia il resto del codice dell'applicazione
COPY . .

# Scarica le dipendenze del modulo
RUN go mod download

# Costruisce l'applicazione Go
RUN go build -o qotd-server .

# Fase di runtime
FROM alpine:latest

# Installazione delle dipendenze necessarie
RUN apk --no-cache add sqlite

WORKDIR /app

# Copia l'eseguibile dall'immagine di build
COPY --from=build /app/qotd-server /app/qotd-server

# Copia il file del database se esiste o crea una nuova directory per il database
COPY --from=build /app/quotes.db /app/quotes.db

# Esponi la porta 17 per il server QOTD
EXPOSE 17

# Comando di esecuzione dell'applicazione
CMD ["./qotd-server"]
