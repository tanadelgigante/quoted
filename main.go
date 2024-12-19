package main

import (
	"quotd"
)

func main() {
    // Inizializza il database
    db := quoted.InitDatabase()
    defer db.Close()

    // Crea il server QOTD
    server := quoted.NewQOTDServer(17, db)

    // Gestisce segnali per un arresto graceful
    go func() {
        sigChan := make(chan os.Signal, 1)
        signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
        <-sigChan
        server.Stop()
        os.Exit(0)
    }()

    // Avvia il server
    server.Start()
}
