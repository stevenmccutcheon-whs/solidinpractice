package nearearth

import (
	"encoding/json"
	"fmt"
	"os"
)

type ImpactFileService struct {
	filepath string
}

func (s ImpactFileService) GetAll() (Impacts, error) {
	return s.readEventsFromFile()
}

type FileServiceOption func(*ImpactFileService)

func WithTestFilePath(filepath string) FileServiceOption {
	return func(s *ImpactFileService) {
		s.filepath = filepath
	}
}

func NewImpactFileService(opts ...FileServiceOption) (ImpactFileService, error) {
	filepath := "nearearth.json"
	defaultService := ImpactFileService{filepath: filepath}
	for _, opt := range opts {
		opt(&defaultService)
	}
	_, err := os.Stat(defaultService.filepath)
	if err != nil {
		return defaultService, fmt.Errorf("file does not exist: %w", err)
	}
	return defaultService, nil
}

func (s ImpactFileService) readEventsFromFile() (Impacts, error) {
	// Read the file
	// Parse the JSON
	// Return the events
	f, err := os.Open(s.filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var res Impacts
	err = json.NewDecoder(f).Decode(&res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
