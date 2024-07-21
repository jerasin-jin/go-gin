package repository

import (
	"fmt"
	"math"
	"strings"

	"github.com/Jerasin/app/model"
	"github.com/Jerasin/app/request"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ProductCategoryRepositoryInterface interface {
	FindAllProductCategory() ([]model.ProductCategory, error)
	PaginationProductCategory(imit int, offset int, search string, sortField string, sortValue string, field map[string]interface{}) ([]model.ProductCategory, error)
	FindOneProductCategory(condition model.ProductCategory) (model.ProductCategory, error)
	FindProductCategoryById(id int) (model.ProductCategory, error)
	Save(product *model.ProductCategory) (model.ProductCategory, error)
	DeleteProductById(id int) error
	Count() (int64, error)
	TotalPage(pageSize int) (int64, error)
	Update(id int, update *request.UpdateProductCategory) error
}

type ProductCategoryRepository struct {
	db *gorm.DB
}

func (p ProductCategoryRepository) FindAllProductCategory() ([]model.ProductCategory, error) {
	var productCategory []model.ProductCategory
	var err = p.db.Find(&productCategory).Error
	if err != nil {
		log.Error("Got an error finding all couples. Error: ", err)
		return nil, err
	}

	return productCategory, nil
}

func (p ProductCategoryRepository) PaginationProductCategory(imit int, offset int, search string, sortField string, sortValue string, field map[string]interface{}) ([]model.ProductCategory, error) {
	var productCategories []model.ProductCategory

	log.Info("offset", offset)
	log.Info("imit", imit)

	fields := getField(field)

	order := fmt.Sprintf("%s %s", sortField, strings.ToUpper(sortValue))
	fmt.Println("order", order)
	var err error

	if fields == "" {
		err = p.db.Order(order).Offset(offset).Limit(imit).Find(&productCategories).Error
	} else {
		err = p.db.Select(fields).Order(order).Offset(offset).Limit(imit).Find(&productCategories).Error
	}

	if err != nil {
		log.Error("Got an error finding all couples. Error: ", err)
		return nil, err
	}

	return productCategories, nil
}

func (p ProductCategoryRepository) FindOneProductCategory(condition model.ProductCategory) (model.ProductCategory, error) {
	// var user model.ProductCategory

	var err = p.db.First(&condition).Error
	if err != nil {
		log.Error("Got an error finding One couples. Error: ", err)
		return condition, err
	}

	return condition, nil
}

func (p ProductCategoryRepository) FindProductCategoryById(id int) (model.ProductCategory, error) {
	var user model.ProductCategory
	err := p.db.First(&user, id).Error
	if err != nil {
		log.Error("Got and error when find user by id. Error: ", err)
		return model.ProductCategory{}, err
	}
	return user, nil
}

func (p ProductCategoryRepository) Save(product *model.ProductCategory) (model.ProductCategory, error) {
	var err = p.db.Save(product).Error
	if err != nil {
		log.Error("Got an error when save user. Error: ", err)
		return model.ProductCategory{}, err
	}
	return *product, nil
}

func (p ProductCategoryRepository) DeleteProductById(id int) error {
	err := p.db.Delete(&model.ProductCategory{}, id).Error
	if err != nil {
		log.Error("Got an error when delete user. Error: ", err)
		return err
	}
	return nil
}

func (p ProductCategoryRepository) Count() (int64, error) {
	var product model.ProductCategory
	var count int64
	err := p.db.Model(&product).Count(&count).Error
	if err != nil {
		log.Error("Got an error when delete user. Error: ", err)
		return count, err
	}
	return count, err
}

func (p ProductCategoryRepository) TotalPage(pageSize int) (int64, error) {
	var user model.User
	var count int64
	err := p.db.Model(&user).Count(&count).Error
	if err != nil {
		log.Error("Got an error when delete user. Error: ", err)
		return count, err
	}

	totalPage := int64(math.Ceil(float64(count) / float64(pageSize)))
	return totalPage, err
}

func (p ProductCategoryRepository) Update(id int, update *request.UpdateProductCategory) error {
	var productCategory model.ProductCategory
	var err = p.db.Model(&productCategory).Where(id).Updates(update).Error
	if err != nil {
		log.Error("Got an error when save user. Error: ", err)
		return err
	}
	return nil
}

func ProductCategoryRepositoryInit(db *gorm.DB) *ProductCategoryRepository {
	db.AutoMigrate(&model.ProductCategory{})
	return &ProductCategoryRepository{
		db: db,
	}
}
