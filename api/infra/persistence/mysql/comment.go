// Package mysql implements repository package
package mysql

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/oshou/AwesomeMusic-api/api/domain/model"
	"github.com/oshou/AwesomeMusic-api/api/domain/repository"
)

type commentRepository struct {
	db *sqlx.DB
}

var _ repository.ICommentRepository = &commentRepository{}

// NewCommentRepository is constructor for commentRepository
func NewCommentRepository(db *sqlx.DB) repository.ICommentRepository {
	return &commentRepository{
		db: db,
	}
}

func (cr *commentRepository) List(postID int) ([]*model.Comment, error) {
	var cc []*model.Comment

	query := `SELECT
							id,
							user_id,
							post_id,
							comment
						FROM
							comment
						WHERE
							post_id = $1`

	if err := cr.db.Select(&cc, query, postID); err != nil {
		return nil, errors.WithStack(err)
	}

	return cc, nil
}

func (cr *commentRepository) GetByID(commentID int) (*model.Comment, error) {
	var c model.Comment

	query := `SELECT
							id,
							user_id,
							post_id,
							comment
						FROM
							comment
						WHERE
							id = $1`

	if err := cr.db.Get(&c, query, commentID); err != nil {
		return nil, errors.WithStack(err)
	}

	return &c, nil
}

func (cr *commentRepository) Add(postID, userID int, comment string) (*model.Comment, error) {
	query := `INSERT INTO
							comment(post_id, user_id, comment)
						VALUES
							($1, $2, $3)`

	result, err := cr.db.Exec(query, postID, userID, comment)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	c := model.Comment{
		UserID:  userID,
		PostID:  postID,
		Comment: comment,
	}
	i64, _ := result.LastInsertId()
	c.ID = int(i64)

	return &c, nil
}
