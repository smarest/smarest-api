package entity

type OrderList struct {
	Orders []Order
}

func NewOrderList(orderList []Order) *OrderList {
	return &OrderList{orderList}
}

func CreateEmptyOrderList() OrderList {
	return OrderList{make([]Order, 0)}
}

func (l *OrderList) ToSlice(fields string) []interface{} {
	var result []interface{} = make([]interface{}, len(l.Orders))
	for i, item := range l.Orders {
		result[i] = item.ToSlice(fields)
	}
	return result
}

func (l *OrderList) Add(item Order) {
	l.Orders = append(l.Orders, item)
}

func (l *OrderList) IsEmpty() bool {
	return len(l.Orders) == 0
}

func (l *OrderList) FindByID(id int64) *Order {
	for _, item := range l.Orders {
		if item.ID == id {
			return &item
		}
	}
	return nil
}

func (l *OrderList) FilterByIDNotIn(ids []int64) OrderList {
	result := CreateEmptyOrderList()
	for _, order := range l.Orders {
		if _, contain := l.isArrayContain(ids, order.ID); !contain {
			result.Add(order)
		}
	}
	return result
}

func (l *OrderList) isArrayContain(array []int64, val int64) (int, bool) {
	for i, item := range array {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func (l *OrderList) GetDistinctProductIDsAndTableIDs() ([]int64, []int64) {
	productKeys := make(map[int64]bool)
	tableKeys := make(map[int64]bool)
	productResult := make([]int64, 0)
	tableResult := make([]int64, 0)
	for _, u := range l.Orders {
		if _, value := productKeys[u.ProductID]; !value {
			productResult = append(productResult, u.ProductID)
			productKeys[u.ProductID] = true
		}
		if _, value := tableKeys[u.TableID]; !value {
			tableResult = append(tableResult, u.TableID)
			tableKeys[u.TableID] = true
		}
	}
	return productResult, tableResult
}
