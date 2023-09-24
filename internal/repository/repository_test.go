package repository

import (
	"testing"

	"github.com/aerostatas/interaction-go-kata/test"
	"github.com/stretchr/testify/suite"
)

type RepositoryTestSuite struct {
	test.DBTestSuite

	repository *repository[Language]
}

func (suite *RepositoryTestSuite) SetupSuite() {
	suite.DBTestSuite.SetupSuite()

	suite.repository = &repository[Language]{db: suite.DB}

	if err := Migrate(suite.DB); err != nil {
		suite.FailNowf("RepositoryTestSuite SetupSuite", "migrate DB: %w", err)
	}
}

func (suite *RepositoryTestSuite) TearDownSuite() {
	suite.DBTestSuite.TearDownSuite()
}

func TestRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(RepositoryTestSuite))
}

func (suite *RepositoryTestSuite) Test_Actions() {
	var err error
	var languages []Language

	language := Language{
		Name: "Test_Actions",
	}

	err = suite.repository.Add(&language)
	suite.Require().NoError(err)

	languages, err = suite.repository.List()
	suite.Require().NoError(err)
	suite.Assert().Len(languages, 1)
	suite.Assert().Equal(language.Name, languages[0].Name)

	err = suite.repository.Remove(&languages[0])
	suite.Require().NoError(err)

	languages, err = suite.repository.List()
	suite.Require().NoError(err)
	suite.Assert().Len(languages, 0)
}
