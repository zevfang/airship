package adding

type Service interface {
	AddArticle(article ...Article)
}

type Repository interface {
	AddArticle(article Article) error
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) AddArticle(a ...Article) {
	for _, data := range a {
		_ = s.r.AddArticle(data)
	}
}
