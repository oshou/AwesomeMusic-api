package service

import (
	"github.com/oshou/AwesomeMusic-api/db"
	"github.com/oshou/AwesomeMusic-api/entity"
)

type PostService struct{}

type Post entity.Post

func (ps PostService) GetAll() ([]Post, error) {
	var pp []Post

	query := `SELECT
							id,
							user_id,
							title,
							url,
							message
						FROM
							post`

	conn := db.DBcon()
	if err := conn.Select(&pp, query); err != nil {
		return nil, err
	}

	return pp, nil
}

func (ps PostService) GetByID(postID int) (Post, error) {
	var p Post

	query := `SELECT
							id,
							user_id,
							title, url,
							message
						FROM
							post
						WHERE id = ?`

	conn := db.DBcon()
	if err := conn.Get(&p, query, postID); err != nil {
		return p, err
	}

	return p, nil
}

func (ps PostService) GetByTagID(tagID int) ([]Post, error) {
	var pp []Post

	query := `SELECT
							p.id,
							p.user_id,
							p.title,
							p.url,
							p.message
						FROM
							post AS p
						INNER JOIN post_tag AS pt ON pt.post_id = p.id
						INNER JOIN tag AS t ON pt.tag_id= t.id
						WHERE
							t.id = ?`

	conn := db.DBcon()
	if err := conn.Select(&pp, query, tagID); err != nil {
		return nil, err
	}

	return pp, nil
}

func (ps PostService) GetByUserID(userID int) ([]Post, error) {
	var pp []Post

	query := `SELECT
							p.id,
							p.user_id,
							p.title,
							p.url,
							p.message
						FROM
							post AS p
						INNER JOIN user AS u ON u.id = p.user_id
						WHERE
							u.id = ?`

	conn := db.DBcon()
	if err := conn.Select(&pp, query, userID); err != nil {
		return nil, err
	}

	return pp, nil
}

func (ps PostService) Add(userID int, title, url, message string) (Post, error) {
	var p = Post{
		UserID:  userID,
		Title:   title,
		URL:     url,
		Message: message,
	}

	query := `INSERT INTO
							post(user_id, title, url, message)
						VALUES
							(?, ?, ?, ?)`

	conn := db.DBcon()
	result, err := conn.Exec(query, userID, title, url, message)

	if err != nil {
		return p, err
	}

	i64, _ := result.LastInsertId()
	p.ID = int(i64)

	return p, nil
}

func (ps PostService) DeleteByID(postID int) error {
	query := `DELETE FROM
							post
						WHERE
							id = ?`

	conn := db.DBcon()
	if _, err := conn.Exec(query, postID); err != nil {
		return err
	}

	return nil
}
