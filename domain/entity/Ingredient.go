package entity

import (
	"strings"

	"github.com/smarest/smarest-common/domain/value"
)

type Ingredient struct {
	ID              int64          `db:"id" json:"id"`
	CategoryID      int64          `db:"category_id" json:"categoryId"`
	UnitID          int64          `db:"unit_id" json:"unitId"`
	Name            string         `db:"name" json:"name"`
	Description     *string        `db:"description" json:"description"`
	Image           *string        `db:"image" json:"image"`
	ReferencePrice  int64          `db:"reference_price" json:"referencePrice"`
	Available       bool           `db:"available" json:"available"`
	Creator         string         `db:"creator" json:"creator"`
	CreatedDate     value.DateTime `db:"created_date" json:"createdDate"`
	Updater         string         `db:"updater" json:"updater"`
	LastUpdatedDate value.DateTime `db:"last_updated_date" json:"lastUpdatedDate"`
}

func (item *Ingredient) ToSlice(fields string) interface{} {
	if fields == "" {
		return *item
	}
	result := make(map[string]interface{})
	// Loop over the parts from the string.
	for _, field := range strings.Split(fields, ",") {
		switch field {
		case "id":
			result[field] = item.ID
		case "categoryId":
			result[field] = item.CategoryID
		case "unitId":
			result[field] = item.UnitID
		case "name":
			result[field] = item.Name
		case "description":
			result[field] = item.Description
		case "image":
			result[field] = item.Image
		case "referencePrice":
			result[field] = item.ReferencePrice
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

type IngredientList struct {
	list []Ingredient
}

func NewIngredientList(list []Ingredient) *IngredientList {
	return &IngredientList{list}
}

func (l *IngredientList) GetAvailable() *IngredientList {
	return l.FilterBy(func(item Ingredient) bool { return item.Available })
}

func (l *IngredientList) FilterBy(filter IngredientFilter) *IngredientList {
	list := make([]Ingredient, 0)
	for _, u := range l.list {
		if filter(u) {
			list = append(list, u)
		}
	}
	return NewIngredientList(list)
}

type IngredientFilter func(item Ingredient) bool

func (l *IngredientList) ToSlice(fields string) []interface{} {
	var result []interface{} = make([]interface{}, len(l.list))
	for i, item := range l.list {
		result[i] = item.ToSlice(fields)
	}
	return result
}

func (l *IngredientList) ToArray() []Ingredient {
	return l.list
}
