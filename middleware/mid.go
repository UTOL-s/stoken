package sToken

import (
	"net/http"
	
	"github.com/ankorstore/yokai/config"
	"github.com/labstack/echo/v4"
	"github.com/supertokens/supertokens-golang/supertokens"
)

type Middleware struct {
	config *config.Config
}

func NewMiddleware(config *config.Config) *Middleware {
	return &Middleware{config}
}

func (mid *Middleware) Handle() echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			supertokens.Middleware(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
				if err := hf(c); err != nil {
					c.Error(err)
				}
			})).ServeHTTP(c.Response(), c.Request())
			return nil
		}
	}
}
