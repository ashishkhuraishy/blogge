package services

import (
	"errors"
	"os"
	"time"

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
		userID,
		isAdmin,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 365).Unix(),
			Issuer:    s.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	tokenGen := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenGen.SignedString([]byte(s.secretKey))
	if err != nil {
		panic(err)
	}

	return token
}

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, errors.New("invalid token")
		}

		return []byte(s.secretKey), nil
	})
}
