package main

import "github.com/zevfang/airship/cmd/spider/app"

func main() {
	w := app.NewCollector()
	w.QueryHot()
}
