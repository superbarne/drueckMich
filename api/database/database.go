package database;

import "gopkg.in/mgo.v2"

var session, sessionErr = mgo.Dial("mongodb://localhost")
var DB = session.DB("HA18DB_barne_juergensen_590328")