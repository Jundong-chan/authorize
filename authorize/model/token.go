package model

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

//token 信息的预配置
type TokenConfig struct {
	//Method    string //加密算法
	Audience  string //接受对象
	ExpiresAt int64  //过期时间（时间戳）
	IssuedAt  int64  //签发时间
	Issuer    string //签发者
	NotBefore int64  //生效时间
}

//签名的paylod 内容
type Claims struct {
	UserId      string
	UserName    string
	Authorities []string
	Standclaim  jwt.StandardClaims
}

func (claims Claims) Valid() error {
	return claims.Standclaim.Valid()
}
func InitTokenConfig() (tokencof *TokenConfig) {
	var tokenconfig TokenConfig
	tokenconfig.Audience = "users"
	tokenconfig.ExpiresAt = time.Now().Add(time.Hour * 2).Unix()
	tokenconfig.IssuedAt = time.Now().Unix()
	tokenconfig.Issuer = "seckill-shop"
	tokenconfig.NotBefore = time.Now().Unix()
	return &tokenconfig
}

func NewClaims(userdetail *UserDetail, tokenconf *TokenConfig) *Claims {
	claims := Claims{
		UserId:      userdetail.UserId,
		UserName:    userdetail.UserName,
		Authorities: userdetail.Authorities,
		Standclaim: jwt.StandardClaims{
			Audience:  tokenconf.Audience,
			ExpiresAt: tokenconf.ExpiresAt,
			IssuedAt:  tokenconf.IssuedAt,
			Issuer:    tokenconf.Issuer,
			NotBefore: tokenconf.NotBefore,
		},
	}
	return &claims
}

//刷新 claims 的时间
func (claims Claims) RefreshClaims() *Claims {
	//如果还有剩不到一小时，就刷新
	expires := claims.Standclaim.ExpiresAt
	if time.Now().Add(time.Hour).Unix() > expires {
		//修改token的过期时间
		claims.Standclaim.ExpiresAt = time.Now().Add(time.Hour * 2).Unix()
		claims.Standclaim.NotBefore = time.Now().Unix()
		claims.Standclaim.IssuedAt = time.Now().Unix()
	}
	return &claims
}
