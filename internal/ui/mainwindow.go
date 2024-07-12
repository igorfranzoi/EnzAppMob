package ui

import (
	"enzappmob/internal/ui/login"
	"enzappmob/internal/utils"
	"fmt"

	"fyne.io/fyne/app"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v2"
)

var appVersion string

func InitializeApp() {

	//Disponibiliza o arquivo de log do aplicativo
	utils.CreateLog()

	// Determina o idioma a ser utilizado (default: inglês)
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)

	//Constrói o app inicial, utilizado por todas as aplicaçãoes
	enzTechApp := app.New()

	loginReturn := login.LoginScreen(&enzTechApp, bundle)

	if loginReturn {
		fmt.Printf("app-version %s", appVersion)
	}
}
