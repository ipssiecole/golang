package store

import (
	"api/model"

	"github.com/globalsign/mgo"
	uuid "github.com/satori/go.uuid"
)

const collection = "tasks"

type mongo struct {
	session *mgo.Session
}

func NewMongoManager(session *mgo.Session) Manager {
	s := session.Copy()
	defer s.Close()

	return &mongo{session: s}
}

func (m *mongo) Create(t *model.Task) error {
	if t.ID == "" {
		id, err := uuid.NewV4()
		if err != nil {
			return err
		}
		t.ID = id.String()
	}

	session := m.session.Copy()
	defer session.Close()
	c := session.DB("").C(collection)

	return c.Insert(t)
}

// TODO Check limit
func (m *mongo) All(offset, limit int) ([]model.Task, error) {
	var err error
	tasks := []model.Task{}

	session := m.session.Copy()
	defer session.Close()
	c := session.DB("").C(collection)

	// total == 100
	// offset = 5
	// limit = 10
	if offset > NoPaging && limit > NoPaging && limit > offset {
		err = c.Find(nil).Skip(offset).Limit(limit).All(&tasks)
	} else {
		err = c.Find(nil).All(&tasks)
	}

	return tasks, err
}
