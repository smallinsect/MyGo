package TestMongoDB

import (
	"github.com/globalsign/mgo"

	// "cmt1-server/conf"
	"testcode/TestMongoDB/mongodb"
)

var dbname string

type MongoDb struct {
}

var MongoDBModule *MongoDb = new(MongoDb)

func (m *MongoDb) Init() {
	dbname = Conf.MongoDB.DbName
	mongodb.Init()
}

func (m *MongoDb) Destroy() {
}

func (m *MongoDb) Run(closeSig chan bool) {

}

func Init() {
	dbname = Conf.MongoDB.DbName
	mongodb.Init()
}

func IsEmpty(collection string) bool {
	return mongodb.IsEmpty(dbname, collection)
}

func Count(collection string, query interface{}) (int, error) {
	return mongodb.Count(dbname, collection, query)
}

func Insert(collection string, docs ...interface{}) error {
	return mongodb.Insert(dbname, collection, docs...)
}

func FindOne(collection string, query, selector, result interface{}) error {
	return mongodb.FindOne(dbname, collection, query, selector, result)
}

func FindAll(collection string, query, selector, result interface{}) error {
	return mongodb.FindAll(dbname, collection, query, selector, result)
}

func FindAllSortLimit(collection string, query, selector, result interface{}, sorts []string, limit int) error {
	return mongodb.FindAllSortLimit(dbname, collection, query, selector, result, sorts, limit)
}

func FindPage(collection string, page, limit int, query, selector, result interface{}) error {
	return mongodb.FindPage(dbname, collection, page, limit, query, selector, result)
}

func FindIter(dbname, collection string, query interface{}) *mgo.Iter {
	return mongodb.FindIter(dbname, collection, query)
}

func Update(collection string, selector, update interface{}) error {
	return mongodb.Update(dbname, collection, selector, update)
}

func Upsert(collection string, selector, update interface{}) error {
	return mongodb.Upsert(dbname, collection, selector, update)
}

func UpdateAll(collection string, selector, update interface{}) error {
	return mongodb.UpdateAll(dbname, collection, selector, update)
}

func Remove(collection string, selector interface{}) error {
	return mongodb.Remove(dbname, collection, selector)
}

func RemoveAll(collection string, selector interface{}) error {
	return mongodb.RemoveAll(dbname, collection, selector)
}

//insert one or multi documents
func BulkInsert(collection string, docs ...interface{}) (*mgo.BulkResult, error) {
	return mongodb.BulkInsert(dbname, collection, docs...)
}

func BulkRemove(collection string, selector ...interface{}) (*mgo.BulkResult, error) {
	return mongodb.BulkRemove(dbname, collection, selector...)
}

func BulkRemoveAll(collection string, selector ...interface{}) (*mgo.BulkResult, error) {
	return mongodb.BulkRemoveAll(dbname, collection, selector...)
}

func BulkUpdate(collection string, pairs ...interface{}) (*mgo.BulkResult, error) {
	return mongodb.BulkUpdate(dbname, collection, pairs...)
}

func BulkUpdateAll(collection string, pairs ...interface{}) (*mgo.BulkResult, error) {
	return mongodb.BulkUpdateAll(dbname, collection)
}

func BulkUpsert(collection string, pairs ...interface{}) (*mgo.BulkResult, error) {
	return mongodb.BulkUpsert(dbname, collection, pairs)
}

func PipeAll(collection string, pipeline, result interface{}, allowDiskUse bool) error {
	return mongodb.PipeAll(dbname, collection, pipeline, result, allowDiskUse)
}

func PipeOne(collection string, pipeline, result interface{}, allowDiskUse bool) error {
	return mongodb.PipeOne(dbname, collection, pipeline, result, allowDiskUse)
}

func PipeIter(collection string, pipeline interface{}, allowDiskUse bool) *mgo.Iter {
	return mongodb.PipeIter(dbname, collection, pipeline, allowDiskUse)
}

func Explain(collection string, pipeline, result interface{}) error {
	return mongodb.Explain(dbname, collection, pipeline, result)
}
func GridFSCreate(prefix, name string) (*mgo.GridFile, error) {
	return mongodb.GridFSCreate(dbname, prefix, name)
}

func GridFSFindOne(prefix string, query, result interface{}) error {
	return mongodb.GridFSFindOne(dbname, prefix, query, result)
}

func GridFSFindAll(prefix string, query, result interface{}) error {
	return mongodb.GridFSFindAll(dbname, prefix, query, result)
}

func GridFSOpen(prefix, name string) (*mgo.GridFile, error) {
	return mongodb.GridFSOpen(dbname, prefix, name)
}

func GridFSRemove(prefix, name string) error {
	return mongodb.GridFSRemove(dbname, prefix, name)
}
