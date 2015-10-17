// Application with GoLang and MongoDB
// Copyright 2015 - Dênis Mendes. All rights reserved.
// Author: Dênis Mendes <denisffmendes@gmail.com>

package service_log

import (
        "log"
        "gopkg.in/mgo.v2"
        "gopkg.in/mgo.v2/bson"
        "application/models"
        "application/messages"
)

func Insert(Log model_log.Log, MongoDBDataBase *mgo.Database)  {
        collection := MongoDBDataBase.C("log")

        err := collection.Insert(Log)
        if err != nil {
                log.Fatalf("%s - %s\n", message_error.InsertDocumentError, err)
        }
}

func ListAll(MongoDBDataBase *mgo.Database) ([]model_log.Log) {
        collection := MongoDBDataBase.C("log")

        var logs []model_log.Log
        err := collection.Find(bson.M{}).All(&logs)
        if err != nil {
                log.Fatalf("%s - %s\n", message_error.ListAllDocumentsError, err)
        }

        return logs
}
