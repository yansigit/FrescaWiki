package model

import (
	"github.com/Kamva/mgm"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"html/template"
	"time"
)

type Arg struct {
	Doc    Doc
	Recent []Recent
}

type Doc struct {
	mgm.DefaultModel				`bson:",inline"`
	Title            string      	`json:"title" bson:"title"`
	Body             string         `json:"body" bson:"body"`
	RenderedBody     template.HTML  `json:"renderedBody" bson:"renderedBody"`
}

type Recent struct {
	Title string
	Date  time.Time
}

func NewDoc(title string, body string) *Doc {
	return &Doc{
		Title: title,
		Body:  body,
	}
}

func NewRecent(title string, date time.Time) *Recent {
	return &Recent{
		Title: title,
		Date:  date,
	}
}

func GetRecent() (recent []Recent, err error) {
	var docs []Doc
	coll := mgm.Coll(&Doc{})
	findOptions := options.Find()
	findOptions.SetLimit(10)
	findOptions.SetSort(bson.M{"_id":1})
	err = coll.SimpleFind(&docs, bson.M{}, options.Find())

	if err == nil {
		for _, doc := range docs {
			newRecent := NewRecent(doc.Title, doc.UpdatedAt)
			recent = append(recent, *newRecent)
		}
	}

	return
}

func SearchDocByTitle(title string) (doc *Doc, err error) {
	doc = &Doc{}
	coll := mgm.Coll(doc)
	err = coll.First(bson.M{"title": title}, doc)
	return
}