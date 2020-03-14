package service

import (
	"github.com/oshou/AwesomeMusic-api/db"
	"github.com/oshou/AwesomeMusic-api/entity"
)

type CommentService struct{}

type Comment entity.Comment

func (cs CommentService) GetAll(postID int) ([]Comment, error) {
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
	if err := conn.Select(&cc, query, postID); err != nil {
		return nil, err
	}

	return cc, nil
}

func (cs CommentService) GetByID(commentID int) (Comment, error) {
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
	if err := conn.Get(&c, query, commentID); err != nil {
		return c, err
	}

	return c, nil
}

func (cs CommentService) Add(postID, userID int, comment string) (Comment, error) {
	var c = Comment{
		UserID:  userID,
		PostID:  postID,
		Comment: comment,
	}

	query := `INSERT INTO
							comment(post_id, user_id, comment)
						VALUES
							(?, ?, ?)`

	conn := db.DBConn()
	result, err := conn.Exec(query, postID, userID, comment)

	if err != nil {
		return c, err
	}

	i64, _ := result.LastInsertId()
	c.ID = int(i64)

	return c, nil
}
