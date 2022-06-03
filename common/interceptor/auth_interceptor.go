package interceptor

import (
	"context"
	"crypto/rsa"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthInterceptor struct {
	tokenPrefix string
	restrictedPaths map[string]bool
	publicKey       *rsa.PublicKey
}

func NewAuthInterceptor(tokenPrefix string, restrictedPaths map[string]bool, publicKey *rsa.PublicKey) *AuthInterceptor {
	return &AuthInterceptor{
		tokenPrefix: tokenPrefix,
		restrictedPaths: restrictedPaths,
		publicKey:       publicKey,
	}
}

func (interceptor *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		ctx, err := interceptor.authorize(ctx, info.FullMethod)
		if err != nil {
			return nil, err
		}

		return handler(ctx, req)
	}
}

func (interceptor *AuthInterceptor) authorize(ctx context.Context, method string) (context.Context, error) {
	isAuthRequired, ok := interceptor.restrictedPaths[method]
	if !ok {
		return ctx, nil
	}

	token, err := grpc_auth.AuthFromMD(ctx, interceptor.tokenPrefix)
	if err != nil {
		if isAuthRequired {
			return nil, err
		} else {
			return ctx, nil
		}
	}

	claims, err := interceptor.verifyToken(token)
	if err != nil && isAuthRequired {
		return ctx, status.Errorf(codes.Unauthenticated, "unauthorized")
	}

	ctx = context.WithValue(ctx, TokenKey{}, token)
	ctx = context.WithValue(ctx, CurrentUserKey{}, claims.Subject)

	return ctx, nil
}

func (interceptor *AuthInterceptor) verifyToken(accessToken string) (*jwt.StandardClaims, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&jwt.StandardClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodRSA)
			if !ok {
				return nil, fmt.Errorf("unexpected token signing method")
			}

			return interceptor.publicKey, nil
		},
	)

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}

type CurrentUserKey struct{}
type TokenKey struct{}