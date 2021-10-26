package http

import (
	fc "dapr-workshop-go/fine-collection-service/internal/fine_collection"

	"github.com/labstack/echo/v4"
)

func MapRoutes(commGroup *echo.Group, h fc.Handlers) {
	commGroup.POST("collectfine", h.CollectFine())
}
