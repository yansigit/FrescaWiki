package setup

import (
	"frescawiki/app/model"
	"github.com/Kamva/mgm"
	"github.com/kataras/iris/v12"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Initialize() *iris.Application {
	app := iris.New()
	developing := false

	_ = mgm.SetDefaultConfig(nil, "Wiki_DB", options.Client().ApplyURI("mongodb://localhost:27017"))
	_, err := mgm.Coll(&model.Doc{}).Indexes().CreateOne(mgm.Ctx(), mongo.IndexModel{
		Keys:    bson.M{"title": 1},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		panic(err)
	}

	if developing {
		tmpl := iris.Pug("./templates", ".pug")
		app.RegisterView(tmpl)

		app.HandleDir("/static", "./assets")
	} else {
		tmpl := iris.Pug("./templates", ".pug").Binary(Asset, AssetNames)
		app.RegisterView(tmpl)

		app.HandleDir("/static", "./assets", iris.DirOptions{
			Asset:      Asset,
			AssetInfo:  AssetInfo,
			AssetNames: AssetNames,
			Gzip:       true,
		})
	}

	return app
}
