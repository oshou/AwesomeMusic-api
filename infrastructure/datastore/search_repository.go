package infrastructure

import (
	"github.com/jmoiron/sqlx"
	"github.com/oshou/AwesomeMusic-api/domain/model"
	"github.com/oshou/AwesomeMusic-api/domain/repository"
)

type searchRepository struct {
	DB *sqlx.DB
}

var _ repository.SearchRepository = searchRepository{}

func NewSearchRepository(db *sqlx.DB) repository.SearchRepository {
	return &searchRepository{
		DB: db,
	}
}

func (sr *searchRepository) GetByPostTitle(q string) ([]*model.Post, error) {
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
							title LIKE ?`

	if err := sr.DB.Select(&pp, query, "%"+q+"%"); err != nil {
		return nil, err
	}

	return pp, nil
}

func (sr *searchRepository) GetByUserName(q string) ([]*model.Post, error) {
	var pp []*model.Post

	query := `SELECT
							id,
							user_id,
							title,
							url,
							message
						FROM
							user
						WHERE
							name LIKE ?`

	if err := sr.DB.Select(&pp, query, "%"+q+"%"); err != nil {
		return nil, err
	}

	return pp, nil
}

func (sr *searchRepository) GetByTagName(q string) ([]*model.Post, error) {
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
							t.name LIKE ?`

	if err := sr.DB.Select(&pp, query, "%"+q+"%"); err != nil {
		return nil, err
	}

	return pp, nil
}
