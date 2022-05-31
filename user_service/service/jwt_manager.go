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

func NewJWTManager(privateKey, publicKey string) (*JWTManager, error) {
	parsedPrivateKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKey))
	if err != nil {
		return nil, err
	}
	parsedPublicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKey))
	if err != nil {
		return nil, err
	}

	return &JWTManager{
		privateKey:          parsedPrivateKey,
		publicKey:           parsedPublicKey,
		accessTokenDuration: 15 * time.Minute,
	}, nil
}

func (manager *JWTManager) GenerateAccessToken(user *model.User) (string, error) {
	claims := jwt.StandardClaims{
			Subject:   user.Email,
			ExpiresAt: time.Now().Add(manager.accessTokenDuration).Unix(),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodRS256,
		claims,
	)

	return token.SignedString(manager.privateKey)
}
