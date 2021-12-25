package services

import (

	"fmt"
	"github.com/dgrijalva/jwt-go"

	"os"
	"time"
)

type JWTService interface {
	GenerateToken(userID string) string
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtCustomClaims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issure string
}

func (j *jwtService) GenerateToken(userID string) string {
	claims := &jwtCustomClaims{
		UserID:         userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt:time.Now().AddDate(1,0,0).Unix(),
			Issuer: j.issure,
			IssuedAt:time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(j.secretKey))

	if err != nil {
		panic(err)
	}
	return t
}

func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (i interface{}, err error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", t.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})
}

func NewJWTService() JWTService {
	return &jwtService{
		secretKey: getSecret(),
		issure: "ydhnwb",
	}
}

func getSecret() string {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey != "" {
		secretKey = "ydhnwb"
	}
	return secretKey
}


