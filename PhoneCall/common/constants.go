package common

import "os"

var User = os.Getenv("USER")
var Password = os.Getenv("PASSWORD")
var NameDB = os.Getenv("NAME_DATABASE")
var Port = os.Getenv("PORT")
var LinkRabbit = os.Getenv("KEY_RABBITMQ")
