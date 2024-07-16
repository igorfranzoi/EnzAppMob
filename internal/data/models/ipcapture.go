package models

/*
{
    "status": "success",
    "country": "Brazil",
    "countryCode": "BR",
    "region": "SP",
    "regionName": "São Paulo",
    "city": "São Paulo",
    "zip": "68515",
    "lat": -23.5335,
    "lon": -46.6359,
    "timezone": "America/Sao_Paulo",
    "isp": "Vivo",
    "org": "TELEF�NICA BRASIL S.A",
    "as": "AS27699 TELEFÔNICA BRASIL S.A",
    "query": "187.75.77.219"
}
*/

type IpCapture struct {
	IpStatus           string  `json:"ipStatus"`           //status
	CountryDescription string  `json:"countryDescription"` //country
	CountryAcronym     string  `json:"countryAcronym"`     //countryCode
	Region             string  `json:"region"`             //region
	RegionName         string  `json:"regionName"`         //regionName
	CityName           string  `json:"cityName"`           //city
	ZipCode            string  `json:"zip"`                //city
	Latitude           float64 `json:"latitude"`           //lat
	Longitude          float64 `json:"longitude"`          //long
	TimeZoneLocal      string  `json:"timezone"`           //timezone
	ProviderInternet   string  `json:"internetServer"`     //isp
	Company            string  `json:"Company"`            //org
	AS                 string  `json:"as"`                 //as
	NumberIp           string  `json:"ipNumber"`           //query
}
