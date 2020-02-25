package TestMongoDB

import (
	//"fmt"
	"testing"
	"time"

	"github.com/globalsign/mgo/bson"
)

type Data struct {
	Id      bson.ObjectId `bson:"_id"`
	Title   string        `bson:"title"`
	Des     string        `bson:"des"`
	Content string        `bson:"content"`
	Date    time.Time     `bson:"date"`
}

func TestDb(t *testing.T) {
	// testconf.InitTestConf()
	InitTestConf()
	data := &Data{
		Id:      bson.NewObjectId(),
		Title:   "博客的标题 2",
		Des:     "博客描述信息 2",
		Content: "博客的具体内容 2",
		Date:    time.Now(),
	}

	err := Insert("blog", data)
	if err != nil {
		t.Log("insert one doc", err)
	}

	data.Title = "sdjlkfjalkfjla"
	data.Des = "jfalkjfl"
	err = Update("blog", bson.M{"_id": data.Id}, data)
	if err != nil {
		t.Log("update one err", err)
	}

	h := &Heros{Uid: bson.NewObjectId(), Heros: make(map[int64]*Hero)}
	err = Upsert("heros", bson.M{"_id": h.Uid}, h)
	if err != nil {
		t.Log("Upsert err", err)
	}
	t.Log("Upsert suc ")
}
