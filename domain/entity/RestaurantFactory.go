package entity

import (
	"strings"
)

type RestaurantFactory struct{}

func NewRestaurantFactory() RestaurantFactory {
	return RestaurantFactory{}
}

func (f *RestaurantFactory) CreateList(items []Restaurant, fields string) []interface{} {
	var result []interface{} = make([]interface{}, len(items))
	for i, item := range items {
		result[i] = f.Create(&item, fields)
	}
	return result
}
func (f *RestaurantFactory) Create(item *Restaurant, fields string) map[string]interface{} {
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
		case "address":
			result[field] = item.Address
		case "phone":
			result[field] = item.Phone
		case "areaVersion":
			result[field] = item.AreaVersion
		case "tableVersion":
			result[field] = item.TableVersion
		case "categoryVersion":
			result[field] = item.CategoryVersion
		case "productVersion":
			result[field] = item.ProductVersion
		case "unitVersion":
			result[field] = item.UnitVersion
		case "commentVersion":
			result[field] = item.CommentVersion
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
