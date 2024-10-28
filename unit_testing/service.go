package main

import (
	"errors"
)

// DataService interface with methods to be mocked
type DataService interface {
	GetData(id string) (string, error)
	SaveData(id string, value string) error
}

// Service struct that uses DataService
type Service struct {
	dataService DataService
}

// NewService constructor
func NewService(ds DataService) *Service {
	return &Service{dataService: ds}
}

// Methods to Unit Test

func (s *Service) FetchData(id string) (string, error) {
	data, err := s.dataService.GetData(id)
	if err != nil {
		return "", errors.New("data not found")
	}
	return data, nil
}

func (s *Service) UpdateData(id, value string) error {
	return s.dataService.SaveData(id, value)
}

func (s *Service) FetchAndUpdateData(id, newValue string) (string, error) {
	data, err := s.FetchData(id)
	if err != nil {
		return "", err
	}
	if err := s.UpdateData(id, newValue); err != nil {
		return "", err
	}
	return data, nil
}

func (s *Service) IsDataEqual(id, compareValue string) (bool, error) {
	data, err := s.FetchData(id)
	if err != nil {
		return false, err
	}
	return data == compareValue, nil
}
