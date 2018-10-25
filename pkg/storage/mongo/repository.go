package mongo

import (
	"github.com/zevfang/airship/pkg/adding"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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
func (s *Storage) AddArticle(article adding.Article) error {
	newA := Article{
		Title:   article.Title,
		Date:    article.Date,
		Content: article.Content,
	}
	return s.db.DB("airship").C("article").Insert(newA)
}

// 查询文章
func (s *Storage) FindArticle(title string) ([]adding.Article, error) {
	result := make([]adding.Article, 0)
	err := s.db.DB("airship").C("article").Find(bson.M{"title": title}).All(result)
	return result, err
}
