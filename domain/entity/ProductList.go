package entity

type ProductList struct {
	Products []Product
}

func NewProductList(productList []Product) ProductList {
	return ProductList{productList}
}

func CreateEmptyProductList() ProductList {
	return ProductList{make([]Product, 0)}
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
