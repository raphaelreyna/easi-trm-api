package main

import (
	"fmt"
	"net/http"
	trm "github.com/raphaelreyna/easi-trm-api"
	"os"
)

func main() {
	retCode := 1
	defer func() {
		os.Exit(retCode)
	}()

// Grab developer key from user
	if len(os.Args) == 1 {
		fmt.Printf("Usage: %s DEVELOPER_KEY\n" +
			"An example application showcasing the Go client for the EASI The Right Move API.\n",
			os.Args[0],
		)
		return
	}
	dk := os.Args[1]
	c := trm.NewAPIClient(dk, http.DefaultClient)

	gr := &trm.ReportRequest{
		FeatureType:   trm.Quick,
		GeographyType: trm.ZIPCode,
		GeographyValue: "91701",
		ReportID:      "AZ001",
	}

	rr, err := c.GetReport(gr, "", nil)
	if err != nil {
		panic(err)
	}

	for _, field := range rr {
		fmt.Printf("%+v\n", *field)
	}
}
