package jwts

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	secretKey = os.Getenv("JWT_SECRET_KEY")
	tokenExpire = 3 * time.Minute // you may change this if you want
)

// GenerateJWTToken generates a new jwt token for user
func GenerateJWTToken(userID int64) (string, error) {

	now := time.Now()
	userIDStr := fmt.Sprintf("%d", userID)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userid": userIDStr,
		"iat": now.Unix(),
		"nbf": now.Unix(),
		"exp": now.Add(tokenExpire).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// VerifyJWTToken verifies jwt token (used in middleware)
func VerifyJWTToken(tokenString string) (string, error) {

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(secretKey), nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, ok := claims["userid"].(string)
		if !ok {
			return "", fmt.Errorf("invalid userID type")
		}
		return userID, nil
	} 

	return "", fmt.Errorf("invalid token")
}