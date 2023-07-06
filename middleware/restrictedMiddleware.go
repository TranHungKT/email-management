package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/TranHungKT/email_management/models"
	"github.com/adam-hanna/jwt-auth/jwt"
	"github.com/gin-gonic/gin"
)

var restrictedRoute jwt.Auth
var HMACKey []byte

func InitRestrictedRoute() {
	authErr := jwt.New(&restrictedRoute, jwt.Options{
		SigningMethodString:   "HS256",
		RefreshTokenValidTime: 72 * time.Hour,
		AuthTokenValidTime:    15 * time.Hour,
		Debug:                 false,
		IsDevEnv:              true,
		HMACKey:               []byte("My super secret key!"),
	})
	if authErr != nil {
		log.Println("Error initializing the JWT's!")
		log.Fatal(authErr)
	}
}

func RestrictedFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := restrictedRoute.Process(c.Writer, c.Request)

		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.Next()
	}
}

func HandleToken(ctx *gin.Context, user models.User) {
	claims := jwt.ClaimsType{}
	claims.CustomClaims = make(map[string]interface{})
	claims.CustomClaims["email"] = user.Email
	claims.CustomClaims["firstName"] = user.FirstName
	claims.CustomClaims["lastName"] = user.LastName

	err := restrictedRoute.IssueNewTokens(ctx.Writer, &claims)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Failed to create token")
		return
	}
}
