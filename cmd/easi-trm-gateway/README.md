# EASI - The Right Move API JSON Gateway

Currently, the only way to consume the [EASI - The Right Move API](https://www.easidemographics.com/MobileApp/APIInfo.asp) is via their official C# client or through a SOAP API.
This is meant to act as a JSON <-> SOAP gateway for the API.
Users simply make JSON requests against this gateway rather than SOAP requests to the EASI API server(s) and will receive a JSON response.

## Features:
* Visibility: simple and full health checks along with extensive logging let you see whats going on.
* Flexibility: supports setting default developer keys and transparently passing along HTTP headers.
* Statelessness: deploys easily, starts up fast and scales quickly.
* Visibility: extensive logging so you can see exactly whats going on when it matters most.
* Ease of Deployment: as a single executable, there's not much to deploy; check out the [Docker image](https://hub.docker.com/r/raphaelreyna/easi-trm-gateway) to get started even faster.
* Speed: benchmarks show that on average, requests made through the gateway only incur a 10ms penalty.


### Obtaining
There are several options here:
* Build from source (requires a working [Go](https://go.dev) installation):
```bash
git clone https://github.com/raphaelreyna/easi-trm-api
cd easi-trm-api/cmd/easi-trm-gateway
make
sudo mv easi-trm-gateway /bin/easi-trm-api 
```

* Using the prebuilt Docker image:
```bash
docker pull raphaelreyna/easi-trm-gateway
```

* Download a prebuilt executable from the [releases page](https://github.com/raphaelreyna/easi-trm-api/releases)

### How To Use:
This gateway expects a flat JSON object containing fields with the same name as the elements in the pReptRequest element in the SOAP XML.
The developer key is given by the user by setting the "DevKey" field in the JSON object.
Requests for reports must be directed at the `/report` path.

#### Health Checks:
Health checks requests should be directed at either `/health` or `/health/full`.

```
Usage of easi-trm-gateway:
  -dev-key string
    	Path to file containing a valid dev key.
    	If set, the dev key in the file will be used as the default dev key for all requests from all clients.
  -host string
    	A reachable hostname.
  -internal-net string
    	CIDR string describing the network allowed to request health checks.
  -port string
    	Port to bind to. (default "80")
  -tls-cert string
    	Path to the key file to use for TLS/SSL/HTTPS.
  -tls-key string
    	Path to the key file to use for TLS/SSL/HTTP.
  -v	Show version information.

```

## Example JSON request:
Sending the following JSON body the the `/report` path as an HTTP POST request ...
```javascript
{
    "FeatureType": "Quick",
    "GeographyType": "ZIPCode",
    "GeographyValue": "91701",
    "ReportID": "AZ001", 
    "DevKey": "VALID_DEV_KEY",
}
```
...will result in the JSON response shown below.
In the event that there is a non-connectivity related issue with the request, there will be a message in the `EasiReportMessage` field explaining what went wrong.
```javascript
{
  "EasiReportMessage": "",
  "EasiReportFieldData": [
    {
      "Description": "Households, 1 Person",
      "Name": "HH1",
      "Rank": "14,338",
      "Score": "B-",
      "Value": "127"
    },
    {
      "Description": "Families, Married with Children Under 18",
      "Name": "FAMMAR18",
      "Rank": "35,314",
      "Score": "E ",
      "Value": "22"
    },
    {
      "Description": "Retired Workers Profile",
      "Name": "RETIRED",
      "Rank": "31,412",
      "Score": "D-",
      "Value": "51"
    },
    {
      "Description": "Above Average Education",
      "Name": "AB_AV_EDU",
      "Rank": "1,762",
      "Score": "A",
      "Value": "186"
    },
    {
      "Description": "EASI Total Crime Index (US Avg=100; A=High)",
      "Name": "TOTCRIME",
      "Rank": "29,598",
      "Score": "D-",
      "Value": "84"
    },
    {
      "Description": "Expensive Homes",
      "Name": "EXP_HOMESM",
      "Rank": "3,982",
      "Score": "A ",
      "Value": "180"
    },
    {
      "Description": "Housing, Median Value Owner Households ($)",
      "Name": "MEDVALOCC",
      "Rank": "4,039",
      "Score": "A ",
      "Value": "328,784"
    },
    {
      "Description": "Population (4/1/2010)",
      "Name": "POP10",
      "Rank": "29,475",
      "Score": "D-",
      "Value": "10,554"
    },
    {
      "Description": "Households (4/1/2010)",
      "Name": "HH10",
      "Rank": "32,001",
      "Score": "E",
      "Value": "3,691"
    },
    {
      "Description": "Population, Median Age",
      "Name": "MEDAGE",
      "Rank": "37,391",
      "Score": "E-",
      "Value": "24.8"
    },
    {
      "Description": "Families",
      "Name": "FAMILIES",
      "Rank": "36,709",
      "Score": "E ",
      "Value": "1,706"
    },
    {
      "Description": "Household Income, Total ($)",
      "Name": "TOTHHINC",
      "Rank": "19,300",
      "Score": "C ",
      "Value": "366,930,011"
    },
    {
      "Description": "Household Income, Median ($)",
      "Name": "MEDHHINC",
      "Rank": "30,120",
      "Score": "D-",
      "Value": "49,022"
    },
    {
      "Description": "Household Income, Average ($)",
      "Name": "AVGHHINC",
      "Rank": "13,846",
      "Score": "B-",
      "Value": "88,609"
    },
    {
      "Description": "Personal Income, Per Capita ($)",
      "Name": "PERCAPINC",
      "Rank": "18,124",
      "Score": "C",
      "Value": "33,316"
    },
    {
      "Description": "White Population, Alone",
      "Name": "WHPOPA",
      "Rank": "22,243",
      "Score": "C-",
      "Value": "9,452"
    },
    {
      "Description": "Black Population, Alone",
      "Name": "BLPOPA",
      "Rank": "20,990",
      "Score": "C ",
      "Value": "144"
    },
    {
      "Description": "Asian Population, Alone",
      "Name": "ASPOPA",
      "Rank": "3,603",
      "Score": "A ",
      "Value": "844"
    },
    {
      "Description": "Hispanic Population",
      "Name": "HISPPOP",
      "Rank": "14,303",
      "Score": "B-",
      "Value": "690"
    },
    {
      "Description": "Housing, Units",
      "Name": "HOUSEUNITS",
      "Rank": "32,855",
      "Score": "E",
      "Value": "4,344"
    },
    {
      "Description": "Housing, Owner Occupied",
      "Name": "OOCCHH",
      "Rank": "33,030",
      "Score": "E",
      "Value": "2,078"
    },
    {
      "Description": "Housing, Renter Occupied",
      "Name": "ROCCHH",
      "Rank": "5,803",
      "Score": "A-",
      "Value": "2,063"
    },
    {
      "Description": "Housing, Vacant Units",
      "Name": "VACUNIT",
      "Rank": "34,409",
      "Score": "E",
      "Value": "203"
    },
    {
      "Description": "Employees, Total (by Place of Work)",
      "Name": "NAIC_T_EMP",
      "Rank": "10,067",
      "Score": "B",
      "Value": "3,819"
    }
  ]
}
```
