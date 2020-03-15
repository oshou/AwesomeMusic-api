package datastore

import (
	"github.com/jmoiron/sqlx"
	"github.com/oshou/AwesomeMusic-api/domain/model"
)

type SearchStore struct {
	DB *sqlx.DB
}

func (ss SearchStore) GetByPostTitle(q string) ([]model.Post, error) {
	var pp []model.Post

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

	if err := ss.DB.Select(&pp, query, "%"+q+"%"); err != nil {
		return nil, err
	}

	return pp, nil
}

func (ss SearchStore) GetByUserName(q string) ([]model.Post, error) {
	var pp []model.Post

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

	if err := ss.DB.Select(&pp, query, "%"+q+"%"); err != nil {
		return nil, err
	}

	return pp, nil
}

func (ss SearchStore) GetByTagName(q string) ([]model.Post, error) {
	var pp []model.Post

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

	if err := ss.DB.Select(&pp, query, "%"+q+"%"); err != nil {
		return nil, err
	}

	return pp, nil
}
