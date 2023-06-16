package stays

import "mime/multipart"

type Core struct {
	ID 						string 			`json:"id" form:"id"`
	Name 					string			`json:"name" form:"name"`
	Price					float64			`json:"price" form:"price"`
	Description		string			`json:"description" form:"description"`
	Latitude			float64			`json:"latitude" form:"latitude"`
	Longitude			float64			`json:"longitude" form:"longitude"`
	City					string			`json:"city" form:"city"`
	Bedrooms			int			`json:"bedrooms" form:"bedrooms"`
	Bathrooms			int			`json:"bathrooms" form:"bathrooms"`
	TV						int			`json:"tv" form:"tv"`
	Wifi					int			`json:"wifi" form:"wifi"`
	Pool					int			`json:"pool" form:"pool"`
	Rating				float64			`json:"rating" form:"rating"`
	User					Users 	`json:"owner" form:"owner"`
	StayImages		[]string `json:"images" form:"images"`
	StayReviews 	[]StayReviews `json:"reviews" form:"reviews"`
}

type Users struct {
	ID						string			`json:"id" form:"id"`
	FullName			string			`json:"fullname" form:"fullname"`
	Email					string			`json:"email" form:"email"`
}

type StayReviews struct {
	User 					Users 			`json:"user" form:"user"`
	Review				string			`json:"review" form:"review"`
	Rating				uint			`json:"rating" form:"rating"`
}

type CoreStayRequest struct {
	Name 					string			`json:"name" form:"name" validate:"required"`
	Price					float64			`json:"price" form:"price" validate:"required"`
	Description		string			`json:"description" form:"description"`
	Latitude			float64			`json:"latitude" form:"latitude"`
	Longitude			float64			`json:"longitude" form:"longitude"`
	City					string			`json:"city" form:"city"`
	Bedrooms			int			`json:"bedrooms" form:"bedrooms"`
	Bathrooms			int			`json:"bathrooms" form:"bathrooms"`
	TV						int			`json:"tv" form:"tv"`
	Wifi					int			`json:"wifi" form:"wifi"`
	Pool					int			`json:"pool" form:"pool"`
	Rating				float64			`json:"rating" form:"rating"`
	UserID				uint				`json:"userId" form:"userId"`	
	Image					*multipart.FileHeader			`form:"image"`
	ImageURI			string
}

type CoreStayImageRequest struct {
	Image					*multipart.FileHeader 			`form:"image" validate:"required"`
}

type CoreStayReviewRequest struct {
	UserID				uint				
	StayID				string
	Review				string						`json:"review" form:"review" validate:"required"`
	Rating				uint							`json:"rating" form:"review" validate:"required"`
}

type StayDataInterface interface {
	Insert(stayData CoreStayRequest) (stayId string, err error)
	Select(stayId string) (stay Core, err error)
	SelectAll() (allStays []Core, err error)
	Update(stayId string, stayData CoreStayRequest) error
	Delete(stayId string) error
	InsertStayImage(stayId string, imageUrl string) error
	DeleteStayImage(stayId string, imageId uint) error
	InsertStayReview(reviewData CoreStayReviewRequest) error
	DeleteStayReview(stayId string, reviewId uint) error
	GetStayImageURI(stayId string, imageId uint) (imageUrl string)
}

type StayServiceInterface interface {
	AddStay(stayData CoreStayRequest) (stayId string, err error)
	GetStay(stayId string) (stay Core, err error)
	GetAllStays() (stays []Core, err error)
	EditStay(stayId string, stayData CoreStayRequest) error
	DeleteStay(stayId string) error
	AddStayImage(stayId string, image CoreStayImageRequest) error
	DeleteStayImage(stayId string, imageId uint) error
	AddStayReview(reviewData CoreStayReviewRequest) error
}