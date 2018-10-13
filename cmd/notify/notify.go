package main

import (
	"github.com/deckarep/gosx-notifier"
	"log"
)

func main() {
	ShowNotification("a", "b", "v", "http://www.baidu.com")
}

// 桌面通知
func ShowNotification(title, subtitle, content, link string) {
	note := gosxnotifier.NewNotification(content)
	note.Title = title
	note.Subtitle = subtitle
	note.Link = link
	//setting
	note.Sound = gosxnotifier.Purr
	note.Group = "com.unique.airship.identifier"
	note.Sender = "com.apple.Safari"
	note.AppIcon = "images/gopher.png"
	if err := note.Push(); err != nil {
		log.Println("Uh oh!")
	}
}
