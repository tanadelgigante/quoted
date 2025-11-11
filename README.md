# Qotd

## Overview
**Qotd** è un server per il protocollo QOTD (Quote of the Day) scritto in Go. Fornisce una citazione del giorno ogni volta che viene effettuata una connessione al server.

## Features
- **Protocollo QOTD**: Implementa il protocollo QOTD per fornire citazioni del giorno.
- **Facile da configurare**: Configurazione semplice e immediata.
- **Leggero e veloce**: Scritto in Go per prestazioni ottimali.

## Application Information
- **Name**: Qotd
- **Version**: 1.0.0
- **Author**: @ilgigante77
- **Website**: [https://github.com/tanadelgigante/quoted](https://github.com/tanadelgigante/quoted)

## Getting Started

### Prerequisites
- Go 1.23+
- Docker (opzionale per il deployment containerizzato)

### Installation

1. **Clone the repository**:
    ```bash
    git clone https://github.com/tanadelgigante/quoted.git
    cd quoted
    ```

2. **Build the application**:
    ```bash
    go build -o qotd-server main.go
    ```

### Configuration

1. **Database Setup**:
   Assicurati di avere un file `quotes.db` nella directory di lavoro. Questo file dovrebbe contenere le citazioni che il server fornirà.

### Running the Application

1. **Run Locally**:
    ```bash
    ./qotd-server
    ```

2. **Using Docker**:
   Crea un Dockerfile per il server:
    ```dockerfile
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
    ```

   Build and run the Docker container:
    ```bash
    docker build -t quoted .
    docker run -p 17:17 quoted
    ```

### Usage

#### API Endpoints
- **QOTD Request**:
    ```bash
    telnet localhost 17
    ```
    Response:
    ```
    "La citazione del giorno è..."
    ```

#### Quotes DB

- **Inizializzazione database**:
     ```sql
    BEGIN TRANSACTION;

    CREATE TABLE IF NOT EXISTS quotes (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    text TEXT NOT NULL,
    author TEXT,
    created_at DATETIME NOT NULL DEFAULT (datetime('now'))
    );

    CREATE INDEX IF NOT EXISTS idx_quotes_author ON quotes(author);

    -- Esempi
    INSERT INTO quotes(text, author) VALUES ('Fletto i muscoli e sono nel vuoto', 'Rat-Man');

    COMMIT;
     ```
- **Template per le insert**:
    ```sql
    -- Template per inserire una citazione (usare binding :text e :author dal client oppure sostituire i valori)
    INSERT INTO quotes(text, author, created_at) VALUES (:text, :author, datetime('now'));
    -- Esempio CLI:
    -- sqlite3 quotes.db "INSERT INTO quotes(text,author) VALUES('Una nuova citazione','Autore');"
    ```

### Debugging

- Usa i log dell'applicazione per monitorare le connessioni e le risposte del server. Cerca messaggi `[INFO]` e `[DEBUG]` nell'output della console.

### Contributing
Le contribuzioni sono benvenute! Fai un fork del repository e invia pull request per miglioramenti o correzioni di bug.

### License
Questo progetto è concesso in licenza sotto la GPL 3.0 License. Vedi il file [LICENSE](LICENSE) per i dettagli.

### Disclaimer
Questo progetto è rilasciato "as-is" e l'autore non è responsabile per danni, errori o uso improprio.

## Contact
Per maggiori informazioni, visita [https://github.com/tanadelgigante/quoted](https://github.com/tanadelgigante/quoted).


