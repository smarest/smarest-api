package entity

import (
	"strings"

	"github.com/smarest/smarest-common/domain/value"
)

type Table struct {
	ID              int64           `db:"id" json:"id"`
	AreaID          int64           `db:"area_id" json:"areaID"`
	Name            string          `db:"name" json:"name"`
	Description     *string         `db:"description" json:"description"`
	Available       bool            `db:"available" json:"available"`
	Creator         string          `db:"creator" json:"creator"`
	CreatedDate     value.DateTime  `db:"created_date" json:"createdDate"`
	Updater         *string         `db:"updater" json:"updater"`
	LastUpdatedDate *value.DateTime `db:"last_updated_date" json:"lastUpdatedDate"`
}

func (item *Table) ToSlice(fields string) interface{} {
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
		case "description":
			result[field] = item.Description
		case "areaID":
			result[field] = item.AreaID
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

type TableList struct {
	Tables []Table
}

func NewTableList(tableList []Table) *TableList {
	return &TableList{tableList}
}

func CreateEmptyTableList() *TableList {
	return &TableList{make([]Table, 0)}
}

func (l *TableList) Add(table Table) {
	l.Tables = append(l.Tables, table)
}

func (l *TableList) FindByID(tableID int64) *Table {
	for _, table := range l.Tables {
		if table.ID == tableID {
			return &table
		}
	}
	return nil
}

func (l *TableList) GetAvailable() *TableList {
	return l.FilterBy(func(item Table) bool { return item.Available })
}

func (l *TableList) FilterBy(filter TableFilter) *TableList {
	list := make([]Table, 0)
	for _, u := range l.Tables {
		if filter(u) {
			list = append(list, u)
		}
	}
	return NewTableList(list)
}

type TableFilter func(item Table) bool

func (l *TableList) ToSlice(fields string) []interface{} {
	var result []interface{} = make([]interface{}, len(l.Tables))
	for i, item := range l.Tables {
		result[i] = item.ToSlice(fields)
	}
	return result
}

func (l *TableList) ToArray() []Table {
	return l.Tables
}
