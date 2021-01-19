package article_test

import (
	"log"
	"testing"

	article "github.com/shenyi-tw/golib/dao/article"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc/status"
)

const (
	DatabaseHost     = "127.0.0.1"
	DatabasePort     = "5432"
	DatabaseName     = "postgres"
	DatabaseUser     = "postgres"
	DatabasePassword = "pass"
	MaxDatabaseRetry = 5
)

var db *gorm.DB

type UserTestSuite struct {
	suite.Suite
	User article.Conn
}

func (suite *UserTestSuite) SetupSuite() {
	var db article.Conn
	db = article.CreateConn(DatabaseUser, DatabasePassword, DatabaseHost, DatabasePort, DatabaseName)
	suite.User = db
}

func (suite *UserTestSuite) TestCreate() {
	t := suite.Require()
	us := []string{
		"t1",
		"t2",
	}
	err := suite.User.Create(us)
	if err != nil {
		st := status.Convert(err)
		log.Println(st.Code())
		log.Println(st.Message())
	}

	t.NoError(err, "failed to create article")
}
func (suite *UserTestSuite) TestGet() {
	t := suite.Require()

	u, err := suite.User.Get()
	log.Println(u)
	if err != nil {
		st := status.Convert(err)
		log.Println(st.Code())
		log.Println(st.Message())
	}
	t.NotZero(u, "article id must not be zero")
	t.NoError(err, "failed to create article")
}

func (suite *UserTestSuite) TestSaveProxy() {
	u, _ := suite.User.Get()
	for idx := range u {
		u[idx].Retry += 1
	}
	suite.User.Save(u)
}

func TestUserTestSuite(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}
