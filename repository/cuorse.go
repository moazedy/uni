package repository

import (
	"log"
	"uni/dataModels"
	"uni/repository/couchbaseQueries"

	"github.com/couchbase/gocb/v2"
)

// CuorseRepoInterface is an interface to cuorse in repo to intract with upper layers
type CuorseRepoInterface interface {
	WriteNewCuorse(cuorseData dataModels.Cuorse) (*string, error)
	ReadCuorse(cuorseId string) (*dataModels.Cuorse, error)
	GetCuorseList(count int) ([]dataModels.Cuorse, error)
}

// cuorse is a struct to hold interface methods and intract to lower layers
type cuorse struct {
	// session is interface of lower layer , in this case db
	session *gocb.Cluster
}

// NewCuorseesRepoInterface is a constructor func to create an instance instance for using upper layers in case
func NewCuorsesRepoInterface() CuorseRepoInterface {
	// creating an instance of interface
	var cInterface CuorseRepoInterface

	// creating an instance from cuorse
	c1 := &cuorse{
		session: cluster,
		// session: sql.NewData()
	}

	// telling the c1 to conform the interface
	cInterface = c1

	return cInterface
}

func (c *cuorse) WriteNewCuorse(cuorseData dataModels.Cuorse) (*string, error) {
	_, err := c.session.Query(couchbaseQueries.WriteNewCuorseQuery, &gocb.QueryOptions{
		PositionalParameters: []interface{}{
			cuorseData.Id,
			cuorseData,
		},
	})
	if err != nil {
		log.Println("error on course wirte query execution, error: ", err.Error())
		return nil, err
	}

	return &cuorseData.Id, nil
}

func (c *cuorse) ReadCuorse(cuorseId string) (*dataModels.Cuorse, error) {
	res, err := c.session.Query(couchbaseQueries.ReadCourseQuery, &gocb.QueryOptions{
		PositionalParameters: []interface{}{
			cuorseId,
		},
	})
	if err != nil {
		log.Println("error on read cuorse query exec, error: ", err.Error())
		return nil, err
	}

	var co dataModels.Cuorse
	err = res.One(&co)
	if err != nil {
		if err == gocb.ErrNoResult {
			return nil, nil
		}

		log.Println("error on reading cuorse data, error: ", err.Error())
		return nil, err
	}

	return &co, nil
}

func (c *cuorse) GetCuorseList(count int) ([]dataModels.Cuorse, error) {
	res, err := c.session.Query(couchbaseQueries.GetCuorseListQuery, &gocb.QueryOptions{
		PositionalParameters: []interface{}{
			count,
		},
	})
	if err != nil {
		log.Println("error on GetCuorseList query execution, error: ", err.Error())
		return nil, err
	}

	var cuorses []dataModels.Cuorse

	for res.Next() {
		var c dataModels.Cuorse
		err := res.Row(&c)
		if err != nil {
			if err == gocb.ErrNoResult {
				return cuorses, nil
			}
			return nil, err
		}

		cuorses = append(cuorses, c)
	}

	return cuorses, nil
}
