package entity

import (
	"strings"

	"github.com/smarest/smarest-common/domain/value"
)

type Category struct {
	ID              int64          `db:"id" json:"id"`
	Name            string         `db:"name" json:"name"`
	Description     *string        `db:"description" json:"description"`
	Available       bool           `db:"available" json:"available"`
	Image           *string        `db:"image" json:"image"`
	Type            string         `db:"type" json:"type"`
	Creator         string         `db:"creator" json:"creator"`
	CreatedDate     value.DateTime `db:"created_date" json:"createdDate"`
	Updater         string         `db:"updater" json:"updater"`
	LastUpdatedDate value.DateTime `db:"last_updated_date" json:"lastUpdatedDate"`
}

func (item *Category) ToSlice(fields string) map[string]interface{} {
	result := make(map[string]interface{})
	// Loop over the parts from the string.
	for _, field := range strings.Split(fields, ",") {
		switch field {
		case "id":
			result[field] = item.ID
		case "name":
			result[field] = item.Name
		case "description":
			result[field] = item.Description
		case "image":
			result[field] = item.Image
		case "type":
			result[field] = item.Type
		case "available":
			result[field] = item.Available
		case "creator":
			result[field] = item.Creator
		case "createdDate":
			result[field] = item.CreatedDate
		case "updater":
			result[field] = item.Updater
		case "lastUpdatedDate":
			result[field] = item.LastUpdatedDate
		default:
		}
	}
	return result
}

type CategoryList struct {
	Categories []Category
}

func NewCategoryList(items []Category) CategoryList {
	return CategoryList{items}
}

func (l *CategoryList) ToSlice(fields string) []interface{} {
	var result []interface{} = make([]interface{}, len(l.Categories))
	for i, item := range l.Categories {
		result[i] = item.ToSlice(fields)
	}
	return result
}

func (l *CategoryList) ToArray() []Category {
	return l.Categories
}

func (l *CategoryList) GetAvailable() CategoryList {
	return l.FilterBy(func(cate Category) bool { return cate.Available })
}

func (l *CategoryList) FilterBy(filter CategoryFilter) CategoryList {
	list := make([]Category, 0)
	for _, u := range l.Categories {
		if filter(u) {
			list = append(list, u)
		}
	}
	return NewCategoryList(list)
}

type CategoryFilter func(cate Category) bool
