package config_test

import (
	"sekolahbeta/database/config"
	"testing"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func Init() {
	err := godotenv.Load("../.env")
	if err != nil {
		logrus.Println(".env not found, using global variable")
	}
}

func TestConnection(t *testing.T) {
	Init()
	db, err := config.OpenConn()

	defer func() {
		errRow := db.Close()
		if errRow != nil {
			logrus.Error(errRow)
		}
	}()

	assert.Nil(t, err)
}
