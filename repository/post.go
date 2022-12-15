package repository

import (
	"Blog/models"
	"database/sql"
	"time"
)

func GetAllPost(db *sql.DB) (errs error, results []models.Post) {
	sql := `SELECT post.id, post.title, category.name, text, account.full_name, post.created_at
	FROM public.post 
	JOIN account ON account.id = post.id_user
	JOIN category ON category.id = post.id_category`

	rows, errs := db.Query(sql)
	if errs != nil {
		return errs, results
	}

	for rows.Next() {
		var post = models.Post{}

		errs = rows.Scan(&post.ID, &post.Title, &post.ID_Category, &post.Text, &post.ID_User, &post.CreatedAt)
		if errs != nil {
			return errs, results
		}

		results = append(results, post)
	}

	return errs, results
}

func InsertPost(db *sql.DB, post models.Post) (uid int, err error) {
	sql := `INSERT INTO public.post(
		title, id_category, text, id_user, created_at)
		VALUES ($1, $2, $3, $4, $5) RETURNING id;`

	err = db.QueryRow(sql, post.Title, post.ID_Category, post.Text, post.ID_User, time.Now()).Scan(&uid)

	return uid, err
}

func GetbyIDPost(db *sql.DB, sql string) (result *sql.Row) {

	row := db.QueryRow(sql)
	
	return row
}

func DeletePost(db *sql.DB, pid int) error {

	sqlStetement := `DELETE from post WHERE id = $1;`

	_, err := db.Exec(sqlStetement, pid)
	if err != nil {
		return err
	}

	return nil
}

func UpdatePost(db *sql.DB, post models.Post) (err error) {
	sql := "UPDATE post SET title = $1, id_category = $2, text = $3 WHERE id = $4"

	errs := db.QueryRow(sql, post.Title, post.ID_Category, post.Text, post.ID)

	return errs.Err()
}
