package repository

import (
	"fmt"

	"gorm.io/gorm"
)

var models = []interface{}{
	&Event{},
	&Language{},
	&VideoQuality{},
	&AudioQuality{},
	&EventLanguage{},
	&EventVideoQuality{},
	&EventAudioQuality{},
	&EventInvitee{},
}

var defaultLanguages = []*Language{
	{Name: "Lithuanian"},
	{Name: "French"},
	{Name: "English"},
}

var defaultVideoQualities = []*VideoQuality{
	{Quality: "360p"},
	{Quality: "720p"},
	{Quality: "1080p"},
}

var defaultAudioQualities = []*AudioQuality{
	{Quality: "Low"},
	{Quality: "Medium"},
	{Quality: "High"},
}

// Migrate runs DB schema migration
func Migrate(db *gorm.DB) error {
	for _, model := range models {
		if err := db.AutoMigrate(model); err != nil {
			return fmt.Errorf("DB migrate: %w", err)
		}
	}
	return nil
}

// Seed creates default DB data if tables are empty
func Seed(db *gorm.DB) error {
	if err := seedDefaults[Language](db, defaultLanguages); err != nil {
		return err
	}

	if err := seedDefaults[VideoQuality](db, defaultVideoQualities); err != nil {
		return err
	}

	if err := seedDefaults[AudioQuality](db, defaultAudioQualities); err != nil {
		return err
	}

	return nil
}

func seedDefaults[T any](db *gorm.DB, defaults []*T) error {
	if len(defaults) < 1 {
		return nil
	}

	var count int64

	if err := db.Model(defaults[0]).Count(&count).Error; err != nil {
		return fmt.Errorf("count check: %w", err)
	}

	if count < 1 {
		if err := db.Create(defaults).Error; err != nil {
			return fmt.Errorf("create defaults: %w", err)
		}
	}

	return nil
}
