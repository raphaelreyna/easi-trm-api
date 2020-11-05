package trm

import "encoding/xml"

const apiURL = "https://therightmove.easidemographics.com/EasiAPIServiceSite/EasiAPI.svc"

type envelope struct {
	XMLName xml.Name `xml:"s:Envelope"`
	XMLNSs  string   `xml:"xmlns:s,attr"`
	Body    body     `xml:"s:Body"`
}

type body struct {
	XMLName xml.Name    `xml:"s:Body"`
	Content interface{} `xml:"GetReportResponse"`
}

type getReport struct {
	XMLName xml.Name `xml:"http://tempuri.org/ GetReport"`

	PReptRequest *ReportRequest `xml:"pReptRequest,omitempty" json:"pReptRequest,omitempty"`

	PDeveloperKey string `xml:"pDeveloperKey,omitempty" json:"pDeveloperKey,omitempty"`
}

type reportResponse struct {
	XMLName xml.Name `xml:"http://tempuri.org/ GetReportResponse"`

	GetReportResult ReportResponse `xml:"GetReportResult,omitempty" json:"GetReportResult,omitempty"`
}

type DatasetType string

const (
	DefaultDataset DatasetType = "Default"

	Census DatasetType = "Census"

	Updated DatasetType = "Updated"
)

type FeatureType string

const (
	DefaultFeature FeatureType = "Default"

	Quick FeatureType = "Quick"

	Ring FeatureType = "Ring"

	ThreeRing FeatureType = "ThreeRing"

	Polygon FeatureType = "Polygon"
)

type GeographyType string

const (
	DefaultGeography GeographyType = "Default"

	ZIPCode GeographyType = "ZIPCode"

	City GeographyType = "City"

	County GeographyType = "County"

	CBSA GeographyType = "CBSA"

	State GeographyType = "State"

	US GeographyType = "US"

	BlockGroup GeographyType = "BlockGroup"

	CensusTract GeographyType = "CensusTract"
)
