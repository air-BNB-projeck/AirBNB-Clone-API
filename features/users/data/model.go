package data

type Users struct {
	ID 				string  `gorm:"type:varchar(50);primaryKey"`
	FullName	string 	`gorm:"type:varchar(100);notNull"`
	Email			string 	`gorm:"type:varchar(50);unique:notNull"`
	Password 	string 	`gorm:"type:varchar(50);notNull"`
	Phone			string	`gorm:"type:varchar(50);unique"`
	Birth			string	`gorm:"type:date"`
	Address		string	`gorm:"type:text"`
}