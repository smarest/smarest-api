package entity

import (
	"strings"

	"github.com/smarest/smarest-common/domain/value"
)

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

func (item *Area) ToSlice(fields string) interface{} {
	if fields == "" {
		return *item
	}
	result := make(map[string]interface{})
	// Loop over the parts from the string.
	for _, field := range strings.Split(fields, ",") {
		switch field {
		case "id":
			result[field] = item.ID
		case "name":
			result[field] = item.Name
		case "restaurantId":
			result[field] = item.RestaurantID
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

type AreaList struct {
	list []Area
}

func NewAreaList(areaList []Area) *AreaList {
	return &AreaList{areaList}
}

func (l *AreaList) GetAvailable() *AreaList {
	return l.FilterBy(func(item Area) bool { return item.Available })
}

func (l *AreaList) FilterBy(filter AreaFilter) *AreaList {
	list := make([]Area, 0)
	for _, u := range l.list {
		if filter(u) {
			list = append(list, u)
		}
	}
	return NewAreaList(list)
}

type AreaFilter func(item Area) bool

func (l *AreaList) ToSlice(fields string) []interface{} {
	var result []interface{} = make([]interface{}, len(l.list))
	for i, item := range l.list {
		result[i] = item.ToSlice(fields)
	}
	return result
}

func (l *AreaList) ToArray() []Area {
	return l.list
}
