package trm

import (
	"encoding/xml"
)

type ReportRequest struct {
	XMLName xml.Name `xml:"pReptRequest" json:"-"`

	XMLNSa string `xml:"xmlns:a,attr" json:"-"`

	XMLNSi string `xml:"xmlns:i,attr" json:"-"`

	Address string `xml:"a:Address,omitempty" json:"Address,omitempty"`

	BlockGroup string `xml:"a:BlockGroup,omitempty" json:"BlockGroup,omitempty"`

	CBSA string `xml:"a:CBSA,omitempty" json:"CBSA,omitempty"`

	CarrierRoute string `xml:"a:CarrierRoute,omitempty" json:"CarrierRoute,omitempty"`

	CensusTract string `xml:"a:CensusTract,omitempty" json:"CensusTract,omitempty"`

	City string `xml:"a:City,omitempty" json:"City,omitempty"`

	CongressionalDistrict string `xml:"a:CongressionalDistrict,omitempty" json:"CongressionalDistrict,omitempty"`

	County string `xml:"a:County,omitempty" json:"County,omitempty"`

	Dataset string `xml:"a:Dataset,omitempty" json:"Dataset,omitempty"`

	DatasetType DatasetType `xml:"a:DatasetType,omitempty" json:"DatasetType,omitempty"`

	DatasetYear int32 `xml:"a:DatasetYear,omitempty" json:"DatasetYear,omitempty"`

	FeatureType FeatureType `xml:"a:FeatureType,omitempty" json:"FeatureType,omitempty"`

	GeographyType GeographyType `xml:"a:GeographyType,omitempty" json:"GeographyType,omitempty"`

	GeographyValue string `xml:"a:GeographyValue,omitempty" json:"GeographyValue,omitempty"`

	Latitude float64 `xml:"a:Latitude,omitempty" json:"Latitude,omitempty"`

	Longitude float64 `xml:"a:Longitude,omitempty" json:"Longitude,omitempty"`

	Radius float64 `xml:"a:Radius,omitempty" json:"Radius,omitempty"`

	Radius2 float64 `xml:"a:Radius2,omitempty" json:"Radius2,omitempty"`

	Radius3 float64 `xml:"a:Radius3,omitempty" json:"Radius3,omitempty"`

	ReportID string `xml:"a:ReportID,omitempty" json:"ReportID,omitempty"`

	ReportTag string `xml:"a:ReportTag,omitempty" json:"ReportTag,omitempty"`

	ReportType string `xml:"a:ReportType,omitempty" json:"ReportType,omitempty"`

	Reserved string `xml:"a:Reserved,omitempty" json:"-"`

	Reserved2 string `xml:"a:Reserved2,omitempty" json:"-"`

	Reserved3 string `xml:"a:Reserved3,omitempty" json:"-"`

	SchoolDistrict string `xml:"a:SchoolDistrict,omitempty" json:"SchoolDistrict,omitempty"`

	State string `xml:"a:State,omitempty" json:"State,omitempty"`

	US string `xml:"a:US,omitempty" json:"US,omitempty"`

	UserData string `xml:"a:UserData,omitempty" json:"UserData,omitempty"`

	ZIPCode string `xml:"a:ZIPCode,omitempty" json:"ZIPCode,omitempty"`
}
