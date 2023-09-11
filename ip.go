package main

import "time"

// IP struct to hold IP information
type IP struct {
	Address string    `bson:"address"`
	Time    time.Time `bson:"time"`
}
