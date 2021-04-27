package entity

import "github.com/smarest/smarest-common/domain/value"

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
