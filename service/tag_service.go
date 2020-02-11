package service

import (
	"github.com/oshou/AwesomeMusic-api/db"
	"github.com/oshou/AwesomeMusic-api/entity"
)

type TagService struct{}

type Tag entity.Tag
type PostTag entity.PostTag

func (ts TagService) GetAll() ([]Tag, error) {
	var tt []Tag
	query := `SELECT
							id,
							name
						FROM
							tag`
	conn := db.DBConn()
	if err := conn.Select(tt, query); err != nil {
		return nil, err
	}
	return tt, nil
}

func (ts TagService) GetById(tag_id int) (Tag, error) {
	var t Tag
	query := `SELECT
							id,
							name
						FROM
							tag
						WHERE
							id = ?`
	conn := db.DBConn()
	if err := conn.Get(t, query, tag_id); err != nil {
		return t, err
	}
	return t, nil
}

func (ts TagService) GetByName(tag_name string) ([]Tag, error) {
	var tt []Tag
	query := `SELECT
							id,
							name
						FROM
							tag
						WHERE
							name LIKE ?`
	conn := db.DBConn()
	if err := conn.Select(tt, query, "%"+tag_name+"%"); err != nil {
		return tt, err
	}
	return tt, nil
}

func (ts TagService) GetByPostId(post_id int) ([]Tag, error) {
	var tt []Tag
	query := `SELECT
							t.id,
							t.name
						FROM
							tag AS tag
						INNER JOIN post_tag AS pt
							ON t.id = pt.tag_id
						WHERE
							pt.post_id = ?`
	conn := db.DBConn()
	if err := conn.Select(tt, query, post_id); err != nil {
		return tt, err
	}
	return tt, nil
}

func (ts TagService) Add(tag_name string) (Tag, error) {
	var t = Tag{
		Name: tag_name,
	}
	query := `INSERT INTO
							tag(name)
						VALUES
							(?)`
	conn := db.DBConn()
	result, err := conn.Exec(query, tag_name)
	if err != nil {
		return t, err
	}
	i64, _ := result.LastInsertId()
	t.ID = int(i64)
	return t, nil
}

func (ts TagService) Attach(post_id, tag_id int) (PostTag, error) {
	var pt = PostTag{
		PostID: post_id,
		TagID:  tag_id,
	}
	query := `INSERT INTO
							post_tag(post_id, tag_id)
						VALUES
							(?, ?)`
	conn := db.DBConn()
	_, err := conn.Exec(query, post_id, tag_id)
	if err != nil {
		return pt, err
	}
	return pt, nil
}
