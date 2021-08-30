package utils

import (
	"errors"
	"log"
	"strconv"
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
	// GetCuorseList gets list of cuorses by given size
	GetCuorseList(count string) ([]dataModels.Cuorse, error)
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

func (c *cuorse) GetCuorseList(count string) ([]dataModels.Cuorse, error) {
	cInt, err := strconv.Atoi(count)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("invalid count value")
	}

	cuorses, err := c.repo.GetCuorseList(cInt)
	if err != nil {
		return nil, errors.New("internal server error")
	}

	return cuorses, nil
}
