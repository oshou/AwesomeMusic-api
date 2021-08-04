// Package postgres is repository implementation package
package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/oshou/AwesomeMusic-api/api/domain/model"
	"github.com/oshou/AwesomeMusic-api/api/domain/repository"
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

func (tr *tagRepository) List() ([]*model.Tag, error) {
	const query = `
		SELECT
				id
			, name
		FROM
			public.tags
		ORDER BY
			id
	`

	var tt []*model.Tag
	if err := tr.db.Select(&tt, query); err != nil {
		return nil, errors.WithStack(err)
	}

	return tt, nil
}

func (tr *tagRepository) ListByName(tagName string) ([]*model.Tag, error) {
	const query = `
		SELECT
				id
			, name
		FROM
			public.tags
		WHERE
			name LIKE $1
		ORDER BY
			id
	`

	var tt []*model.Tag
	if err := tr.db.Select(&tt, query, "%"+tagName+"%"); err != nil {
		return nil, errors.WithStack(err)
	}

	return tt, nil
}

func (tr *tagRepository) ListByPostID(postID int) ([]*model.Tag, error) {
	const query = `
		SELECT
				t.id
			, t.name
		FROM
			public.tags AS t
		INNER JOIN public.post_tag AS pt
			ON t.id = pt.tag_id
		WHERE
			pt.post_id = $1
		ORDER BY
			t.id
	`

	var tt []*model.Tag
	if err := tr.db.Select(&tt, query, postID); err != nil {
		return nil, errors.WithStack(err)
	}

	return tt, nil
}

func (tr *tagRepository) GetByID(tagID int) (*model.Tag, error) {
	const query = `
		SELECT
				id
			, name
		FROM
			public.tags
		WHERE
			id = $1
	`

	var t model.Tag
	if err := tr.db.Get(&t, query, tagID); err != nil {
		return nil, errors.WithStack(err)
	}

	return &t, nil
}

func (tr *tagRepository) Add(tagName string) (*model.Tag, error) {
	const query = `
		INSERT INTO
			public.tags(name)
		VALUES
			($1)
		RETURNING
			id
	`

	t := model.Tag{
		Name: tagName,
	}

	if err := tr.db.QueryRow(query, tagName).Scan(&t.ID); err != nil {
		return nil, errors.WithStack(err)
	}

	return &t, nil
}

func (tr *tagRepository) Attach(postID, tagID int) (*model.PostTag, error) {
	const query = `
		INSERT INTO
			public.post_tag(post_id, tag_id)
		VALUES
			($1, $2)
	`

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
