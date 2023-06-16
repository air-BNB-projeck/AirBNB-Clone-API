package data

import (
	"alta/air-bnb/app/helper"
	"alta/air-bnb/features/stays"
	"fmt"

	"gorm.io/gorm"
)

type StayData struct {
	db *gorm.DB
}

// DeleteStayImage implements stays.StayDataInterface
func (repo *StayData) DeleteStayImage(stayId string, imageId uint) error {
	if tx := repo.db.Where("id = ? && stay_id = ?", imageId, stayId).Delete(&StayImages{}); tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (repo *StayData) GetStayImageURI(stayId string, imageId uint) (imageUri string) {
	var imageData StayImages
	if tx := repo.db.Select("image_url").Where("id = ? && stay_id = ?", imageId, stayId).First(&imageData); tx.Error != nil {
		return ""
	}
	return imageData.ImageUrl
}

// DeleteStayReview implements stays.StayDataInterface
func (*StayData) DeleteStayReview(stayId string, reviewId uint) error {
	panic("unimplemented")
}

// InsertStayReview implements stays.StayDataInterface
func (repo *StayData) InsertStayReview(stayReviewData stays.CoreStayReviewRequest) error {
	stayReviewDataMap := CoreReviewToModels(stayReviewData)
	if tx := repo.db.Create(&stayReviewDataMap); tx.Error != nil {
		return tx.Error
	}
	return nil
}

// InsertStayImage implements stays.StayDataInterface
func (repo *StayData) InsertStayImage(stayId string, imageUrl string) error {
	var stayImageData = StayImages{
		ImageUrl: imageUrl,
		StayID:   stayId,
	}
	if tx := repo.db.Create(&stayImageData); tx.Error != nil {
		return tx.Error
	}
	return nil
}

// Insert implements stays.StayDataInterface
func (repo *StayData) Insert(stayData stays.CoreStayRequest) (stayId string, err error) {
	var id = helper.GenerateNewId()
	var stayCoreMap = CoreStayToModel(stayData)
	stayCoreMap.ID = id
	if tx := repo.db.Create(&stayCoreMap); tx.Error != nil {
		return "", tx.Error
	}
	var stayImages = StayImages{
		ImageUrl: stayData.ImageURI,
		StayID:   stayCoreMap.ID,
	}
	if tx := repo.db.Create(&stayImages); tx.Error != nil {
		return "", tx.Error
	}
	return stayCoreMap.ID, nil
}

// SelectAll implements stays.StayDataInterface
func (repo *StayData) SelectAll() (allStays []stays.Core, err error) {
	var staysData []Stays
	if tx := repo.db.Preload("User").Preload("StaysImages").Preload("StayReviews").Find(&staysData); tx.Error != nil {
		return nil, tx.Error
	}
	var staysCoreMap []stays.Core
	for _, stay := range staysData {
		stayMap := ModelStayToCore(stay)
		fmt.Println(stayMap.StayReviews)
		stayMap.User = stays.Users{
			ID:       stay.User.ID,
			FullName: stay.User.FullName,
			Email:    stay.User.Email,
		}
		for _, stayImage := range stay.StaysImages {
			stayMap.StayImages = append(stayMap.StayImages, stayImage.ImageUrl)
		}
		for _, stayReview := range stay.StayReviews {
			stayMap.StayReviews = append(stayMap.StayReviews, stays.StayReviews{
				User: stays.Users(stayReview.User),
				Review: stayReview.Review,
				Rating: stayReview.Rating,
			})
		}
		staysCoreMap = append(staysCoreMap, stayMap)
	}
	return staysCoreMap, nil
}

// Select implements stays.StayDataInterface
func (repo *StayData) Select(stayId string) (stay stays.Core, err error) {
	var stayData Stays
	if tx := repo.db.Where("id = ?", stayId).Preload("User").Preload("StaysImages").First(&stayData); tx.Error != nil {
		return stays.Core{}, tx.Error
	}
	var stayDataMap = ModelStayToCore(stayData)
	stayDataMap.User = stays.Users{
		ID:       stayData.User.ID,
		FullName: stayData.User.FullName,
		Email:    stayData.User.Email,
	}
	for _, stayImage := range stayData.StaysImages {
		stayDataMap.StayImages = append(stayDataMap.StayImages, stayImage.ImageUrl)
	}
	return stayDataMap, nil
}

// Update implements stays.StayDataInterface
func (repo *StayData) Update(stayId string, stayData stays.CoreStayRequest) error {
	var stay Stays
	if tx := repo.db.Where("id = ?", stayId).First(&stay); tx.Error != nil {
		return tx.Error
	}
	var stayDataMap = CoreStayToModel(stayData)
	if tx := repo.db.Model(&stay).Updates(stayDataMap); tx.Error != nil {
		return tx.Error
	}
	return nil
}

// Delete implements stays.StayDataInterface
func (repo *StayData) Delete(stayId string) error {
	if tx := repo.db.Where("id = ?", stayId).Delete(&Stays{}); tx.Error != nil {
		return tx.Error
	}
	return nil
}

func New(db *gorm.DB) stays.StayDataInterface {
	return &StayData{
		db: db,
	}
}
