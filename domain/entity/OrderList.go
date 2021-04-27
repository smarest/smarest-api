package entity

type OrderList struct {
	Orders []Order
}

func NewOrderList(orderList []Order) OrderList {
	return OrderList{orderList}
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

func (l *OrderList) FindByID(id int64) *Order {
	for _, item := range l.Orders {
		if item.ID == id {
			return &item
		}
	}
	return nil
}

func (l *OrderList) FilterByNotIn(ids []int64) OrderList {
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


type OrderDetailList struct {
	OrderDetails []OrderDetail
}

func NewOrderDetailList(items []OrderDetail) OrderDetailList {
	return OrderDetailList{items}
}

func (l *OrderDetailList) ToSlice(fields string) []interface{} {
	var result []interface{} = make([]interface{}, len(l.OrderDetails))
	for i, item := range l.OrderDetails {
		result[i] = item.ToSlice(fields)
	}
	return result
}

func (l *OrderDetailList) ToArray() []OrderDetail {
	return l.OrderDetails
}
