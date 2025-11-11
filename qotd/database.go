package qotd

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

type Quote struct {
	Text   string
	Author string
}

func InitDatabase() *sql.DB {
	// Apre il database SQLite
	db, err := sql.Open("sqlite", "./quotes.db")
	if err != nil {
		log.Fatalf("Errore apertura database: %v", err)
	}

	// Crea la tabella se non esiste (solo id, text, author)
	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS quotes (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            text TEXT NOT NULL,
            author TEXT NOT NULL
        )
    `)
	if err != nil {
		log.Fatalf("Errore creazione tabella: %v", err)
	}

	// Inserisce citazioni di esempio se la tabella è vuota
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM quotes").Scan(&count)
	if err != nil {
		log.Fatalf("Errore conteggio citazioni: %v", err)
	}

	if count == 0 {
		quotes := []Quote{
			{Text: "La vita è quello che ti accade mentre sei intento a fare altri piani.", Author: "John Lennon"},
			{Text: "Non importa quanto vai piano, l'importante è non fermarsi.", Author: "Confucio"},
			{Text: "Il successo è la somma di piccoli sforzi, ripetuti giorno dopo giorno.", Author: "Robert Collier"},
		}

		tx, err := db.Begin()
		if err != nil {
			log.Fatalf("Errore inizio transazione: %v", err)
		}

		for _, q := range quotes {
			_, err := tx.Exec("INSERT INTO quotes (text, author) VALUES (?, ?)",
				q.Text, q.Author)
			if err != nil {
				tx.Rollback()
				log.Fatalf("Errore inserimento citazione: %v", err)
			}
		}

		err = tx.Commit()
		if err != nil {
			log.Fatalf("Errore commit transazione: %v", err)
		}
	}

	return db
}

func getRandomQuote(db *sql.DB) (Quote, error) {
	var quote Quote

	// Query per selezionare una citazione casuale (solo text e author)
	row := db.QueryRow(`
        SELECT text, author
        FROM quotes
        ORDER BY RANDOM()
        LIMIT 1
    `)

	err := row.Scan(&quote.Text, &quote.Author)
	if err != nil {
		return Quote{}, err
	}

	return quote, nil
}
