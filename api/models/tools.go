package models

import (
	"api/db"
)

type Tools struct {
	Id int
	Icon string
	Cname string
	Code string
	Title string
	Description string
	Status int8
}

/*
获取工具
 */
func GetToolsByCName(name string) ([]Tools, error) {
	list := make([]Tools, 0)

	if name == "" {
		err := db.GetMysqlLink().Find(&list)
		if err != nil {
			return nil, err
		}
	} else {
		err := db.GetMysqlLink().Where("cname = ?", name).Find(&list)
		if err != nil {
			return nil, err
		}
	}

	return list, nil
}
