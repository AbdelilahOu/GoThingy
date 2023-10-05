package application

import (
	"net/http"

	"github.com/AbdelilahOu/GoThingy/handler"
	"github.com/AbdelilahOu/GoThingy/repository"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (a *App) loadRoutes() *chi.Mux {
	// create chi router
	router := chi.NewRouter()
	//  use logger middleware
	router.Use(middleware.Logger)
	// basinc route
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	// sub routes
	router.Route("/orders", a.loadOrderRoutes)
	//
	return router
}

func (a *App) loadOrderRoutes(router chi.Router) {
	// get order handler
	orderHandler := &handler.Order{
		Repo: &repository.OrderRepo{
			DB: a.db,
		},
	}
	// attach routes
	router.Post("/", orderHandler.Create)
	router.Get("/", orderHandler.GetAll)
	router.Get("/{id}", orderHandler.GetByID)
	router.Put("/{id}", orderHandler.UpdateByID)
	router.Delete("/{id}", orderHandler.DeleteByID)
}

func (a *App) loadProductRoutes(router chi.Router) {
	// get product handler
	productHandler := &handler.Product{
		Repo: &repository.ProductRepo{
			DB: a.db,
		},
	}
	// attach routes
	router.Post("/", productHandler.Create)
	router.Get("/", productHandler.GetAll)
	router.Get("/{id}", productHandler.GetByID)
	router.Put("/{id}", productHandler.UpdateByID)
	router.Delete("/{id}", productHandler.DeleteByID)
}

func (a *App) LoadClientsRoutes(router chi.Router) {
	clientHandler := &handler.Client{
		Repo: &repository.ClientRepo{
			DB: a.db,
		},
	}
	router.Post("/", clientHandler.Create)
	router.Get("/", clientHandler.GetAll)
	router.Get("/{id}", clientHandler.GetByID)
	router.Put("/{id}", clientHandler.UpdateByID)
	router.Delete("/{id}", clientHandler.DeleteByID)
}
