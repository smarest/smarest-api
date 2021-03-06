package entity

import (
	"sort"
	"strings"

	"github.com/smarest/smarest-common/domain/value"
)

// display entity
type OrderDetail struct {
	ID            int64           `json:"id"`
	WaiterID      string          `json:"waiterID"`
	ChefID        *string         `json:"chefID"`
	OrderNumberID int64           `json:"orderNumberID"`
	TableID       int64           `json:"tableID"`
	ProductID     int64           `json:"productID"`
	Count         int64           `json:"count"`
	Comments      string          `json:"comments"`
	OrderTime     value.DateTime  `json:"orderTime"`
	FinishTime    *value.DateTime `json:"finishTime"`
	Status        bool            `json:"status"`
	Price         int64           `json:"price"`
	TableName     string          `json:"tableName"`
	ProductName   string          `json:"productName"`
}

func (item *OrderDetail) ToSlice(fields string) interface{} {
	if fields == "" {
		return *item
	}
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

func NewOrderDetail(order *Order, table *Table, product *Product) OrderDetail {
	return OrderDetail{
		ID:            order.ID,
		WaiterID:      order.WaiterID,
		ChefID:        order.ChefID,
		OrderNumberID: order.OrderNumberID,
		TableID:       order.TableID,
		ProductID:     order.ProductID,
		Count:         order.Count,
		Comments:      order.Comments,
		OrderTime:     order.OrderTime,
		FinishTime:    order.FinishTime,
		Status:        order.Status,
		Price:         order.Price,
		TableName:     table.Name,
		ProductName:   product.Name,
	}
}

type OrderDetailList struct {
	OrderDetails []OrderDetail
}

func NewOrderDetailList(items []OrderDetail) *OrderDetailList {
	return &OrderDetailList{items}
}

func CreateEmptyOrderDetailList() *OrderDetailList {
	return &OrderDetailList{make([]OrderDetail, 0)}
}

func (l *OrderDetailList) Add(item OrderDetail) {
	l.OrderDetails = append(l.OrderDetails, item)
}

func (l *OrderDetailList) ToSlice(fields string) interface{} {
	if fields == "" {
		return l.OrderDetails
	}
	var result []interface{} = make([]interface{}, len(l.OrderDetails))
	for i, item := range l.OrderDetails {
		result[i] = item.ToSlice(fields)
	}
	return result
}

func (l *OrderDetailList) SortByProductID() *OrderDetailList {
	sort.Slice(l.OrderDetails, func(i, j int) bool { return l.OrderDetails[i].ProductID < l.OrderDetails[j].ProductID })
	return l
}

func (l *OrderDetailList) SortByProductName() *OrderDetailList {
	sort.Slice(l.OrderDetails, func(i, j int) bool {
		obj1 := l.OrderDetails[i]
		obj2 := l.OrderDetails[j]
		return strings.Compare(obj1.ProductName, obj2.ProductName) < 0 || obj1.ProductID < obj2.ProductID || obj1.OrderTime.Unix() < obj2.OrderTime.Unix()
	})
	return l
}

func (l *OrderDetailList) GroupByProductIDAndToSlice(fields string) interface{} {
	groupMap := make(map[int64][]interface{})
	for _, u := range l.OrderDetails {
		group, value := groupMap[u.ProductID]
		if !value {
			group = make([]interface{}, 0)
		}
		group = append(group, u.ToSlice(fields))
		groupMap[u.ProductID] = group
	}

	result := make([]interface{}, 0)
	for _, value := range groupMap {
		result = append(result, value)
	}
	return result
}

func (l *OrderDetailList) SortByOrderTime() *OrderDetailList {
	sort.Slice(l.OrderDetails, func(i, j int) bool { return l.OrderDetails[i].OrderTime.Unix() < l.OrderDetails[j].OrderTime.Unix() })
	return l
}

// display entity
type OrderDetailGroupingByProductID struct {
	ProductID    int64         `json:"productID"`
	OrderDetails []OrderDetail `json:"orderDetails"`
}
