package service

//这是操作商品的接口
import (
	model "CommoditySpike/server/model"
	"errors"
	"fmt"
	"log"

	"github.com/gohouse/gorose"
)

type UserService interface {
	CreateUser(user *model.User) error
	GetUserList() ([]gorose.Data, error)
	QureyUser(feild string, condition string, value interface{}) ([]gorose.Data, error)
	CheckUser(feild string, username string, password string) ([]gorose.Data, error)
}

type UserServiceMidWare func(UserService) UserService

type UserServiceImpl struct {
}

func (p UserServiceImpl) CreateUser(user *model.User) error {
	usermodel := model.NewUserModel()
	fmt.Println(user)
	if _, erro := p.QureyUser("email", "=", user.Email); erro == nil {
		return errors.New("user has existed")
	}
	err := usermodel.CreateUser(user)
	if err != nil {
		log.Printf("UserService.CreateUser Error:%v", err)
		return err
	}
	return nil
}

func (p UserServiceImpl) GetUserList() ([]gorose.Data, error) {
	usermodel := model.NewUserModel()
	list, err := usermodel.GetUserList()
	if err != nil {
		log.Printf("UserService.GetUserList Error:%v", err)
		return nil, err
	}
	return list, nil
}

func (p UserServiceImpl) QureyUser(feild string, condition string, value interface{}) ([]gorose.Data, error) {
	usermodel := model.NewUserModel()
	result, err := usermodel.QureyUserByCondition(feild, condition, value)
	if err != nil {
		log.Printf("UserService.QureyUser Error:%v", err)
		return nil, err
	}
	return result, nil
}

//验证用户传入的信息是否正确，正确则返回用户全部信息
func (p UserServiceImpl) CheckUser(feild string, username string, password string) ([]gorose.Data, error) {
	usermodel := model.NewUserModel()
	result, err := usermodel.CheckUser(feild, username, password)
	if err != nil {
		return nil, err
	}
	return result, nil
}
