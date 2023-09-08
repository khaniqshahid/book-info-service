package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/khaniqshahid/book-details-service/domain"
	"github.com/khaniqshahid/book-details-service/service"
)

func Start() {

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
	log.Fatal(http.ListenAndServe("localhost:8081", router))
}
