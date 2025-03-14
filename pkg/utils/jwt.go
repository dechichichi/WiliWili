package utils

import (
	"fmt"
	"time"
	"wiliwili/config"
	"wiliwili/pkg/constants"
	"wiliwili/pkg/errno"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	Type   int64 `json:"type"`
	UserID int64 `json:"user_id"`
	jwt.StandardClaims
}

// 根据tokentype的不同返回过期时间不同的token
func CreateToken(tokenType int64, uid int64) (string, error) {
	if config.Server == nil {
		return "", errno.Errorf(errno.ErrConfigParseCode, "config.Server is nil")
	}
	expireTime := time.Now().Add(getTokenTTL(tokenType))
	claims := Claims{
		Type:   tokenType,
		UserID: uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    config.Server.Name,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	key, err := PraisePrivateKey(config.Server.Secret)
	if err != nil {
		return "", errno.Errorf(errno.ErrTokenGenCode, "failed to generate token: %w", err)
	}
	signToken, err := token.SignedString(key)
	if err != nil {
		return "", errno.Errorf(errno.ErrTokenGenCode, "failed to generate token: %w", err)
	}
	return signToken, nil
}

// 创建accessToken和refreshToken
func CreateAllToken(uid int64) (string, string, error) {
	accessToken, err := CreateToken(constants.TypeAccessToken, uid)
	if err != nil {
		return "", "", errno.Errorf(errno.ErrTokenGenCode, "failed to generate token: %w", err)
	}
	refreshToken, err := CreateToken(constants.TypeRefreshToken, uid)
	if err != nil {
		return "", "", errno.Errorf(errno.ErrTokenGenCode, "failed to generate token: %w", err)
	}
	return accessToken, refreshToken, nil
}

// 检查token是否有效,返回token的类型和用户id
func CheckToken(token string) (int64, int64, error) {
	if config.Server == nil {
		return 0, 0, errno.Errorf(errno.ErrConfigParseCode, "config.Server is nil")
	}
	if token == "" {
		return 0, 0, errno.Errorf(errno.ErrTokenInvalidCode, "token is empty")
	}
	unverifiedClaims, err := parseUnverifiedClaims(token)
	if err != nil {
		return 0, 0, errno.Errorf(errno.ErrTokenInvalidCode, "failed to parse token: %w", err)
	}
	secret, err := PraisePublicKey(config.Server.PublicKey)
	if err != nil {
		return 0, 0, errno.Errorf(errno.ErrTokenInvalidCode, "failed to parse token: %w", err)
	}
	verifiedClaims, err := VerifyToken(token, secret)
	if err != nil {
		return unverifiedClaims.Type, 0, errno.Errorf(errno.ErrTokenVerifyCode, "failed to verify token: %w", err)
	}
	return verifiedClaims.Type, verifiedClaims.UserID, nil

}

// 验证token并且返回claims
func VerifyToken(token string, key interface{}) (*Claims, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})
	if err != nil {
		return nil, errno.Errorf(errno.ErrTokenInvalidCode, "failed to parse token: %w", err)
	}
	if claims, ok := parsedToken.Claims.(*Claims); ok && parsedToken.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}

// Create时传入的key为私钥，Verify时传入的key为公钥
func PraisePrivateKey(key string) (interface{}, error) {
	privatekey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(key))
	if err != nil {
		return nil, err
	}
	return privatekey, nil
}

func PraisePublicKey(key string) (interface{}, error) {
	publickey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(key))
	if err != nil {
		return nil, err
	}
	return publickey, nil
}

// getTokenTTL 根据 token 类型返回过期时间
func getTokenTTL(tokenType int64) time.Duration {
	switch tokenType {
	case constants.TypeAccessToken:
		return constants.AccessTokenTTL
	case constants.TypeRefreshToken:
		return constants.RefreshTokenTTL
	default:
		return 0
	}
}

func parseUnverifiedClaims(token string) (*Claims, error) {
	tokenStruct, _, err := new(jwt.Parser).ParseUnverified(token, &Claims{})
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}
	claims, ok := tokenStruct.Claims.(*Claims)
	if !ok {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}
	return claims, nil
}
