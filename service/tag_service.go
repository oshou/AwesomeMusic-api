package service

import (
	"github.com/oshou/AwesomeMusic-api/db"
	"github.com/oshou/AwesomeMusic-api/entity"
)

type TagService struct{}

type Tag entity.Tag
type PostTag entity.PostTag

func (ts TagService) GetAll() ([]Tag, error) {

	var t []Tag

	stmt := db.GetDBConn()
	stmt = stmt.Table("tag")
	stmt = stmt.Select("id,name")
	if err := stmt.Find(&t).Error; err != nil {
		return nil, err
	}
	return t, nil
}

func (ts TagService) GetById(tag_id int) (Tag, error) {

	var t Tag

	stmt := db.GetDBConn()
	stmt = stmt.Table("tag")
	stmt = stmt.Select("id,name")
	stmt = stmt.Where("id = ?", tag_id)
	if err := stmt.First(&t).Error; err != nil {
		return t, err
	}
	return t, nil
}

func (ts TagService) Add(name string) (Tag, error) {

	var t Tag
	t.Name = name

	stmt := db.GetDBConn()
	stmt = stmt.Table("tag")
	if err := stmt.Create(&t).Error; err != nil {
		return t, err
	}
	return t, nil
}

func (ts TagService) Attach(post_id, tag_id int) (PostTag, error) {

	var pt PostTag
	pt.PostID = post_id
	pt.TagID = tag_id

	stmt := db.GetDBConn()
	stmt = stmt.Table("post_tag")
	if err := stmt.Create(&pt).Error; err != nil {
		return pt, err
	}
	return pt, nil
}
