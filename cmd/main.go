package main

import (
	"os"

	"github.com/mahdifr17/sirka-test/handler"
	"github.com/mahdifr17/sirka-test/repository"
	"github.com/mahdifr17/sirka-test/usecase"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	server := newServer()

	app.Post("/DisplayUser", server.GetUser)
	app.Get("/DisplayAllUser", server.GetAllUser)

	app.Listen(":3000")
}

func newServer() *handler.Server {
	dbDsn := os.Getenv("DATABASE_URL")
	// dbDsn := "postgres://postgres:cicilaja@localhost:5432/test_case?sslmode=disable"

	var repo repository.RepositoryInterface = repository.NewRepository(repository.NewRepositoryOptions{
		Dsn: dbDsn,
	})
	usecase := usecase.NewUseCase(repo)
	return handler.NewServer(usecase)
}
