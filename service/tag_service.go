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

	conn := db.DBcon()
	if err := conn.Select(&tt, query); err != nil {
		return nil, err
	}

	return tt, nil
}

func (ts TagService) GetByID(tagID int) (Tag, error) {
	var t Tag

	query := `SELECT
							id,
							name
						FROM
							tag
						WHERE
							id = ?`

	conn := db.DBcon()
	if err := conn.Get(&t, query, tagID); err != nil {
		return t, err
	}

	return t, nil
}

func (ts TagService) GetByName(tagName string) ([]Tag, error) {
	var tt []Tag

	query := `SELECT
							id,
							name
						FROM
							tag
						WHERE
							name LIKE ?`

	conn := db.DBcon()
	if err := conn.Select(&tt, query, "%"+tagName+"%"); err != nil {
		return tt, err
	}

	return tt, nil
}

func (ts TagService) GetByPostID(postID int) ([]Tag, error) {
	var tt []Tag

	query := `SELECT
							t.id,
							t.name
						FROM
							tag AS t
						INNER JOIN post_tag AS pt
							ON t.id = pt.tag_id
						WHERE
							pt.post_id = ?`

	conn := db.DBcon()
	if err := conn.Select(&tt, query, postID); err != nil {
		return tt, err
	}

	return tt, nil
}

func (ts TagService) Add(tagName string) (Tag, error) {
	var t = Tag{
		Name: tagName,
	}

	query := `INSERT INTO
							tag(name)
						VALUES
							(?)`

	conn := db.DBcon()
	result, err := conn.Exec(query, tagName)

	if err != nil {
		return t, err
	}

	i64, _ := result.LastInsertId()
	t.ID = int(i64)

	return t, nil
}

func (ts TagService) Attach(postID, tagID int) (PostTag, error) {
	var pt = PostTag{
		PostID: postID,
		TagID:  tagID,
	}

	query := `INSERT INTO
							post_tag(post_id, tag_id)
						VALUES
							(?, ?)`

	conn := db.DBcon()
	_, err := conn.Exec(query, postID, tagID)

	if err != nil {
		return pt, err
	}

	return pt, nil
}
