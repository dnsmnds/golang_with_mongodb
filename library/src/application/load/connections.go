// Application with GoLang and MongoDB
// Copyright 2015 - Dênis Mendes. All rights reserved.
// Author: Dênis Mendes <denisffmendes@gmail.com>

package connections

import (
    "gopkg.in/mgo.v2"
    "time"
    "log"
    //--------------------------
    // Application
    //--------------------------
    "application/configurations"
    "application/messages"
    //--------------------------
)

func ConnectWithMongoDB() (*mgo.Database)  {
    dialInfo := &mgo.DialInfo{
        Addrs:    []string{configurations.MongoDBHost},
        Timeout:  60 * time.Second,
        Database: configurations.MongoDBDataBase,
        Username: configurations.MongoDBUsername,
        Password: configurations.MongoDBPassword,
    }

    session, err := mgo.DialWithInfo(dialInfo)

    if err != nil {
        log.Fatalf("%s - %s\n", message_error.ConnectionFailed, err)
    }

    session.SetSafe(&mgo.Safe{})

    database := session.DB(configurations.MongoDBDataBase)

    return database
}
