package app

import (
	"akim/internal/config"
	"akim/internal/controller/http/api"
	"akim/internal/domain/usecases"
	mongo "akim/internal/infrastructure/repository/mongodb"
	mysql "akim/internal/infrastructure/repository/mysql"
	"akim/utility/server"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg *config.Config) error {
	mysqlDB := mysql.NewMySQLRepository(cfg)
	relationalDatabaseUseCase := usecases.NewRepository(mysqlDB)

	mongoDB, err := mongo.NewMongoDBRepository(cfg)
	if err != nil {
		log.Fatal("troubles with mongo:", err)
	}
	mongodbUseCase := usecases.NewFuzzyArtifactUseCase(mongoDB, mysqlDB)

	router := mux.NewRouter()

	fs := http.FileServer(http.Dir("../static"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	router.HandleFunc("/search", api.FindHandler(relationalDatabaseUseCase)).Methods("POST")
	router.HandleFunc("/fuzzy", api.FuzzyHandler(mongodbUseCase)).Methods("POST")
	router.HandleFunc("/", serveIndex).Methods("GET")
	http.Handle("/", router)

	s := server.New(router, "80")

	wait(s)

	return nil
}

func wait(s *server.Server) {
	log.Println("App started!")

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

	select {
	case i := <-ch:
		log.Println("shutdown signal: " + i.String())
	case err := <-s.Notify():
		log.Fatal(err, "wait - server.Notify")
	}
	log.Println("App is stopping...")
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../internal/infrastructure/presenters/model.html")
}
