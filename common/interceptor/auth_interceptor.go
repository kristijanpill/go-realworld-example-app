package interceptor

import (
	"context"
	"crypto/rsa"
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type AuthInterceptor struct {
	restrictedPaths map[string]bool
	publicKey       *rsa.PublicKey
}

func NewAuthInterceptor(restrictedPaths map[string]bool, publicKey *rsa.PublicKey) *AuthInterceptor {
	return &AuthInterceptor{
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
	_, ok := interceptor.restrictedPaths[method]
	if !ok {
		return ctx, nil
	}

	var values []string
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ctx, status.Errorf(codes.Unauthenticated, "unauthorized")
	}

	values = md.Get("Authorization")
	if len(values) == 0 {
		return ctx, status.Errorf(codes.Unauthenticated, "unauthorized")
	}

	authHeader := values[0]
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 {
		return ctx, status.Errorf(codes.Unauthenticated, "unauthorized")
	}

	token := parts[1]
	claims, err := interceptor.verifyToken(token)
	if err != nil {
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