package util

import (
	"slices"

	"github.com/Jerasin/app/constant"
	"github.com/Jerasin/app/pkg"
)

type CustomError struct {
	message string
}

func (c CustomError) Error() string {
	return c.message
}

func FindElementByCondition[T any](slice []T, condition func(T) bool) (*T, error) {
	index := slices.IndexFunc(slice, condition)

	if index != -1 {
		return &slice[index], nil
	} else {
		pkg.PanicException(constant.BadRequest)
		return nil, CustomError{
			message: "FindElementByCondition Error",
		}
	}

}
