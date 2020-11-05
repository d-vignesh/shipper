package main

import (
	pb "github.com/d-vignesh/shipper/user-service/proto/user"
	"github.com/dgrijalva/jwt-go"
)

var (
	// secure key string used as a salt when hashing our tokens.
	key = []byte("superSecretKey")
)

// CustomClaims is our custom metadata, which will be hashed
// and sent as the second segment in our JWT
type CustomClaims struct {
	User *pb.User
	jwt.StandardClaims
}

type TokenService struct {
	repo Repository
}

// Decode a token string into a token object
func (ts *TokenService) Decode(token string) (*CustomClaims, error) {

	// parse the token
	tokenType, err := jwt.ParseWithClaims(string(key), &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	// Validate the token and return the custom claims
	if claims, ok := tokenType.Claims.(*CustomClaims); ok && tokenType.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

// Encode a claim into JWT
func (ts *TokenService) Encode(user *pb.User) (string, error) {
	// create the claims
	claims := CustomClaims {
		user,
		jwt.StandardClaims {
			ExpiresAt: 15000,
			Issuer:    "shipper.user.service",
		},
	}
	
	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// sign token and return
	return token.SignedString(key)
}