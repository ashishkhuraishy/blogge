package middleware

import (
	"os"

	"github.com/dgrijalva/jwt-go"
)

// JWTService will generate and validate a token to authenticate
// a user before processing request
type JWTService interface {
	GenerateToken(int64, bool) string
	ValidateToken(string) (*jwt.Token, error)
}

type authClaims struct {
	UserID  int64 `json:"user_id"`
	IsAdmin bool  `json:"is_admin"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

// JWTAuthService is used to generate a public constructor for
// creating an instance of the service
func JWTAuthService() JWTService {
	return &jwtService{
		secretKey: getSecretKey(),
		issuer:    "Ashish Khuraishy",
	}
}

func getSecretKey() string {
	sercretKey := os.Getenv("JWT_SECRET_KEY")
	if sercretKey == "" {
		sercretKey = "SECRET_KEY"
	}

	return sercretKey
}

func (s *jwtService) GenerateToken(userID int64, isAdmin bool) string {
	claims := &authClaims{
		UserID:  userID,
		IsAdmin: isAdmin,
	}
	return ""
}

func (s *jwtService) ValidateToken(string) (*jwt.Token, error) {
	return nil, nil
}
