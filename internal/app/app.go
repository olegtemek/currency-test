package app

import (
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/olegtemek/currency-task/docs"
	"github.com/olegtemek/currency-task/internal/config"
	"github.com/olegtemek/currency-task/internal/handler"
	"github.com/olegtemek/currency-task/internal/repository"
	"github.com/olegtemek/currency-task/internal/service"
	"github.com/olegtemek/currency-task/internal/storage"
	httpSwagger "github.com/swaggo/http-swagger"
)

type handlerI interface {
	Save(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type App struct {
	Log      *slog.Logger
	Cfg      *config.Config
	Handler  *mux.Router
	Handlers handlerI
}

func New(log *slog.Logger, cfg *config.Config) *App {

	db, err := storage.NewPostgresConnect(log, cfg.DbUrl)
	if err != nil {
		panic("cannot connect to db")
	}

	repository := repository.New(log, db)
	service := service.New(log, repository)
	handlers := handler.New(log, service)

	return &App{
		Log:      log,
		Cfg:      cfg,
		Handler:  mux.NewRouter(),
		Handlers: handlers,
	}
}

func (h *App) Init() *http.Server {

	h.Handler.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	h.InitAllRoutes()

	srv := &http.Server{
		Addr:    h.Cfg.ServerAddr,
		Handler: h.Handler,
	}

	return srv
}

func (h *App) InitAllRoutes() {

	h.Handler.HandleFunc("/currency/save/{date}", h.Handlers.Save)

	h.Handler.HandleFunc("/currency/{date}", h.Handlers.Get)
	h.Handler.HandleFunc("/currency/{date}/{code}", h.Handlers.Get)

}
