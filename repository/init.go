package repository

import (
	"log"
	"time"

	"github.com/couchbase/gocb/v2"
)

var cluster *gocb.Cluster

func Init() {
	//sql.Init()
	couchbaseConnection()

}

func couchbaseConnection() {
	clu, err := gocb.Connect(
		"127.0.0.1:8091",
		gocb.ClusterOptions{
			Username: "admin",
			Password: "adminadmin",
		})
	if err != nil {
		log.Println("error on connecting with cluster")
		panic(err)
	}
	// Waiting for the cluster to be ready
	err = clu.WaitUntilReady(time.Second*25,
		&gocb.WaitUntilReadyOptions{
			DesiredState: gocb.ClusterStateOnline,
		},
	)
	if err != nil {
		log.Println("error on wahit until ready")
		panic(err)
	}
	// Pinging cluster to ensure query service working
	pingResult, err := clu.Ping(&gocb.PingOptions{
		ServiceTypes: []gocb.ServiceType{gocb.ServiceTypeQuery},
	})
	if err != nil {
		panic(err)
	}
	log.Println("couchbase connection ... ")
	log.Printf("%#v \n", pingResult.Services)

	cluster = clu
}
