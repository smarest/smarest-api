package persistence

import (
	"database/sql"
	"strconv"
	"strings"

	"github.com/smarest/smarest-api/domain/entity"
	"github.com/smarest/smarest-api/domain/repository"
	"github.com/smarest/smarest-common/infrastructure/persistence"
	"gopkg.in/gorp.v3"
)

type OrderRepositoryImpl struct {
	*persistence.DAOImpl
}

func updateFilter(col *gorp.ColumnMap) bool {
	return col.ColumnName != "order_number_id"
}

func NewOrderRepository(dbMap *gorp.DbMap) repository.OrderRepository {
	dbMap.AddTableWithName(entity.Order{}, "order").SetKeys(true, "id")
	dbMap.AddTableWithName(entity.OrderNumber{}, "order_number").SetKeys(true, "id")
	return &OrderRepositoryImpl{persistence.NewDAOImpl("`order`", dbMap)}
}

func (r *OrderRepositoryImpl) FindByAreaIDAndGroupByOrderNumberID(id int64) ([]entity.OrderGroupByOrderNumberID, error) {
	var items []entity.OrderGroupByOrderNumberID
	_, err := r.DAOImpl.Select(&items, "SELECT order_number_id,GROUP_CONCAT(DISTINCT(name)) as table_name, sum(count) as count_sum,sum(price) as price_sum FROM (SELECT o.order_number_id,t.name,o.count,price*count AS price,o.order_time FROM "+r.Table+" o INNER JOIN `table` t ON t.id = o.table_id WHERE t.area_id=?) AS t GROUP BY order_number_id ORDER BY order_time DESC", id)

	if err != nil {
		return nil, err
	}

	return items, nil
}

func (r *OrderRepositoryImpl) FindByOrderNumberID(orderNumberID int64) (*entity.OrderList, error) {
	var items []entity.Order
	_, err := r.DAOImpl.Select(&items, "SELECT * FROM "+r.Table+" WHERE order_number_id=? ORDER BY order_time DESC", orderNumberID)
	return entity.NewOrderList(items), err
}

func (r *OrderRepositoryImpl) FindByRestaurantID(restaurantID int64) (*entity.OrderList, error) {
	var items []entity.Order
	_, err := r.DAOImpl.Select(&items, "SELECT o.* FROM `order_number` n LEFT JOIN "+r.Table+" o ON n.id=o.order_number_id WHERE n.status=1 AND o.status=0 AND n.restaurant_id=?", restaurantID)
	return entity.NewOrderList(items), err
}

func (r *OrderRepositoryImpl) RegisterOrder(order entity.Order) (int64, error) {
	//sql := "INSERT INTO " + r.Table
	//sql += "(waiter_id,chef_id,order_number_id,table_id,product_id,count,comments,order_time,status,price) VALUES "
	//sql += "(?,?,?,?,?,?,?,now(),?,?)"

	//return r.DAOImpl.InsertBySQL(sql, order.WaiterID, order.ChefID, order.OrderNumberID, order.TableID, order.ProductID, order.Count, order.Comments, order.Status, order.Price)
	return order.ID, r.DAOImpl.Insert(&order)
}

func (r *OrderRepositoryImpl) UpdateOrder(order entity.Order) (int64, error) {
	return r.DAOImpl.UpdateColumns(updateFilter, &order)
}

func (r *OrderRepositoryImpl) DeleteByOrderNumberIDAndIDNotIn(orderNumberID int64, ids []int64) (int64, error) {
	sqlQuery := "DELETE FROM " + r.Table + " WHERE order_number_id=?"
	if len(ids) > 0 {
		idStrings := make([]string, len(ids))
		for i, id := range ids {
			idStrings[i] = strconv.FormatInt(id, 10)
		}
		sqlQuery += " AND id NOT IN (" + strings.Join(idStrings, ",") + ")"
	}

	return r.DAOImpl.DeleteBySQL(sqlQuery, orderNumberID)
}

func (r *OrderRepositoryImpl) RegisterOrderNumber(restaurantID int64) (*entity.OrderNumber, error) {
	var orderNumber *entity.OrderNumber
	err := r.DbMap.SelectOne(&orderNumber, "SELECT * FROM `order_number` WHERE status=0 ORDER BY id ASC LIMIT 1")
	if err != nil {
		if err == sql.ErrNoRows {
			orderNumber = &entity.OrderNumber{RestaurantID: restaurantID, Status: true}
			//return r.DAOImpl.InsertBySQL("INSERT INTO `order_number` (status) VALUES(?)", true)
			return orderNumber, r.DAOImpl.Insert(orderNumber)
		}
		return nil, err
	}
	//return r.DAOImpl.UpdateBySQL("UPDATE `order_number` SET status=? WHERE id=?", true, orderNumber.ID)
	orderNumber.Status = true
	_, err = r.DAOImpl.Update(orderNumber)
	return orderNumber, err
}
func (r *OrderRepositoryImpl) DeleteOrderNumber(orderNumber *entity.OrderNumber) (int64, error) {
	orderNumber.Status = false
	return r.DAOImpl.Update(orderNumber)
	//return r.DAOImpl.Delete("UPDATE `order_number` SET status=0 WHERE id=?", orderNumberID)
}

func (r *OrderRepositoryImpl) FindOrderNumber(orderNumberID int64) (*entity.OrderNumber, error) {
	var result entity.OrderNumber
	return &result, r.DbMap.SelectOne(&result, "SELECT * FROM `order_number` WHERE id=?", orderNumberID)
}
