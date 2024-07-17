package utils

import (
	"fmt"
	"os"
	"strings"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/rs/zerolog/log"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v2"
)

// Definir as mensagens conforme a linguagem localizada pelo IP
func GetLanguageUser(mapIP map[string]interface{}) language.Tag {

	countryCode, _ := mapIP["countryCode"].(string)

	switch {
	case strings.ToUpper(countryCode) == "BR":
		return language.BrazilianPortuguese
	case strings.ToUpper(countryCode) == "US":
		return language.AmericanEnglish
	default:
		return language.BrazilianPortuguese
	}

}

// Inicializar as mensagens
func InitMessages(bundle *i18n.Bundle, languageUser language.Tag, fileYaml string) error {

	fileMessages := GetPathFileLanguage(languageUser, fileYaml)

	// Carrega arquivo de mensagens
	if err := LoadMessages(bundle, languageUser, fileMessages); err != nil {
		log.Warn().Msg(fmt.Sprintf("Error loading messages: %v\n ", err))
	}

	return nil
}

// Carrega os arquivos de mensagens/language
func LoadMessages(bundle *i18n.Bundle, languageUser language.Tag, filePath string) error {
	fileMessages, err := os.Open(filePath)

	if err != nil {
		return err
	}

	defer fileMessages.Close()

	// Decodifica o arquivo YAML em um mapa genérico
	msgMap := make(map[string]interface{})

	if err := yaml.NewDecoder(fileMessages).Decode(&msgMap); err != nil {
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

			bundle.AddMessages(languageUser, &defineMessage)

		case map[interface{}]interface{}:
			// Mensagem com pluralização
			defineMessage.ID = key

			if strValue, ok := msg["text"].(string); ok {
				defineMessage.One = strValue
			}

			if strValue, ok := msg["text"].(string); ok {
				defineMessage.Zero = strValue
			}

			bundle.AddMessages(languageUser, &defineMessage)

		default:
			return fmt.Errorf("unsupported message format for key %s", key)
		}
	}

	return nil
}

// Definir as mensagens conforme a linguagem localizada pelo IP
func GetPathFileLanguage(selectLanguage language.Tag, fileYaml string) string {

	// Definindo o mapa com os caminhos dos arquivos
	paths := map[string]string{
		"en_us": "locale/en-us/screens/" + fileYaml,
		"pt_br": "locale/pt-br/screens/" + fileYaml,
	}

	switch {
	case selectLanguage == language.BrazilianPortuguese:
		return paths["pt_br"]
	case selectLanguage == language.AmericanEnglish:
		return paths["en_us"]
	}

	return ""
}

// Função para localizar uma mensagem no bundle
/*func Localize(bundle *i18n.Bundle, userLanguage language.Tag, messageID string) string {
	localizer := i18n.NewLocalizer(bundle, userLanguage)
	msg, _ := localizer.Localize(&i18n.LocalizeConfig{MessageID: messageID})
	return msg
}*/
