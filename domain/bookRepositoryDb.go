package domain

import (
	"database/sql"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/khaniqshahid/book-details-service/errs"
	"github.com/khaniqshahid/book-details-service/logger"
)

type BookRepositoryDb struct {
	client *sqlx.DB
}

func (d BookRepositoryDb) FindAll() ([]Book, *errs.AppError) {

	var err error
	books := make([]Book, 0)

	findAllSql := "select book_id, title, author, publisher, price, issued_at, description from books"

	err = d.client.Select(&books, findAllSql)

	if err != nil {
		logger.Error("Error while quering books table :" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error ")
	}

	// err = sqlx.StructScan(rows, &books)
	// if err != nil {
	// 	logger.Error("Error while scanning books table :" + err.Error())
	// 	return nil, errs.NewUnexpectedError("Unexpected database error")
	// }
	// for rows.Next() {
	// 	var b Book
	// 	err := rows.Scan(&b.Id, &b.Title, &b.Author, &b.Publisher, &b.Price, &b.IssuedAt, &b.Description)
	// 	if err != nil {
	// 		logger.Error("Error while scanning books table " + err.Error())
	// 		return nil, errs.NewUnexpectedError("Unexpected database error")
	// 	}
	// 	books = append(books, b)
	// }
	return books, nil
}

func (d BookRepositoryDb) ById(id int) (*Book, *errs.AppError) {

	findBookSql := "select book_id, title, author, publisher, price, issued_at, description from books where book_id = ?"
	row := d.client.QueryRow(findBookSql, id)
	var b Book
	err := row.Scan(&b.Id, &b.Title, &b.Author, &b.Publisher, &b.Price, &b.IssuedAt, &b.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			// In case the book id does not exist in the database
			return nil, errs.NewNotFoundError("Book not found")
		} else {
			logger.Error("Error while scanning book " + err.Error())
			// In case of an unexpected error e.g status 500 Internal server error) aong with the error message
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}
	return &b, nil
}

// Helper Function
func NewBookRepositoryDb() BookRepositoryDb {

	//OS ENVnvironment variable
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dataSource := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?parseTime=true"
	// client, err := sql.Open("mysql", "admin:P@ssword1@tcp(localhost:3306)/bookdetails?parseTime=true")
	// client, err := sqlx.Open("mysql", "root:admin@tcp(localhost:3306)/bookdetails?parseTime=true")
	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return BookRepositoryDb{client}
}
