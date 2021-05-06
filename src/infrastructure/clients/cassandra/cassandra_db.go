package cassandra

import (
	"github.com/esequielvirtuoso/bookstore_oauth_api/src/pkg/logger"
	"github.com/gocql/gocql"
)

var (
	session *gocql.Session
)

func init() {
	// Connect to Cassandra cluster
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum

	var err error
	if session, err = cluster.CreateSession(); err != nil {
		logger.Error("error when trying to establish a session to db", err)
		panic(err)
	}
}

func GetSession() *gocql.Session{
	return session
}