package models

import (
	"api/db"
)

type Category struct {
	Id int
	Name string
	Code string
	Status int8
}

/*
获取所有分类
 */
func GetAllCategory() ([]Category, error) {
	list := make([]Category, 0)

	err := db.GetMysqlLink().Find(&list)

	if err != nil {
		return nil, err
	}
	return list, nil
}
