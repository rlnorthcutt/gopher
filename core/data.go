package core

import (
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/models"
)

// DataService wraps PocketBase DAO for common data operations.
type DataService struct {
	pb *pocketbase.PocketBase
}

// NewDataService returns a new instance of DataService.
func NewDataService(pb *pocketbase.PocketBase) *DataService {
	return &DataService{pb: pb}
}

// GetByID returns a single record by ID.
func (d *DataService) GetByID(collection, id string) (*models.Record, error) {
	return d.pb.Dao().FindRecordById(collection, id)
}

// GetAll returns all records for a given collection (with optional pagination later).
func (d *DataService) GetAll(collection string) ([]*models.Record, error) {
	return d.pb.Dao().FindRecordsByFilter(collection, "", "", "", 0, 100)
}

// Create inserts a new record from a field map.
func (d *DataService) Create(collection string, fields map[string]interface{}) (*models.Record, error) {
	record := models.NewRecord(d.pb.Dao().ModelCollection(collection))
	for k, v := range fields {
		record.Set(k, v)
	}
	if err := d.pb.Dao().SaveRecord(record); err != nil {
		return nil, err
	}
	return record, nil
}

// Update modifies an existing record with new fields.
func (d *DataService) Update(collection, id string, updates map[string]interface{}) (*models.Record, error) {
	record, err := d.GetByID(collection, id)
	if err != nil {
		return nil, err
	}
	for k, v := range updates {
		record.Set(k, v)
	}
	if err := d.pb.Dao().SaveRecord(record); err != nil {
		return nil, err
	}
	return record, nil
}

// Delete removes a record from the collection.
func (d *DataService) Delete(collection, id string) error {
	record, err := d.GetByID(collection, id)
	if err != nil {
		return err
	}
	return d.pb.Dao().DeleteRecord(record)
}
