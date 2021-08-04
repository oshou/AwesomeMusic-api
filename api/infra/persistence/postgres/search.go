// Package postgres is repository implementation package
package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/oshou/AwesomeMusic-api/api/domain/model"
	"github.com/oshou/AwesomeMusic-api/api/domain/repository"
)

type searchRepository struct {
	db *sqlx.DB
}

var _ repository.ISearchRepository = &searchRepository{}

// NewSearchRepository is constructor for searchRepository
func NewSearchRepository(db *sqlx.DB) repository.ISearchRepository {
	return &searchRepository{
		db: db,
	}
}

func (sr *searchRepository) ListByTitle(q string) ([]*model.Post, error) {
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
			title LIKE $1
	`

	var pp []*model.Post
	if err := sr.db.Select(&pp, query, "%"+q+"%"); err != nil {
		return nil, errors.WithStack(err)
	}

	return pp, nil
}

func (sr *searchRepository) ListByUserName(q string) ([]*model.Post, error) {
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
			u.name LIKE $1
	`

	var pp []*model.Post
	if err := sr.db.Select(&pp, query, "%"+q+"%"); err != nil {
		return nil, errors.WithStack(err)
	}

	return pp, nil
}

func (sr *searchRepository) ListByTagName(q string) ([]*model.Post, error) {
	const query = `
		SELECT
				p.id
			,	p.user_id
			,	p.title
			,	p.url
			,	p.message
		FROM
			public.posts AS p
		INNER JOIN public.post_tag AS pt
			ON pt.post_id = p.id
		INNER JOIN public.tags AS t
			ON pt.tag_id = t.id
		WHERE
			t.name LIKE $1
	`

	var pp []*model.Post
	if err := sr.db.Select(&pp, query, "%"+q+"%"); err != nil {
		return nil, errors.WithStack(err)
	}

	return pp, nil
}
