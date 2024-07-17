package ui

import (
	"enzappmob/internal/data/database"
	"enzappmob/internal/ui/appconfig"
	"enzappmob/internal/ui/login"
	"enzappmob/internal/utils"

	"fyne.io/fyne/app"
	"fyne.io/fyne/dialog"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v2"
)

const appVersion string = "0.0.1"

func InitializeApp() {

	var err error
	var appConfig appconfig.AppGeneral

	//Disponibiliza o arquivo de log do aplicativo
	utils.CreateLog()

	//Verifica a conexão com a internet
	if utils.CheckInternetConnection() {
		appConfig.MapIPConnect = utils.GetAppIP()

		//Se existe conexão, a aplicação deve saber que a validação de usuário/senhas e os processos podem ser realizados de forma on-line
		if appConfig.MapIPConnect != nil {
			appConfig.SignalConnect = true
		} else {
			appConfig.SignalConnect = false
		}
	}

	//Instância a linguagem que será utilizada pelo APP
	appConfig.CurrentLanguage = utils.GetLanguageUser(appConfig.MapIPConnect)

	//Realiza a conexão com o banco de dados
	appConfig.GormConnect, err = database.Connect()

	// Determina o idioma a ser utilizado (default: inglês)
	appConfig.Bundle = i18n.NewBundle(appConfig.CurrentLanguage)
	appConfig.Bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)

	//Constrói o app inicial, utilizado por todas as aplicaçãoes
	enzTechApp := app.New()

	if err != nil {
		dialog.ShowInformation("Alert", "Problem connect database.", app.New().NewWindow(""))
	}

	appConfig.MainApp = &enzTechApp

	//loginRetu\rn := login.LoginScreen(&enzTechApp, bundle, ctrInternetConnection, dbConnect)
	loginReturn := login.LoginScreen(&appConfig)

	if loginReturn {
		log.Info().Str("app-version", appVersion).Msg("App-Version")
	}
}
