// Application with GoLang and MongoDB
// Copyright 2015 - Dênis Mendes. All rights reserved.
// Author: Dênis Mendes <denisffmendes@gmail.com>

package model_log

import (
        "time"
)

type (
        Log struct {
                Action     string `bson:"action"`
                CreateAt   time.Time `bson:"createAt"`
        }
)
