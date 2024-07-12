package utils

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var logUseFile bool = true
var logUsePath string = "logs/"
var logUseName string = "logapp"

func CreateLog() bool {

	var retFun bool = true

	// Criar um diretório com permissões padrão (0777 no Unix)
	if err := os.Mkdir(logUsePath, 0755); err != nil {
		// Tratar erro se não for possível criar o diretório
		return retFun
	}

	if logUseFile {

		// Obter a data atual
		dateToday := time.Now()

		monthString := fmt.Sprintf("%02d", dateToday.Month())
		dayString := fmt.Sprintf("%02d", dateToday.Day())

		logUseName += "_" + strconv.Itoa(dateToday.Year()) + "_" + monthString + "_" + dayString

		logDir := logUsePath + logUseName

		// Abrir ou criar o arquivo de log
		fileLog, err := os.OpenFile(logDir, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

		if err != nil {
			log.Fatal().Err(err).Msg("Falha ao abrir/criar o arquivo de log")
		}

		defer fileLog.Close()

		// Configurar zerolog para escrever no arquivo
		logger := zerolog.New(fileLog).With().Timestamp().Logger()

		//Define o output geral para também utilizar o zerolog
		//log.SetOutput(logger)
		logger.Info().Msg("iniciando aplicação")

		/*
			// Exemplo de logging
			log.Info().Msg("Mensagem de informação")
			log.Warn().Str("animal", "gato").Int("size", 10).Msg("Um animal apareceu")
			log.Error().Str("animal", "leão").Int("size", 100).Msg("Animal perigoso")
			log.Fatal().Str("animal", "tigre").Int("size", 80).Msg("Animal fatal encontrado")
		*/
	}

	return retFun
}
