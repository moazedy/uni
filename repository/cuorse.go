package repository

import (
	"errors"
	"uni/dataModels"
	"uni/sql"
)

// CuorseDbInterface is an interface to cuorse in repo to intract with upper layers
type CuorseDbInterface interface {
	WriteNewCuorse(cuorseData dataModels.Cuorse) (*string, error)
	ReadCuorse(cuorseId string) (*dataModels.Cuorse, error)
}

// cuorse is a struct to hold interface methods and intract to lower layers
type cuorse struct {
	// session is interface of lower layer , in this case db
	session sql.Data
}

// NewCuorseesDbInterface is a constructor func to create an instance instance for using upper layers in case
func NewCuorsesDbInterface() CuorseDbInterface {
	// creating an instance of interface
	var cInterface CuorseDbInterface

	// creating an instance from cuorse
	c1 := &cuorse{
		session: sql.RealData,
	}

	// telling the c1 to conform the interface
	cInterface = c1

	return cInterface
}

func (c *cuorse) WriteNewCuorse(cuorseData dataModels.Cuorse) (*string, error) {

	c.session.Cuorses[cuorseData.Id] = cuorseData

	return &cuorseData.Id, nil
}

func (c *cuorse) ReadCuorse(cuorseId string) (*dataModels.Cuorse, error) {

	val, exist := c.session.Cuorses[cuorseId]

	if !exist {
		return nil, errors.New("the cuorse does not exist!")
	}

	return &val, nil
}