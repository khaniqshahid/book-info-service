package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/khaniqshahid/book-details-service/domain"
	"github.com/khaniqshahid/book-details-service/logger"
	"github.com/khaniqshahid/book-details-service/service"
)

func sanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" ||
		os.Getenv("SERVER_PORT") == "" {
		logger.Fatal("Environment variable not defined")
		// Run following in your linux terminal
		// export SERVER_ADDRESS=localhost && export SERVER_PORT=8081 export DB_USER=root && export DB_PASS=admin && export DB_HOST=localhost && export DB_PORT=3306 && export DB_NAME=bookdetails
	}
}

func Start() {

	sanityCheck()
	router := mux.NewRouter()
	//Wiring up with the help of a handler BookInfoHandler which connects bookInfo service-interface(Primary) ---> Domain--->(Secondary)RepoInterface ---> Repository
	// Injected stub to test the complete wiring of the interfaces
	// bh := BookInfoHandler{service.NewCustomBookService(domain.NewBookRepositoryStub())}
	// Injected DB connection
	bh := BookInfoHandler{service.NewCustomBookService(domain.NewBookRepositoryDb())}
	// Start Routes
	router.HandleFunc("/books", bh.getAllBooks).Methods(http.MethodGet)
	router.HandleFunc("/books/{book_id:[0-9]+}", bh.GetBookById).Methods(http.MethodGet)

	// Start Server

	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	// log.Fatal(http.ListenAndServe("localhost:8081", router))
	logger.Info("Starting server on " + address + " port:" + port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}
