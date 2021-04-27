package entity

import "github.com/smarest/smarest-common/domain/value"

type RestaurantProduct struct {
	RestaurantID          int64          `db:"restaurant_id" json:"restaurantID"`
	ProductID             int64          `db:"product_id" json:"productID"`
	Price                 string         `db:"price" json:"price"`
	CategoryID            int64          `db:"category_id" json:"categoryId"`
	UnitID                int64          `db:"unit_id" json:"unitId"`
	Description           string         `db:"description" json:"description"`
	Image                 string         `db:"image" json:"image"`
	DefaultStatus         int            `db:"default_status" json:"defaultStatus"`
	ReferencePrice        int64          `db:"reference_price" json:"referencePrice"`
	QuantityOnSingleOrder int            `db:"quantity_on_single_order" json:"quantityOnSingleOrder"`
	Available             bool           `db:"available" json:"available"`
	Creator               string         `db:"creator" json:"creator"`
	CreatedDate           value.DateTime `db:"created_date" json:"createdDate"`
	Updater               string         `db:"updater" json:"updater"`
	LastUpdatedDate       value.DateTime `db:"last_updated_date" json:"lastUpdatedDate"`
}
