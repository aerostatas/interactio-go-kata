package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/aerostatas/interaction-go-kata/internal/service"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"gopkg.in/guregu/null.v3"
)

type MockEventService struct {
	mock.Mock
}

func (r *MockEventService) CreateEvent(params service.EventCreate) (*service.Event, error) {
	args := r.Called(params)
	event, ok := args.Get(0).(*service.Event)
	if !ok {
		log.Fatal("expected return arg to be of type service.Event")
	}
	return event, args.Error(1)
}

type EventHandlerTestSuite struct {
	suite.Suite
	service *MockEventService
	handler *EventHandler
}

func (suite *EventHandlerTestSuite) SetupTest() {
	suite.service = new(MockEventService)
	suite.handler = NewEventHandler(suite.service)
}

func TestEventHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(EventHandlerTestSuite))
}

func (suite *EventHandlerTestSuite) Test_CreateEvent() {
	dateString := "2023-04-20T14:00:00Z"
	date, err := time.Parse(time.RFC3339, dateString)
	if err != nil {
		suite.FailNow(err.Error())
	}

	req := service.EventCreate{
		Name:         "test event",
		Date:         date,
		Description:  null.NewString("test description", true),
		Languages:    []string{"Lithuanian", "English"},
		VideoQuality: []string{"360p", "720p"},
		AudioQuality: []string{"Low", "High"},
		Invitees:     []string{"one@email.com", "two@email.com"},
	}

	created := &service.Event{
		ID:           1,
		Name:         req.Name,
		Description:  req.Description,
		Date:         req.Date,
		Languages:    req.Languages,
		VideoQuality: req.VideoQuality,
		AudioQuality: req.AudioQuality,
		Invitees:     req.Invitees,
	}

	suite.service.On("CreateEvent", req).Return(created, nil)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/api/events", structJSONBuffer(req))

	suite.handler.CreateEvent().ServeHTTP(w, r)

	var decoded service.Event
	if err := json.NewDecoder(w.Result().Body).Decode(&decoded); err != nil {
		suite.FailNow(err.Error())
	}

	suite.Assert().Equal(http.StatusCreated, w.Result().StatusCode)
	suite.Assert().Equal(created, &decoded)
}

func structJSONBuffer(r any) io.Reader {
	b := new(bytes.Buffer)

	if err := json.NewEncoder(b).Encode(r); err != nil {
		log.Fatal(err)
	}

	return b
}
