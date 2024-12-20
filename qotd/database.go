package qotd

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

type Quote struct {
	ID       int
	Text     string
	Author   string
	Category string
}

func InitDatabase() *sql.DB {
	// Apre il database SQLite
	db, err := sql.Open("sqlite", "./quotes.db")
	if err != nil {
		log.Fatalf("Errore apertura database: %v", err)
	}

	// Crea la tabella se non esiste
	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS quotes (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            text TEXT NOT NULL,
            author TEXT NOT NULL,
            category TEXT
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
			{Text: "La vita è quello che ti accade mentre sei intento a fare altri piani.", Author: "John Lennon", Category: "Filosofia"},
			{Text: "Non importa quanto vai piano, l'importante è non fermarsi.", Author: "Confucio", Category: "Motivazione"},
			{Text: "Il successo è la somma di piccoli sforzi, ripetuti giorno dopo giorno.", Author: "Robert Collier", Category: "Successo"},
		}

		tx, err := db.Begin()
		if err != nil {
			log.Fatalf("Errore inizio transazione: %v", err)
		}

		for _, quote := range quotes {
			_, err := tx.Exec("INSERT INTO quotes (text, author, category) VALUES (?, ?, ?)",
				quote.Text, quote.Author, quote.Category)
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

	// Query per selezionare una citazione casuale
	row := db.QueryRow(`
        SELECT id, text, author, category 
        FROM quotes 
        ORDER BY RANDOM() 
        LIMIT 1
    `)

	err := row.Scan(&quote.ID, &quote.Text, &quote.Author, &quote.Category)
	if err != nil {
		return Quote{}, err
	}

	return quote, nil
}
