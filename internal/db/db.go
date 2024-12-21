package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() error {
	// In production or advanced usage, load from environment variables or a config file
	host := "localhost"
	port := 5432
	user := "postgres"       // or "movieadmin" if you created that user
	password := "qwerty123!" // if needed
	dbname := "cvwo_web_forum"

	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	var err error
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}

	// Check if db is reachable
	err = DB.Ping()
	if err != nil {
		return err
	}

	fmt.Println("Connected to Postgres successfully!")
	return nil
}

func Migrate() error {
	createUsersTable := `
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        username VARCHAR(100) UNIQUE NOT NULL,
        score INT DEFAULT 0,
        is_anonymous BOOLEAN DEFAULT FALSE
    );
    `

	createThreadsTable := `
    CREATE TABLE IF NOT EXISTS threads (
        id SERIAL PRIMARY KEY,
        title VARCHAR(255) NOT NULL,
        topic VARCHAR(255) NOT NULL, -- e.g. "Action", "Rom-com"
        creator_id INT REFERENCES users(id),
        theme_preference VARCHAR(50) DEFAULT 'light'
    );
    `

	createPostsTable := `
    CREATE TABLE IF NOT EXISTS posts (
        id SERIAL PRIMARY KEY,
        content TEXT NOT NULL,
        thread_id INT REFERENCES threads(id),
        author_id INT REFERENCES users(id),
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );
    `

	_, err := DB.Exec(createUsersTable)
	if err != nil {
		return err
	}
	_, err = DB.Exec(createThreadsTable)
	if err != nil {
		return err
	}
	_, err = DB.Exec(createPostsTable)
	if err != nil {
		return err
	}

	return nil
}
