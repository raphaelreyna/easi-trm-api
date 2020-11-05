package main

import (
	"math/rand"
	"time"
	"os"
	"log"
	"fmt"
)

type ctxKey struct{}

var key ctxKey = ctxKey{}

var Version string = "INVALID BUILD"

func main() {
	retCode := 1
	defer func() {
		os.Exit(retCode)
	}()
	rand.Seed(time.Now().UnixNano())

	var version bool
	s, err := serverFromFlags(nil, &version)
	if err != nil {
		log.Println(err)
		return
	}

	// Check if the user just wants the version info
	if version {
		fmt.Printf("%s version: %s\n", os.Args[0], Version)
		retCode = 0
		return
	}

	if err := s.ListenAndServe(); err != nil {
		log.Println(err)
		return
	}

	retCode = 0
}
