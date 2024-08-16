package services

import (
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

var publicKey *rsa.PublicKey

func getPublicKey(kid string) (*rsa.PublicKey, error) {
	if publicKey != nil {
		return publicKey, nil
	}

	keysPath := fmt.Sprintf("https://%s/oauth/v2/keys", viper.GetString("AUTH_URL"))
	resp, err := http.Get(keysPath)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var jwks struct {
		Keys []struct {
			Kid string `json:"kid"`
			Use string `json:"use"`
			Kty string `json:"kty"`
			Alg string `json:"alg"`
			N   string `json:"n"`
			E   string `json:"e"`
		} `json:"keys"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&jwks); err != nil {
		return nil, err
	}

	for _, key := range jwks.Keys {
		if key.Kid == kid {
			nBytes, err := base64.RawURLEncoding.DecodeString(key.N)
			if err != nil {
				return nil, err
			}
			eBytes, err := base64.RawURLEncoding.DecodeString(key.E)
			if err != nil {
				return nil, err
			}
			e := big.NewInt(0).SetBytes(eBytes).Int64()

			pubKey := &rsa.PublicKey{
				N: big.NewInt(0).SetBytes(nBytes),
				E: int(e),
			}
			publicKey = pubKey
			return pubKey, nil
		}
	}
	return nil, fmt.Errorf("unable to find matching key")
}

func getKeyID(tokenString string) (string, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return "", err
	}
	if header, ok := token.Header["kid"].(string); ok {
		return header, nil
	}
	return "", fmt.Errorf("no kid found in token header")
}

func VerifyIDToken(idToken string) (jwt.MapClaims, error) {
	kid, err := getKeyID(idToken)
	if err != nil {
		return nil, err
	}

	publicKey, err := getPublicKey(kid)
	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(idToken, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}