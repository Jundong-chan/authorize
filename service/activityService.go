package service

//这是操作秒杀活动的接口
import (
	model "CommoditySpike/server/model"
	"log"

	"github.com/gohouse/gorose"
)

type ActivityService interface {
	CreateActivity(activity *model.Activity) error
	GetActivityList() ([]gorose.Data, error)
}

type ActivityServiceMidWare func(ActivityService) ActivityService

type ActivityServiceImpl struct {
}

func (a ActivityServiceImpl) CreateActivity(activity *model.Activity) error {
	activitymodel := model.NewActivityModel()
	err := activitymodel.CreateActivity(activity)
	if err != nil {
		log.Printf("ActivityService.CreateActivity Error:%v", err)
		return err
	}
	return nil
}

func (a ActivityServiceImpl) GetActivityList() ([]gorose.Data, error) {
	activitymodel := model.NewActivityModel()
	list, err := activitymodel.GetActivityList()
	if err != nil {
		log.Printf("ActivityService.GetActivityList Error:%v", err)
		return nil, err
	}
	return list, nil
}
