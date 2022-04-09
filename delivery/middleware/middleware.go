package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"strings"
	"time"
)

var ApplicationName = "Dating App"
var JwtSigningMethod = jwt.SigningMethodHS256
var JwtSignatureKey = []byte("k3yJWtee")

type MyClaim struct {
	jwt.StandardClaims
	Username string     `json:"username"`
	JoinDate *time.Time `json:"join_date"`
}

type authHeader struct {
	Authorization string `header:"Authorization"`
}

func AuthTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/dating/login" || c.Request.URL.Path == "/dating/api/verification" {
			c.Next()
		} else {
			h := authHeader{}
			if err := c.ShouldBindHeader(&h); err != nil {
				c.JSON(401, gin.H{
					"message": "Unauthorized",
				})
				c.Abort()
				return
			}
			tokenString := strings.Replace(h.Authorization, "Bearer ", "", -1)
			fmt.Println(tokenString)
			if tokenString == "" {
				c.JSON(401, gin.H{
					"message": "Unauthorized",
				})
				c.Abort()
				return
			}
			token, err := ParseToken(tokenString)
			if err != nil {
				c.JSON(401, gin.H{
					"message": "Unauthorized",
				})
				c.Abort()
				return
			}
			fmt.Println(token)
			if token["iss"] == ApplicationName {
				c.Next()
			} else {
				c.JSON(401, gin.H{
					"message": "Unauthorized",
				})
				c.Abort()
				return
			}
		}
	}
}

func ParseToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Signing method invalid")
		} else if method != JwtSigningMethod {
			return nil, fmt.Errorf("Signing method invalid")
		}

		return JwtSignatureKey, nil
	})
	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, err
	}
	return claim, nil
}

func GenerateToken(UserName string, JoinDate *time.Time) (string, error) {
	claims := MyClaim{
		StandardClaims: jwt.StandardClaims{
			Issuer: ApplicationName,
		},
		Username: UserName,
		JoinDate: JoinDate,
	}
	token := jwt.NewWithClaims(
		JwtSigningMethod,
		claims,
	)
	return token.SignedString(JwtSignatureKey)
}
