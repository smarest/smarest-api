package service

import (
	"fmt"

	"github.com/smarest/smarest-api/domain/entity"
	"github.com/smarest/smarest-api/domain/repository"
	"github.com/smarest/smarest-common/domain/entity/exception"
)

type AreaService struct {
	AreaRepository repository.AreaRepository
}

func NewAreaService(areaRepository repository.AreaRepository) *AreaService {
	return &AreaService{areaRepository}
}

func (s *AreaService) FindAvailableByRestaurantID(restID int64) (*entity.AreaList, *exception.Error) {
	var results, err = s.AreaRepository.FindAvailableByRestaurantID(restID)
	if err != nil {
		return nil, exception.CreateErrorWithRootCause(exception.CodeSystemError, fmt.Sprintf("Can not get Areas. restaurantID=[%d]", restID), err)
	}

	return results, nil
}
