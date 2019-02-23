package storage

import (
	"encoding/json"
	"io/ioutil"

	"io"

	"github.com/dohernandez/travels-budget/internal/domain"
)

type fileActivityFinder struct {
	reader io.Reader
}

// NewFileActivityFinder creates an instance of domain.ActivityFinder with file data source
func NewFileActivityFinder(reader io.Reader) domain.ActivityFinder {
	return &fileActivityFinder{
		reader: reader,
	}
}

// FindAll reads all activities from json file
func (f *fileActivityFinder) FindAll() ([]domain.Activity, error) {
	var activities []domain.Activity

	byteValue, _ := ioutil.ReadAll(f.reader)

	err := json.Unmarshal(byteValue, &activities)
	if err != nil {
		return nil, err
	}

	return activities, nil
}
