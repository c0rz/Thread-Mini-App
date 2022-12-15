package repository

import (
	"Blog/models"
	"database/sql"
	"time"
)

func InsertComment(db *sql.DB, comment models.Comment) (err error) {
	sql := `INSERT INTO comment (id_post, id_user, comment, created_at)
		VALUES ($1, $2, $3, $4);`

	errs := db.QueryRow(sql, comment.ID_Post, comment.ID_User, comment.Comment, time.Now())

	return errs.Err()
}

func GetComment(db *sql.DB, pid string) (comments []models.Comment, errs error) {
	sql := `SELECT account.full_name, comment.comment
	FROM comment
	JOIN account ON account.id = comment.id_user
	WHERE comment.id_post = '` + pid + `';`

	rows, errs := db.Query(sql)
	if errs != nil {
		return comments, errs
	}

	for rows.Next() {
		var comment = models.Comment{}

		errs = rows.Scan(&comment.ID_User, &comment.Comment)
		if errs != nil {
			return comments, errs
		}

		comments = append(comments, comment)
	}

	return comments, errs
}
