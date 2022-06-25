package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/nachogiljaldo/openapi-example/docs"
	"github.com/nachogiljaldo/openapi-example/pkg/controller"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @title Swagger User API
// @version 1.0
// @description API for User operations.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v1
func main() {

	docs.SwaggerInfo.Host = "locahost:8080"
	r := chi.NewRouter()

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), //The url pointing to API definition
	))

	r.Delete("/users/{id}", controller.UpdateUser)
	r.Get("/users", controller.GetUsers)
	r.Get("/users/{id}", controller.GetUser)
	r.Post("/users", controller.CreateUser)
	r.Put("/users", controller.UpdateUser)

	http.ListenAndServe(":8080", r)
}
