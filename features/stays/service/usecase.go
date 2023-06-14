package service

import (
	"alta/air-bnb/app/helper"
	"alta/air-bnb/features/stays"
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
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
	image, errGetImage := stayData.Image.Open()
	if errGetImage != nil {
		return "", errors.New("failed to open file: " + errGetImage.Error())
	}
	defer image.Close()
	imageKey := stayData.Image.Filename + "_" + helper.GenerateNewId()
	_, errUpload := helper.UploaderS3().PutObject(&s3.PutObjectInput{
		Bucket: aws.String("alta-airbnb"),
		Key: aws.String(imageKey),
		Body: image,
	})
	if errUpload != nil {
		return "", errors.New("failed to upload file image: " + errUpload.Error())
	}
	stayData.ImageURI = "https://alta-airbnb.s3.ap-southeast-3.amazonaws.com/" + imageKey
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
