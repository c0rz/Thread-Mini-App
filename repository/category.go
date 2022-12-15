package repository

import (
	"Blog/models"
	"database/sql"
	"time"
)

func GetAllCategory(db *sql.DB) (err error, results []models.Category) {
	sql := "SELECT * FROM category"

	rows, err := db.Query(sql)
	if err != nil {
		return err, nil
	}

	for rows.Next() {
		var category = models.Category{}

		err = rows.Scan(&category.ID, &category.Name, &category.CreatedAt)
		if err != nil {
			return err, nil
		}

		results = append(results, category)
	}

	return nil, results
}

func InsertCategory(db *sql.DB, category models.Category) (err error){	
	sql := `INSERT INTO category(
		name, created_at)
		VALUES ($1, $2);`

	
	errs := db.QueryRow(sql, category.Name, time.Now())

	return errs.Err()
}

func DeleteCategory(db *sql.DB, aID int) (err error) {

	sql := `DELETE FROM category WHERE id = $1` 

	errs := db.QueryRow(sql, aID)

	return errs.Err()
}

func GetbyIDCategory(db *sql.DB, aID string) (err error) {
	sql := "SELECT FROM category WHERE id = $1"

	errs := db.QueryRow(sql, aID)

	return errs.Err()

}