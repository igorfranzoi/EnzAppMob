package login

import (
	"enzappmob/internal/ui/appconfig"
	"enzappmob/internal/utils"
	"errors"
	"fmt"
	"os"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/container"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/rs/zerolog/log"
)

var languageFile string = "login.yaml"

var (
	ErrInvalidLogin = errors.New("Login failed: Invalid credentials")
)

// Tela de Login
// func LoginScreen(mainApp *fyne.App, bundle *i18n.Bundle, connected bool, gormConnect *gorm.DB) bool {
func LoginScreen(mainApp *appconfig.AppGeneral) bool {

	var loginReturn bool = true

	//Iniciando tela de login
	log.Info().Msg("Iniciando Login Screen")

	//Inicializar mensagens para a tela de login
	utils.InitMessages(mainApp.Bundle, mainApp.CurrentLanguage, languageFile)

	//Constrói o app inicial, utilizado por todas as aplicações
	windowLogin := (*mainApp.MainApp).NewWindow("EnzTech-Mobile")
	windowLogin.Resize(fyne.NewSize(400, 400))

	// Carregar a imagem de fundo
	backgroundImage := canvas.NewImageFromFile("images/mainBackgroundImage.jpg")
	if backgroundImage != nil {
		// Redimensionando a imagem para cobrir toda a janela
		backgroundImage.FillMode = canvas.ImageFillStretch
	} else {
		log.Warn().Msg("Erro ao carregar a imagem de fundo")
	}

	// Obter o diretório de trabalho atual
	completePath, err := os.Getwd()
	if err != nil {
		log.Warn().Msg("Erro ao obter o diretório de trabalho: " + err.Error())
	} else {
		//Exemplo de ícone (pode ser substituído)
		rscImg, _ := utils.ResourcePathLoad(completePath + "/images/enztech_favicon.png")
		windowLogin.SetIcon(rscImg)
	}

	entryUserName := widget.NewEntry()
	entryPassword := widget.NewPasswordEntry()

	//bundle  = i18n.Localizer

	fieldForm := widget.NewForm(
		widget.NewFormItem("UserName", entryUserName),
		//widget.NewFormItem(mainApp.Bundle, entryUserName),
		widget.NewFormItem("Password", entryPassword),
	)

	// working on cancel and submit functions of form
	fieldForm.OnCancel = func() {
		lblButton := widget.NewLabel("")
		lblButton.Text = "Canceled"
		lblButton.Refresh()

		(*mainApp.MainApp).Quit()
	}

	fieldForm.OnSubmit = func() {
		strUsername := entryUserName.Text
		strPassword := entryPassword.Text

		// Lógica de validação do usuário e senha
		if strUsername == "admin" && strPassword == "admin" {
			// Login válido, podemos fechar a janela de login
			windowLogin.Close()

			// Chamando a próxima tela (exemplo: menu)
			//MenuScreen(mainApp, bundle) // Implemente a função MenuScreen

			// Definindo o retorno como verdadeiro
			loginReturn = true
		} else {
			log.Warn().Str("UserName", strUsername).Msg("Um animal apareceu")

			dialog.ShowError(ErrInvalidLogin, windowLogin)
		}
	}

	// Botão com ícone para ir para a tela de configuração
	btnSettings := widget.NewButtonWithIcon("", theme.SettingsIcon(), func() {
		fmt.Print("testes")
	})

	loginForm := container.NewVBox(
		container.NewCenter(container.NewHBox(
			container.NewCenter(widget.NewIcon(theme.HomeIcon())),
		)),
		fieldForm,
		layout.NewSpacer(),
		container.NewHBox(
			layout.NewSpacer(), // Espaçador à esquerda
			container.NewMax(
				layout.NewSpacer(), // Espaçador
				container.NewHBox(
					layout.NewSpacer(), // Espaçador entre o conteúdo e o botão
					btnSettings,        // Botão Settings
				),
				layout.NewSpacer(), // Espaçador
			),
		),
	)

	windowLogin.SetContent(container.NewVBox(backgroundImage, loginForm))

	// Interceptando o evento de fechamento da janela para encerrar a aplicação
	windowLogin.SetCloseIntercept(func() {
		log.Info().Msg("Fechando a aplicação - SetCloseIntercept")

		(*mainApp.MainApp).Quit() // Garante que a aplicação seja fechada corretamente
	})

	windowLogin.ShowAndRun()

	//Finalizando tela de login
	log.Info().Msg("Finalizando Login Screen")

	return loginReturn
}
