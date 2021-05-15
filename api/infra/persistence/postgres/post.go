// Package postgres implements repository package
package postgres

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
	const query = `
		SELECT
				id
			, user_id
			, title
			, url
			, message
		FROM
			public.posts
		ORDER BY
			id
	`

	var pp []*model.Post
	if err := pr.db.Select(&pp, query); err != nil {
		return nil, errors.WithStack(err)
	}

	return pp, nil
}

func (pr *postRepository) ListByTagID(tagID int) ([]*model.Post, error) {
	const query = `
		SELECT
				p.id
			, p.user_id
			, p.title
			, p.url
			, p.message
		FROM
			public.posts AS p
		INNER JOIN public.post_tag AS pt
			ON pt.post_id = p.id
		INNER JOIN public.tags AS t
			ON pt.tag_id = t.id
		WHERE
			t.id = $1
		ORDER BY
			p.id
	`

	var pp []*model.Post
	if err := pr.db.Select(&pp, query, tagID); err != nil {
		return nil, errors.WithStack(err)
	}

	return pp, nil
}

func (pr *postRepository) ListByUserID(userID int) ([]*model.Post, error) {
	const query = `
		SELECT
				p.id
			, p.user_id
			, p.title
			, p.url
			, p.message
		FROM
			public.posts AS p
		INNER JOIN public.users AS u
			ON u.id = p.user_id
		WHERE
			u.id = $1
		ORDER BY
			p.id
	`

	var pp []*model.Post
	if err := pr.db.Select(&pp, query, userID); err != nil {
		return nil, errors.WithStack(err)
	}

	return pp, nil
}

func (pr *postRepository) GetByID(postID int) (*model.Post, error) {
	const query = `
		SELECT
				id
			, user_id
			, title
			, url
			, message
		FROM
			public.posts
		WHERE
			id = $1
	`

	var p model.Post
	if err := pr.db.Get(&p, query, postID); err != nil {
		return nil, errors.WithStack(err)
	}

	return &p, nil
}

func (pr *postRepository) Add(userID int, title, url, message string) (*model.Post, error) {
	const query = `
		INSERT INTO
			public.posts(user_id, title, url, message)
		VALUES
			($1, $2, $3, $4)
		RETURNING
			id
	`

	p := model.Post{
		UserID:  userID,
		Title:   title,
		URL:     url,
		Message: message,
	}

	err := pr.db.QueryRow(query, userID, title, url, message).Scan(&p.ID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &p, nil
}

func (pr *postRepository) DeleteByID(postID int) error {
	const query = `
		DELETE FROM
			public.posts
		WHERE
			id = $1
		ORDER BY
			id
	`

	if _, err := pr.db.Exec(query, postID); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
