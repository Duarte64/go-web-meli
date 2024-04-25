package users

import (
	"encoding/json"
	"testing"

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
	repository := NewRepository(&store)
	var us, err = repository.GetAll()

	expectFirstId := uint(1)
	expectSecondId := uint(2)

	assert.Nil(t, err)
	assert.Len(t, us, 2)
	assert.Equal(t, us[0].ID, expectFirstId, "o primeiro ID deve ser igual a 1")
	assert.Equal(t, us[1].ID, expectSecondId, "o segundo ID deve ser igual a 2")
}

func TestLastId(t *testing.T) {
	store := StoreStub{readWasCalled: false}
	repository := NewRepository(&store)
	var id, err = repository.LastId()

	expectLastId := uint(2)

	assert.Nil(t, err)
	assert.Equal(t, id, expectLastId, "o ultimo ID deve ser igual a 2")
}

func TestUpdate(t *testing.T) {
	store := StoreStub{readWasCalled: false}
	repository := NewRepository(&store)

	assert.False(t, store.readWasCalled)

	expectNameAfter := "After Update"
	us, err := repository.Update(1, "After Update", "Test", "test@test.com", 22, 1.7, true)

	assert.Nil(t, err)
	assert.True(t, store.readWasCalled)
	assert.Equal(t, us.Name, expectNameAfter, "After Update")
}

func TestUpdateNotFound(t *testing.T) {
	store := StoreStub{readWasCalled: false}
	repository := NewRepository(&store)

	assert.False(t, store.readWasCalled)

	_, err := repository.Update(30, "After Update", "Test", "test@test.com", 22, 1.7, true)

	assert.Error(t, err)
}

func TestPatch(t *testing.T) {
	store := StoreStub{readWasCalled: false}
	repository := NewRepository(&store)

	expectNameAfter := "After Update"
	us, err := repository.Patch(1, "After Update", 45)

	assert.Nil(t, err)
	assert.Equal(t, us.Lastname, expectNameAfter, "devem ser iguais")
}

func TestPatchNotFound(t *testing.T) {
	store := StoreStub{readWasCalled: false}
	repository := NewRepository(&store)

	_, err := repository.Patch(20, "After Update", 45)

	assert.Error(t, err)
}

func TestDelete(t *testing.T) {
	store := StoreStub{readWasCalled: false}
	repository := NewRepository(&store)

	err := repository.Delete(uint(1))

	assert.Nil(t, err)
}

func TestDeleteError(t *testing.T) {
	store := StoreStub{readWasCalled: false}
	repository := NewRepository(&store)

	err := repository.Delete(uint(20))

	assert.Error(t, err)
}

func TestGetById(t *testing.T) {
	store := StoreStub{readWasCalled: false}
	repository := NewRepository(&store)

	us, err := repository.GetById(uint(1))

	assert.NoError(t, err)
	assert.Equal(t, us.ID, uint(1), "devem ser iguais")
}

func TestGetByIdError(t *testing.T) {
	store := StoreStub{readWasCalled: false}
	repository := NewRepository(&store)

	_, err := repository.GetById(uint(20))

	assert.Error(t, err)
}

func TestStore(t *testing.T) {
	store := StoreStub{readWasCalled: false}
	repository := NewRepository(&store)

	us, err := repository.Store(uint(3), "test", "test", "test", "test", 22, 1.7, true)

	assert.NoError(t, err)
	assert.Equal(t, us.ID, uint(3), "devem ser iguais")
}
