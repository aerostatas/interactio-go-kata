package repository

import (
	"testing"

	"github.com/aerostatas/interaction-go-kata/test"
	"github.com/stretchr/testify/suite"
)

type SchemaTestSuite struct {
	test.DBTestSuite

	factory *RepositoryFactory
}

func (suite *SchemaTestSuite) SetupSuite() {
	suite.DBTestSuite.SetupSuite()

	suite.factory = NewRepositoryFactory(suite.DB)

	if err := Migrate(suite.DB); err != nil {
		suite.FailNowf("SchemaTestSuite SetupSuite", "migrate DB: %w", err)
	}

	if err := Seed(suite.DB); err != nil {
		suite.FailNowf("SchemaTestSuite SetupSuite", "seed DB: %w", err)
	}
}

func (suite *SchemaTestSuite) TearDownSuite() {
	suite.DBTestSuite.TearDownSuite()
}

func TestSchemaTestSuite(t *testing.T) {
	suite.Run(t, new(SchemaTestSuite))
}

func (suite *SchemaTestSuite) Test_Migrate() {
	var tableNames []string
	err := suite.DB.Raw("SELECT name FROM sqlite_master WHERE type = ?", "table").Pluck("name", &tableNames).Error
	if err != nil {
		suite.FailNow(err.Error())
	}

	suite.Assert().Len(tableNames, len(models))
}

func (suite *SchemaTestSuite) Test_Seed() {
	suite.Assert().NoError(Seed(suite.DB))

	languages, err := suite.factory.Languages().List()
	if err != nil {
		suite.FailNow(err.Error())
	}

	languageNames := []string{}
	for _, language := range languages {
		languageNames = append(languageNames, language.Name)
	}
	for _, language := range defaultLanguages {
		suite.Assert().Contains(languageNames, language.Name)
	}

	videoQualities, err := suite.factory.VideoQualities().List()
	if err != nil {
		suite.FailNow(err.Error())
	}

	videoQuals := []string{}
	for _, videoQuality := range videoQualities {
		videoQuals = append(videoQuals, videoQuality.Quality)
	}
	for _, videoQuality := range defaultVideoQualities {
		suite.Assert().Contains(videoQuals, videoQuality.Quality)
	}

	audioQualities, err := suite.factory.AudioQualities().List()
	if err != nil {
		suite.FailNow(err.Error())
	}

	audioQuals := []string{}
	for _, audioQuality := range audioQualities {
		audioQuals = append(audioQuals, audioQuality.Quality)
	}
	for _, audioQuality := range defaultAudioQualities {
		suite.Assert().Contains(audioQuals, audioQuality.Quality)
	}
}
