package store

import (
	"api/model"
	"fmt"
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	uuid "github.com/satori/go.uuid"
)

const collection = "tasks"

type mongo struct {
	session *mgo.Session
}

func NewMongoManager(session *mgo.Session) Manager {
	return &mongo{session: session}
}

func (m *mongo) Create(t *model.Task) error {
	if t.ID == "" {
		id, err := uuid.NewV4()
		if err != nil {
			return err
		}
		t.ID = id.String()
		t.CreatedAt = time.Now()
	}

	session := m.session.Copy()
	defer session.Close()
	c := session.DB("").C(collection)

	fmt.Println(t)
	return c.Insert(t)
}

func (m *mongo) Find(ID string) (*model.Task, error) {
	if _, err := uuid.FromString(ID); err != nil {
		return nil, err
	}

	session := m.session.Copy()
	defer session.Close()

	task := &model.Task{}
	err := session.DB("").C(collection).FindId(ID).One(task)

	return task, err
}

func (m *mongo) Delete(ID string) error {
	if _, err := uuid.FromString(ID); err != nil {
		return err
	}

	session := m.session.Copy()
	defer session.Close()

	return session.DB("").C(collection).RemoveId(ID)
}

func (m *mongo) Update(t *model.Task) error {
	session := m.session.Copy()
	defer session.Close()

	return session.DB("").C(collection).UpdateId(t.ID, t)
}

func (m *mongo) FindByStatus(flag bool) ([]model.Task, error) {
	session := m.session.Copy()
	defer session.Close()

	tasks := []model.Task{}

	c := session.DB("").C(collection).Find(bson.M{"status": flag})
	err := c.All(&tasks)

	return tasks, err
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
