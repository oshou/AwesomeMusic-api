package datastore

import (
	"github.com/jmoiron/sqlx"
	"github.com/oshou/AwesomeMusic-api/domain/model"
)

type CommentStore struct {
	DB *sqlx.DB
}

func (cs *CommentStore) GetAll(postID int) ([]model.Comment, error) {
	var cc []model.Comment

	query := `SELECT
							id,
							user_id,
							post_id,
							comment
						FROM
							comment
						WHERE
							post_id = ?`

	if err := cs.DB.Select(&cc, query, postID); err != nil {
		return nil, err
	}

	return cc, nil
}

func (cs *CommentStore) GetByID(commentID int) (model.Comment, error) {
	var c model.Comment

	query := `SELECT
							id,
							user_id,
							post_id,
							comment
						FROM
							comment
						WHERE
							id = ?`

	if err := cs.DB.Get(&c, query, commentID); err != nil {
		return c, err
	}

	return c, nil
}

func (cs *CommentStore) Add(postID, userID int, comment string) (model.Comment, error) {
	var c = model.Comment{
		UserID:  userID,
		PostID:  postID,
		Comment: comment,
	}

	query := `INSERT INTO
							comment(post_id, user_id, comment)
						VALUES
							(?, ?, ?)`

	result, err := cs.DB.Exec(query, postID, userID, comment)

	if err != nil {
		return c, err
	}

	i64, _ := result.LastInsertId()
	c.ID = int(i64)

	return c, nil
}
