package repository

import (
	"bytes"
	"fmt"
	"math"
	"strings"

	"github.com/Jerasin/app/constant"
	"github.com/Jerasin/app/pkg"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type BaseRepositoryInterface interface {
	ClientDb() *gorm.DB
	Pagination(p PaginationModel) (result interface{}, Error error)
	Create(tx *gorm.DB, model interface{}) error
	IsExits(tx *gorm.DB, model interface{}, query interface{}, args ...interface{}) error
	FindOne(tx *gorm.DB, model interface{}, query interface{}, args ...interface{}) error
	Update(id int, model interface{}, update interface{}) error
	TotalPage(model interface{}, pageSize int) (int64, error)
	Delete(model interface{}, id int) error
}

type BaseRepository struct {
	db *gorm.DB
}

type PaginationModel struct {
	Limit     int
	Offset    int
	Search    string
	SortField string
	SortValue string
	Field     map[string]interface{}
	Dest      interface{}
}

func BaseRepositoryInit(db *gorm.DB) *BaseRepository {

	return &BaseRepository{
		db: db,
	}
}

func getField(field map[string]interface{}) string {
	b := new(bytes.Buffer)
	index := 0
	for key := range field {
		// fmt.Println("key", key)
		if index > 0 {
			fmt.Fprintf(b, ",%s", strings.ToLower(key))
		} else {
			fmt.Fprintf(b, "%s", strings.ToLower(key))
		}

		index += 1
	}
	return b.String()

}
func (b BaseRepository) ClientDb() *gorm.DB {
	return b.db
}

func (b BaseRepository) Pagination(p PaginationModel) (result interface{}, Error error) {
	var err error
	order := fmt.Sprintf("%s %s", p.SortField, strings.ToUpper(p.SortValue))
	fields := getField(p.Field)

	if fields == "" {
		err = b.db.Order(order).Offset(p.Offset).Limit(p.Limit).Find(&p.Dest).Error
	} else {
		err = b.db.Select(fields).Order(order).Offset(p.Offset).Limit(p.Limit).Find(&p.Dest).Error
	}

	if err != nil {
		log.Error("Got an error finding all couples. Error: ", err)
		return nil, err
	}

	return p.Dest, nil
}

func (b BaseRepository) Create(tx *gorm.DB, model interface{}) error {
	db := b.db

	if tx != nil {
		db = tx
	}

	var err = db.Save(model).Error
	if err != nil {
		log.Error("Got an error when save Error: ", err)
		return err
	}
	return nil
}

func (b BaseRepository) IsExits(tx *gorm.DB, model interface{}, query interface{}, args ...interface{}) error {
	db := b.db

	if tx != nil {
		db = tx
	}

	var err error
	if query != nil {
		err = db.Where(query, args).First(model).Error
	} else {
		err = db.First(model).Error
	}

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}

		log.Error("Got an error when findOne Error: ", err)
		return err
	}

	pkg.PanicException(constant.DataIsExit)
	return nil
}

func (b BaseRepository) FindOne(tx *gorm.DB, model interface{}, query interface{}, args ...interface{}) error {
	db := b.db

	if tx != nil {
		db = tx
	}
	var err error
	if query == nil || args == nil {
		log.Error("Got an error when findOne required query")
		pkg.PanicException(constant.RequiredQuery)
	}

	err = db.Where(query, args).First(model).Error

	if err != nil {
		log.Error("Got an error when findOne Error: ", err)
		return err
	}

	return nil
}

func (b BaseRepository) Update(id int, model interface{}, update interface{}) error {
	var err = b.db.Model(model).Where(id).Updates(update).Error
	if err != nil {
		log.Error("Got an error when save user. Error: ", err)
		return err
	}
	return nil
}

func (b BaseRepository) TotalPage(model interface{}, pageSize int) (int64, error) {
	var count int64
	err := b.db.Model(model).Count(&count).Error
	if err != nil {
		log.Error("Got an error when delete user. Error: ", err)
		return count, err
	}

	totalPage := int64(math.Ceil(float64(count) / float64(pageSize)))
	return totalPage, err
}

func (b BaseRepository) Delete(model interface{}, id int) error {
	err := b.db.Delete(model, id).Error
	if err != nil {
		log.Error("Got an error when delete user. Error: ", err)
		return err
	}
	return nil
}
