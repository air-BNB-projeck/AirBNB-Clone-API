package service

import (
	"alta/air-bnb/features/stays"
	"errors"

	"github.com/go-playground/validator/v10"
)

type StayService struct {
	data stays.StayDataInterface
	validator *validator.Validate
}

// AddStay implements stays.StayServiceInterface
func (service *StayService) AddStay(stayData stays.CoreStayRequest) (stayId string, err error) {
	if errValidate := service.validator.Struct(&stayData); errValidate != nil {
		return "", errors.New("validation error: " + errValidate.Error())
	}
	stayId, errInsert := service.data.Insert(stayData)
	if errInsert != nil {
		return "", errInsert
	}
	return stayId, nil
}

// GetAllStays implements stays.StayServiceInterface
func (service *StayService) GetAllStays() (stays []stays.Core, err error) {
	allStays, errGetAll := service.data.SelectAll()
	if errGetAll != nil {
		return nil, errGetAll
	}
	return allStays, nil 
}

// GetStay implements stays.StayServiceInterface
func (service *StayService) GetStay(stayId string) (stay stays.Core, err error) {
	stayData, errGet := service.data.Select(stayId);
	if errGet != nil {
		return stays.Core{}, errGet
	}
	return stayData, nil
}

// EditStay implements stays.StayServiceInterface
func (service *StayService) EditStay(stayId string, stayData stays.CoreStayRequest) error {
	if errValidate := service.validator.Struct(&stayData); errValidate != nil {
		return errors.New("validation error: " + errValidate.Error())
	}
	if errUpdate := service.data.Update(stayId, stayData); errUpdate != nil {
		return errUpdate
	}
	return nil
}

// DeleteStay implements stays.StayServiceInterface
func (service *StayService) DeleteStay(stayId string) error {
	if errDelete := service.data.Delete(stayId); errDelete != nil {
		return errDelete
	}
	return nil
}

func New(data stays.StayDataInterface) stays.StayServiceInterface {
	return &StayService{
		data: data,
		validator: validator.New(),
	}
}
