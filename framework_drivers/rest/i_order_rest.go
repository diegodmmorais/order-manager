package framework

import "github.com/labstack/echo"

type IOrderRest interface {
	Save(c echo.Context) (err error)
}
