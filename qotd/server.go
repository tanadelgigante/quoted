package qotd

import (
	"database/sql"
	"fmt"
	"log"
	"net"
)

type QOTDServer struct {
	listener net.Listener
	db       *sql.DB
}

func NewQOTDServer(port int, db *sql.DB) *QOTDServer {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Errore creazione listener: %v", err)
	}

	return &QOTDServer{
		listener: listener,
		db:       db,
	}
}

func (s *QOTDServer) Start() {
	log.Printf("Server QOTD in ascolto su :%d", 17)

	for {
		// Accetta connessioni in entrata
		conn, err := s.listener.Accept()
		if err != nil {
			log.Printf("Errore accettazione connessione: %v", err)
			continue
		}

		// Gestisce ogni connessione in un goroutine separato
		go s.handleConnection(conn)
	}
}

func (s *QOTDServer) handleConnection(conn net.Conn) {
	defer conn.Close()

	// Recupera una citazione casuale
	quote, err := getRandomQuote(s.db)
	if err != nil {
		log.Printf("Errore recupero citazione: %v", err)
		conn.Write([]byte("Impossibile recuperare una citazione.\n"))
		return
	}

	// Formatta la citazione secondo lo standard QOTD
	quoteMessage := fmt.Sprintf("\"%s\"\n\t- %s\r\n", quote.Text, quote.Author)

	// Invia la citazione
	_, err = conn.Write([]byte(quoteMessage))
	if err != nil {
		log.Printf("Errore invio citazione: %v", err)
	}
}

func (s *QOTDServer) Stop() {
	s.listener.Close()
	s.db.Close()
	log.Println("Server QOTD fermato")
}
