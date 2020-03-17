package main

import (
	"frescawiki/app/controller"
	"frescawiki/app/setup"
	"github.com/kataras/iris/v12"
)

func main() {
	app := setup.Initialize()

	app.Get("/", func(ctx iris.Context) {
		ctx.Redirect("/w/IT스터디위키:대문")
	})

	app.Get("/w/{doc_name:path}", controller.Index)
	app.Get("/edit/{doc_name:path}", controller.EditGet)
	app.Get("/remove/{doc_name:path}", controller.RemoveGet)
	app.Get("/recent", controller.Recent)
	app.Get("/random", controller.Random)

	app.Post("/edit/{doc_name:path}", controller.EditPost)
	app.Post("/remove/{doc_name:path}", controller.RemovePost)
	app.Post("/search", controller.Search)

	// http://localhost:8080
	app.Run(iris.Addr(":8080"))
}

