package service

import (
	"github.com/oshou/AwesomeMusic-api/db"
	//"github.com/oshou/AwesomeMusic-api/entity"
)

type SearchService struct{}

//type Post entity.Post

func (ss SearchService) GetByUserId(q string) ([]Post, error) {

	var p []Post

	stmt := db.GetDBConn()
	stmt = stmt.Table("post")
	stmt = stmt.Select("post.id,post.url,post.message")
	stmt = stmt.Joins("INNER JOIN user on user.id = post.user_id")
	stmt = stmt.Where("user.id = ?", q)
	if err := stmt.Find(&p).Error; err != nil {
		return nil, err
	}
	return p, nil
}

func (ss SearchService) GetByUserName(q string) ([]Post, error) {

	var p []Post

	stmt := db.GetDBConn()
	stmt = stmt.Table("post")
	stmt = stmt.Select("post.id,post.url,post.message")
	stmt = stmt.Joins("INNER JOIN user on user.id = post.user_id")
	stmt = stmt.Where("user.name LIKE ?", "%"+q+"%")
	if err := stmt.Find(&p).Error; err != nil {
		return nil, err
	}
	return p, nil
}

func (ss SearchService) GetByTagId(q string) ([]Post, error) {

	var p []Post

	stmt := db.GetDBConn()
	stmt = stmt.Table("post")
	stmt = stmt.Select("post.id,post.url,post.message")
	stmt = stmt.Joins("INNER JOIN post_tag on post_tag.post_id = post.id")
	stmt = stmt.Joins("INNER JOIN tag on post_tag.tag_id = tag.id")
	stmt = stmt.Where("tag.id = ?", q)
	if err := stmt.Find(&p).Error; err != nil {
		return nil, err
	}
	return p, nil
}

func (ss SearchService) GetByTagName(q string) ([]Post, error) {

	var p []Post

	stmt := db.GetDBConn()
	stmt = stmt.Table("post")
	stmt = stmt.Select("post.id,post.url,post.message")
	stmt = stmt.Joins("INNER JOIN post_tag on post_tag.post_id = post.id")
	stmt = stmt.Joins("INNER JOIN tag on post_tag.tag_id = tag.id")
	stmt = stmt.Where("tag.name LIKE ?", "%"+q+"%")
	if err := stmt.Find(&p).Error; err != nil {
		return nil, err
	}
	return p, nil
}
