package main

import (
	"client-server/internal/server/apihandler"
	"client-server/internal/server/dbservice"
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	mux := IniciaServeMux()
	http.ListenAndServe(":8080", mux)
}

func IniciaServeMux() *http.ServeMux {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {

	}

	db.AutoMigrate(&dbservice.Cotacao{})

	cotacaoRepo := dbservice.NewCotacaoRepo(db)
	cotacaoApiHandler := apihandler.NewCotacaoApiHandler(*cotacaoRepo)

	mux := http.NewServeMux()
	mux.HandleFunc("/cotacao", cotacaoApiHandler.GetCotacaoHandler)

	return mux
}
