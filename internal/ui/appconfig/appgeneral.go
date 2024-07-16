package appconfig

import (
	"fyne.io/fyne"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)

type AppGeneral struct {
	MainApp         *fyne.App
	Bundle          *i18n.Bundle
	CurrentLanguage language.Tag
	MapIPConnect    map[string]interface{}
	SignalConnect   bool
	GormConnect     *gorm.DB
}
