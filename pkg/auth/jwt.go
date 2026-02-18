package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type JWTClaims struct {
	Subject   string `json:"sub"`
	IssuedAt  int64  `json:"iat"`
	ExpiresAt int64  `json:"exp"`
}

var (
	errInvalidToken     = errors.New("invalid token")
	errExpiredToken     = errors.New("token expired")
	errInvalidSignature = errors.New("invalid token signature")
)

func GetEnv(key, fallback string) string {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return fallback
	}
	return value
}

func JWTSecret() string {
	return GetEnv("JWT_SECRET", "change-me-in-production")
}

func GenerateToken(subject, secret string, ttl time.Duration) (string, time.Time, error) {
	now := time.Now().UTC()
	expiresAt := now.Add(ttl)

	header := map[string]string{"alg": "HS256", "typ": "JWT"}
	claims := JWTClaims{Subject: subject, IssuedAt: now.Unix(), ExpiresAt: expiresAt.Unix()}

	headerBytes, err := json.Marshal(header)
	if err != nil {
		return "", time.Time{}, err
	}
	claimsBytes, err := json.Marshal(claims)
	if err != nil {
		return "", time.Time{}, err
	}

	headerSegment := base64.RawURLEncoding.EncodeToString(headerBytes)
	claimsSegment := base64.RawURLEncoding.EncodeToString(claimsBytes)
	unsigned := fmt.Sprintf("%s.%s", headerSegment, claimsSegment)
	signature := sign(unsigned, secret)

	return fmt.Sprintf("%s.%s", unsigned, signature), expiresAt, nil
}

func ParseToken(tokenString, secret string) (JWTClaims, error) {
	parts := strings.Split(tokenString, ".")
	if len(parts) != 3 {
		return JWTClaims{}, errInvalidToken
	}

	unsigned := parts[0] + "." + parts[1]
	expectedSignature := sign(unsigned, secret)
	if !hmac.Equal([]byte(parts[2]), []byte(expectedSignature)) {
		return JWTClaims{}, errInvalidSignature
	}

	payloadBytes, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return JWTClaims{}, errInvalidToken
	}

	var claims JWTClaims
	if err := json.Unmarshal(payloadBytes, &claims); err != nil {
		return JWTClaims{}, errInvalidToken
	}
	if claims.ExpiresAt <= time.Now().UTC().Unix() {
		return JWTClaims{}, errExpiredToken
	}
	if claims.Subject == "" {
		return JWTClaims{}, errInvalidToken
	}

	return claims, nil
}

func sign(unsignedToken, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(unsignedToken))
	return base64.RawURLEncoding.EncodeToString(h.Sum(nil))
}
