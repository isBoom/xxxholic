package model

var VideoTypes  = [...]string {"anime","movie","funny","other"}
type VideoType struct {
	ID        uint `gorm:"primary_key"`
	VideoType string `json:"videoType" gorm:"unique varchar(20)"`
}