package main

import (
	"database/sql"

	"log"
	"os"

	_ "github.com/lib/pq"
)

type Application struct {
	errorLog    *log.Logger
	infoLog     *log.Logger
	userRepo    UserRepository
	templateDir string
}

func connectToDatabase(name string) (*sql.DB, error) {
	db, err := sql.Open("postgres", name)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

func main() {

	// mux := http.NewServeMux()
	// mux.HandleFunc("/", home)
	// mux.HandleFunc("/about", about)
	// mux.HandleFunc("/contact", contact)
	// commented because we are using the Application struct to handle requests now

	db, err := connectToDatabase("host=localhost port=5432 user=postgres password=admin dbname=go_db sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	app := &Application{
		errorLog:    log.New(os.Stderr, "Error\t", log.Ltime|log.LstdFlags|log.Lmicroseconds|log.Lshortfile),
		infoLog:     log.New(os.Stderr, "Error\t", log.Ltime|log.LstdFlags),
		userRepo:    NewSQLUserRepository(db),
		templateDir: "./templates",
	}

	log.Println("Listening on localhost:8080")
	if err := app.serve(); err != nil {
		log.Print(err)
	}

}
