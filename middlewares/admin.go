package middlewares

import (
	"net/http"

	"github.com/pushpaldev/base-api/helpers"
	"github.com/pushpaldev/base-api/store"
	"gopkg.in/gin-gonic/gin.v1"
)

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := store.Current(c)

		if !user.Admin {
			c.AbortWithError(http.StatusUnauthorized, helpers.ErrorWithCode("admin_required", "The user is not administrator"))
			return
		}
	}
}
