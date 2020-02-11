package service

import (
	"github.com/oshou/AwesomeMusic-api/db"
	"github.com/oshou/AwesomeMusic-api/entity"
)

type CommentService struct{}

type Comment entity.Comment

func (cs CommentService) GetAll(post_id int) ([]Comment, error) {
	var cc []Comment
	query := `SELECT
							id,
							user_id,
							post_id,
							comment
						FROM
							comment
						WHERE
							post_id = ?`
	conn := db.DBConn()
	if err := conn.Select(cc, query, post_id); err != nil {
		return nil, err
	}
	return cc, nil
}

func (cs CommentService) GetById(comment_id int) (Comment, error) {
	var c Comment
	query := `SELECT
							id,
							user_id,
							post_id,
							comment
						FROM
							comment
						WHERE
							id = ?`
	conn := db.DBConn()
	if err := conn.Get(c, query, comment_id); err != nil {
		return c, err
	}
	return c, nil
}

func (cs CommentService) Add(post_id, user_id int, comment string) (Comment, error) {
	var c = Comment{
		UserID:  user_id,
		PostID:  post_id,
		Comment: comment,
	}
	conn := db.DBConn()
	query := `INSERT INTO
							comment(post_id, user_id, comment)
						VALUES
							(?, ?, ?)`
	result, err := conn.Exec(query, post_id, user_id, comment)
	if err != nil {
		return c, err
	}
	i64, _ := result.LastInsertId()
	c.ID = int(i64)
	return c, nil
}
