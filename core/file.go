package core

import (
	"io"
	"mime/multipart"
	"path/filepath"
	"strings"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/models"
)

// FileService handles file uploads, retrieval, and deletion.
type FileService struct {
	pb *pocketbase.PocketBase
}

// NewFileService returns a new file manager using the PB instance.
func NewFileService(pb *pocketbase.PocketBase) *FileService {
	return &FileService{pb: pb}
}

// Upload saves a file to a record field in the given collection.
func (fs *FileService) Upload(collection, recordID, field string, fileHeader *multipart.FileHeader) (*models.Record, error) {
	dao := fs.pb.Dao()
	record, err := dao.FindRecordById(collection, recordID)
	if err != nil {
		return nil, err
	}

	file, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	filename := fileHeader.Filename
	ext := filepath.Ext(filename)

	savedFile, err := fs.pb.Files().Upload(file, filename, ext)
	if err != nil {
		return nil, err
	}

	record.Set(field, savedFile)
	if err := dao.SaveRecord(record); err != nil {
		return nil, err
	}

	return record, nil
}

// UploadMultiple allows multiple file uploads (same field) to a record.
func (fs *FileService) UploadMultiple(collection, recordID, field string, files []*multipart.FileHeader) (*models.Record, error) {
	dao := fs.pb.Dao()
	record, err := dao.FindRecordById(collection, recordID)
	if err != nil {
		return nil, err
	}

	var filenames []string

	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			continue
		}
		defer file.Close()

		ext := filepath.Ext(fileHeader.Filename)
		saved, err := fs.pb.Files().Upload(file, fileHeader.Filename, ext)
		if err == nil {
			filenames = append(filenames, saved)
		}
	}

	record.Set(field, filenames)
	if err := dao.SaveRecord(record); err != nil {
		return nil, err
	}

	return record, nil
}

// GetURL returns the public URL to the stored file.
func (fs *FileService) GetURL(record *models.Record, field string) string {
	return fs.pb.Files().GetUrl(record, field)
}

// GetThumbnailURL returns a resized image (if image field).
func (fs *FileService) GetThumbnailURL(record *models.Record, field string, width, height int) string {
	return fs.pb.Files().GetUrl(record, field) + "?thumb=" + strings.Join([]string{
		"fit:cover",
		"width:" + itoa(width),
		"height:" + itoa(height),
	}, ",")
}

// Delete removes the file from a record field.
func (fs *FileService) Delete(collection, recordID, field string) error {
	dao := fs.pb.Dao()
	record, err := dao.FindRecordById(collection, recordID)
	if err != nil {
		return err
	}

	record.Set(field, nil)
	return dao.SaveRecord(record)
}
