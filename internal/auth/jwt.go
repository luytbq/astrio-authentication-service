package auth

import (
	"crypto/sha256"
	"errors"
	"log"
	"math/rand"
	"strings"
	"time"

	"encoding/base64"

	"github.com/golang-jwt/jwt/v5"
	"github.com/luytbq/astrio-authentication-service/config"
)

func HashPassword(password string) (hashedPassword, salt string) {
	salt = randStringRunes(64)
	hashedPassword = hashWithSalt(password, salt)
	return
}

func hashWithSalt(password, salt string) string {
	hash := sha256.New()

	hashInBytes := hash.Sum([]byte(password + salt))

	return base64.URLEncoding.EncodeToString(hashInBytes)
}

func VerifyPassword(password, hashedPassword, salt string) bool {
	hashedPassword2 := hashWithSalt(password, salt)
	return strings.Compare(hashedPassword, hashedPassword2) == 0
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func CreateJWTToken(claims map[string]any) (string, error) {
	key := config.App.HMAC_KEY

	mapClaims := jwt.MapClaims(claims)
	mapClaims["iat"] = jwt.NumericDate{Time: time.Now()}
	mapClaims["exp"] = jwt.NumericDate{Time: time.Now().Add(time.Hour * 12)}

	v, _ := mapClaims.GetExpirationTime()
	log.Printf("exp: %+v", v)
	log.Printf("exp: %+v", mapClaims["exp"])

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)
	str, err := token.SignedString([]byte(key))

	if err != nil {
		return "", err
	}
	return str, nil
}

func ParseJWTToken(tokenString string) (jwt.MapClaims, error) {
	key := []byte(config.App.HMAC_KEY)

	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return key, nil
	})

	if err != nil {
		return nil, err
	}

	// Check if the token is valid and extract claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
