package repository

import (
	"database/sql"
	"time"
)

type Language struct {
	ID        uint   `gorm:"primarykey"`
	Name      string `gorm:"uniqueIndex"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type VideoQuality struct {
	ID        uint   `gorm:"primarykey"`
	Quality   string `gorm:"uniqueIndex"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AudioQuality struct {
	ID        uint   `gorm:"primarykey"`
	Quality   string `gorm:"uniqueIndex"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Event struct {
	ID             uint `gorm:"primarykey"`
	Name           string
	Description    sql.NullString
	Date           time.Time
	Languages      []Language     `gorm:"many2many:event_languages"`
	VideoQualities []VideoQuality `gorm:"many2many:event_video_qualities"`
	AudioQualities []AudioQuality `gorm:"many2many:event_audio_qualities"`
	EventInvitees  []EventInvitee
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type EventLanguage struct {
	ID         uint `gorm:"primarykey"`
	EventID    uint
	LanguageID uint
	Event      Event    `gorm:"foreignKey:EventID"`
	Language   Language `gorm:"foreignKey:LanguageID"`
	CreatedAt  time.Time
}

type EventVideoQuality struct {
	ID             uint `gorm:"primarykey"`
	EventID        uint
	VideoQualityID uint
	Event          Event        `gorm:"foreignKey:EventID"`
	VideoQuality   VideoQuality `gorm:"foreignKey:VideoQualityID"`
	CreatedAt      time.Time
}

type EventAudioQuality struct {
	ID             uint `gorm:"primarykey"`
	EventID        uint
	AudioQualityID uint
	Event          Event        `gorm:"foreignKey:EventID"`
	AudioQuality   AudioQuality `gorm:"foreignKey:AudioQualityID"`
	CreatedAt      time.Time
}

type EventInvitee struct {
	ID        uint `gorm:"primarykey"`
	EventID   uint
	Email     string
	Event     Event `gorm:"foreignKey:EventID"`
	CreatedAt time.Time
}
