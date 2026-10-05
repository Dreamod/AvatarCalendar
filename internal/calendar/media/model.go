package media

import (
	string2 "aiCalendarApp/pkg/string"
	"database/sql/driver"
	"time"

	"gorm.io/gorm"
)

type Media struct {
	gorm.Model
	Src     string `json:"date"`
	Preview string `json:"preview"`
	Type    Type   `json:"type:status"`
	DayId   string `json:"day_id" gorm:"index"`
}

type Type string

const (
	Image Type = "image"
	Video Type = "video"
)

func (p *Type) Scan(value interface{}) error {
	*p = Type(value.(byte))
	return nil
}

func (s Type) Value() (driver.Value, error) {
	return string(s), nil
}

func NewMedia(date time.Time, src string, mediaType Type) *Media {
	media := &Media{
		Src:     src,
		Type:    mediaType,
		Preview: src,
	}
	media.GenerateDayId(date)
	return media
}

func (media *Media) GenerateDayId(date time.Time) {
	media.DayId = string2.GetMD5Hash(date.Format("02-01-2006"))
}
