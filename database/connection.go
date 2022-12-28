package database

import "github.com/gocql/gocql"

type DBConnection struct {
	cluster *gocql.ClusterConfig
	session *gocql.Session
}

var connection DBConnection

func SetUpDBConnection() {

	connection.cluster = gocql.NewCluster("127.0.0.1:9042")

	connection.cluster.Consistency = gocql.Quorum
	connection.cluster.Keyspace = "chatapp"
	connection.session, _ = connection.cluster.CreateSession()

}

func ExecuteQuery(query string) {
	println("///////////////////////////////////////####")
	if err := connection.session.Query(query).Exec(); err != nil {
		println("///////////////////////////////////////")
		println("error occur", &err)
	}
}
