package http

import (
	fc "dapr-workshop-go/fine-collection-service/internal/fine_collection"

	"github.com/labstack/echo/v4"
)

type Subscription struct {
	PubsubName string `json:"pubsubName"`
	Topic      string `json:"topic"`
	Route      string `json:"route"`
}

func MapRoutes(commGroup *echo.Group, h fc.Handlers) {
	commGroup.POST("collectfine", h.CollectFine())

	// commGroup.GET("dapr/subcribe", func(c echo.Context) error {
	// 	subscription := &Subscription{
	// 		PubsubName: "pubsub",
	// 		Topic:      "speedingviolations",
	// 		Route:      "/collectfine",
	// 	}

	// 	return c.JSONPretty(http.StatusOK, subscription, "")
	// })
}
