package entity

import "github.com/smarest/smarest-common/domain/value"

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

type TableList struct {
	Tables []Table
}

func NewTableList(tableList []Table) TableList {
	return TableList{tableList}
}

func CreateEmptyTableList() TableList {
	return TableList{make([]Table, 0)}
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
