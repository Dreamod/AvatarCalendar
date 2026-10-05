package day

import (
	"aiCalendarApp/internal/calendar/media"
	string2 "aiCalendarApp/pkg/string"
	"time"

	"gorm.io/gorm"
)

type Day struct {
	gorm.Model
	Date        time.Time     `json:"date"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Id          string        `json:"id" gorm:"uniqueIndex,primaryKey"`
	Media       []media.Media `json:"medias"`
}

func NewDay(date time.Time, title string, description string) *Day {
	day := &Day{
		Date:        date,
		Title:       title,
		Description: description,
	}
	day.GenerateId()
	return day
}

func (day *Day) GenerateId() {
	day.Id = string2.GetMD5Hash(day.Date.Format("02-01-2006"))
}
