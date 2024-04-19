package collectingapi

import (
	"blinders/packages/collecting"

	"github.com/gofiber/fiber/v2"
)

type Service struct {
	App       *fiber.App
	Collector *collecting.EventCollector
}

func NewCollectingService(app *fiber.App, collector *collecting.EventCollector) *Service {
	return &Service{
		App:       app,
		Collector: collector,
	}
}

func (s *Service) InitRoute() error {
	s.App.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("service healthy")
	})

	s.App.Post("/", s.HandlePushEvent)
	s.App.Get("/", s.HandleGetEvent)
	return nil
}