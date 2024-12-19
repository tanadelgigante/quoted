# Utilizza una singola fase di build e runtime
FROM golang:1.21-alpine

WORKDIR /app

# Copia i file del modulo Go
COPY go.mod ./

# Copia il resto del codice dell'applicazione
COPY . .

# Scarica le dipendenze del modulo
RUN go mod download

# Esponi la porta 17 per il server QOTD
EXPOSE 17

# Comando di esecuzione dell'applicazione
CMD ["go", "run", "."]
