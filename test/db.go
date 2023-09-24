package test

import (
	"fmt"
	"os"
	"time"

	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DBTestSuite struct {
	suite.Suite

	DSN string
	DB  *gorm.DB
}

func (suite *DBTestSuite) SetupSuite() {
	var err error
	suite.DSN = getDSN()
	suite.DB, err = gorm.Open(sqlite.Open(suite.DSN), &gorm.Config{})
	if err != nil {
		suite.FailNow("DBTestSuite SetupSuite", err)
	}
}

func (suite *DBTestSuite) TearDownSuite() {
	if err := os.Remove(suite.DSN); err != nil {
		suite.FailNow("DBTestSuite TearDownSuite", err)
	}
}

func getDSN() string {
	return fmt.Sprintf("test_%d.sqlite", time.Now().Unix())
}
