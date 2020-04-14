package datastore

import (
	"github.com/jmoiron/sqlx"
	"github.com/oshou/AwesomeMusic-api/domain/model"
	"github.com/oshou/AwesomeMusic-api/domain/repository"
)

type tagRepository struct {
	DB *sqlx.DB
}

var _ repository.ITagRepository = (*tagRepository)(nil)

func NewTagRepository(db *sqlx.DB) repository.ITagRepository {
	return &tagRepository{
		DB: db,
	}
}

func (tr *tagRepository) GetAll() ([]*model.Tag, error) {
	var tt []*model.Tag

	query := `SELECT
							id,
							name
						FROM
							tag`

	if err := tr.DB.Select(&tt, query); err != nil {
		return nil, err
	}

	return tt, nil
}

func (tr *tagRepository) GetByID(tagID int) (*model.Tag, error) {
	t := &model.Tag{}

	query := `SELECT
							id,
							name
						FROM
							tag
						WHERE
							id = ?`

	if err := tr.DB.Get(t, query, tagID); err != nil {
		return t, err
	}

	return t, nil
}

func (tr *tagRepository) GetByName(tagName string) ([]*model.Tag, error) {
	var tt []*model.Tag

	query := `SELECT
							id,
							name
						FROM
							tag
						WHERE
							name LIKE ?`

	if err := tr.DB.Select(&tt, query, "%"+tagName+"%"); err != nil {
		return nil, err
	}

	return tt, nil
}

func (tr *tagRepository) GetByPostID(postID int) ([]*model.Tag, error) {
	var tt []*model.Tag

	query := `SELECT
							t.id,
							t.name
						FROM
							tag AS t
						INNER JOIN post_tag AS pt
							ON t.id = pt.tag_id
						WHERE
							pt.post_id = ?`

	if err := tr.DB.Select(&tt, query, postID); err != nil {
		return tt, err
	}

	return tt, nil
}

func (tr *tagRepository) Add(tagName string) (*model.Tag, error) {
	t := &model.Tag{
		Name: tagName,
	}
	query := `INSERT INTO
							tag(name)
						VALUES
							(?)`

	result, err := tr.DB.Exec(query, tagName)
	if err != nil {
		return t, err
	}

	i64, _ := result.LastInsertId()
	t.ID = int(i64)

	return t, nil
}

func (tr *tagRepository) Attach(postID, tagID int) (*model.PostTag, error) {
	query := `INSERT INTO
							post_tag(post_id, tag_id)
						VALUES
							(?, ?)`

	_, err := tr.DB.Exec(query, postID, tagID)

	if err != nil {
		return nil, err
	}

	var pt = &model.PostTag{
		PostID: postID,
		TagID:  tagID,
	}

	return pt, nil
}
