package auth

import (
	"fmt"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Auth interface {
	CreateToken(id string) (string, error)
	DeleteToken(token string) error
	VerifyToken(token string) (int, error)
}

type jwtAuth struct {
	secretKey []byte
}

func NewJWTAuth(secretKey string) Auth {
	return &jwtAuth{
		secretKey: []byte(secretKey),
	}
}

func (ja *jwtAuth) CreateToken(id string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(ja.secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (ja *jwtAuth) DeleteToken(token string) error {
	return nil
}

func (ja *jwtAuth) VerifyToken(token string) (int, error) {
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return ja.secretKey, nil
	})
	if err != nil {
		return 0, err
	}

	if claims, ok := jwtToken.Claims.(jwt.MapClaims); ok && jwtToken.Valid {
		userID, ok := claims["id"].(string)
		if !ok {
			return 0, fmt.Errorf("invalid user id in token")
		}

		id, err := strconv.Atoi(userID)
		if err != nil {
			return 0, fmt.Errorf("invalid user id in token")
		}

		return id, nil
	} else {
		return 0, fmt.Errorf("invalid token")
	}
}
