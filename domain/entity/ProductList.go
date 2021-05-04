package entity

type ProductList struct {
	Products []Product
}

func NewProductList(productList []Product) *ProductList {
	return &ProductList{productList}
}

func CreateEmptyProductList() *ProductList {
	return &ProductList{make([]Product, 0)}
}

func (l *ProductList) Add(product Product) {
	l.Products = append(l.Products, product)
}

func (l *ProductList) FindByID(productID int64) *Product {
	for _, product := range l.Products {
		if product.ID == productID {
			return &product
		}
	}
	return nil
}

func (l *ProductList) GetAvailable() *ProductList {
	return l.FilterBy(func(item Product) bool { return item.Available })
}

func (l *ProductList) FilterByCategoryID(categoryID int64) *ProductList {
	return l.FilterBy(func(item Product) bool { return item.CategoryID == categoryID })

}

func (l *ProductList) FilterBy(filter ProductFilter) *ProductList {
	list := make([]Product, 0)
	for _, u := range l.Products {
		if filter(u) {
			list = append(list, u)
		}
	}
	return NewProductList(list)
}

type ProductFilter func(item Product) bool

func (list *ProductList) ToSlice(fields string) []interface{} {
	var result []interface{} = make([]interface{}, len(list.Products))
	for i, item := range list.Products {
		result[i] = item.ToSlide(fields)
	}
	return result
}

func (l *ProductList) ToArray() []Product {
	return l.Products
}
