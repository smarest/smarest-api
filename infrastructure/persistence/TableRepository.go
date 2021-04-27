package persistence

import (
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

func (r *TableRepositoryImpl) FindByID(id int64) (*entity.Table, error) {
	var item entity.Table
	return &item, r.DAOImpl.FindByID(id, &item)
}
func (r *TableRepositoryImpl) FindByAreaID(areaID int64) ([]entity.Table, error) {
	var items []entity.Table
	_, err := r.DbMap.Select(&items, "SELECT * FROM "+r.Table+" WHERE area_id=?", areaID)
	return items, err

}
func (r *TableRepositoryImpl) FindAll() ([]entity.Table, error) {
	var items []entity.Table
	_, err := r.DAOImpl.FindAll(&items)
	return items, err
}
func (r *TableRepositoryImpl) FindByIDs(tableIDs []int64) (entity.TableList, error) {
	var items []entity.Table
	tableIDStrings := make([]string, len(tableIDs))
	for i, tableID := range tableIDs {
		tableIDStrings[i] = strconv.FormatInt(tableID, 10)
	}
	_, err := r.DAOImpl.Select(&items, "SELECT * FROM "+r.Table+" WHERE id IN ("+strings.Join(tableIDStrings, ",")+")")
	return entity.NewTableList(items), err
}
