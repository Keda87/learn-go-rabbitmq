package app

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-playground/validator/v10"

	"github.com/Keda87/learn-go-rabbitmq/service/api/health"
	"github.com/Keda87/learn-go-rabbitmq/service/api/publisher"
	"github.com/Keda87/learn-go-rabbitmq/service/config"
)

type App struct {
	router   *chi.Mux
	validate *validator.Validate
}

func New(conf *config.Config) *App {
	app := &App{}

	app.initValidator()
	app.initQueue()
	app.initRoutes()

	return app
}

func (a *App) initRoutes() {
	a.router = chi.NewRouter()

	// setup middlewares.
	a.router.Use(middleware.RequestID)
	a.router.Use(middleware.RealIP)
	a.router.Use(middleware.Logger)
	a.router.Use(middleware.Recoverer)

	// controller definitions.
	healthController := health.NewController()
	publisherController := publisher.NewController(a.validate)

	// setup routes.
	a.router.Get("/", healthController.HandlerHealth)
	a.router.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Route("/publishers", func(r chi.Router) {
				r.Post("/", publisherController.HandlerPublishMessage)
			})
		})
	})
}

func (a *App) initValidator() {
	a.validate = validator.New()
}

func (a *App) initQueue() {

}

func (a *App) Start() {
	quitChan := make(chan os.Signal)
	signal.Notify(quitChan,
		syscall.SIGTERM, // stopped by: `KILL 9`
		syscall.SIGINT,  // stopped by: `Ctrl + C`
	)

	go func() {
		defer func() {
			quitChan <- syscall.SIGTERM
		}()

		if err := http.ListenAndServe(":5000", a.router); err != nil {
			fmt.Println(err)
		}
	}()

	fmt.Println("Listening to port: 5000")
	<-quitChan
	fmt.Println("Shutting down...")
}

func (a *App) Stop() {
	fmt.Println("Shutdown gracefully")
}
