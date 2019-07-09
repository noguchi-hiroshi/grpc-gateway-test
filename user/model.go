package user

import (
	"errors"
	"github.com/thoas/go-funk"
	"time"
)

type Model interface {
	Create(email string, password string) (*Entity, error)
	Find(id int64) (*Entity, error)
}

type model struct {
	entities []*Entity
}

func NewModel() Model {
	entities := make([]*Entity, 0)
	return &model{entities: entities}
}

func (m *model) Find(id int64) (*Entity, error) {
	ret := funk.Find(m.entities, func (e *Entity) bool {
		return e.ID == id
	})
	if ret == nil {
		return nil, errors.New("not found user")
	}
	switch entity := ret.(type) {
	case *Entity: return entity, nil
	default: return nil, errors.New("cast error")
	}
}

func (m *model) Create(email string, password string) (*Entity, error) {
	id := time.Now().Unix()
	entity := Entity{ID: id, Email: email, Password: password}
	m.entities = append(m.entities, &entity)
	return &entity, nil
}