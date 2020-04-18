// Package postgres is repository implementation package
package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/oshou/AwesomeMusic-api/domain/model"
	"github.com/oshou/AwesomeMusic-api/domain/repository"
)

type postRepository struct {
	db *sqlx.DB
}

var _ repository.IPostRepository = (*postRepository)(nil)

// NewPostRepository is constructor for postRepository
func NewPostRepository(db *sqlx.DB) repository.IPostRepository {
	return &postRepository{
		db: db,
	}
}

func (pr *postRepository) GetAll() ([]*model.Post, error) {
	var pp []*model.Post

	query := `SELECT
							id,
							user_id,
							title,
							url,
							message
						FROM
							post`

	if err := pr.db.Select(&pp, query); err != nil {
		return nil, err
	}

	return pp, nil
}

func (pr *postRepository) GetByID(postID int) (*model.Post, error) {
	p := &model.Post{}

	query := `SELECT
							id,
							user_id,
							title,
							url,
							message
						FROM
							post
						WHERE
							id = ?`

	if err := pr.db.Get(p, query, postID); err != nil {
		return nil, err
	}

	return p, nil
}

func (pr *postRepository) GetByTagID(tagID int) ([]*model.Post, error) {
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
							t.id = ?`

	if err := pr.db.Select(&pp, query, tagID); err != nil {
		return nil, err
	}

	return pp, nil
}

func (pr *postRepository) GetByUserID(userID int) ([]*model.Post, error) {
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
							u.id = ?`

	if err := pr.db.Select(&pp, query, userID); err != nil {
		return nil, err
	}

	return pp, nil
}

func (pr *postRepository) Add(userID int, title, url, message string) (*model.Post, error) {
	query := `INSERT INTO
							post(user_id, title, url, message)
						VALUES
							(?, ?, ?, ?)`

	result, err := pr.db.Exec(query, userID, title, url, message)
	if err != nil {
		return nil, err
	}

	var p = &model.Post{
		UserID:  userID,
		Title:   title,
		URL:     url,
		Message: message,
	}

	i64, _ := result.LastInsertId()
	p.ID = int(i64)

	return p, nil
}

func (pr *postRepository) DeleteByID(postID int) error {
	query := `DELETE FROM
							post
						WHERE
							id = ?`

	if _, err := pr.db.Exec(query, postID); err != nil {
		return err
	}

	return nil
}
