package service

import (
	model "github.com/Jundong-chan/authorize/authorize/model"
	"errors"

	"github.com/dgrijalva/jwt-go"
)

type TokenService interface {
	//创建token
	CreateToken(claims *model.Claims) (tokenstring string, err error)
	//验证传入的token,有效则返回claims(payload内容)
	ParseToken(token string) (*model.Claims, error)
	//验证token的合法性，并返回新的token,以及claims
	ParseTokenAndReToken(token string) (newtoken string, claims *model.Claims, err error)
}

var jwtSecret = []byte("dfsdfueh2euw7d")
var Method = "HS256"

type TokenServiceImpl struct {
}

//传入claims生成一个token
func (tokensvcimp TokenServiceImpl) CreateToken(claims *model.Claims) (tokenstring string, err error) {
	token := jwt.NewWithClaims(jwt.GetSigningMethod(Method), claims)
	tokenstring, err = token.SignedString([]byte(jwtSecret))
	return
}

//验证token，如果过期或者未到签发时间则会返回对应的错误,正确则返回claims
func (tokensvcimp TokenServiceImpl) ParseToken(token string) (*model.Claims, error) {
	//第三个参数是用来返回签名密钥的，用于验证
	tokenclaim, err := jwt.ParseWithClaims(token, &model.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenclaim != nil {
		// 从tokenClaims中获取到Claims对象，并使用断言，将该对象转换为我们自己定义的Claims
		claims, ok := tokenclaim.Claims.(*model.Claims)
		if ok && tokenclaim.Valid {
			return claims, nil
		}
	}
	return nil, errors.New("ParseToken Faild:" + err.Error())
}

//传入token 返回正确的claim和token
func (tokensvcimp TokenServiceImpl) ParseTokenAndReToken(token string) (newtoken string, claims *model.Claims, err error) {
	//先验证合法性
	claims, err = tokensvcimp.ParseToken(token)
	if err != nil {
		return "", nil, err
	}
	//刷新claims
	nclaims := claims.RefreshClaims()

	//生成新的token
	ntoken, err := tokensvcimp.CreateToken(nclaims)
	if err != nil {
		return "", nil, err
	}
	return ntoken, nclaims, nil

}
