package handler

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/dqx0/GoHalves/go/usecase"
	"github.com/gin-gonic/gin"

	"fmt"
	"net/http"
	"os"
	"time"
)

type ISessionHandler interface {
	Login() gin.HandlerFunc
	Logout() gin.HandlerFunc
	CheckSession(c *gin.Context)
	IsLoggedIn(c *gin.Context) bool
}
type sessionHandler struct {
	bu usecase.IBaseUsecase
}

func NewSessionHandler(bu usecase.IBaseUsecase) ISessionHandler {
	return &sessionHandler{bu}
}
func (sc *sessionHandler) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.PostForm("user_id")
		password := c.PostForm("password")

		su := sc.bu.GetSessionUsecase()
		ok, err := su.Login(username, password)
		if !ok || err != nil {
			c.Redirect(http.StatusFound, "/account")
			return
		}

		secretKey := os.Getenv("JWT_SECRET_KEY")
		if secretKey == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Secret key not configured"})
			return
		}
		expirationTime := time.Now().Add(time.Hour)
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"userId": username,
			"exp":    expirationTime.Unix(),
		})

		tokenString, err := token.SignedString([]byte(secretKey))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
			return
		}

		c.Header("Set-Cookie", "jwtToken="+tokenString+"; Path='/'; Domain=localhost; Max-Age=3600;")
		c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
	}
}
func (sc *sessionHandler) Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.SetCookie(
			"jwtToken",
			"",
			-1,
			"/",
			"localhost",
			false,
			true,
		)
		c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
	}
}
func (sc *sessionHandler) CheckSession(c *gin.Context) {
	tokenString, err := c.Cookie("jwtToken")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization cookie is required"})
		return
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// トークンの有効期限をリセット
		newToken, err := sc.resetTokenExpiration(claims)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not regenerate token"})
			c.SetCookie(
				"jwtToken",
				"",
				-1,
				"/",
				"localhost",
				false,
				true,
			)
			return
		}
		account, _ := sc.bu.GetAccountUsecase().GetAccountByUserId(claims["userId"].(string))
		c.SetCookie("jwtToken", newToken, 3600, "/", "localhost", false, false)
		c.Set("userId", account.ID)
	}
}

func (sc *sessionHandler) resetTokenExpiration(claims jwt.MapClaims) (string, error) {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		return "", fmt.Errorf("secret key not configured")
	}

	expirationTime := time.Now().Add(time.Hour) // 新しい有効期限
	claims["exp"] = expirationTime.Unix()

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := newToken.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
func (sc *sessionHandler) IsLoggedIn(c *gin.Context) bool {
	tokenString, err := c.Cookie("jwtToken")
	if err != nil {
		return false
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		return false
	}
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true
	}
	return false
}
