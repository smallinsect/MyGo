package testdemo

import (
	"fmt"
	"testing"

	"github.com/globalsign/mgo/bson"
)

// bson.ObjectId 和 string 互相转换
func TestMongoDBFunc(t *testing.T) {
	// 创建id
	uid := bson.NewObjectId
	fmt.Println(uid)
	// bson.ObjectId -> string
	str := uid.Hex()
	fmt.Println(str)
	// string -> bson.ObjectId
	fmt.Println(bson.ObjectIdHex(str))

	t.Logf("test ChannelFunc succ")
}
