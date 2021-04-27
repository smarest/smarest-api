package entity

import (
	"strings"
)

type IngredientFactory struct{}

func NewIngredientFactory() IngredientFactory {
	return IngredientFactory{}
}

func (f *IngredientFactory) CreateList(items []Ingredient, fields string) []interface{} {
	var result []interface{} = make([]interface{}, len(items))
	for i, item := range items {
		result[i] = f.Create(&item, fields)
	}
	return result
}
func (f *IngredientFactory) Create(item *Ingredient, fields string) map[string]interface{} {
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
