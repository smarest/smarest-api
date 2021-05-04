package service

import (
	"database/sql"
	"fmt"

	"github.com/smarest/smarest-api/domain/entity"
	"github.com/smarest/smarest-api/domain/repository"
	"github.com/smarest/smarest-common/domain/entity/exception"
)

type TableService struct {
	AreaRepository  repository.AreaRepository
	TableRepository repository.TableRepository
}

func NewTableService(areaRepository repository.AreaRepository, tableRepository repository.TableRepository) *TableService {
	return &TableService{areaRepository, tableRepository}
}

func (s *TableService) FindAvailableByRestaurantIDAndAreaID(restaurantID int64, areaID int64) (*entity.TableList, *exception.Error) {
	area, aErr := s.AreaRepository.FindAvailableByID(areaID)
	if aErr != nil {
		if aErr == sql.ErrNoRows {
			return nil, exception.CreateError(exception.CodeNotFound, fmt.Sprintf("AreaID=[%d] not found. ", areaID))
		}
		return nil, exception.CreateErrorWithRootCause(exception.CodeSystemError, fmt.Sprintf("Can not get AreaID=[%d]", areaID), aErr)
	}

	if area.RestaurantID != restaurantID {
		return nil, exception.CreateError(exception.CodeValueInvalid, fmt.Sprintf("AreaID=[%d] not belong to RestaurantID=[%d]. ", areaID, restaurantID))
	}

	var result, err = s.TableRepository.FindAvailableByAreaID(areaID)
	if err != nil {
		return nil, exception.CreateErrorWithRootCause(exception.CodeSystemError, fmt.Sprintf("Can not get Tables. areaID=[%d]", areaID), err)
	}
	return result, nil
}
