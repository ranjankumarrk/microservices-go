package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type Book struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Author      string  `json:"author"`
	Owner       string  `json:"owner"`
	PublishYear int     `json:"publishYear"`
	Price       float64 `json:"price"`
}

type User struct {
	ID       int    `json:"id"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"userName"`
	jwt.RegisteredClaims
}

var db *sql.DB

func main() {
	db, err := sql.Open("mysql", "root:my-secret-pw@tcp(127.0.0.1:3306)/bookswap")
	if err != nil {
		log.Fatal("Error while connecting to database:", err)
	}
	log.Println("Database connected successfully")
	defer db.Close()

	router := mux.NewRouter()

	router.HandleFunc("/api/books", getAllBooksHandler).Methods("GET")

	router.HandleFunc("/api/v1/login", loginHandler).Methods("POST")

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}

/*
 */
func loginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Login successful"})
}

// 1. get books
func getAllBooksHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("select id, title, author, owner, publishYear, price from books")
	if err != nil {
		log.Print("Error while getting books:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	books := []Book{}
	for rows.Next() {
		var book Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Owner, &book.PublishYear, &book.Price); err != nil {
			log.Print("Error while scanning the row:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		books = append(books, book)
	}
	if err = rows.Err(); err != nil {
		log.Print("Error while getting books:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}
