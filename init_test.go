package weedharvester

import (
	"flag"
	"log"
)

var masterURL = flag.String("master", "http://docker:9333", "The url where to find the master")
var filerURL = flag.String("filer", "http://docker:8888", "The url where to find the filer")

func init() {
	flag.Parse()
	log.Printf("Init: masterURL: %s filerURL: %s", *masterURL, *filerURL)
}
