package entity

import (
	"strings"

	"github.com/smarest/smarest-common/domain/value"
)

type Product struct {
	ID                    int64          `db:"id" json:"id"`
	RestaurantGroupID     int64          `db:"restaurant_group_id" json:"restaurantGroupID"`
	Name                  string         `db:"name" json:"name"`
	CategoryID            int64          `db:"category_id" json:"categoryId"`
	Price                 int64          `db:"price" json:"price"`
	UnitID                int64          `db:"unit_id" json:"unitId"`
	Description           *string        `db:"description" json:"description"`
	Image                 string         `db:"image" json:"image"`
	DefaultStatus         bool           `db:"default_status" json:"defaultStatus"`
	QuantityOnSingleOrder int64          `db:"quantity_on_single_order" json:"quantityOnSingleOrder"`
	Available             bool           `db:"available" json:"available"`
	Creator               string         `db:"creator" json:"creator"`
	CreatedDate           value.DateTime `db:"created_date" json:"createdDate"`
	Updater               string         `db:"updater" json:"updater"`
	LastUpdatedDate       value.DateTime `db:"last_updated_date" json:"lastUpdatedDate"`
}

func (item *Product) ToSlide(fields string) interface{} {
	if fields == "" {
		return *item
	}
	result := make(map[string]interface{})
	// Loop over the parts from the string.
	for _, field := range strings.Split(fields, ",") {
		switch field {
		case "id":
			result[field] = item.ID
		case "restaurantGroupID":
			result[field] = item.RestaurantGroupID
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
