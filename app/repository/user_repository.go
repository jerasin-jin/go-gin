package repository

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/Jerasin/app/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	FindAllUser(imit int, offset int, search string, sortField string, sortValue string, field map[string]interface{}) ([]model.User, error)
	FindOneUser(user model.User, query map[interface{}]interface{}, field map[string]interface{}) (model.User, error)
	FindUserById(id int) (model.User, error)
	Save(user *model.User) (model.User, error)
	DeleteUserById(id int) error
	Count() (int64, error)
}

type UserRepository struct {
	db *gorm.DB
}

func getField(field map[string]interface{}) string {
	b := new(bytes.Buffer)
	index := 0
	for key := range field {

		if index > 0 {
			fmt.Fprintf(b, ",%s", strings.ToLower(key))
		} else {
			fmt.Fprintf(b, "%s", strings.ToLower(key))
		}

		index += 1
	}
	return b.String()

}

func (u UserRepository) FindAllUser(imit int, offset int, search string, sortField string, sortValue string, field map[string]interface{}) ([]model.User, error) {
	var users []model.User

	log.Info("offset", offset)
	log.Info("imit", imit)

	fields := getField(field)

	order := fmt.Sprintf("%s %s", sortField, strings.ToUpper(sortValue))

	fmt.Println("order", order)
	fmt.Println("fields", fields)

	var err error

	fmt.Println("fields ==", fields == "")

	if fields == "" {
		err = u.db.Order(order).Offset(offset).Limit(imit).Find(&users).Error
	} else {
		err = u.db.Select(fields).Order(order).Offset(offset).Limit(imit).Find(&users).Error
	}

	if err != nil {
		log.Error("Got an error finding all couples. Error: ", err)
		return nil, err
	}

	return users, nil
}

func (u UserRepository) FindOneUser(user model.User, query map[interface{}]interface{}, field map[string]interface{}) (model.User, error) {

	fields := getField(field)

	var err error

	if fields == "" {
		err = u.db.Where(query).First(&user).Error
	} else {
		err = u.db.Select(fields).Where(query).First(&user).Error
	}

	if err != nil {
		log.Error("Got an error finding One couples. Error: ", err)
		return user, err
	}

	fmt.Println("user", user)

	return user, nil
}

func (u UserRepository) FindUserById(id int) (model.User, error) {
	var user model.User
	err := u.db.First(&user, id).Error
	if err != nil {
		log.Error("Got and error when find user by id. Error: ", err)
		return model.User{}, err
	}
	return user, nil
}

func (u UserRepository) Save(user *model.User) (model.User, error) {
	var err = u.db.Save(user).Error
	if err != nil {
		log.Error("Got an error when save user. Error: ", err)
		return model.User{}, err
	}
	return *user, nil
}

func (u UserRepository) DeleteUserById(id int) error {
	err := u.db.Delete(&model.User{}, id).Error
	if err != nil {
		log.Error("Got an error when delete user. Error: ", err)
		return err
	}
	return nil
}

func (u UserRepository) Count() (int64, error) {
	var user model.User
	var count int64
	err := u.db.Model(&user).Count(&count).Error
	if err != nil {
		log.Error("Got an error when delete user. Error: ", err)
		return count, err
	}
	return count, err
}

func UserRepositoryInit(db *gorm.DB) UserRepositoryInterface {
	db.AutoMigrate(&model.User{})
	return &UserRepository{
		db: db,
	}
}
