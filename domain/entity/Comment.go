package entity

import "github.com/smarest/smarest-common/domain/value"

type Comment struct {
	ID              int64          `db:"id" json:"id"`
	Name            string         `db:"name" json:"name"`
	Description     *string        `db:"description" json:"description"`
	Available       bool           `db:"available" json:"available"`
	Creator         string         `db:"creator" json:"creator"`
	CreatedDate     value.DateTime `db:"created_date" json:"createdDate"`
	Updater         string         `db:"updater" json:"updater"`
	LastUpdatedDate value.DateTime `db:"last_updated_date" json:"lastUpdatedDate"`
}
