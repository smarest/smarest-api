package entity

import "github.com/smarest/smarest-common/domain/value"

type Restaurant struct {
	ID              int64           `db:"id" json:"id"`
	AccessKey       string          `db:"access_key" json:"-"`
	Name            string          `db:"name" json:"name"`
	Description     *string         `db:"description" json:"description"`
	Image           *string         `db:"image" json:"image"`
	Address         *string         `db:"address" json:"address"`
	Phone           *string         `db:"phone" json:"phone"`
	Available       bool            `db:"available" json:"available"`
	AreaVersion     int64           `db:"area_version" json:"areaVersion"`
	TableVersion    int64           `db:"table_version" json:"tableVersion"`
	CategoryVersion int64           `db:"category_version" json:"categoryVersion"`
	ProductVersion  int64           `db:"product_version" json:"productVersion"`
	UnitVersion     int64           `db:"unit_version" json:"unitVersion"`
	CommentVersion  int64           `db:"comment_version" json:"commentVersion"`
	Creator         string          `db:"creator" json:"creator"`
	CreatedDate     value.DateTime  `db:"created_date" json:"createdDate"`
	Updater         *string         `db:"updater" json:"updater"`
	LastUpdatedDate *value.DateTime `db:"last_updated_date" json:"lastUpdatedDate"`
}
