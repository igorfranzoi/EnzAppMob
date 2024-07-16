package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

var (
	strGoogle string = "https://www.google.com"
	strIpApi  string = "http://ip-api.com"
	strRegBr  string = "https://registro.br"
)

// Endereços para verificar a conexão
var urlSlice = []string{strGoogle, strIpApi, strRegBr}

// Estrutura ip-api que corresponde à estrutura do JSON
type IPstruct struct {
	Status      string `json:"status"`
	Country     string `json:"country"`
	CountryCode string `json:"countryCode"`
	Region      string `json:"region"`
	RegionName  string `json:"regionName"`
	City        string `json:"city"`
	ZipCode     string `json:"zip"`
	Latitude    string `json:"lat"`
	Longitude   string `json:"lon"`
	TimeZone    string `json:"timezone"`
	Isp         string `json:"isp"`
	OrgIsp      string `json:"org"`
	As          string `json:"as"`
	IpNumber    string `json:"query"`
}

func CheckInternetConnection() bool {

	var retCon bool = false

	for _, strURL := range urlSlice {
		for tryCon := 0; tryCon < 3; tryCon++ {
			log.Warn().Msg(fmt.Sprintf("Tentando verificar %s, tentativa %d...\n", strURL, tryCon+1))

			resGet, err := http.Get(strURL)

			if err == nil && resGet.StatusCode == http.StatusOK {
				retCon = true

				log.Info().Msg(fmt.Sprintf("Status de conexão: Conectado à internet (%s)", strURL))

				return retCon
			}

			time.Sleep(time.Second) // Espera 1 segundo entre as tentativas
		}
	}

	return retCon
}

func GetAppIP() (IPstruct, error) {

	var ipReturn IPstruct
	var retVld bool

	// Estrutura para armazenar o JSON deserializado
	var dataAux map[string]interface{}

	resGet, err := http.Get(strIpApi + "/json/")

	if err != nil {
		return ipReturn, err
	}

	defer resGet.Body.Close()

	bodyRet, err := io.ReadAll(resGet.Body)

	if err != nil {
		return ipReturn, err
	}

	json.Unmarshal(bodyRet, &dataAux)

	// Verificando se o atributo "email" existe
	ipReturn.Status, retVld = dataAux["status"].(string)

	if !retVld {
		fmt.Println("O atributo 'email' não existe ou não é uma string")
	}

	return ipReturn, err
}
