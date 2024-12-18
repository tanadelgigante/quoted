# Fase di build
FROM golang:1.21-alpine AS build

# Crea una directory di lavoro
WORKDIR /app

# Copia il modulo Go e i file di dipendenze
COPY go.mod go.sum ./

# Scarica le dipendenze del modulo
RUN go mod download

# Copia il resto del codice dell'applicazione
COPY . .

# Costruisce l'applicazione Go
RUN go build -o qotd-server .

# Fase di runtime
FROM alpine:latest

# Installazione delle dipendenze necessarie
RUN apk --no-cache add sqlite

# Crea una directory di lavoro
WORKDIR /app

# Copia l'eseguibile dall'immagine di build
COPY --from=build /app/qotd-server /app/qotd-server

# Copia il file del database se esiste o crea una nuova directory per il database
COPY --from=build /app/quotes.db /app/quotes.db

# Esponi la porta 17 per il server QOTD
EXPOSE 17

# Comando di esecuzione dell'applicazione
CMD ["./qotd-server"]
