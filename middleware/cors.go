package middleware

import (
	"strings"
	
	"github.com/ankorstore/yokai/config"
	"github.com/labstack/echo/v4"
	"github.com/supertokens/supertokens-golang/supertokens"
)

type CorsMiddleware struct {
	config *config.Config
}

func NewCorsMiddleware(config *config.Config) *CorsMiddleware {
	return &CorsMiddleware{config}
}

func (m *CorsMiddleware) Handle() echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			allowedOrigins := []string{
				"http://localhost:3000",
				"http://localhost:8000",
				"https://www.utol.com",
				"https://www.utol.com.ph",
				"https://portal-admin-v2.utol.com.ph",
				"https://staging-landing-page.utol.com.ph",
				"https://admin-staging-portal.utol.ph",
				"https://staging-admin-v2.utol.com.ph",
				"https://admin-portal.utol.com.ph",
				"https://staging-admin-portal.utol.com.ph",
				"https://staging-landing-v2.utol.com.ph",
				"https://accounting-admin.utol.com.ph",
				"https://staging-accounting-admin.utol.com.ph",
			}
			origin := c.Request().Header.Get("Origin")
			
			isAllowedOrigin := false
			for _, allowedOrigin := range allowedOrigins {
				if origin == allowedOrigin {
					isAllowedOrigin = true
					break
				}
			}
			
			if isAllowedOrigin {
				c.Response().Header().Set("Access-Control-Allow-Origin", origin)
				c.Response().Header().Set("Access-Control-Allow-Credentials", "true")
			}
			
			if c.Request().Method == "OPTIONS" {
				c.Response().Header().Set("Access-Control-Allow-Headers", strings.Join(append([]string{"Content-Type"}, supertokens.GetAllCORSHeaders()...), ","))
				c.Response().Header().Set("Access-Control-Allow-Methods", "*")
				_, err := c.Response().Write([]byte(""))
				if err != nil {
					return err
				}
				return nil
			} else {
				return hf(c)
			}
		}
	}
}
