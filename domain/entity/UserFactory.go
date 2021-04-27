package entity

import (
	"strings"
)

type UserFactory struct{}

func NewUserFactory() UserFactory {
	return UserFactory{}
}

func (f *UserFactory) CreateList(items []User, fields string) []interface{} {
	var result []interface{} = make([]interface{}, len(items))
	for i, item := range items {
		result[i] = f.Create(&item, fields)
	}
	return result
}

func (f *UserFactory) Create(item *User, fields string) map[string]interface{} {
	result := make(map[string]interface{})
	for _, field := range strings.Split(fields, ",") {
		switch field {
		case "userName":
			result[field] = item.UserName
		case "role":
			result[field] = item.Role
		case "name":
			result[field] = item.Name
		case "salaryType":
			result[field] = item.SalaryType
		case "joinedDate":
			result[field] = item.JoinedDate
		case "leftDate":
			result[field] = item.LeftDate
		default:
		}
	}
	return result
}
