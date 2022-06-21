package http

import (
	"github.com/labstack/echo/v4"

	fc "dapr-workshop-go/fine-collection-service/internal/fine_collection"
)

func MapRoutes(commGroup *echo.Group, h fc.Handlers) {
	commGroup.POST("collectfine", h.CollectFine())
}
