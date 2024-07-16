package database

import (
	"enzappmob/internal/data/models"
	"os"

	"github.com/rs/zerolog/log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// var dbTypeGorm *gorm.DB
var dataBasePath string = "enzDatabase/"
var dataBaseName string = "enztechapp.db"

var (
	MsgErrConnOpen    = "Erro de conexão"
	MsgErrCreateTable = "Erro de criação da tabela..."
	MsgCreateTable    = "Tabela criada"
	MsgErrVldPath     = "Erro ao verificar existência do diretório:"
	MsgErrCreatePath  = "Erro ao verificar existência do diretório:"
)

func Connect() (*gorm.DB, error) {

	// Criar um diretório com permissões padrão (0777 no Unix)
	if _, err := os.Stat(dataBasePath); os.IsNotExist(err) {
		if err = os.Mkdir(dataBasePath, 0755); err != nil {
			log.Warn().Str("Error:", err.Error()).Msg(MsgErrCreatePath)

			return nil, err
		}
	} else if err != nil {
		log.Warn().Str("Erro:", err.Error()).Msg(MsgErrVldPath)

		return nil, err
	}

	dataBaseLocal := dataBasePath + dataBaseName

	// Conectar ao banco de dados SQLite
	dbOpenConn, err := gorm.Open(sqlite.Open(dataBaseLocal), &gorm.Config{})

	if err != nil {

		log.Warn().Str("Error gorm:", err.Error()).Msg(MsgErrConnOpen)

		return nil, err
	}

	// Configuração de logging, configurações adicionais, etc.
	//dbTypeGorm = dbOpenConn

	return dbOpenConn, nil
}

func CloseConnect(dataBase *gorm.DB) error {

	dbInstance, err := dataBase.DB()

	if err != nil {
		return err
	}

	err = dbInstance.Close()

	if err != nil {
		return err
	}

	return nil
}

func CreateEntity(dataBase *gorm.DB) {

	err := createIpCapture(dataBase)

	if err != nil {
		log.Warn().Str("Error gorm:", err.Error()).Msg(MsgErrCreateTable)
	} else {
		log.Info().Msg(MsgCreateTable + "IpCapture")
	}

}

func createIpCapture(accessBase *gorm.DB) error {
	// Verificar se a tabela User já existe no banco de dados
	if !accessBase.Migrator().HasTable(&models.IpCapture{}) {
		// Se a tabela não existir, criar a tabela
		err := accessBase.AutoMigrate(&models.IpCapture{})

		if err != nil {
			return err
		}
	}

	return nil
}
