package infrastructure

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thn-lee/01-task-management-api/internal/handlers"
	tasksHandler "github.com/thn-lee/01-task-management-api/pkg/tasks/handler"
	tasksRepo "github.com/thn-lee/01-task-management-api/pkg/tasks/repository"
	tasksUsecase "github.com/thn-lee/01-task-management-api/pkg/tasks/usecase"
	// "github.com/thn-lee/01-task-management-api/pkg/tasks"
)

// SetupRoutes is the Router for GoFiber App
func (s *Server) SetupRoutes(app *fiber.App) {

	// Prepare a static middleware to serve the built React files.
	app.Static("/", "./web/build")

	// API routes group
	groupApiV1 := app.Group("/api/v:version?", handlers.ApiLimiter)
	{
		groupApiV1.Get("/", handlers.Index())
	}

	// App Repository
	taskRepository := tasksRepo.NewTaskRepository(s.MainDbConn)

	// // auto migrate DB only on main process
	// if !fiber.IsChild() {
	// }

	// App Services
	taskUsecase := tasksUsecase.NewTaskUsecase(taskRepository)

	// App Routes
	tasksHandler.NewTaskHandler(groupApiV1.Group("/tasks"), taskUsecase)

	// Prepare a fallback route to always serve the 'index.html', had there not be any matching routes.
	app.Static("*", "./web/build/index.html")
}
