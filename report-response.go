package trm

import "encoding/xml"

type ReportResponse struct {
	XMLName xml.Name `xml:"s:Body>GetReportResponse>GetReportResult"`

	XMLNSa string `xml:"xmlns:a,attr"`

	XMLNSi string `xml:"xmlns:i,attr"`

	AreaCount int32 `xml:"a:AreaCount,omitempty" json:"AreaCount,omitempty"`

	AreaCount2 int32 `xml:"a:AreaCount2,omitempty" json:"AreaCount2,omitempty"`

	AreaCount3 int32 `xml:"a:AreaCount3,omitempty" json:"AreaCount3,omitempty"`

	BlockGroupCode string `xml:"a:BlockGroupCode,omitempty" json:"BlockGroupCode,omitempty"`

	BlockGroupCount int32 `xml:"a:BlockGroupCount,omitempty" json:"BlockGroupCount,omitempty"`

	BusinessOnly int32 `xml:"a:BusinessOnly,omitempty" json:"BusinessOnly,omitempty"`

	CBSA string `xml:"a:CBSA,omitempty" json:"CBSA,omitempty"`

	CBSACode string `xml:"a:CBSACode,omitempty" json:"CBSACode,omitempty"`

	CensusTractCode string `xml:"a:CensusTractCode,omitempty" json:"CensusTractCode,omitempty"`

	City string `xml:"a:City,omitempty" json:"City,omitempty"`

	CityCode string `xml:"a:CityCode,omitempty" json:"CityCode,omitempty"`

	County string `xml:"a:County,omitempty" json:"County,omitempty"`

	DatasetType string `xml:"a:DatasetType,omitempty" json:"DatasetType,omitempty"`

	DatasetYear int32 `xml:"a:DatasetYear,omitempty" json:"DatasetYear,omitempty"`

	DominantProfile string `xml:"a:DominantProfile,omitempty" json:"DominantProfile,omitempty"`

	FIPS string `xml:"a:FIPS,omitempty" json:"FIPS,omitempty"`

	FeatureType string `xml:"a:FeatureType,omitempty" json:"FeatureType,omitempty"`

	Fields FieldDataSlice `xml:"a:Fields,omitempty" json:"Fields,omitempty"`

	GeographyType string `xml:"a:GeographyType,omitempty" json:"GeographyType,omitempty"`

	Instance string `xml:"a:Instance,omitempty" json:"Instance,omitempty"`

	Latitude float64 `xml:"a:Latitude,omitempty" json:"Latitude,omitempty"`

	Longitude float64 `xml:"a:Longitude,omitempty" json:"Longitude,omitempty"`

	Msg string `xml:"a:Msg,omitempty" json:"Msg,omitempty"`

	MsgID int32 `xml:"a:MsgID,omitempty" json:"MsgID,omitempty"`

	PO string `xml:"a:PO,omitempty" json:"PO,omitempty"`

	Radius float64 `xml:"a:Radius,omitempty" json:"Radius,omitempty"`

	Radius2 float64 `xml:"a:Radius2,omitempty" json:"Radius2,omitempty"`

	Radius3 float64 `xml:"a:Radius3,omitempty" json:"Radius3,omitempty"`

	ReportDescription string `xml:"a:ReportDescription,omitempty" json:"ReportDescription,omitempty"`

	ReportID string `xml:"a:ReportID,omitempty" json:"ReportID,omitempty"`

	ReportName string `xml:"a:ReportName,omitempty" json:"ReportName,omitempty"`

	ReportYear int32 `xml:"a:ReportYear,omitempty" json:"ReportYear,omitempty"`

	Reserved string `xml:"a:Reserved,omitempty" json:"Reserved,omitempty"`

	Reserved2 string `xml:"a:Reserved2,omitempty" json:"Reserved2,omitempty"`

	Reserved3 string `xml:"a:Reserved3,omitempty" json:"Reserved3,omitempty"`

	State string `xml:"a:State,omitempty" json:"State,omitempty"`

	UserData string `xml:"a:UserData,omitempty" json:"UserData,omitempty"`

	ZIPCode string `xml:"a:ZIPCode,omitempty" json:"ZIPCode,omitempty"`
}

type FieldDataSlice struct {
	XMLName xml.Name `xml:"a:Fields"`

	FieldData []FieldData `xml:"a:EasiAPIReport.FieldData,omitempty" json:"EasiAPIReport.FieldData,omitempty"`
}

type FieldData struct {
	XMLName xml.Name `xml:"a:EasiAPIReport.FieldData" json:"-"`

	Description string `xml:"a:EasiDescription,omitempty" json:"Description,omitempty"`

	Name string `xml:"a:EasiName,omitempty" json:"Name,omitempty"`

	Percent string `xml:"a:EasiPercent,omitempty" json:"Percent,omitempty"`

	Rank string `xml:"a:EasiRank,omitempty" json:"Rank,omitempty"`

	Score string `xml:"a:EasiScore,omitempty" json:"Score,omitempty"`

	Value string `xml:"a:EasiValue,omitempty" json:"Value,omitempty"`

	Value2 string `xml:"a:EasiValue2,omitempty" json:"Value2,omitempty"`

	Value3 string `xml:"a:EasiValue3,omitempty" json:"Value3,omitempty"`

	Growth string `xml:"a:EasiGrowth,omitempty" json:"Growth,omitempty"`

	GrowthForecast string `xml:"a:EasiGrowthForecast,omitempty" json:"GrowthForecast,omitempty"`
}
