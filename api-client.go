package trm

import (
	"bytes"
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"

	"github.com/beevik/etree"
)

// APIClient handles communicating with the EASI servers
type APIClient struct {
	devKey     string
	httpClient *http.Client
}

// NewAPIClient initializes an APIClient into a ready-to-use state
func NewAPIClient(devKey string, c *http.Client) *APIClient {
	return &APIClient{
		devKey:     devKey,
		httpClient: c,
	}
}

// (envelope appropriately wraps the ReportRequest as SOAP XML
func (c *APIClient) envelope(r *ReportRequest, devKey string) *envelope {
	// Set the xmlns:a and xmlns:i attributes
	r.XMLNSa = "http://schemas.datacontract.org/2004/07/EasiAPI"
	r.XMLNSi = "http://www.w3.org/2001/XMLSchema-instance"

	dk := c.devKey
	if devKey != "" { dk = devKey }

	return &envelope{
		XMLNSs: "http://schemas.xmlsoap.org/soap/envelope/",
		Body: body{
			Content: &getReport{
				PReptRequest:  r,
				PDeveloperKey: dk,
			},
		},
	}
}

// GetReportContext requests a report from the EASI servers using a context
func (c *APIClient) GetReportContext(ctx context.Context, request *ReportRequest, devKey string, h http.Header) ([]*FieldData, string, error) {
	// Create the XML for the SOAP request
	env := c.envelope(request, devKey)
	out, err := xml.Marshal(&env)
	if err != nil {
		return nil, "", err
	}

	// Create the http request to the API server
	req, err := http.NewRequest("POST", apiURL, bytes.NewReader(out))
	if err != nil {
		return nil, "", err
	}
	if h != nil { req.Header = h }
	req.Header.Set("Content-Type", "text/xml; charset=utf-8")
	req.Header.Set("SOAPAction", "http://tempuri.org/IEasiAPIService/GetReport")

	// Make the request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, "", err
	}


	// Make sure we got an OK status from the server
	if resp.StatusCode != http.StatusOK {
		return nil, "", fmt.Errorf("error from server: %s", resp.Status)
	}

	defer resp.Body.Close()

	return extractFieldData(resp.Body)
}

// GetReport requests a report from the EASI servers
func (c *APIClient) GetReport(request *ReportRequest, devKey string, h http.Header) ([]*FieldData, string, error) {
	return c.GetReportContext(
		context.Background(),
		request, devKey, h,
	)
}

var msgPath = etree.MustCompilePath(".//a:Msg")
var fieldsPath = etree.MustCompilePath("[EasiDescription]")
// extractFieldData essentially unmarshals the XML SOAP response from the EASI servers into
// a slice of *FieldData.
// Ideally, we would using encoding/xml from the standard library to unmarshal but
// it's only producing zero-values for some reason...
// Using the github.com/beevik/etree package is ***currently*** a work-around
func extractFieldData(r io.Reader) ([]*FieldData, string, error) {
	// Grab the xml document from the the io.Reader
	doc := etree.NewDocument()
	if _, err := doc.ReadFrom(r); err != nil {
		return nil, "", err
	}

	// Grab the root element of the document
	re := doc.Root()

	// Extract the message from EASI
	var msg string
	msgEl := re.FindElementPath(msgPath)
	if msgEl != nil {
		msg = msgEl.Text()
	}

	// Grab all elements in the XML document that contain child with the name 'EasiDescription'
	// These elements are the EasiReport.FieldData elements
	els := re.FindElementsPath(fieldsPath)

	// Create the slice of FieldData pointers we'll be returning
	fields := []*FieldData{}

	// Range over the EasiReport.FieldData elements and extract their data into a struct
	for _, el := range els {
		fd := &FieldData{}
		for _, child := range el.ChildElements() {
			text := child.Text()
			if text == "" {
				continue
			}
			switch child.Tag {
			case "EasiDescription":
				fd.Description = text
			case "EasiName":
				fd.Name = text
			case "EasiPercent":
				fd.Percent = text
			case "EasiRank":
				fd.Rank = text
			case "EasiScore":
				fd.Score = text
			case "EasiValue":
				fd.Value = text
			case "EasiValue2":
				fd.Value2 = text
			case "EasiValue3":
				fd.Value3 = text
			case "EasiGrowth":
				fd.Growth = text
			case "EasiForecast":
				fd.GrowthForecast = text
			default:
				continue
			}
		}
		fields = append(fields, fd)
	}

	return fields, msg, nil
}
