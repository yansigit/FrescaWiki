package controller

import (
	"bytes"
	"frescawiki/app/model"
	"github.com/Kamva/mgm"
	"github.com/kataras/iris/v12"
	"github.com/yuin/goldmark"
	"go.mongodb.org/mongo-driver/mongo"
	"html/template"
)

func Search(ctx iris.Context) {
	docName := ctx.PostValue("search")
	ctx.Redirect("/w/" + docName, 302)
}

func Recent(ctx iris.Context) {
	ctx.View("recent.pug")
}

func Random(ctx iris.Context) {
	docName := "랜덤"
	ctx.Redirect("/w/" + docName, 302)
}

func EditGet(ctx iris.Context) {
	docName := ctx.Params().GetString("doc_name")
	doc, _ := model.SearchDocByTitle(docName)
	doc.Title = docName
	args := &model.Arg{
		Doc:    *doc,
		Recent: []model.Recent{},
	}
	ctx.View("edit.pug", args)
}

func EditPost(ctx iris.Context) {
	title := ctx.PostValue("title")
	body := ctx.PostValue("body")

	var buf bytes.Buffer
	if err := goldmark.Convert([]byte(body), &buf); err != nil {
		panic(err)
	}

	doc, err := model.SearchDocByTitle(title)
	if doc != nil {
		if err == mongo.ErrNoDocuments {
			doc = model.NewDoc(title, body)
			doc.RenderedBody = template.HTML(buf.String())
			err = mgm.Coll(doc).Create(doc)
			if err != nil {
				panic(err)
			}
		} else {
			doc.Body = body
			doc.RenderedBody = template.HTML(buf.String())
			err = mgm.Coll(doc).Update(doc)
			if err != nil {
				panic(err)
			}
		}
	}

	ctx.Redirect("/w/"+title, 302)
}

func RemoveGet(ctx iris.Context) {
	docName := ctx.Params().GetString("doc_name")
	ctx.View("remove.pug", docName)
}

func RemovePost(ctx iris.Context) {

}

func Index(ctx iris.Context) {
	docName := ctx.Params().GetString("doc_name")
	doc, _ := model.SearchDocByTitle(docName)

	doc.Title = docName
	recent, _ := model.GetRecent()
	args := &model.Arg{
		Doc:    *doc,
		Recent: recent,
	}
	ctx.View("index.pug", args)
}