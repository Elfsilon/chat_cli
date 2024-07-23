package utils

import (
	"crypto"
	"errors"
	"fmt"
	"io"
	"math"
	"math/rand"
	"server/pkg/utils/role"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UserID int `json:"uid"`
	Role   int `json:"role"`
	jwt.StandardClaims
}

var errorClaims = Claims{
	UserID: -1,
	Role:   -1,
}

var (
	ErrInvalidClaims = errors.New("invalid jwt claims: unable to assert to Claims type")
	ErrUserIdMissing = errors.New("userID claim is missing")
	ErrRoleMissing   = errors.New("role claim is missing")
)

func errUnexpectedSigningMethod(alg interface{}) error {
	return fmt.Errorf("unexpected signing method: %v", alg)
}

type TokenManager struct {
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
	Secret          []byte
}

func NewTokenManager(accessTokenTTL, refreshTokenTTL time.Duration, secret []byte) *TokenManager {
	return &TokenManager{
		AccessTokenTTL:  accessTokenTTL,
		RefreshTokenTTL: refreshTokenTTL,
		Secret:          secret,
	}
}

// Token is invalid if there is an error
func (s *TokenManager) GetAccessTokenClaims(tokenString string) (Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errUnexpectedSigningMethod(token.Header["alg"])
		}
		return s.Secret, nil
	})

	if err != nil {
		return errorClaims, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return errorClaims, ErrInvalidClaims
	}

	return *claims, nil
}

func (s *TokenManager) GenerateAccessToken(userID int, role role.Role) (string, error) {
	issuedAt := time.Now().UTC()
	expiresAt := issuedAt.Add(s.AccessTokenTTL)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		UserID: userID,
		Role:   int(role),
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  issuedAt.Unix(),
			ExpiresAt: expiresAt.Unix(),
		},
	})

	return token.SignedString(s.Secret)
}

func (s *TokenManager) GenerateServiceToken() (string, error) {
	issuedAt := time.Now().UTC()
	expiresAt := issuedAt.Add(s.AccessTokenTTL)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  issuedAt.Unix(),
			ExpiresAt: expiresAt.Unix(),
		},
	})

	return token.SignedString(s.Secret)
}

func (s *TokenManager) GenerateRefreshToken() (string, time.Time) {
	dt := time.Duration(rand.Int63n(math.MaxInt))
	t := time.Now().Add(dt * time.Millisecond)
	strTime := strconv.FormatInt(t.Unix(), 10)

	hash := crypto.MD5.New()
	io.WriteString(hash, strTime)
	hstr := fmt.Sprintf("%x", hash.Sum(nil))

	expiresAt := time.Now().Add(s.RefreshTokenTTL)
	return hstr, expiresAt
}
