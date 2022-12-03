package models

type User struct {
	Name              string `json:"name,omitempty"`
	WalletAddress     string `gorm:"primary_key" json:"walletAddress"`
	ProfilePictureUrl string `json:"profilePictureUrl,omitempty"`
	Country           string `json:"country,omitempty"`
}
