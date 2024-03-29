// Package postgres implements repository package
package postgres

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
	const query = `
		SELECT
				id,
			, user_id
			, post_id
			, comment
		FROM
			public.comments
		WHERE
			post_id = $1
		ORDER BY
			id
	`

	var cc []*model.Comment
	if err := cr.db.Select(&cc, query, postID); err != nil {
		return nil, errors.WithStack(err)
	}

	return cc, nil
}

func (cr *commentRepository) GetByID(commentID int) (*model.Comment, error) {
	const query = `
		SELECT
				id
			, user_id
			, post_id
			, comment
		FROM
			public.comments
		WHERE
			id = $1
	`

	var c model.Comment
	if err := cr.db.Get(&c, query, commentID); err != nil {
		return nil, errors.WithStack(err)
	}

	return &c, nil
}

func (cr *commentRepository) Add(postID, userID int, comment string) (*model.Comment, error) {
	const query = `
		INSERT INTO
				public.comments(post_id, user_id, comment)
		VALUES
			($1, $2, $3)
		RETURNING
			id
	`

	c := model.Comment{
		UserID:  userID,
		PostID:  postID,
		Comment: comment,
	}

	if err := cr.db.QueryRow(query, postID, userID, comment).Scan(&c.ID); err != nil {
		return nil, errors.WithStack(err)
	}

	return &c, nil
}
