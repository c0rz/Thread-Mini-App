package repository

import (
	"Blog/models"
	"database/sql"
	"time"
)


func GetAllAccount(db *sql.DB) (err error, results []models.Account) {
	sql := "SELECT * FROM account"

	rows, err := db.Query(sql)
	if err != nil {
		return err, nil
	}

	for rows.Next() {
		var account = models.Account{}

		err = rows.Scan(&account.ID, &account.Full_Name, &account.Email, &account.Level, &account.CreatedAt)
		if err != nil {
			return err, nil
		}

		results = append(results, account)
	}

	return nil, results
}

func InsertAccount(db *sql.DB, account models.Account) (err error) {
	sql := `INSERT INTO account(
		full_name, email, password, level, created_at)
		VALUES ($1, $2, $3, $4, $5);`

	errs := db.QueryRow(sql, account.Full_Name, account.Email, account.Password, account.Level, time.Now())

	return errs.Err()
}

func DeleteAccount(db *sql.DB, sql string, aID string) (err bool) {

	_ = db.QueryRow(sql, aID)

	return true
}

func GetbyIDAccount(db *sql.DB, sql string) (result *sql.Row) {

	row := db.QueryRow(sql)
	
	return row
}
