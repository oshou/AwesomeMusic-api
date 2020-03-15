package datastore

import (
	"github.com/jmoiron/sqlx"
	"github.com/oshou/AwesomeMusic-api/domain/model"
)

type TagStore struct {
	DB *sqlx.DB
}

func (ts *TagStore) GetAll() ([]model.Tag, error) {
	var tt []model.Tag
	query := `SELECT
							id,
							name
						FROM
							tag`

	if err := ts.DB.Select(&tt, query); err != nil {
		return nil, err
	}

	return tt, nil
}

func (ts *TagStore) GetByID(tagID int) (model.Tag, error) {
	var t model.Tag

	query := `SELECT
							id,
							name
						FROM
							tag
						WHERE
							id = ?`

	if err := ts.DB.Get(&t, query, tagID); err != nil {
		return t, err
	}

	return t, nil
}

func (ts *TagStore) GetByName(tagName string) ([]model.Tag, error) {
	var tt []model.Tag

	query := `SELECT
							id,
							name
						FROM
							tag
						WHERE
							name LIKE ?`

	if err := ts.DB.Select(&tt, query, "%"+tagName+"%"); err != nil {
		return tt, err
	}

	return tt, nil
}

func (ts *TagStore) GetByPostID(postID int) ([]model.Tag, error) {
	var tt []model.Tag

	query := `SELECT
							t.id,
							t.name
						FROM
							tag AS t
						INNER JOIN post_tag AS pt
							ON t.id = pt.tag_id
						WHERE
							pt.post_id = ?`

	if err := ts.DB.Select(&tt, query, postID); err != nil {
		return tt, err
	}

	return tt, nil
}

func (ts *TagStore) Add(tagName string) (model.Tag, error) {
	var t = model.Tag{
		Name: tagName,
	}

	query := `INSERT INTO
							tag(name)
						VALUES
							(?)`

	result, err := ts.DB.Exec(query, tagName)

	if err != nil {
		return t, err
	}

	i64, _ := result.LastInsertId()
	t.ID = int(i64)

	return t, nil
}

func (ts *TagStore) Attach(postID, tagID int) (model.PostTag, error) {
	var pt = model.PostTag{
		PostID: postID,
		TagID:  tagID,
	}

	query := `INSERT INTO
							post_tag(post_id, tag_id)
						VALUES
							(?, ?)`

	_, err := ts.DB.Exec(query, postID, tagID)

	if err != nil {
		return pt, err
	}

	return pt, nil
}
