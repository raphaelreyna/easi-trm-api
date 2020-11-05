package trm

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"testing"
)

func TestRequestMarshal(t *testing.T) {
	// Create the APICLient to test
	c := NewAPIClient("DEV_KEY", nil)

	rr := &ReportRequest{
		FeatureType:   Ring,
		GeographyType: BlockGroup,
		Latitude:      47.64013,
		Longitude:     -122.129731,
		Radius:        2.5,
		ReportID:      "ARC01",
	}

	// Create the XML for the SOAP request
	env := c.envelope(rr, "")
	out, err := xml.Marshal(&env)
	if err != nil {
		t.Fatalf(err.Error())
	}

	master, err := ioutil.ReadFile("request.xml")
	if err != nil {
		t.Fatalf(err.Error())
	}

	if !bytes.Equal(out, master) {
		t.Fail()
	}
}
