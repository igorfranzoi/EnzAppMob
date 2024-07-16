package utils

import (
	"fmt"
	"os"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/rs/zerolog/log"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v2"
)

var (
	en_us_message = "locales/en-us/screens/login.yaml"
	pt_br_message = "locales/pt-br/screens/login.yaml"
)

// Carrega os arquivos de mensagens/language
func LoadMessages(bundle *i18n.Bundle, filename string) error {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Print("primeiro if")
		return err
	}
	defer file.Close()

	// Decodifica o arquivo YAML em um mapa genérico
	msgMap := make(map[string]interface{})

	if err := yaml.NewDecoder(file).Decode(&msgMap); err != nil {
		fmt.Print("segundo if")
		return err
	}

	// Incluir as mensagens ao bundle
	for key, value := range msgMap {

		var defineMessage i18n.Message

		switch msg := value.(type) {
		case string:
			// Mensagem simples
			defineMessage.ID = key
			defineMessage.One = msg
			defineMessage.Zero = msg

			bundle.AddMessages(language.English, &defineMessage)

		case map[interface{}]interface{}:
			// Mensagem com pluralização
			defineMessage.ID = key
			defineMessage.One = msg["one"].(string)
			defineMessage.Zero = msg["other"].(string)

			bundle.AddMessages(language.English, &defineMessage)

		default:
			return fmt.Errorf("unsupported message format for key %s", key)
		}
	}

	return nil
}

// Inicializar as mensagens
func InitMessages(bundle *i18n.Bundle /*, filename string*/) error {

	filename := en_us_message

	// Carrega arquivos de mensagens
	if err := LoadMessages(bundle, filename); err != nil {
		log.Warn().Msg(fmt.Sprintf("Error loading messages: %v\n ", err))
	}

	/*if err := utils.LoadMessages(bundle, pt_br_message); err != nil {
		log.Warn().Msg(fmt.Sprintf("Error loading messages: %v\n", err))
	}*/

	return nil
}
