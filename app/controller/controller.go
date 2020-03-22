package controller

import (
	"bytes"
	"frescawiki/app/model"
	"github.com/Kamva/mgm"
	"github.com/kataras/iris/v12"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"go.mongodb.org/mongo-driver/mongo"
	"html/template"
	"path"
	"strings"
)

func Search(ctx iris.Context) {
	docName := ctx.PostValue("search")
	ctx.Redirect("/w/"+docName, 302)
}

func Recent(ctx iris.Context) {
	ctx.View("recent.pug")
}

func Random(ctx iris.Context) {
	docName := "랜덤"
	ctx.Redirect("/w/"+docName, 302)
}

func EditGet(ctx iris.Context) {
	docName := ctx.Params().GetString("doc_name")
	doc, _ := model.SearchDocByTitle(docName)
	doc.Title = docName
	recent, _ := model.GetRecent()

	ctx.ViewData("Doc", *doc)
	ctx.ViewData("Recent", recent)
	ctx.View("edit.pug")
}

func EditPost(ctx iris.Context) {
	title := ctx.PostValue("title")
	body := ctx.PostValue("body")

	if strings.TrimSpace(body) == "" {
		doc, err := model.SearchDocByTitle(title)
		if err != mongo.ErrNoDocuments {
			err = mgm.Coll(doc).Delete(doc)
			ctx.Redirect("/w/"+title, 302)
			return
		}
	}

	var buf bytes.Buffer
	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	)
	if err := md.Convert([]byte(body), &buf); err != nil {
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

func Index(ctx iris.Context) {
	docName := ctx.Params().GetString("doc_name")
	doc, _ := model.SearchDocByTitle(docName)

	doc.Title = docName
	recent, _ := model.GetRecent()

	var parentDoc string
	if strings.Contains(docName, "/") {
		parentDoc = path.Dir(docName)
	}

	ctx.ViewData("Doc", doc)
	ctx.ViewData("ParentDoc", parentDoc)
	ctx.ViewData("Recent", recent)
	ctx.View("index.pug")
}
