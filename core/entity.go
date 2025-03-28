package core

import (
	"errors"

	"github.com/pocketbase/pocketbase/models"
)

// Entity defines a full lifecycle model tied to a PocketBase collection.
type Entity interface {
	TableName() string
	ToPBRecord() *models.Record
	FromPBRecord(*models.Record) error

	BeforeSave() error
	AfterSave() error
	BeforeUpdate() error
	AfterUpdate() error

	Save() error
	Delete() error
}

// entityRuntime is the injected runtime used by all entities.
var entityRuntime *DataService

func RegisterEntityRuntime(data *DataService) {
	entityRuntime = data
}

// BaseEntity provides shared ID and default implementations for hooks and persistence.
type BaseEntity struct {
	ID string
}

func (e *BaseEntity) BeforeSave() error   { return nil }
func (e *BaseEntity) AfterSave() error    { return nil }
func (e *BaseEntity) BeforeUpdate() error { return nil }
func (e *BaseEntity) AfterUpdate() error  { return nil }

func (e *BaseEntity) Save(entity Entity) error {
	if entityRuntime == nil {
		return errors.New("entity runtime not initialized")
	}

	rec := entity.ToPBRecord()
	if e.ID != "" {
		rec.Id = e.ID
		_ = entity.BeforeUpdate()
	} else {
		_ = entity.BeforeSave()
	}

	saved, err := entityRuntime.pb.Dao().SaveRecord(rec)
	if err != nil {
		return err
	}

	e.ID = saved.Id
	_ = entity.FromPBRecord(saved)

	if e.ID != "" {
		_ = entity.AfterUpdate()
	} else {
		_ = entity.AfterSave()
	}

	return nil
}

func (e *BaseEntity) Delete(entity Entity) error {
	if entityRuntime == nil {
		return errors.New("entity runtime not initialized")
	}
	if e.ID == "" {
		return errors.New("cannot delete unsaved entity")
	}
	return entityRuntime.Delete(entity.TableName(), e.ID)
}
