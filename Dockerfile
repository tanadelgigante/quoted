# Fase di build
FROM golang:1.21-alpine AS build

WORKDIR /app

# Imposta le variabili di ambiente per la compilazione cross-platform
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

# Copia il file go.mod e genera il file go.sum
COPY go.mod ./

# Copia il resto del codice dell'applicazione
COPY . .

# Scarica le dipendenze del modulo
RUN go mod download

# Costruisce l'applicazione Go e verifica il tipo di file
RUN go build -o qotd-server . && file qotd-server

# Fase di runtime
FROM alpine:latest

# Installazione delle dipendenze necessarie
RUN apk --no-cache add sqlite

WORKDIR /app

# Copia l'eseguibile dall'immagine di build
COPY --from=build /app/qotd-server /app/qotd-server

# Imposta i permessi di esecuzione sull'eseguibile
RUN chmod +x /app/qotd-server

# Copia il file del database se esiste o crea una nuova directory per il database
COPY --from=build /app/quotes.db /app/quotes.db

# Esponi la porta 17 per il server QOTD
EXPOSE 17

# Comando di esecuzione dell'applicazione
CMD ["./qotd-server"]
