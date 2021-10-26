package finecollection

import "github.com/labstack/echo/v4"

type Handlers interface {
	CollectFine() echo.HandlerFunc
}
