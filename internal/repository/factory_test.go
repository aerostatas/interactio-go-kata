package repository

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type FactoryTestSuite struct {
	suite.Suite

	factory *RepositoryFactory
}

func (suite *FactoryTestSuite) SetupSuite() {
	suite.factory = NewRepositoryFactory(&gorm.DB{})
}

func TestFactoryTestSuite(t *testing.T) {
	suite.Run(t, new(FactoryTestSuite))
}

func (suite *FactoryTestSuite) Test_Events() {
	suite.Assert().IsType(&repository[Event]{}, suite.factory.Events())
}

func (suite *FactoryTestSuite) Test_Languages() {
	suite.Assert().IsType(&repository[Language]{}, suite.factory.Languages())
}

func (suite *FactoryTestSuite) Test_VideoQualities() {
	suite.Assert().IsType(&repository[VideoQuality]{}, suite.factory.VideoQualities())
}

func (suite *FactoryTestSuite) Test_AudioQualities() {
	suite.Assert().IsType(&repository[AudioQuality]{}, suite.factory.AudioQualities())
}
