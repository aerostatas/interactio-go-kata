package handler

import (
	"log"
	"testing"

	"github.com/aerostatas/interaction-go-kata/internal/repository"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockRepository[T any] struct {
	mock.Mock
}

func (r *MockRepository[T]) List() ([]T, error) {
	args := r.Called()
	t, ok := args.Get(0).([]T)
	if !ok {
		log.Fatal("expected return arg to be of type []T")
	}
	return t, args.Error(1)
}

func (r *MockRepository[T]) Add(e *T) error {
	args := r.Called(e)
	return args.Error(0)
}

func (r *MockRepository[T]) Remove(e *T) error {
	args := r.Called(e)
	return args.Error(0)
}

type MockRepositoryFactory struct {
	mock.Mock
}

func (f *MockRepositoryFactory) Events() repository.EventRepository {
	args := f.Called()
	r, ok := args.Get(0).(repository.EventRepository)
	if !ok {
		log.Fatal("expected return arg to be of type repository.EventRepository")
	}
	return r
}

func (f *MockRepositoryFactory) Languages() repository.LanguageRepository {
	args := f.Called()
	r, ok := args.Get(0).(repository.LanguageRepository)
	if !ok {
		log.Fatal("expected return arg to be of type repository.LanguageRepository")
	}
	return r
}

func (f *MockRepositoryFactory) VideoQualities() repository.VideoQualityRepository {
	args := f.Called()
	r, ok := args.Get(0).(repository.VideoQualityRepository)
	if !ok {
		log.Fatal("expected return arg to be of type repository.VideoQualityRepository")
	}
	return r
}

func (f *MockRepositoryFactory) AudioQualities() repository.AudioQualityRepository {
	args := f.Called()
	r, ok := args.Get(0).(repository.AudioQualityRepository)
	if !ok {
		log.Fatal("expected return arg to be of type repository.AudioQualityRepository")
	}
	return r
}

type HandlerFactoryTestSuite struct {
	suite.Suite
	repositoryFactory *MockRepositoryFactory
	factory           *HandlerFactory
}

func (suite *HandlerFactoryTestSuite) SetupTest() {
	suite.repositoryFactory = new(MockRepositoryFactory)
	suite.factory = NewHandlerFactory(suite.repositoryFactory)
}

func TestHandlerFactoryTestSuite(t *testing.T) {
	suite.Run(t, new(HandlerFactoryTestSuite))
}

func (suite *HandlerFactoryTestSuite) Test_EventHandler() {
	mockEventRepository := new(MockRepository[repository.Event])
	mockLanguageRepository := new(MockRepository[repository.Language])
	mockVideoQualityRepository := new(MockRepository[repository.VideoQuality])
	mockAudioQualityRepository := new(MockRepository[repository.AudioQuality])

	suite.repositoryFactory.On("Events").Return(mockEventRepository)
	suite.repositoryFactory.On("Languages").Return(mockLanguageRepository)
	suite.repositoryFactory.On("VideoQualities").Return(mockVideoQualityRepository)
	suite.repositoryFactory.On("AudioQualities").Return(mockAudioQualityRepository)

	mockLanguageRepository.On("List").Return([]repository.Language{{ID: 1, Name: "English"}}, nil)
	mockVideoQualityRepository.On("List").Return([]repository.VideoQuality{{ID: 2, Quality: "1080p"}}, nil)
	mockAudioQualityRepository.On("List").Return([]repository.AudioQuality{{ID: 3, Quality: "High"}}, nil)

	handler, err := suite.factory.EventHandler()
	suite.Require().NoError(err)
	suite.Assert().IsType(&EventHandler{}, handler)
	suite.repositoryFactory.AssertExpectations(suite.T())
}
