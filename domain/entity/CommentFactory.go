package entity

import (
	"strings"
)

type CommentFactory struct{}

func NewCommentFactory() CommentFactory {
	return CommentFactory{}
}

func (f *CommentFactory) CreateList(items []Comment, fields string) []interface{} {
	var result []interface{} = make([]interface{}, len(items))
	for i, item := range items {
		result[i] = f.Create(&item, fields)
	}
	return result
}

func (f *CommentFactory) Create(item *Comment, fields string) map[string]interface{} {
	result := make(map[string]interface{})
	// Loop over the parts from the string.
	for _, field := range strings.Split(fields, ",") {
		switch field {
		case "id":
			result[field] = item.ID
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
