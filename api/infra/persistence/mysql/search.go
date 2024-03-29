// Package mysql implements repository package
package mysql

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
	var pp []*model.Post

	query := `SELECT
							id,
							user_id,
							title,
							url,
							message
						FROM
							post
						WHERE
							title LIKE $1`

	if err := sr.db.Select(&pp, query, "%"+q+"%"); err != nil {
		return nil, errors.WithStack(err)
	}

	return pp, nil
}

func (sr *searchRepository) ListByUserName(q string) ([]*model.Post, error) {
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
							u.name LIKE $1`

	if err := sr.db.Select(&pp, query, "%"+q+"%"); err != nil {
		return nil, errors.WithStack(err)
	}

	return pp, nil
}

func (sr *searchRepository) ListByTagName(q string) ([]*model.Post, error) {
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
							t.name LIKE $1`

	if err := sr.db.Select(&pp, query, "%"+q+"%"); err != nil {
		return nil, errors.WithStack(err)
	}

	return pp, nil
}
