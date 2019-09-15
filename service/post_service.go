package service

import (
	"github.com/oshou/AwesomeMusic-api/db"
	"github.com/oshou/AwesomeMusic-api/entity"
)

type PostService struct{}

type Post entity.Post

func (ps PostService) GetAll() ([]Post, error) {

	var p []Post

	stmt := db.GetDBConn()
	stmt = stmt.Table("post")
	stmt = stmt.Select("id,user_id,title,url,message")
	if err := stmt.Find(&p).Error; err != nil {
		return nil, err
	}
	return p, nil
}

func (ps PostService) Add(user_id int, title, url, message string) (Post, error) {

	var p Post
	p.UserID = user_id
	p.Title = title
	p.URL = url
	p.Message = message

	stmt := db.GetDBConn()
	stmt = stmt.Table("post")
	if err := stmt.Create(&p).Error; err != nil {
		return p, err
	}
	return p, nil
}

func (ps PostService) GetById(post_id int) (Post, error) {

	var p Post

	stmt := db.GetDBConn()
	stmt = stmt.Table("post")
	stmt = stmt.Select("id,user_id,title,url,message")
	stmt = stmt.Where("id = ?", post_id)
	if err := stmt.First(&p).Error; err != nil {
		return p, err
	}
	return p, nil
}

func (ps PostService) GetByTagId(tag_id int) ([]Post, error) {

	var p []Post

	stmt := db.GetDBConn()
	stmt = stmt.Table("post")
	stmt = stmt.Select("post.id,post.user_id,post.title,post.url,post.message")
	stmt = stmt.Joins("INNER JOIN post_tag on post_tag.post_id = post.id")
	stmt = stmt.Joins("INNER JOIN tag on post_tag.tag_id = tag.id")
	stmt = stmt.Where("tag.id = ?", tag_id)
	if err := stmt.Find(&p).Error; err != nil {
		return nil, err
	}
	return p, nil
}

func (ps PostService) GetByUserId(user_id int) ([]Post, error) {

	var p []Post

	stmt := db.GetDBConn()
	stmt = stmt.Table("post")
	stmt = stmt.Select("post.id,post.user_id,post.title,post.url,post.message")
	stmt = stmt.Joins("INNER JOIN user on user.id = post.user_id")
	stmt = stmt.Where("user.id = ?", user_id)
	if err := stmt.Find(&p).Error; err != nil {
		return nil, err
	}
	return p, nil
}

func (ps PostService) DeleteById(post_id int) error {

	var p Post

	stmt := db.GetDBConn()
	stmt = stmt.Table("post")
	stmt = stmt.Where("id = ?", post_id)
	if err := stmt.Delete(&p).Error; err != nil {
		return err
	}
	return nil
}
