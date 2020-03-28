package app

import (
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type ClaimsJWT struct {
	jwt.StandardClaims
	Username string `json:"username"`
	Name     string `json:"name"`
	Level    string `json:"level"`
}

func CheckAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("Authorization")

		if !strings.Contains(authorizationHeader, "Bearer") {
			c.JSON(500, INVALID_ATHORIZATION)
			c.Abort()
			return
		}

		tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

		_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

			if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				c.JSON(500, ResultData{
					Code:    999,
					Success: false,
					Data:    "Signing method invalid",
				})
				c.Abort()
				return nil, nil
			} else if method != JWT_SIGNING_METHOD {
				c.JSON(500, ResultData{
					Code:    999,
					Success: false,
					Data:    "Signing method invalid",
				})
				c.Abort()
				return nil, nil
			}

			return JWT_SIGNATURE_KEY, nil
		})

		if err != nil {
			c.JSON(500, ResultData{
				Code:    999,
				Success: false,
				Data:    err.Error(),
			})
			c.Abort()
		}

		c.Next()

		// _, ok := token.Claims.(jwt.MapClaims)
		// if !ok || !token.Valid {
		// 	c.JSON(500, ResultData{
		// 		Code:    999,
		// 		Success: false,
		// 		Data:    err.Error(),
		// 	})
		// }

	}
}
