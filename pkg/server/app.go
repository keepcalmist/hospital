package app

import (
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/keepcalmist/hospital/pkg/doctors"
	"github.com/keepcalmist/hospital/pkg/interfases"
	"github.com/keepcalmist/hospital/pkg/middleware"
	"github.com/keepcalmist/hospital/pkg/repository"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
)

type app struct {
	router  *mux.Router
	db      *sqlx.DB
	routing map[string]interfases.Connector
	quit    chan os.Signal
	server  *http.Server
}

func InitApp(cfg *viper.Viper) *app {
	app := &app{
		routing: make(map[string]interfases.Connector),
	}
	var err error
	app.db, err = sqlx.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			cfg.Get("DB_HOST"), cfg.Get("DB_PORT"), cfg.Get("DB_USER"), cfg.Get("DB_PASSWORD"), cfg.Get("DB_NAME")))
	if err != nil {
		log.Println(err)
		log.Fatal("Cant connect to database")
	}

	doctorRepo := repository.NewDoctorRepo(app.db)
	dcts := doctors.NewService(doctorRepo)

	app.router = mux.NewRouter()
	app.addRoute(dcts.Path(), dcts)
	app.router.Handle("/doctors", dcts.Router())

	err = app.checkRepeatingPath()
	if err != nil {
		log.Fatal(err)
	}

	app.mount()
	app.addMiddleware()
	return app
}

func (s *app) mount() {
	for path, router := range s.routing {
		s.router.Handle(path, router.Router())
	}
}

func (s *app) addMiddleware() {
	s.router.Use(mux.CORSMethodMiddleware(s.router))
	s.router.Use(middleware.Logger)
}

func (s *app) ListenAndServe(cfg *viper.Viper) {
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

func (s *app) addRoute(path string, router interfases.Connector) {
	s.routing[path] = router
}

func (s *app) checkRepeatingPath() error {
	checkMap := make(map[string]int)
	for key, _ := range s.routing {
		checkMap[key]++
		if checkMap[key] > 1 {
			return errors.New("Extra path detected ")
		}
	}
	return nil
}
