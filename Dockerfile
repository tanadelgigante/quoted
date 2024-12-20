# Fase di build
FROM golang:1.23.4 AS build

WORKDIR /app

# Copia i file go.mod e go.sum
COPY go.mod .
COPY go.sum .

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
