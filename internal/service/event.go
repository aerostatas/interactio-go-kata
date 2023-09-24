package service

import (
	"fmt"
	"time"

	"github.com/aerostatas/interaction-go-kata/internal/repository"
	"gopkg.in/guregu/null.v3"
)

type EventRepository interface {
	Add(event *repository.Event) error
}

type EventService struct {
	eventRepository  EventRepository
	languageMap      map[string]uint
	videoQualityMap  map[string]uint
	audioQualityMap  map[string]uint
	maxEventInvitees uint
}

func NewEventService(
	eventRepository EventRepository,
	languageMap map[string]uint,
	videoQualityMap map[string]uint,
	audioQualityMap map[string]uint,
	maxEventInvitees uint,
) *EventService {
	return &EventService{
		eventRepository,
		languageMap,
		videoQualityMap,
		audioQualityMap,
		maxEventInvitees,
	}
}

type EventCreate struct {
	Name         string      `json:"name" validate:"required" example:"EU Summit"`
	Date         time.Time   `json:"date" validate:"required" example:"2023-04-20T14:00:00Z"`
	Languages    []string    `json:"languages" swaggertype:"array,string" validate:"required" example:"Lithuanian,French"`
	VideoQuality []string    `json:"videoQuality" swaggertype:"array,string" validate:"required" example:"720p,1080p"`
	AudioQuality []string    `json:"audioQuality" swaggertype:"array,string" validate:"required" example:"Low,High"`
	Invitees     []string    `json:"invitees" swaggertype:"array,string" validate:"required" example:"one@email.com,two@email.com"`
	Description  null.String `json:"description" swaggertype:"string" validate:"optional" example:"EU Summit 2023"`
}

type Event struct {
	ID           uint        `json:"id" example:"123"`
	Name         string      `json:"name" validate:"required" example:"EU Summit"`
	Date         time.Time   `json:"date" validate:"required" example:"2023-04-20T14:00:00Z"`
	Languages    []string    `json:"languages" swaggertype:"array,string" validate:"required" example:"Lithuanian,French"`
	VideoQuality []string    `json:"videoQuality" swaggertype:"array,string" validate:"required" example:"720p,1080p"`
	AudioQuality []string    `json:"audioQuality" swaggertype:"array,string" validate:"required" example:"Low,High"`
	Invitees     []string    `json:"invitees" swaggertype:"array,string" validate:"required" example:"one@email.com,two@email.com"`
	Description  null.String `json:"description" swaggertype:"string" validate:"optional" example:"EU Summit 2023"`
}

// CreateEvent validates provided event data and creates it if the data is valid
func (s *EventService) CreateEvent(params EventCreate) (*Event, error) {
	event := repository.Event{
		Name:        params.Name,
		Description: params.Description.NullString,
		Date:        params.Date,
	}

	if len(params.Invitees) > int(s.maxEventInvitees) {
		return nil, fmt.Errorf("invitees [max %d]: %w", s.maxEventInvitees, ErrInvalidAmount)
	}

	for _, inviteeEmail := range params.Invitees {
		event.EventInvitees = append(event.EventInvitees, repository.EventInvitee{Email: inviteeEmail})
	}

	for _, language := range params.Languages {
		id, ok := s.languageMap[language]
		if !ok {
			return nil, fmt.Errorf("language [%s]: %w", language, ErrInvalidValue)
		}
		event.Languages = append(event.Languages, repository.Language{ID: id})
	}

	for _, videoQuality := range params.VideoQuality {
		id, ok := s.videoQualityMap[videoQuality]
		if !ok {
			return nil, fmt.Errorf("video quality [%s]: %w", videoQuality, ErrInvalidValue)
		}
		event.VideoQualities = append(event.VideoQualities, repository.VideoQuality{ID: id})
	}

	for _, audioQuality := range params.AudioQuality {
		id, ok := s.audioQualityMap[audioQuality]
		if !ok {
			return nil, fmt.Errorf("audio quality [%s]: %w", audioQuality, ErrInvalidValue)
		}
		event.AudioQualities = append(event.AudioQualities, repository.AudioQuality{ID: id})
	}

	if err := s.eventRepository.Add(&event); err != nil {
		return nil, fmt.Errorf("event create: %w", err)
	}

	return &Event{
		ID:           event.ID,
		Name:         event.Name,
		Date:         event.Date,
		Languages:    params.Languages,
		VideoQuality: params.VideoQuality,
		AudioQuality: params.AudioQuality,
		Invitees:     params.Invitees,
		Description:  params.Description,
	}, nil
}
