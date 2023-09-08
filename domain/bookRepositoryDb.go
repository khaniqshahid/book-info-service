package domain

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type BookRepositoryDb struct {
	client *sql.DB
}

func (d BookRepositoryDb) FindAll() ([]Book, error) {
	// findAllSql := "select book_id, title, author, publisher, price, issued_at, description from books"
	findAllSql := "select book_id, title, author, publisher, price, issued_at, description from books"

	rows, err := d.client.Query(findAllSql)
	if err != nil {
		log.Println("Error while quering books table " + err.Error())
		return nil, err
	}
	books := make([]Book, 0)
	for rows.Next() {
		var b Book
		err := rows.Scan(&b.Id, &b.Title, &b.Author, &b.Publisher, &b.Price, &b.IssuedAt, &b.Description)
		if err != nil {
			log.Println("Error while scanning books table " + err.Error())
			return nil, err
		}
		books = append(books, b)
	}
	return books, nil
}

func (d BookRepositoryDb) ById(id int) (*Book, error) {

	findBookSql := "select book_id, title, author, publisher, price, issued_at, description from books where book_id = ?"
	row := d.client.QueryRow(findBookSql, id)
	var b Book
	err := row.Scan(&b.Id, &b.Title, &b.Author, &b.Publisher, &b.Price, &b.IssuedAt, &b.Description)
	if err != nil {
		log.Println("Error while scanning book " + err.Error())
		return nil, err
	}
	return &b, nil
}

// Helper Function
func NewBookRepositoryDb() BookRepositoryDb {

	client, err := sql.Open("mysql", "admin:P@ssword1@tcp(localhost:3306)/bookdetails?parseTime=true")
	if err != nil {
		panic(err)
	}
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return BookRepositoryDb{client}
}
