package main

import (
	"fmt"
	"github.com/c12s/magnetar/config"
	"github.com/c12s/magnetar/service"
	"github.com/c12s/magnetar/storage/etcd"
	"github.com/c12s/magnetar/storage/influx"
	"github.com/c12s/magnetar/sync/nats"
	"time"
)

func main() {
	conf, err := config.ConfigFile()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	sync, err := nats.New(conf.Sync, conf.Health)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	db, err := etcd.New(conf.Endpoints, 10*time.Second)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	mdb, err := influx.New(conf.Metrics)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	service.Run(db, mdb, sync, conf.Address)
}
