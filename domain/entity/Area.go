package entity

import "github.com/smarest/smarest-common/domain/value"

type Area struct {
	ID              int64          `db:"id" json:"id"`
	Name            string         `db:"name" json:"name"`
	RestaurantID    int64          `db:"restaurant_id" json:"restaurantID"`
	Available       bool           `db:"available" json:"available"`
	Creator         string         `db:"creator" json:"creator"`
	CreatedDate     value.DateTime `db:"created_date" json:"createdDate"`
	Updater         string         `db:"updater" json:"updater"`
	LastUpdatedDate value.DateTime `db:"last_updated_date" json:"lastUpdatedDate"`
}
