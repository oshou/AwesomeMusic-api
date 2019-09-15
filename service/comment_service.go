package service

import (
	"github.com/oshou/AwesomeMusic-api/db"
	"github.com/oshou/AwesomeMusic-api/entity"
)

type CommentService struct{}

type Comment entity.Comment

func (cs CommentService) GetAll(post_id int) ([]Comment, error) {

	var c []Comment

	stmt := db.GetDBConn()
	stmt = stmt.Table("comment")
	stmt = stmt.Select("id,user_id,comment")
	stmt = stmt.Where("post_id = ?", post_id)
	if err := stmt.Find(&c).Error; err != nil {
		return nil, err
	}
	return c, nil
}

func (cs CommentService) GetById(comment_id int) (Comment, error) {

	var c Comment

	stmt := db.GetDBConn()
	stmt = stmt.Table("comment")
	stmt = stmt.Select("id,user_id,comment")
	stmt = stmt.Where("id = ?", comment_id)
	if err := stmt.First(&c).Error; err != nil {
		return c, err
	}
	return c, nil
}

func (cs CommentService) Add(post_id, user_id int, comment string) (Comment, error) {

	var c Comment
	c.PostID = post_id
	c.UserID = user_id
	c.Comment = comment

	stmt := db.GetDBConn()
	stmt = stmt.Table("comment")
	if err := stmt.Create(&c).Error; err != nil {
		return c, err
	}
	return c, nil
}
