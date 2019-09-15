package service

import (
	"github.com/oshou/AwesomeMusic-api/db"
)

type SearchService struct{}

//type Post entity.Post

func (ss SearchService) GetByPostTitle(q string) ([]Post, error) {

	var p []Post

	stmt := db.GetDBConn()
	stmt = stmt.Table("post")
	stmt = stmt.Select("id,user_id,title,url,message")
	stmt = stmt.Where("post.title LIKE ?", "%"+q+"%")
	if err := stmt.Find(&p).Error; err != nil {
		return nil, err
	}
	return p, nil
}

func (ss SearchService) GetByUserName(q string) ([]Post, error) {

	var p []Post

	stmt := db.GetDBConn()
	stmt = stmt.Table("user")
	stmt = stmt.Select("post.id,post.user_id,post.title,post.url,post.message")
	stmt = stmt.Where("user.name LIKE ?", "%"+q+"%")
	if err := stmt.Find(&p).Error; err != nil {
		return nil, err
	}
	return p, nil
}

func (ss SearchService) GetByTagName(q string) ([]Post, error) {

	var p []Post

	stmt := db.GetDBConn()
	stmt = stmt.Table("post")
	stmt = stmt.Select("post.id,post.user_id,post.title,post.url,post.message")
	stmt = stmt.Joins("INNER JOIN post_tag on post_tag.post_id = post.id")
	stmt = stmt.Joins("INNER JOIN tag on post_tag.tag_id = tag.id")
	stmt = stmt.Where("tag.name LIKE ?", "%"+q+"%")
	if err := stmt.Find(&p).Error; err != nil {
		return nil, err
	}
	return p, nil
}
