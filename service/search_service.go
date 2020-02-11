package service

import (
	"github.com/oshou/AwesomeMusic-api/db"
)

type SearchService struct{}

//type Post entity.Post

func (ss SearchService) GetByPostTitle(q string) ([]Post, error) {
	var pp []Post
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
	conn := db.DBConn()
	if err := conn.Select(&pp, query, "%"+q+"%"); err != nil {
		return nil, err
	}
	return pp, nil
}

func (ss SearchService) GetByUserName(q string) ([]Post, error) {
	var pp []Post
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
	conn := db.DBConn()
	if err := conn.Select(&pp, query, "%"+q+"%"); err != nil {
		return nil, err
	}
	return pp, nil
}

func (ss SearchService) GetByTagName(q string) ([]Post, error) {
	var pp []Post
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
	conn := db.DBConn()
	if err := conn.Select(&pp, query, "%"+q+"%"); err != nil {
		return nil, err
	}
	return pp, nil
}
