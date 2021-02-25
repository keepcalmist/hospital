package app

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
)

type app struct {
	router  *mux.Router
	db      *sql.DB
	routing map[string]*mux.Router
	quit    chan os.Signal
	server  *http.Server
}

func InitApp(cfg viper.Viper) *app {
	app := &app{}
	var err error
	app.db, err = sql.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			cfg.Get("DB_HOST"), cfg.Get("DB_PORT"), cfg.Get("DB_USER"), cfg.Get("DB_PASSWORD"), cfg.Get("DB_NAME")))
	if err != nil {
		log.Fatal("Cant connect to database")
	}

	app.router = mux.NewRouter()
}

func (s *app) Mount() {
	for path, router := range s.routing {
		s.router.Handle(path, router)
	}
}

func (s *app) AddMiddleware() {
	s.router.Use(mux.CORSMethodMiddleware(s.router))
}

func (s *app) ListenAndServe(cfg viper.Viper) {
	s.server = &http.Server{
		Addr:    cfg.GetString("SERVER_ADDRESS"),
		Handler: s.router,
	}
	s.quit = make(chan os.Signal, 1)
	go func() {
		if err := s.server.ListenAndServe(); err != nil {
			log.Fatal("FATAL CLOSE CONNECTION")
		}
	}()
	<-s.quit
}
