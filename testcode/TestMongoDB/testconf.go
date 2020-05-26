package TestMongoDB

import (
	"log"
	"math/rand"
	"time"
	// "cmt1-server/conf"
	// "cmt1-server/data"
	// "cmt1-server/db"
)

func InitTestConf() {
	// err := conf.Conf.Load("/home/luxingwen/work/server/trunk/cmt1-server/appconf/app.conf")
	err := Conf.Load("E:/MyGo/src/testcode/TestMongoDB/app.conf")
	if err != nil {
		log.Fatal("conf load err:", err)
	}

	Init()
	// err = data.Init("/home/luxingwen/work/server/trunk/cmt1-server/data/json")
	// err = data.Init("/home/luxingwen/work/server/trunk/cmt1-server/data/json")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	rand.Seed(time.Now().UnixNano())
}

func InitTestConf2() {
	err := Conf.Load("E:/MyGo/src/testcode/TestMongoDB/app.conf")
	if err != nil {
		log.Fatal("conf load err:", err)
	}

	Init()
	// err = data.Init("D:\\work\\cmt1\\Miciteam1-Server\\trunk\\cmt1-server\\data\\json")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	rand.Seed(time.Now().UnixNano())
}
