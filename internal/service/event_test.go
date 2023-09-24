package service

import (
	"testing"
	"time"

	"github.com/aerostatas/interaction-go-kata/internal/repository"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"gopkg.in/guregu/null.v3"
)

type MockEventRepository struct {
	mock.Mock
}

func (r *MockEventRepository) Add(event *repository.Event) error {
	args := r.Called(event)
	return args.Error(0)
}

type EventServiceTestSuite struct {
	suite.Suite
	repository *MockEventRepository
	service    *EventService
}

func (suite *EventServiceTestSuite) SetupTest() {
	suite.repository = new(MockEventRepository)
	suite.service = NewEventService(
		suite.repository,
		map[string]uint{"Lithuanian": 1, "French": 2},
		map[string]uint{"720p": 1, "1080p": 2},
		map[string]uint{"Low": 1, "High": 2},
		2,
	)
}

func TestEventServiceTestSuite(t *testing.T) {
	suite.Run(t, new(EventServiceTestSuite))
}

func (suite *EventServiceTestSuite) Test_CreateEvent_rejectsTooManyInvitees() {
	suite.T().Parallel()

	_, err := suite.service.CreateEvent(
		EventCreate{
			"test",
			time.Now(),
			[]string{"French"},
			[]string{"1080p"},
			[]string{"High"},
			[]string{"one@mail.com", "two@mail.com", "three@mail.com"},
			null.NewString("test event", true),
		},
	)

	suite.Assert().ErrorIs(err, ErrInvalidAmount)
	suite.Assert().Equal("invitees [max 2]: invalid amount", err.Error())
}

func (suite *EventServiceTestSuite) Test_CreateEvent_rejectsInvalidLanguage() {
	suite.T().Parallel()

	_, err := suite.service.CreateEvent(
		EventCreate{
			"test",
			time.Now(),
			[]string{"Madeup"},
			[]string{"1080p"},
			[]string{"High"},
			[]string{"one@mail.com"},
			null.NewString("test event", true),
		},
	)

	suite.Assert().ErrorIs(err, ErrInvalidValue)
	suite.Assert().Equal("language [Madeup]: invalid value", err.Error())
}

func (suite *EventServiceTestSuite) Test_CreateEvent_rejectsInvalidVideoQuality() {
	suite.T().Parallel()

	_, err := suite.service.CreateEvent(
		EventCreate{
			"test",
			time.Now(),
			[]string{"French"},
			[]string{"900p"},
			[]string{"High"},
			[]string{"one@mail.com"},
			null.NewString("test event", true),
		},
	)

	suite.Assert().ErrorIs(err, ErrInvalidValue)
	suite.Assert().Equal("video quality [900p]: invalid value", err.Error())
}

func (suite *EventServiceTestSuite) Test_CreateEvent_rejectsInvalidAudioQuality() {
	suite.T().Parallel()

	_, err := suite.service.CreateEvent(
		EventCreate{
			"test",
			time.Now(),
			[]string{"French"},
			[]string{"1080p"},
			[]string{"Invalid"},
			[]string{"one@mail.com"},
			null.NewString("test event", true),
		},
	)

	suite.Assert().ErrorIs(err, ErrInvalidValue)
	suite.Assert().Equal("audio quality [Invalid]: invalid value", err.Error())
}

func (suite *EventServiceTestSuite) Test_CreateEvent() {
	suite.T().Parallel()

	name := "test"
	description := null.NewString("test event", true)
	date := time.Now()
	email1 := "one@mail.com"
	email2 := "two@mail.com"

	suite.repository.On("Add", &repository.Event{
		Name:           name,
		Description:    description.NullString,
		Date:           date,
		Languages:      []repository.Language{{ID: 1}, {ID: 2}},
		VideoQualities: []repository.VideoQuality{{ID: 1}, {ID: 2}},
		AudioQualities: []repository.AudioQuality{{ID: 1}, {ID: 2}},
		EventInvitees:  []repository.EventInvitee{{Email: email1}, {Email: email2}},
	}).Return(nil)

	_, err := suite.service.CreateEvent(
		EventCreate{
			name,
			date,
			[]string{"Lithuanian", "French"},
			[]string{"720p", "1080p"},
			[]string{"Low", "High"},
			[]string{email1, email2},
			description,
		},
	)

	suite.Require().NoError(err)
	suite.repository.AssertExpectations(suite.T())
}
