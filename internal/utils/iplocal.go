package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

const strMsgErro = "GetAppIP - Problema na carga dos dados de localização/IP - : "

var (
	strGoogle string = "https://www.google.com"
	strIpApi  string = "http://ip-api.com"
	strRegBr  string = "https://registro.br"
)

// Endereços para verificar a conexão
var urlSlice = []string{strGoogle, strIpApi, strRegBr}

func CheckInternetConnection() bool {

	var retCon bool = false

	for _, strURL := range urlSlice {
		for tryCon := 0; tryCon < 3; tryCon++ {
			log.Warn().Msg(fmt.Sprintf("Tentando verificar %s, tentativa %d...\n", strURL, tryCon+1))

			getRes, err := http.Get(strURL)

			if err == nil && getRes.StatusCode == http.StatusOK {
				retCon = true

				log.Info().Msg(fmt.Sprintf("Status de conexão: Conectado à internet (%s)", strURL))

				return retCon
			}

			time.Sleep(time.Second) // Espera 1 segundo entre as tentativas
		}
	}

	return retCon
}

func GetAppIP() map[string]interface{} {

	var ipReturn map[string]interface{}

	getRes, err := http.Get(strIpApi + "/json/")

	if err != nil {
		log.Warn().Msg(strMsgErro + err.Error())

		return ipReturn
	}

	defer getRes.Body.Close()

	bodyRet, err := io.ReadAll(getRes.Body)

	if err != nil {
		log.Warn().Msg(strMsgErro + err.Error())

		return ipReturn
	}

	json.Unmarshal(bodyRet, &ipReturn)

	return ipReturn
}
