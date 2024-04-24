package tests

import (
	"encoding/json"
	"testing"

	"github.com/Duarte64/go-web-meli/internal/users"
	"github.com/stretchr/testify/assert"
)

var file = []byte(`[
        {
			"id": 1,
			"name": "Before Update",
			"lastname": "Doe",
			"email": "jane.doe@gmail.com",
			"age": 28,
			"height": 1.7,
			"active": true,
			"created_at": "2019-02-01 00:00:00"
		},
        {
			"id": 2,
			"name": "Gabriel",
			"lastname": "Duarte",
			"email": "gabriel.figueiredo@mercadolivre.com",
			"age": 23,
			"height": 1.7,
			"active": true,
			"created_at": "2024-04-12 11:04:19.42315 -0300 -03 m=+11.688279793"
		}
    ]`)

type StoreStub struct {
	readWasCalled bool
}

func (s *StoreStub) Read(data interface{}) error {
	s.readWasCalled = true
	return json.Unmarshal(file, &data)
}

func (s StoreStub) Write(interface{}) error {
	return nil
}

func TestGetAll(t *testing.T) {
	store := StoreStub{readWasCalled: false}
	repository := users.NewRepository(&store)
	var us, err = repository.GetAll()

	expectFirstId := uint(1)
	expectSecondId := uint(2)

	assert.Nil(t, err)
	assert.Len(t, us, 2)
	assert.Equal(t, us[0].ID, expectFirstId, "o primeiro ID deve ser igual a 1")
	assert.Equal(t, us[1].ID, expectSecondId, "o segundo ID deve ser igual a 2")
}

func TestUpdate(t *testing.T) {
	store := StoreStub{readWasCalled: false}
	repository := users.NewRepository(&store)

	assert.False(t, store.readWasCalled)

	expectNameAfter := "After Update"
	us, err := repository.Update(1, "After Update", "Test", "test@test.com", 22, 1.7, true)

	assert.Nil(t, err)
	assert.True(t, store.readWasCalled)
	assert.Equal(t, us.Name, expectNameAfter, "After Update")
}
