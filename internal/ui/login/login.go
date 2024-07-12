package login

import (
	"enzappmob/internal/utils"
	"fmt"
	"os"

	"fyne.io/fyne"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

var (
	en_us_message = "locales/en-us/screens/login.yaml"
	pt_br_message = "locales/pt-br/screens/login.yaml"
)

func LoginScreen(mainApp *fyne.App, bundle *i18n.Bundle) bool {

	var loginReturn bool = true

	// Carrega arquivos de mensagens
	if err := utils.LoadMessages(bundle, en_us_message); err != nil {
		fmt.Fprintf(os.Stderr, "Error loading messages: %v\n", err)
		//os.Exit(1)
	}

	if err := utils.LoadMessages(bundle, pt_br_message); err != nil {
		fmt.Fprintf(os.Stderr, "Error loading messages: %v\n", err)
		//os.Exit(1)
	}

	//Constrói o app inicial, utilizado por todas as aplicaçãoes

	windowLogin := (*mainApp).NewWindow("EnzTech-Mobile")

	windowLogin.Resize(fyne.NewSize(400, 400))

	label := widget.NewLabel("Hello, Fyne!")

	loginForm := widget.NewForm(
		widget.NewFormItem("UserName", widget.NewEntry()),
		widget.NewFormItem("Password", widget.NewPasswordEntry()),
	)

	windowLogin.SetContent(container.NewVBox(loginForm, label))

	windowLogin.ShowAndRun()

	return loginReturn
}
