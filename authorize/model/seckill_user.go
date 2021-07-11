package model

import (
	"github.com/Jundong-chan/authorize/authorize/pkg/encrypt"
	"github.com/Jundong-chan/authorize/authorize/pkg/mysql"
	"errors"

	"github.com/gohouse/gorose"
)

type User struct {
	UserId     string `gorose:"user_id" json:"userid"`
	UserName   string `gorose:"user_name" json:"username"`
	Password   string `gorose:"password"`
	Salt       string `gorose:"salt"`
	Gender     string `gorose:"gender"`
	Email      string `gorose:"email"`
	Phone      string `gorose:"phone"`
	CreateTime string `gorose:"create_time"`
	UpdateTime string `gorose:"update_time"`
}

type Usermodel struct {
}

func NewUserModel() *Usermodel {
	return &Usermodel{}
}

func (umodel *Usermodel) getTableName() string {
	return "seckill_user"
}

//传入用户账户类型，账户，密码，验证用户,成功则返回用户信息
func (um *Usermodel) CheckUser(feild string, account string, password string) (*UserDetail, error) {
	conn := mysql.DB()
	var user []gorose.Data
	user, err := conn.Table(um.getTableName()).Where(feild, "=", account).Get()
	if err != nil {
		return nil, errors.New("Wrong feild," + err.Error())
	}
	if len(user) == 0 {
		return nil, errors.New("Wrong User Account")
	}

	iscorrect := encrypt.TestifyEncrypt(password, user[0]["salt"], user[0]["password"])
	if !iscorrect {
		return nil, errors.New("Wrong password")
	}
	udetail := UserDetail{
		UserId:      user[0]["user_id"].(string),
		UserName:    user[0]["user_name"].(string),
		Password:    user[0]["password"].(string),
		Gender:      user[0]["gender"].(string),
		Email:       user[0]["email"].(string),
		Phone:       user[0]["phone"].(string),
		Authorities: []string{"read", "write"},
	}
	return &udetail, nil

}
