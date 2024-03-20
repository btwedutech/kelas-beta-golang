package utils

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"sekolahbeta/database/model"
)

func InsertData(conn *sql.DB, bdy model.Car, ctx context.Context) error {
	syn := fmt.Sprintf(`
		INSERT INTO car(id, nama, tipe, tahun)
		VALUES ('%s', '%s', '%s', '%s')
	`, bdy.ID, bdy.Nama, bdy.Tipe, bdy.Tahun)

	// fmt.Println(syn)

	_, err := conn.ExecContext(ctx, syn)
	if err != nil {
		return err
	}

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	return nil
}

func GetByID(conn *sql.DB, id string, ctx context.Context) (model.Car, error) {
	res := model.Car{}
	syn := fmt.Sprintf(`
		SELECT * FROM car WHERE id = %s LIMIT 1
	`, id)

	row, err := conn.QueryContext(ctx, syn)
	if err != nil {
		return model.Car{}, err
	}

	defer func() {
		err := row.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	for row.Next() {
		err = row.Scan(
			&res.ID,
			&res.Nama,
			&res.Tipe,
			&res.Tahun,
		)
		if err != nil {
			return model.Car{}, err
		}
	}

	return res, nil
}
