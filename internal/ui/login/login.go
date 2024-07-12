package login

import (
	"enzappmob/internal/utils"
	"fmt"

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

// Tela de Login
func LoginScreen(mainApp *fyne.App, bundle *i18n.Bundle) bool {

	var loginReturn bool = true

	//Iniciando tela de login
	log.Info().Msg("Iniciando Login Screen")

	// Carrega arquivos de mensagens
	if err := utils.LoadMessages(bundle, en_us_message); err != nil {
		log.Warn().Msg(fmt.Sprintf("Error loading messages: %v\n ", err))
		//os.Exit(1)
	}

	if err := utils.LoadMessages(bundle, pt_br_message); err != nil {
		log.Warn().Msg(fmt.Sprintf("Error loading messages: %v\n", err))
		//os.Exit(1)
	}

	//Constrói o app inicial, utilizado por todas as aplicações
	windowLogin := (*mainApp).NewWindow("EnzTech-Mobile")
	windowLogin.SetIcon(theme.FyneLogo())
	windowLogin.Resize(fyne.NewSize(400, 400))

	// Carregar a imagem de fundo (substitua o caminho pelo seu arquivo de imagem)
	backgroundImage := canvas.NewImageFromFile("images/mainBackgroundImage.jpg")
	if backgroundImage == nil {
		fmt.Println("Erro ao carregar a imagem de fundo")
	}

	// Redimensionar a imagem para cobrir toda a janela
	backgroundImage.FillMode = canvas.ImageFillStretch

	//Exemplo de ícone (pode ser substituído)
	imgSettings := widget.NewIcon(theme.HomeIcon())

	userName := widget.NewEntry()
	frmUserName := widget.NewFormItem("UserName", userName)

	userPassword := widget.NewPasswordEntry()
	frmUserPass := widget.NewFormItem("Password", userPassword)

	// Botão com ícone para ir para a tela de configuração
	btnSettings := widget.NewButtonWithIcon("", theme.SettingsIcon(), func() {
		fmt.Print("testes")
	})

	btnLogin := widget.NewButton("Login", loginValid)

	btnClose := widget.NewButton("Fechar", func() {
		(*mainApp).Quit() // Função para encerrar a aplicação
	})

	fieldForm := widget.NewForm(
		frmUserName,
		frmUserPass,
	)

	fieldForm.Resize(fyne.NewSize(400, 100))

	loginForm := container.NewVBox(
		layout.NewSpacer(),
		container.NewCenter(container.NewHBox(
			container.NewCenter(imgSettings),
		)),
		layout.NewSpacer(),
		container.NewCenter(container.NewHBox(
			fieldForm,
		)),
		layout.NewSpacer(),
		/*container.NewMax(container.NewHBox(
			btnSettings,
		)),*/
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
			layout.NewSpacer(), // Espaçador à direita
		),
		container.NewCenter(container.NewHBox(
			btnLogin,
			btnClose,
		)),
	)

	windowLogin.SetContent(container.NewMax(backgroundImage, loginForm))

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

// Tela de Login
func loginValid() {

}
