package pechox

import (
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/labstack/echo/v4"
)

func GctxMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			new_ctx := gctx.WithCtx(c.Request().Context())
			new_req := c.Request().WithContext(new_ctx)
			new_req.ParseForm()
			c.SetRequest(new_req)
			c.Response().Header().Add("TraceId", gctx.CtxId(new_ctx))
			return next(c)
		}
	}
}
