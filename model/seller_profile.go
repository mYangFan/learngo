package model

const TableSellerProfile string = "seller_profile"

type SellerProfile struct {
	Id           int
	TokenId      int `gorm:"column:tokenId"`
	Region       string
	ProfileId    int    `gorm:"column:profileId"`
	CountryCode  string `gorm:"column:countryCode"`
	CurrencyCode string `gorm:"column:currencyCode"`
}

func (SellerProfile) TableName() string {
	return TableSellerProfile
}
