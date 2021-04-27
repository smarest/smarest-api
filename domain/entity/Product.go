package entity

import "github.com/smarest/smarest-common/domain/value"

type Product struct {
	ID                    int64          `db:"id" json:"id"`
	Name                  string         `db:"name" json:"name"`
	CategoryID            int64          `db:"category_id" json:"categoryId"`
	Price                 int64          `db:"price" json:"price"`
	UnitID                int64          `db:"unit_id" json:"unitId"`
	Description           *string        `db:"description" json:"description"`
	Image                 string         `db:"image" json:"image"`
	DefaultStatus         bool           `db:"default_status" json:"defaultStatus"`
	QuantityOnSingleOrder int            `db:"quantity_on_single_order" json:"quantityOnSingleOrder"`
	Available             bool           `db:"available" json:"available"`
	Creator               string         `db:"creator" json:"creator"`
	CreatedDate           value.DateTime `db:"created_date" json:"createdDate"`
	Updater               string         `db:"updater" json:"updater"`
	LastUpdatedDate       value.DateTime `db:"last_updated_date" json:"lastUpdatedDate"`
}
