package mongo

import (
	"github.com/zevfang/airship/pkg/adding"
	"gopkg.in/mgo.v2"
)

type Storage struct {
	db *mgo.Session
}

func NewStorage() (*Storage, error) {
	var err error
	s := new(Storage)
	s.db, err = mgo.Dial("localhost:27017")
	return s, err
}

// 添加文章
func (s *Storage) AddArticle(a adding.Article) error {
	newA := Article{
		Title:   a.Title,
		Date:    a.Date,
		Content: a.Content,
	}
	return s.db.DB("airship").C("article").Insert(newA)
}
