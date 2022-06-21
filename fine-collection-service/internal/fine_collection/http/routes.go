package http

import (
  "github.com/labstack/echo/v4"

  fc "dapr-workshop-go/fine-collection-service/internal/fine_collection"
  "net/http"
)

type Subscription struct {
  Pubsubname string `json:"pubsubname"`
  Topic      string `json:"topic"`
  Route      string `json:"route"`
}

func MapRoutes(commGroup *echo.Group, h fc.Handlers) {
  commGroup.POST("collectfine", h.CollectFine())

  commGroup.GET("dapr/subcribe", func(c echo.Context) error {
    subscription := &Subscription{
      Pubsubname: "pubsub",
      Topic:      "speedingviolations",
      Route:      "/collectfine",
    }

    return c.JSONPretty(http.StatusOK, subscription, "")
  })
}
