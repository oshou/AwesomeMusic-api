// Package mysql implements repository package
package mysql

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/oshou/AwesomeMusic-api/domain/model"
	"github.com/oshou/AwesomeMusic-api/domain/repository"
)

type tagRepository struct {
	db *sqlx.DB
}

var _ repository.ITagRepository = &tagRepository{}

// NewTagRepository is constructor for tagRepository
func NewTagRepository(db *sqlx.DB) repository.ITagRepository {
	return &tagRepository{
		db: db,
	}
}

func (tr *tagRepository) GetAll() ([]*model.Tag, error) {
	var tt []*model.Tag

	query := `SELECT
							id,
							name
						FROM
							tag`

	if err := tr.db.Select(&tt, query); err != nil {
		return nil, errors.WithStack(err)
	}

	return tt, nil
}

func (tr *tagRepository) GetByID(tagID int) (*model.Tag, error) {
	var t model.Tag

	query := `SELECT
							id,
							name
						FROM
							tag
						WHERE
							id = $1`

	if err := tr.db.Get(&t, query, tagID); err != nil {
		return nil, errors.WithStack(err)
	}

	return &t, nil
}

func (tr *tagRepository) GetByName(tagName string) ([]*model.Tag, error) {
	var tt []*model.Tag

	query := `SELECT
							id,
							name
						FROM
							tag
						WHERE
							name LIKE $1`

	if err := tr.db.Select(&tt, query, "%"+tagName+"%"); err != nil {
		return nil, errors.WithStack(err)
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
							pt.post_id = $1`

	if err := tr.db.Select(&tt, query, postID); err != nil {
		return nil, errors.WithStack(err)
	}

	return tt, nil
}

func (tr *tagRepository) Add(tagName string) (*model.Tag, error) {
	t := model.Tag{
		Name: tagName,
	}
	query := `INSERT INTO
							tag(name)
						VALUES
							($1)`

	result, err := tr.db.Exec(query, tagName)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	i64, _ := result.LastInsertId()
	t.ID = int(i64)

	return &t, nil
}

func (tr *tagRepository) Attach(postID, tagID int) (*model.PostTag, error) {
	query := `INSERT INTO
							post_tag(post_id, tag_id)
						VALUES
							($1, $2)`

	_, err := tr.db.Exec(query, postID, tagID)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	pt := model.PostTag{
		PostID: postID,
		TagID:  tagID,
	}

	return &pt, nil
}
