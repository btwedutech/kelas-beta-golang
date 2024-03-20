package utils_test

import (
	"context"
	"fmt"
	"sekolahbeta/database/config"
	"sekolahbeta/database/model"
	"sekolahbeta/database/utils"
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

func TestCreateDataSuccess(t *testing.T) {
	Init()

	conn, err := config.OpenConn()
	assert.Nil(t, err)

	bdy := model.Car{
		ID:    "1",
		Nama:  "toyota",
		Tipe:  "yaris",
		Tahun: "2018",
	}

	err = utils.InsertData(conn, bdy, context.TODO())
	assert.Nil(t, err)
}

func TestCreateDataFailed(t *testing.T) {
	Init()

	conn, err := config.OpenConn()
	assert.Nil(t, err)

	bdy := model.Car{
		ID:    "1234",
		Nama:  "toyota",
		Tipe:  "yaris",
		Tahun: "2018",
	}

	err = utils.InsertData(conn, bdy, context.TODO())
	assert.Nil(t, err)

	bdy1 := model.Car{
		ID:    "1234",
		Nama:  "toyota",
		Tipe:  "yaris",
		Tahun: "2018",
	}

	err1 := utils.InsertData(conn, bdy1, context.TODO())
	assert.NotNil(t, err1)
}

func TestGetByIDSuccess(t *testing.T) {
	Init()

	conn, err := config.OpenConn()
	assert.Nil(t, err)
	ctx := context.TODO()

	res, err := utils.GetByID(conn, "12345", ctx)
	assert.Nil(t, err)

	fmt.Println(res)
}
