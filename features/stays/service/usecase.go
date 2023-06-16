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
	data      stays.StayDataInterface
	validator *validator.Validate
}

// DeleteStayImage implements stays.StayServiceInterface
func (service *StayService) DeleteStayImage(stayId string, imageId uint) error {
	// imageUrl := service.data.GetStayImageURI(stayId, imageId)
	// // _, errDeleteS3 := helper.UploaderS3().DeleteObject(&s3.DeleteObjectInput{
	// // 	Bucket: aws.String("alta-airbnb"),
	// // 	Key: aws.String(imageUrl),
	// // })
	// // if errDeleteS3 != nil {
	// // 	return errors.New("error delete object in s3: " + errDeleteS3.Error())
	// // }
	if err := service.data.DeleteStayImage(stayId, imageId); err != nil {
		return err
	}
	return nil
}

// AddStayReview implements stays.StayServiceInterface
func (service *StayService) AddStayReview(reviewData stays.CoreStayReviewRequest) error {
	if errValidate := service.validator.Struct(reviewData); errValidate != nil {
		return errors.New("error validation: " + errValidate.Error())
	}
	if err := service.data.InsertStayReview(reviewData); err != nil {
		return err
	}
	return nil
}

// AddStayImage implements stays.StayServiceInterface
func (service *StayService) AddStayImage(stayId string, image stays.CoreStayImageRequest) error {
	if errValidate := service.validator.Struct(image); errValidate != nil {
		return errors.New("error validation: " + errValidate.Error())
	}
	imageData, errGetImage := image.Image.Open()
	if errGetImage != nil {
		return errors.New("failed to open file: " + errGetImage.Error())
	}
	defer imageData.Close()
	imageKey := helper.GenerateNewId() + "_" + image.Image.Filename
	_, errUpload := helper.UploaderS3().PutObject(&s3.PutObjectInput{
		Bucket: aws.String("alta-airbnb"),
		Key:    aws.String(imageKey),
		Body:   imageData,
	})
	if errUpload != nil {
		return errors.New("failed to upload file image: " + errUpload.Error())
	}
	imageUrl := "https://alta-airbnb.s3.ap-southeast-3.amazonaws.com/" + imageKey
	if err := service.data.InsertStayImage(stayId, imageUrl); err != nil {
		return err
	}
	return nil
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
	imageKey := helper.GenerateNewId() + "_" + stayData.Image.Filename
	_, errUpload := helper.UploaderS3().PutObject(&s3.PutObjectInput{
		Bucket: aws.String("alta-airbnb"),
		Key:    aws.String(imageKey),
		Body:   image,
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
	stayData, errGet := service.data.Select(stayId)
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
		data:      data,
		validator: validator.New(),
	}
}
