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
	conn := db.DBConn()
	if err := conn.Select(&pp, query); err != nil {
		return nil, err
	}
	return pp, nil
}

func (ps PostService) GetById(post_id int) (Post, error) {
	var p Post
	query := `SELECT
							id,
							user_id,
							title, url,
							message
						FROM
							post
						WHERE id = ?`
	conn := db.DBConn()
	if err := conn.Get(&p, query, post_id); err != nil {
		return p, err
	}
	return p, nil
}

func (ps PostService) GetByTagId(tag_id int) ([]Post, error) {
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
	conn := db.DBConn()
	if err := conn.Select(&pp, query, tag_id); err != nil {
		return nil, err
	}
	return pp, nil
}

func (ps PostService) GetByUserId(user_id int) ([]Post, error) {
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
	conn := db.DBConn()
	if err := conn.Select(&pp, query, user_id); err != nil {
		return nil, err
	}
	return pp, nil
}

func (ps PostService) Add(user_id int, title, url, message string) (Post, error) {
	var p = Post{
		UserID:  user_id,
		Title:   title,
		URL:     url,
		Message: message,
	}
	conn := db.DBConn()
	query := `INSERT INTO
							post(user_id, title, url, message)
						VALUES
							(?, ?, ?, ?)`
	result, err := conn.Exec(query, user_id, title, url, message)
	if err != nil {
		return p, err
	}
	i64, _ := result.LastInsertId()
	p.ID = int(i64)
	return p, nil
}

func (ps PostService) DeleteById(post_id int) error {
	conn := db.DBConn()
	query := `DELETE FROM post WHERE id = ?`
	if _, err := conn.Exec(query, post_id); err != nil {
		return err
	}
	return nil
}
