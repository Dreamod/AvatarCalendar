package media

import (
	string2 "aiCalendarApp/pkg/string"
	"time"

	"gorm.io/gorm"
)

type Media struct {
	gorm.Model
	Id      int    `json:"id"`
	Src     string `json:"src"`
	Preview string `json:"preview"`
	Type    string `json:"type:status"`
	DayId   string `json:"day_id" gorm:"index"`
}

func NewMedia(date time.Time, src string, preview string, mediaType string) *Media {
	media := &Media{
		Src:     src,
		Preview: preview,
		Type:    mediaType,
	}
	media.GenerateDayId(date)
	return media
}

func (media *Media) GenerateDayId(date time.Time) {
	media.DayId = string2.GetMD5Hash(date.Format("02-01-2006"))
}
