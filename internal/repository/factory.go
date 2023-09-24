package repository

import "gorm.io/gorm"

type Repository[T any] interface {
	List() ([]T, error)
	Add(e *T) error
	Remove(e *T) error
}

type EventRepository Repository[Event]
type LanguageRepository Repository[Language]
type VideoQualityRepository Repository[VideoQuality]
type AudioQualityRepository Repository[AudioQuality]

// RepositoryFactory builds model repositories
type RepositoryFactory struct {
	db *gorm.DB
}

func NewRepositoryFactory(db *gorm.DB) *RepositoryFactory {
	return &RepositoryFactory{db}
}

func (f *RepositoryFactory) Events() EventRepository {
	return &repository[Event]{db: f.db}
}

func (f *RepositoryFactory) Languages() LanguageRepository {
	return &repository[Language]{db: f.db}
}

func (f *RepositoryFactory) VideoQualities() VideoQualityRepository {
	return &repository[VideoQuality]{db: f.db}
}

func (f *RepositoryFactory) AudioQualities() AudioQualityRepository {
	return &repository[AudioQuality]{db: f.db}
}
