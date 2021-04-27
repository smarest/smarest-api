package entity

type OrderRequest struct {
	RestaurantID        *int64               `json:"restaurantID"`
	OrderNumberID       *int64               `json:"orderNumberID"`
	WaiterID            *string              `json:"waiterID"`
	OrderRequestRecords []OrderRequestRecord `json:"data"`
}

type OrderRequestRecord struct {
	ID        *int64 `json:"id"`
	TableID   int64  `json:"tableID"`
	ProductID int64  `json:"productID"`
	Count     int64  `json:"count"`
	Comments  string `json:"comments"`
}

func (l *OrderRequest) GetDistinctProductIDs() []int64 {
	keys := make(map[int64]bool)
	result := make([]int64, 0)
	for _, u := range l.OrderRequestRecords {
		if _, value := keys[u.ProductID]; !value {
			result = append(result, u.ProductID)
			keys[u.ProductID] = true
		}
	}
	return result
}

func (l *OrderRequest) GetDistinctTableIDs() []int64 {
	keys := make(map[int64]bool)
	result := make([]int64, 0)
	for _, u := range l.OrderRequestRecords {
		if _, value := keys[u.TableID]; !value {
			result = append(result, u.TableID)
			keys[u.TableID] = true
		}
	}
	return result
}

func (l *OrderRequest) Discover() ([]OrderRequestRecord, []OrderRequestRecord, []int64) {
	keys := make(map[int64]bool)
	nonNullIDresult := make([]int64, 0)
	nonNullresult := make([]OrderRequestRecord, 0)
	nullResult := make([]OrderRequestRecord, 0)
	for _, u := range l.OrderRequestRecords {
		if u.ID == nil {
			nullResult = append(nullResult, u)
			continue
		}
		if _, duplicate := keys[*u.ID]; !duplicate {
			nonNullresult = append(nonNullresult, u)
			nonNullIDresult = append(nonNullIDresult, *u.ID)
			keys[*u.ID] = true
		}
	}
	return nullResult, nonNullresult, nonNullIDresult
}

func (l *OrderRequest) HasOrder(orderId int64) bool {
	for _, u := range l.OrderRequestRecords {
		if u.ID == nil {
			continue
		}
		if *u.ID == orderId {
			return true
		}
	}
	return false
}
