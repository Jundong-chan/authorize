package service

//这是操作商品的接口
import (
	model "github.com/Jundong-chan/seckill/model"
	"log"
	"strconv"

	"github.com/gohouse/gorose"
)

type ProductService interface {
	CreateProduct(product *model.Product) error
	GetProductList() ([]gorose.Data, error)
	ModifyProduct(product *model.Product) error
	DeleteProduct(id string) error
	QureyProductByUserId(feild string, condition string, value interface{}) ([]gorose.Data, error)
	GetProductCounts() (int, error)
	PageProducts(page int, limit int) ([]gorose.Data, error)
}

type ProductServiceMidWare func(ProductService) ProductService

type ProductServiceImpl struct {
}

func (p ProductServiceImpl) CreateProduct(product *model.Product) error {
	produtmodel := model.NewProductModel()
	err := produtmodel.CreateProduct(product)
	if err != nil {
		log.Printf("ProductService.CraeteProduct Error:%v", err)
		return err
	}
	return nil
}

func (p ProductServiceImpl) GetProductList() ([]gorose.Data, error) {
	produtmodel := model.NewProductModel()
	list, err := produtmodel.GetProductList()
	if err != nil {
		log.Printf("ProductService.GetProductList Error:%v", err)
		return nil, err
	}
	return list, nil
}

func (p ProductServiceImpl) ModifyProduct(product *model.Product) error {
	produtmodel := model.NewProductModel()
	err := produtmodel.ModifyProduct(product)
	if err != nil {
		log.Printf("ProductService.ModifyProduct Error:%v", err)
		return err
	}
	return nil

}

//查询商品总数
func (p ProductServiceImpl) GetProductCounts() (int, error) {
	produtmodel := model.NewProductModel()
	count, err := produtmodel.GetProductCounts()
	if err != nil {
		log.Printf("ProductService.GetproductCounts Error:%v", err)
		return 0, err
	}
	return count, nil
}

func (p ProductServiceImpl) DeleteProduct(id string) error {
	produtmodel := model.NewProductModel()
	err := produtmodel.DeleteProduct(id)
	if err != nil {
		log.Printf("ProductService.DeleteProduct Error:%v", err)
		return err
	}
	return nil
}

func (p ProductServiceImpl) QureyProductByUserId(feild string, condition string, value interface{}) ([]gorose.Data, error) {
	tem := value.(string)
	var values interface{}
	produtmodel := model.NewProductModel()
	if feild == "product_id" {
		values, _ = strconv.Atoi(tem)
	}
	values = value
	result, err := produtmodel.QureyProductByUserId(feild, condition, values)
	if err != nil {
		log.Printf("ProductService.QureyProductByUserId Error:%v", err)
		return nil, err
	}
	return result, nil
}

//分页查询
func (p ProductServiceImpl) PageProducts(page int, limit int) ([]gorose.Data, error) {
	produtmodel := model.NewProductModel()
	result, err := produtmodel.PageProducts(page, limit)
	if err != nil {
		log.Printf("ProductService.PageProducts Error:%v", err)
		return nil, err
	}
	return result, nil
}
