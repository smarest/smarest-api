package entity

import (
	"strings"

	"github.com/smarest/smarest-common/domain/value"
)

type Order struct {
	ID            int64           `db:"id" json:"id"`
	WaiterID      string          `db:"waiter_id" json:"waiterID"`
	ChefID        *string         `db:"chef_id" json:"chefID"`
	OrderNumberID int64           `db:"order_number_id" json:"orderNumberID"`
	TableID       int64           `db:"table_id" json:"tableID"`
	ProductID     int64           `db:"product_id" json:"productID"`
	Count         int64           `db:"count" json:"count"`
	Comments      string          `db:"comments" json:"comments"`
	OrderTime     value.DateTime  `db:"order_time" json:"orderTime"`
	FinishTime    *value.DateTime `db:"finish_time" json:"finishTime"`
	Status        bool            `db:"status" json:"status"`
	Price         int64           `db:"price" json:"price"`
}

func (item *Order) ToSlice(fields string) map[string]interface{} {
	result := make(map[string]interface{})
	// Loop over the parts from the string.
	for _, field := range strings.Split(fields, ",") {

		switch field {
		case "id":
			result[field] = item.ID
		case "waiterID":
			result[field] = item.WaiterID
		case "chefID":
			result[field] = item.ChefID
		case "orderNumberID":
			result[field] = item.OrderNumberID
		case "tableID":
			result[field] = item.TableID
		case "productID":
			result[field] = item.ProductID
		case "count":
			result[field] = item.Count
		case "comments":
			result[field] = item.Comments
		case "orderTime":
			result[field] = item.OrderTime
		case "finishTime":
			result[field] = item.FinishTime
		case "status":
			result[field] = item.Status
		case "price":
			result[field] = item.Price
		default:
		}
	}
	return result
}

type OrderNumber struct {
	ID     int64 `db:"id"`
	Status bool  `db:"status"`
}

// display entity
type OrderDetail struct {
	ID            int64           `db:"id" json:"id"`
	WaiterID      string          `db:"waiter_id" json:"waiterID"`
	ChefID        *string         `db:"chef_id" json:"chefID"`
	OrderNumberID int64           `db:"order_number_id" json:"orderNumberID"`
	TableID       int64           `db:"table_id" json:"tableID"`
	ProductID     int64           `db:"product_id" json:"productID"`
	Count         int64           `db:"count" json:"count"`
	Comments      string          `db:"comments" json:"comments"`
	OrderTime     value.DateTime  `db:"order_time" json:"orderTime"`
	FinishTime    *value.DateTime `db:"finish_time" json:"finishTime"`
	Status        bool            `db:"status" json:"status"`
	Price         int64           `db:"price" json:"price"`
	//innerJoin
	TableName   *string `db:"table_name" json:"tableName"`
	ProductName *string `db:"product_name" json:"productName"`
}

func (item *OrderDetail) ToSlice(fields string) map[string]interface{} {
	result := make(map[string]interface{})
	// Loop over the parts from the string.
	for _, field := range strings.Split(fields, ",") {

		switch field {
		case "id":
			result[field] = item.ID
		case "waiterID":
			result[field] = item.WaiterID
		case "chefID":
			result[field] = item.ChefID
		case "orderNumberID":
			result[field] = item.OrderNumberID
		case "tableID":
			result[field] = item.TableID
		case "productID":
			result[field] = item.ProductID
		case "count":
			result[field] = item.Count
		case "comments":
			result[field] = item.Comments
		case "orderTime":
			result[field] = item.OrderTime
		case "finishTime":
			result[field] = item.FinishTime
		case "status":
			result[field] = item.Status
		case "price":
			result[field] = item.Price
		case "tableName":
			result[field] = item.TableName
		case "productName":
			result[field] = item.ProductName
		default:
		}
	}
	return result
}

type OrderGroupByOrderNumberID struct {
	OrderNumberID int64  `db:"order_number_id" json:"orderNumberID"`
	TableName     string `db:"table_name" json:"tableName"`
	CountSum      int64  `db:"count_sum" json:"countSum"`
	PriceSum      int64  `db:"price_sum" json:"priceSum"`
}
