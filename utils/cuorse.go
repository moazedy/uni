package utils

import (
	"errors"
	"uni/dataModels"
	"uni/repository"

	"github.com/google/uuid"
)

// CuorseInterface an interface for connect with upper layers
type CuorseInterfaceUtils interface {
	// CreateCurse addes a new cuorse to db
	CreateCuorse(name string) (*string, error)
	// GetCuorseData returns data of target Course
	GetCuorseData(cuorseId string) (*dataModels.Cuorse, error)
}

// cuorse is a struct to hold interface methods and connect to lower layers
type cuorse struct {
	repo repository.CuorseRepoInterface
}

// NewCuorseInterface is constructory func for creating an instance of CuorseInterface
func NewCuorseInterfaceUtils() CuorseInterfaceUtils {
	var CuorseInterface CuorseInterfaceUtils

	c := &cuorse{
		repo: repository.NewCuorsesRepoInterface(),
	}

	CuorseInterface = c

	return CuorseInterface
}

func (c *cuorse) CreateCuorse(name string) (*string, error) {

	newCuorse := dataModels.Cuorse{
		Id:   uuid.New().String(),
		Name: name,
	}

	err := newCuorse.Validate()
	if err != nil {
		return nil, err
	}

	id, err := c.repo.WriteNewCuorse(newCuorse)
	if err != nil {
		return nil, err
	}

	return id, nil
}

func (c *cuorse) GetCuorseData(cuorseId string) (*dataModels.Cuorse, error) {

	_, err := uuid.Parse(cuorseId)
	if err != nil {
		return nil, errors.New("invalid id")
	}

	cuorseData, err := c.repo.ReadCuorse(cuorseId)
	if err != nil {
		return nil, err
	}

	return cuorseData, nil
}
