package login

import (
	"enzappmob/internal/utils"
	"fmt"
	"os"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/container"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/rs/zerolog/log"
)

var (
	en_us_message = "locales/en-us/screens/login.yaml"
	pt_br_message = "locales/pt-br/screens/login.yaml"
)

func initApp(bundle *i18n.Bundle) {
	// Carrega arquivos de mensagens
	if err := utils.LoadMessages(bundle, en_us_message); err != nil {
		log.Warn().Msg(fmt.Sprintf("Error loading messages: %v\n ", err))
		//os.Exit(1)
	}

	if err := utils.LoadMessages(bundle, pt_br_message); err != nil {
		log.Warn().Msg(fmt.Sprintf("Error loading messages: %v\n", err))
		//os.Exit(1)
	}

	//Verifica a conexão com a internet
	if utils.CheckInternetConnection() {
		fmt.Sprint(utils.GetAppIP())
	}
}

// Tela de Login
func LoginScreen(mainApp *fyne.App, bundle *i18n.Bundle) bool {

	var loginReturn bool = true

	//Iniciando tela de login
	log.Info().Msg("Iniciando Login Screen")

	//Inicializar
	initApp(bundle)

	//Constrói o app inicial, utilizado por todas as aplicações
	windowLogin := (*mainApp).NewWindow("EnzTech-Mobile")
	windowLogin.Resize(fyne.NewSize(400, 400))

	// Carregar a imagem de fundo (substitua o caminho pelo seu arquivo de imagem)
	backgroundImage := canvas.NewImageFromFile("images/mainBackgroundImage.jpg")
	if backgroundImage != nil {
		// Redimensionar a imagem para cobrir toda a janela
		backgroundImage.FillMode = canvas.ImageFillStretch
	} else {
		log.Warn().Msg("Erro ao carregar a imagem de fundo")
	}

	// Obter o diretório de trabalho atual
	completePath, err := os.Getwd()
	if err != nil {
		fmt.Println("Erro ao obter o diretório de trabalho:", err)
	} else {
		//Exemplo de ícone (pode ser substituído)
		rscImg, _ := utils.ResourcePathLoad(completePath + "/images/enztech_favicon.png")
		windowLogin.SetIcon(rscImg)
	}

	fieldForm := widget.NewForm(
		widget.NewFormItem("UserName", widget.NewEntry()),
		widget.NewFormItem("Password", widget.NewPasswordEntry()),
	)

	// working on cancel and submit functions of form
	fieldForm.OnCancel = func() {
		lblButton := widget.NewLabel("")
		lblButton.Text = "Canceled"
		lblButton.Refresh()
	}
	fieldForm.OnSubmit = func() {
		lblButton := widget.NewLabel("")
		lblButton.Text = "submitted"
		lblButton.Refresh()
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
		container.NewCenter(
			layout.NewSpacer(),
			layout.NewSpacer(),
		),
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
		fmt.Println("Fechando a aplicação")
		(*mainApp).Quit() // Garante que a aplicação seja fechada corretamente
	})

	windowLogin.ShowAndRun()

	//Finalizando tela de login
	log.Info().Msg("Finalizando Login Screen")

	return loginReturn
}
