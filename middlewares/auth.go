package middlewares

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/pushpaldev/base-api/helpers"
	"github.com/pushpaldev/base-api/models"
	"github.com/pushpaldev/base-api/services"
	"github.com/pushpaldev/base-api/store"
	"gopkg.in/gin-gonic/gin.v1"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenReader := c.Request.Header.Get("Authorization")

		authHeaderParts := strings.Split(tokenReader, " ")
		if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
			c.AbortWithError(http.StatusBadRequest, errors.New("Authorization header format must be Bearer {token}"))
			return
		}

		publicKeyFile, _ := ioutil.ReadFile(os.Getenv("BASEAPI_RSA_PUBLIC"))
		publicKey, _ := jwt.ParseRSAPublicKeyFromPEM(publicKeyFile)

		token, err := jwt.Parse(authHeaderParts[1], func(token *jwt.Token) (interface{}, error) {
			return publicKey, nil
		})

		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, errors.New("Error parsing token"))
			return
		}

		if token.Header["alg"] != jwt.SigningMethodRS256.Alg() {
			c.AbortWithError(http.StatusUnauthorized, errors.New("Signing method not valid"))
			return
		}

		if !token.Valid {
			c.AbortWithError(http.StatusUnauthorized, errors.New("Token invalid"))
			return
		}

		claims, _ := token.Claims.(jwt.MapClaims)

		user := &models.User{}

		// Gets the user from the redis store
		err = services.GetRedis(c).GetValueForKey(claims["id"].(string), &user)
		if err != nil {
			user, _ = store.FindUserById(c, claims["id"].(string))
			services.GetRedis(c).SetValueForKey(user.Id, &user)
		}

		// Check if the token is still valid in the database
		if !user.HasToken(claims["token"].(string)) {
			c.AbortWithError(http.StatusUnauthorized, helpers.ErrorWithCode("token_invalidated", "This token isn't valid anymore"))
			return
		}

		c.Set("currentUser", user)
	}
}
