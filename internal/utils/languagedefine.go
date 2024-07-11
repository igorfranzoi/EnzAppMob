package utils

import (
	"os"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"gopkg.in/yaml.v2"
)

// Carrega os arquivos de mensagens/language
func LoadMessages(bundle *i18n.Bundle, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Decodifica o arquivo YAML em um mapa genérico
	msgs := make(map[string]interface{})
	if err := yaml.NewDecoder(file).Decode(&msgs); err != nil {
		return err
	}

	// Adiciona as mensagens ao bundle usando o pacote catalog
	/*for key, value := range msgs {
		switch msg := value.(type) {
		case string:
			// Mensagem simples
			bundle.AddMessages(language.English, &catalog.Message{
				ID:   key,
				One:  msg,
				Zero: msg, // Pode ser necessário definir Zero se aplicável
			})
		case map[interface{}]interface{}:
			// Mensagem com pluralização
			one := msg["one"].(string)
			other := msg["other"].(string)
			bundle.AddMessages(language.English, &catalog.Message{
				ID:    key,
				One:   one,
				Other: other,
			})
		default:
			return fmt.Errorf("unsupported message format for key %s", key)
		}
	}*/

	return nil
}
