package entity

import (
	"strings"

	"github.com/smarest/smarest-common/domain/value"
)

type Comment struct {
	ID              int64          `db:"id" json:"id"`
	ProductID       int64          `db:"product_id" json:"productID"`
	Name            string         `db:"name" json:"name"`
	Description     *string        `db:"description" json:"description"`
	Available       bool           `db:"available" json:"available"`
	Creator         string         `db:"creator" json:"creator"`
	CreatedDate     value.DateTime `db:"created_date" json:"createdDate"`
	Updater         string         `db:"updater" json:"updater"`
	LastUpdatedDate value.DateTime `db:"last_updated_date" json:"lastUpdatedDate"`
}

func (item *Comment) ToSlice(fields string) interface{} {
	if fields == "" {
		return *item
	}
	result := make(map[string]interface{})
	// Loop over the parts from the string.
	for _, field := range strings.Split(fields, ",") {
		switch field {
		case "id":
			result[field] = item.ID
		case "productID":
			result[field] = item.ProductID
		case "name":
			result[field] = item.Name
		case "description":
			result[field] = item.Description
		case "available":
			result[field] = item.Available
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

type CommentList struct {
	list []Comment
}

func NewCommentList(list []Comment) *CommentList {
	return &CommentList{list}
}

func CreateEmptyCommentList() *CommentList {
	return &CommentList{make([]Comment, 0)}
}

func (l *CommentList) Add(item Comment) {
	l.list = append(l.list, item)
}

func (l *CommentList) GetAvailable() *CommentList {
	return l.FilterBy(func(item Comment) bool { return item.Available })
}

func (l *CommentList) FilterByProductID(productID int64) *CommentList {
	return l.FilterBy(func(item Comment) bool { return item.ProductID == productID })
}

func (l *CommentList) FilterBy(filter CommentFilter) *CommentList {
	list := make([]Comment, 0)
	for _, u := range l.list {
		if filter(u) {
			list = append(list, u)
		}
	}
	return NewCommentList(list)
}

type CommentFilter func(item Comment) bool

func (l *CommentList) ToSlice(fields string) []interface{} {
	var result []interface{} = make([]interface{}, len(l.list))
	for i, item := range l.list {
		result[i] = item.ToSlice(fields)
	}
	return result
}

func (l *CommentList) ToArray() []Comment {
	return l.list
}
