// Package middleware is http-middleware package
package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// CORSMaxAgeHour is CORS Lifetime(hour)
const CORSMaxAgeHour = 24

// SetCors is setting for HTTP CORS Policy
func SetCors(e *gin.Engine) gin.HandlerFunc {
	return cors.New(cors.Config{
		// 許可したいHTTPメソッドの一覧
		AllowMethods: []string{
			"GET",
			"HEAD",
			"PUT",
			"PATCH",
			"POST",
			"OPTIONS",
			"DELETE",
		},
		// 許可したいHTTPリクエストヘッダの一覧
		AllowHeaders: []string{
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			"Authorization",
		},
		// 許可したいアクセス元の一覧
		AllowOrigins: []string{
			"*",
		},
		MaxAge: CORSMaxAgeHour * time.Hour,
	})
}
