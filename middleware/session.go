package sToken

import (
	"net/http"
	
	"github.com/ankorstore/yokai/config"
	"github.com/labstack/echo/v4"
	"github.com/supertokens/supertokens-golang/recipe/session"
)

type SessionMiddleware struct {
	config *config.Config
}

func NewSessionMiddleware(config *config.Config) *SessionMiddleware {
	return &SessionMiddleware{
		config: config,
	}
}

func (m SessionMiddleware) Handle() echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			session.VerifySession(nil, func(rw http.ResponseWriter, r *http.Request) {
				c.Set("session", session.GetSessionFromRequestContext(r.Context()))
				
				// Call the handler
				err := hf(c)
				if err != nil {
					c.Error(err)
				}
			})(c.Response(), c.Request())
			
			return nil
		}
	}
}
