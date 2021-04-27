package entity

import (
	"strings"
)

type ProductFactory struct{}

func NewProductFactory() ProductFactory {
	return ProductFactory{}
}

func (f *ProductFactory) CreateList(items []Product, fields string) []interface{} {
	var result []interface{} = make([]interface{}, len(items))
	for i, item := range items {
		result[i] = f.Create(&item, fields)
	}
	return result
}
func (f *ProductFactory) Create(item *Product, fields string) map[string]interface{} {
	result := make(map[string]interface{})
	// Loop over the parts from the string.
	for _, field := range strings.Split(fields, ",") {
		switch field {
		case "id":
			result[field] = item.ID
		case "name":
			result[field] = item.Name
		case "categoryId":
			result[field] = item.CategoryID
		case "unitId":
			result[field] = item.CategoryID
		case "description":
			result[field] = item.Description
		case "image":
			result[field] = item.Image
		case "defaultStatus":
			result[field] = item.DefaultStatus
		case "price":
			result[field] = item.Price
		case "quantityOnSingleOrder":
			result[field] = item.QuantityOnSingleOrder
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
