package resource

type OrderRequestResource struct {
	RestaurantID  *int64                     `json:"restaurantID"`
	OrderNumberID *int64                     `json:"orderNumberID"`
	WaiterID      *string                    `json:"waiterID"`
	Data          []OrderRequestResourceData `json:"data"`
}

type OrderRequestResourceData struct {
	ID        int64  `json:"id"`
	TableID   int64  `json:"tableID"`
	ProductID int64  `json:"productID"`
	Comments  string `json:"comments"`
}
