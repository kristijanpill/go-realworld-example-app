package service

import (
	"crypto/rsa"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/kristijanpill/go-realworld-example-app/user_service/model"
)

type JWTManager struct {
	privateKey          *rsa.PrivateKey
	publicKey           *rsa.PublicKey
	accessTokenDuration time.Duration
}

func NewJWTManager(privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) *JWTManager {
	return &JWTManager{
		privateKey:          privateKey,
		publicKey:           publicKey,
		accessTokenDuration: 15 * time.Minute,
	}
}

func (manager *JWTManager) GenerateAccessToken(user *model.User) (string, error) {
	claims := jwt.StandardClaims{
			Subject:   user.ID.String(),
			ExpiresAt: time.Now().Add(manager.accessTokenDuration).Unix(),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodRS256,
		claims,
	)

	return token.SignedString(manager.privateKey)
}
