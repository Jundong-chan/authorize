package service

//鉴权服务，鉴别用户是否存在并且有效
import (
	model "github.com/Jundong-chan/authorize/authorize/model"
)

type AutorizeService interface {
	CheckUser(feild string, username string, password string) (*model.UserDetail, string, string, error)
}

type AutorizeServiceImpl struct {
}

//传入用户的账户及密码信息，验证用户，并返回用户的基本信息,token,refreshtoken
func (authsvcimpl AutorizeServiceImpl) CheckUser(feild string, username string, password string) (*model.UserDetail, string, error) {
	usermodel := model.NewUserModel()
	userdetail, err := usermodel.CheckUser(feild, username, password)
	if err != nil {
		return userdetail, "", err
	}
	tokensvcimp := TokenServiceImpl{}
	claims := model.NewClaims(userdetail, model.InitTokenConfig())
	token, err := tokensvcimp.CreateToken(claims)
	if err != nil {
		return userdetail, "", err
	}
	return userdetail, token, err
}
