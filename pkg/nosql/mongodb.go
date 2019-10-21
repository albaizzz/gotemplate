package nosql

import (
	"fmt"

	"github.com/globalsign/mgo"
)

// MongoConnection is the struct that act as the connection for MongoDB
type MongoConnection struct {
	User      string
	Password  string
	Host      string
	Port      int
	DBName    string
	mongoConn *mgo.Session
}

var urlFormat = "mongodb://%s:%s@%s:%d/%s"

// InitMongo is the function that used to initialize MongoDB connection
func (m *MongoConnection) CreateSession() {
	var dsn string

	dsn = fmt.Sprintf(urlFormat, m.User, m.Password, m.Host, m.Port, m.DBName)

	mongoConn, err := mgo.Dial(dsn)
	if err != nil {
		panic(fmt.Errorf("failed dial : %v", err))
	}

	mongoConn.SetSafe(&mgo.Safe{WMode: "majority"})
	m.mongoConn = mongoConn
	return
}

/*
GetConnection is the pointer receiver of `MongoConnection` struct.
This function is the function that used to get MongoDB connection.
*/
func (m *MongoConnection) GetConnection() *mgo.Session {
	return m.mongoConn.Copy()
}

func (m *MongoConnection) GetDBName() string {
	return m.DBName
}

/*
Close is the pointer receiver of `MongoConnection` struct.
This function is the function that used to close MongoDB connection.
*/
func (m *MongoConnection) Close() {
	m.mongoConn.Close()
}
