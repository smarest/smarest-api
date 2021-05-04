package entity

import (
	"strings"

	"github.com/smarest/smarest-common/domain/value"
)

type Restaurant struct {
	ID                int64           `db:"id" json:"id"`
	Code              string          `db:"code" json:"code"`
	RestaurantGroupID int64           `db:"restaurant_group_id" json:"restaurantGroupID"`
	Name              string          `db:"name" json:"name"`
	Description       *string         `db:"description" json:"description"`
	Image             *string         `db:"image" json:"image"`
	Address           *string         `db:"address" json:"address"`
	Phone             *string         `db:"phone" json:"phone"`
	Available         bool            `db:"available" json:"available"`
	Creator           string          `db:"creator" json:"creator"`
	CreatedDate       value.DateTime  `db:"created_date" json:"createdDate"`
	Updater           *string         `db:"updater" json:"updater"`
	LastUpdatedDate   *value.DateTime `db:"last_updated_date" json:"lastUpdatedDate"`
}

func (item *Restaurant) ToSlice(fields string) interface{} {
	if fields == "" {
		return *item
	}
	result := make(map[string]interface{})
	// Loop over the parts from the string.
	for _, field := range strings.Split(fields, ",") {
		switch field {
		case "id":
			result[field] = item.ID
		case "code":
			result[field] = item.Code
		case "restaurantGroupID":
			result[field] = item.RestaurantGroupID
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

type RestaurantList struct {
	list []Restaurant
}

func NewRestaurantList(list []Restaurant) *RestaurantList {
	return &RestaurantList{list}
}

func CreateEmptyRestaurantList() *RestaurantList {
	return &RestaurantList{make([]Restaurant, 0)}
}

func (l *RestaurantList) Add(item Restaurant) {
	l.list = append(l.list, item)
}

func (l *RestaurantList) FindByID(id int64) *Restaurant {
	for _, item := range l.list {
		if item.ID == id {
			return &item
		}
	}
	return nil
}

func (l *RestaurantList) GetAvailable() *RestaurantList {
	return l.FilterBy(func(item Restaurant) bool { return item.Available })
}

func (l *RestaurantList) FilterBy(filter RestaurantFilter) *RestaurantList {
	list := make([]Restaurant, 0)
	for _, u := range l.list {
		if filter(u) {
			list = append(list, u)
		}
	}
	return NewRestaurantList(list)
}

type RestaurantFilter func(item Restaurant) bool

func (l *RestaurantList) ToSlice(fields string) []interface{} {
	var result []interface{} = make([]interface{}, len(l.list))
	for i, item := range l.list {
		result[i] = item.ToSlice(fields)
	}
	return result
}

func (l *RestaurantList) ToArray() []Restaurant {
	return l.list
}
