package repo

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/yalagtyarzh/aggregator/pkg/config"
	"github.com/yalagtyarzh/aggregator/pkg/models"
	"time"
)

var (
	ErrInvalidTokenSignature = errors.New("invalid token signature")
	ErrInvalidToken          = errors.New("invalid token")
	ErrValidate              = errors.New("validate token error")
)

type JWTer struct {
	db  IDB
	cfg config.JWTConfig
}

func NewJWTer(db IDB, cfg config.JWTConfig) *JWTer {
	return &JWTer{
		db:  db,
		cfg: cfg,
	}
}

func (l *JWTer) GenerateTokens(userId uuid.UUID, email string) (string, string, error) {
	accessPayload := models.TokenPayload{
		UserID: userId,
		Email:  email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(30 * time.Minute).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	refreshPayload := models.TokenPayload{
		UserID: userId,
		Email:  email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(30 * 24 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessPayload)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshPayload)

	aTokenStr, err := accessToken.SignedString([]byte(l.cfg.SigningKey))
	if err != nil {
		return "", "", err
	}

	rTokenStr, err := refreshToken.SignedString([]byte(l.cfg.SigningKey))
	if err != nil {
		return "", "", err
	}

	return aTokenStr, rTokenStr, nil
}

func (l *JWTer) SaveToken(userId uuid.UUID, refreshToken string) error {
	t, err := l.db.GetToken(userId)
	if err != nil {
		return err
	}

	if t != nil {
		return l.db.UpdateToken(userId, refreshToken)
	}

	return l.db.InsertToken(userId, refreshToken)
}

func (l *JWTer) ValidateAccessToken(token string) (models.TokenPayload, error) {
	return l.ValidateToken(token, l.cfg.AccessSecret)
}

func (l *JWTer) ValidateRefreshToken(token string) (models.TokenPayload, error) {
	return l.ValidateToken(token, l.cfg.RefreshSecret)
}

func (l *JWTer) ValidateToken(token, signingKey string) (models.TokenPayload, error) {
	claims := models.TokenPayload{}

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return models.TokenPayload{}, ErrInvalidTokenSignature
		}
		return models.TokenPayload{}, ErrValidate

	}

	if !tkn.Valid {
		return models.TokenPayload{}, ErrInvalidToken
	}

	return claims, nil
}
