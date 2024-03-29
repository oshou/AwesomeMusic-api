// Package mysql implements repository package
package mysql

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/oshou/AwesomeMusic-api/api/domain/model"
	"github.com/oshou/AwesomeMusic-api/api/domain/repository"
)

type postRepository struct {
	db *sqlx.DB
}

var _ repository.IPostRepository = &postRepository{}

// NewPostRepository is constructor for postRepository
func NewPostRepository(db *sqlx.DB) repository.IPostRepository {
	return &postRepository{
		db: db,
	}
}

func (pr *postRepository) List() ([]*model.Post, error) {
	var pp []*model.Post

	query := `SELECT
							id,
							user_id,
							title,
							url,
							message
						FROM
							post`

	err := pr.db.Select(&pp, query)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return pp, nil
}

func (pr *postRepository) GetByID(postID int) (*model.Post, error) {
	var p model.Post

	query := `SELECT
							id,
							user_id,
							title,
							url,
							message
						FROM
							post
						WHERE
							id = $1`

	if err := pr.db.Get(&p, query, postID); err != nil {
		return nil, errors.WithStack(err)
	}

	return &p, nil
}

func (pr *postRepository) ListByTagID(tagID int) ([]*model.Post, error) {
	var pp []*model.Post

	query := `SELECT
							p.id,
							p.user_id,
							p.title,
							p.url,
							p.message
						FROM
							post AS p
						INNER JOIN post_tag AS pt
							ON pt.post_id = p.id
						INNER JOIN tag AS t
							ON pt.tag_id = t.id
						WHERE
							t.id = $1`

	if err := pr.db.Select(&pp, query, tagID); err != nil {
		return nil, errors.WithStack(err)
	}

	return pp, nil
}

func (pr *postRepository) ListByUserID(userID int) ([]*model.Post, error) {
	var pp []*model.Post

	query := `SELECT
							p.id,
							p.user_id,
							p.title,
							p.url,
							p.message
						FROM
							post AS p
						INNER JOIN user AS u
							ON u.id = p.user_id
						WHERE
							u.id = $1`

	if err := pr.db.Select(&pp, query, userID); err != nil {
		return nil, errors.WithStack(err)
	}

	return pp, nil
}

func (pr *postRepository) Add(userID int, title, url, message string) (*model.Post, error) {
	query := `INSERT INTO
							post(user_id, title, url, message)
						VALUES
							($1, $2, $3, $4)`

	result, err := pr.db.Exec(query, userID, title, url, message)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	p := model.Post{
		UserID:  userID,
		Title:   title,
		URL:     url,
		Message: message,
	}

	i64, _ := result.LastInsertId()
	p.ID = int(i64)

	return &p, nil
}

func (pr *postRepository) DeleteByID(postID int) error {
	query := `DELETE FROM
							post
						WHERE
							id = $1`

	if _, err := pr.db.Exec(query, postID); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
