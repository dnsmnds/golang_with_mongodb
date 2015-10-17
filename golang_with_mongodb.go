// Application with GoLang and MongoDB
// Copyright 2015 - Dênis Mendes. All rights reserved.
// Author: Dênis Mendes <denisffmendes@gmail.com>

package main

import (
	"flag"
	"fmt"
	"time"
	"application/load"
	"application/models"
	"application/services"
)

func main() {
 	mongoDBDataBase := connections.ConnectWithMongoDB()

	flag.Parse()

	argumentOption := flag.Arg(0)
	switch argumentOption {
		case "i":
			argumentLogAction := flag.Arg(1)

			Log := model_log.Log{Action: argumentLogAction, CreateAt: time.Now() }
			service_log.Insert(Log, mongoDBDataBase)
		case "l":
			fmt.Printf("Found document:\n %+v\n", service_log.ListAll(mongoDBDataBase))
	}
}
