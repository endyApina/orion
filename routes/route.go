package routes

import (
	"orion/controller"
	"os"

	// _ "orion/docs" //swagger docs, you should import it

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

//SetupRouter handles application routing
func SetupRouter(appPort string) *chi.Mux {
	router := chi.NewRouter()

	// Setup CORS
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	router.Use(middleware.Logger)
	// router.Use(middleware.RequestID)
	// router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)
	router.Get("/", controller.InitFunction)
	router.Post("/subscribe", controller.Subscribe)

	hostAddress := os.Getenv("localhost_address")
	if os.Getenv("app_mode") == "prod" {
		hostAddress = os.Getenv("prod_host_address")
	}

	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(hostAddress+appPort+"/swagger/doc.json"),
		// httpSwagger.URL(hostAddress+"/swagger/doc.json"),
	))

	return router
}
