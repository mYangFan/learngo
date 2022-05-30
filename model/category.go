package model

const TableCategory string = "category"

type Category struct {
	Id          int
	CategoryId  int    `gorm:"column:categoryId"`
	ParentId    int    `gorm:"column:parentId"`
	Region      string `gorm:"column:region"`
	CountryCode string `gorm:"column:countryCode"`
	Name        string `gorm:"column:name"`
	Target      int    `gorm:"column:target"`
}

func (Category) TableName() string {
	return TableCategory
}
