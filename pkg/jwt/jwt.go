package jwt

import (
	"fmt"
	"os"
	"path"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	privateKey string
	publicKey  string
}

type CustomClaims struct {
	ID  int `json:"id"`
	Sub int `json:"sub"`
	jwt.RegisteredClaims
}

func New() *JWT {
	pub, priv := getKeys()
	return &JWT{
		privateKey: priv,
		publicKey:  pub,
	}
}

func (j *JWT) VerifyToken(rawToken string) (*CustomClaims, error) {
	key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(j.publicKey))
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %w", err)
	}

	claims := &CustomClaims{}

	token, err := jwt.ParseWithClaims(rawToken, claims, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return key, nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("token is invalid")
	}

	return claims, nil
}

func getKeys() (string, string) {
	cwd, err := os.Getwd()
	if err != nil {
		panic(fmt.Errorf("failed to get current working directory: %w", err))
	}

	pubKey, err := os.ReadFile(path.Join(cwd, "keys", "public_key.pem"))
	if err != nil {
		panic(fmt.Errorf("failed to read public key: %w", err))
	}

	privKey, err := os.ReadFile(path.Join(cwd, "keys", "private_key.pem"))
	if err != nil {
		panic(fmt.Errorf("failed to read private key: %w", err))
	}

	return string(pubKey), string(privKey)
}
