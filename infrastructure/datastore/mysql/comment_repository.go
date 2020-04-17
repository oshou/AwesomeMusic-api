package mysql

import (
	"github.com/jmoiron/sqlx"
	"github.com/oshou/AwesomeMusic-api/domain/model"
	"github.com/oshou/AwesomeMusic-api/domain/repository"
)

type commentRepository struct {
	db *sqlx.DB
}

var _ repository.ICommentRepository = (*commentRepository)(nil)

func NewCommentRepository(db *sqlx.DB) repository.ICommentRepository {
	return &commentRepository{
		db: db,
	}
}

func (cr *commentRepository) GetAll(postID int) ([]*model.Comment, error) {
	var cc []*model.Comment

	query := `SELECT
							id,
							user_id,
							post_id,
							comment
						FROM
							comment
						WHERE
							post_id = ?`

	if err := cr.db.Select(&cc, query, postID); err != nil {
		return nil, err
	}

	return cc, nil
}

func (cr *commentRepository) GetByID(commentID int) (*model.Comment, error) {
	c := &model.Comment{}

	query := `SELECT
							id,
							user_id,
							post_id,
							comment
						FROM
							comment
						WHERE
							id = ?`

	if err := cr.db.Get(c, query, commentID); err != nil {
		return c, err
	}

	return c, nil
}

func (cr *commentRepository) Add(postID, userID int, comment string) (*model.Comment, error) {
	query := `INSERT INTO
							comment(post_id, user_id, comment)
						VALUES
							(?, ?, ?)`

	result, err := cr.db.Exec(query, postID, userID, comment)
	if err != nil {
		return nil, err
	}

	c := &model.Comment{
		UserID:  userID,
		PostID:  postID,
		Comment: comment,
	}
	i64, _ := result.LastInsertId()
	c.ID = int(i64)

	return c, nil
}
