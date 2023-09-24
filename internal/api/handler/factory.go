package handler

import (
	"fmt"

	"github.com/aerostatas/interaction-go-kata/internal/repository"
	"github.com/aerostatas/interaction-go-kata/internal/service"
)

type RepositoryFactory interface {
	Events() repository.EventRepository
	Languages() repository.LanguageRepository
	VideoQualities() repository.VideoQualityRepository
	AudioQualities() repository.AudioQualityRepository
}

// HandlerFactory builds the handlers for HTTP requests
type HandlerFactory struct {
	repositoryFactory RepositoryFactory
}

func NewHandlerFactory(repositoryFactory RepositoryFactory) *HandlerFactory {
	return &HandlerFactory{
		repositoryFactory,
	}
}

func (f *HandlerFactory) EventHandler() (*EventHandler, error) {
	eventRepository := f.repositoryFactory.Events()
	languageRepository := f.repositoryFactory.Languages()
	videoQualityRepository := f.repositoryFactory.VideoQualities()
	audioQualityRepository := f.repositoryFactory.AudioQualities()

	languageMap := make(map[string]uint)
	videoQualityMap := make(map[string]uint)
	audioQualityMap := make(map[string]uint)

	languages, err := languageRepository.List()
	if err != nil {
		return nil, fmt.Errorf("language list: %w", err)
	}

	videoQualities, err := videoQualityRepository.List()
	if err != nil {
		return nil, fmt.Errorf("video quality list: %w", err)
	}

	audioQualities, err := audioQualityRepository.List()
	if err != nil {
		return nil, fmt.Errorf("audio quality list: %w", err)
	}

	for _, language := range languages {
		languageMap[language.Name] = language.ID
	}

	for _, videoQuality := range videoQualities {
		videoQualityMap[videoQuality.Quality] = videoQuality.ID
	}

	for _, audioQuality := range audioQualities {
		audioQualityMap[audioQuality.Quality] = audioQuality.ID
	}

	eventService := service.NewEventService(
		eventRepository,
		languageMap,
		videoQualityMap,
		audioQualityMap,
		100,
	)

	return NewEventHandler(eventService), nil
}
