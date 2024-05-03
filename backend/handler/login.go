package handler

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/dqx0/GoHalves/go/usecase"
	"github.com/gin-gonic/gin"

	"golang.org/x/crypto/bcrypt"

	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

type ISessionHandler interface {
	Login() gin.HandlerFunc
	CheckSession(c *gin.Context)
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

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			return
		}
		password = string(hashedPassword)

		su := sc.bu.GetSessionUsecase()
		ok, err := su.Login(c, username, password)
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

		c.JSON(http.StatusOK, gin.H{"token": tokenString})
	}
}
func (sc *sessionHandler) CheckSession(c *gin.Context) {
	authorizationHeader := c.GetHeader("Authorization")
	if authorizationHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
		return
	}

	tokenString := strings.TrimPrefix(authorizationHeader, "Bearer ")
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
			return
		}
		c.Header("Authorization", "Bearer "+newToken)
		c.Set("userId", claims["userId"])
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
