package exploreapi

import (
	"blinders/packages/auth"
	"blinders/packages/db"

	"github.com/gofiber/fiber/v2"
)

type Manager struct {
	App     *fiber.App
	Auth    auth.Manager
	DB      *db.MongoManager
	Service *Service
}

func NewManager(app *fiber.App, auth auth.Manager, db *db.MongoManager, service *Service) *Manager {
	return &Manager{
		App:     app,
		Auth:    auth,
		DB:      db,
		Service: service,
	}
}

func (m *Manager) InitRoute() {
	m.App.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("service healthy")
	})

	// Temporarily expose this method, it must be call internal, or we will listen to user update-match-information event
	m.App.Post("/explore", m.Service.HandleAddUserMatch)

	matchRoute := m.App.Group("/explore", auth.FiberAuthMiddleware(m.Auth, m.DB.Users))
	matchRoute.Get("/suggest", m.Service.HandleGetMatches)
}
