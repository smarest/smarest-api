package persistence

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/smarest/smarest-api/domain/entity"
	"github.com/smarest/smarest-api/domain/repository"

	"github.com/smarest/smarest-common/infrastructure/persistence"
	"gopkg.in/gorp.v3"
)

type TableRepositoryImpl struct {
	*persistence.DAOImpl
}

func NewTableRepository(dbMap *gorp.DbMap) repository.TableRepository {
	return &TableRepositoryImpl{persistence.NewDAOImpl("`table`", dbMap)}
}

func (r *TableRepositoryImpl) FindAvailableByAreaID(areaID int64) (*entity.TableList, error) {
	var items []entity.Table
	_, err := r.DbMap.Select(&items, "SELECT * FROM "+r.Table+" WHERE area_id=? AND available=1", areaID)
	return entity.NewTableList(items), err

}
func (r *TableRepositoryImpl) FindAvailableByRestaurantIDAndIDs(restID int64, tableIDs []int64) (*entity.TableList, error) {
	if len(tableIDs) == 0 {
		return entity.CreateEmptyTableList(), nil
	}
	var items []entity.Table
	tableIDStrings := make([]string, len(tableIDs))
	for i, tableID := range tableIDs {
		tableIDStrings[i] = strconv.FormatInt(tableID, 10)
	}
	_, err := r.DbMap.Select(&items, "SELECT t.id,t.area_id,t.name,t.description,t.available,t.creator,t.created_date,t.updater,t.last_updated_date FROM "+r.Table+" as t LEFT JOIN `area` as a ON t.area_id=a.id WHERE a.restaurant_id=? AND t.id IN("+strings.Join(tableIDStrings, ",")+") AND t.available=1 AND a.available=1 ", restID)
	return entity.NewTableList(items), err
}

func (r *TableRepositoryImpl) FindByIDs(tableIDs []int64) (*entity.TableList, error) {
	if len(tableIDs) == 0 {
		return entity.CreateEmptyTableList(), nil
	}
	var items []entity.Table
	tableIDStrings := make([]string, len(tableIDs))
	for i, tableID := range tableIDs {
		tableIDStrings[i] = fmt.Sprint(tableID)
	}
	_, err := r.DbMap.Select(&items, "SELECT * FROM "+r.Table+" WHERE id IN("+strings.Join(tableIDStrings, ",")+")")
	return entity.NewTableList(items), err

}
