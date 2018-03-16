package store

import (
	"time"

	"github.com/globalsign/mgo"
)

type StoreType int

const (
	MongoType StoreType = iota
	MockType
)

func CreateManager(t StoreType) (Manager, error) {
	switch t {
	case MongoType:
		timeout := 5 * time.Second
		url := "mongodb://mongodb/todo"

		session, err := mgo.DialWithTimeout(url, timeout)
		if err != nil {
			return nil, err
		}

		session.SetPoolLimit(30)
		return NewMongoManager(session), nil

	default:
		return NewMockManager(), nil
	}
}
