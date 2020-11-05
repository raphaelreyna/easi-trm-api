package main

import (
	"fmt"
	"net/http"
	trm "github.com/raphaelreyna/easi-trm-api"
	"os"
	"encoding/json"
)

func main() {
	retCode := 1
	defer func() {
		os.Exit(retCode)
	}()

// Grab developer key from user
	if len(os.Args) != 3 {
		fmt.Printf("Usage: %s DEVELOPER_KEY US_ZIP_CODE\n" +
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
		GeographyValue: os.Args[2],
		ReportID:      "AZ001",
	}

	rr, msg, err := c.GetReport(gr, "", nil)
	if err != nil {
		panic(err)
	} else if msg != "" {
		fmt.Println(msg)
	}

	output := map[string]interface{}{
		"EasiMessage": msg,
		"EasiReportFields": rr,
	}

	data, err := json.MarshalIndent(output, "", "  ")
	if err != nil { panic(err) }

	os.Stdout.Write(append(data, 10)) // Write to stdout with a newline (0x10)
	retCode = 0
}

