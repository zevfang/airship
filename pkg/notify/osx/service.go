package osx

import "github.com/deckarep/gosx-notifier"

type Service interface {
	ShowNotification(...Message) error
}

type Repository interface {
	ShowNotification(Message) error
}

type service struct {
	msg Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) ShowNotification(msg ...Message) error {

	var errs error
	for _, m := range msg {
		note := gosxnotifier.NewNotification(m.Content)
		note.Title = m.Title
		note.Subtitle = m.Subtitle
		note.Link = m.Link
		//setting
		note.Sound = gosxnotifier.Purr
		note.Group = "com.unique.airship.identifier"
		note.Sender = "com.apple.Safari"
		note.AppIcon = "images/gopher.png"
		if err := note.Push(); err != nil {
			errs.Error() += ";" + err.Error()
		}
	}
	return errs
}
