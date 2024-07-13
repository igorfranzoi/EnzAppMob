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
	strIpApi  string = "https://ip-api.com"
	strRegBr  string = "https://registro.br"
)

// Endereços para verificar a conexão
var urlSlice = []string{strGoogle, strIpApi, strRegBr}

type IP struct {
	Query string
}

func CheckInternetConnection() bool {

	var retCon bool = false

	for _, strURL := range urlSlice {
		for tryCon := 0; tryCon < 3; tryCon++ {
			log.Warn().Msg(fmt.Sprintf("Tentando verificar %s, tentativa %d...\n", strURL, tryCon+1))

			resGet, err := http.Get(strURL)

			if err == nil && resGet.StatusCode == http.StatusOK {
				log.Info().Msg(fmt.Sprintf("Status de conexão: Conectado à internet (%s)", strURL))

				return retCon
			}

			time.Sleep(time.Second) // Espera 1 segundo entre as tentativas
		}
	}

	return retCon
}

func GetAppIP() string {

	var ipReturn IP

	resGet, err := http.Get(strIpApi + "/json/")

	if err != nil {
		return err.Error()
	}

	defer resGet.Body.Close()

	bodyRet, err := io.ReadAll(resGet.Body)

	if err != nil {
		return err.Error()
	}

	json.Unmarshal(bodyRet, &ipReturn)

	return ipReturn.Query
}
