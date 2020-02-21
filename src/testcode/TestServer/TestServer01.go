package testserver

import (
	"fmt"
)

var cache *Cache

// ProfileFunc ProfileFunc
func TestServerFunc() {
	fmt.Println("TestServerFunc ...")
	cache := NewCache()
	cache.AddTable("user1", 10, NewUser)
	cache.AddTable("user2", 10, NewUser)
	cache.AddTable("user3", 10, NewUser)
	cache.AddTable("user4", 10, NewUser)
	cache.AddTable("user5", 10, NewUser)
	cache.AddTable("user6", 10, NewUser)
	cache.AddTable("user7", 10, NewUser)
	cache.AddTable("user8", 10, NewUser)
	cache.AddTable("user9", 10, NewUser)
	cache.AddTable("user10", 10, NewUser)
	cache.AddTable("user11", 10, NewUser)
	cache.AddTable("user12", 10, NewUser)
	cache.AddTable("user13", 10, NewUser)
	cache.AddTable("user14", 10, NewUser)
	cache.AddTable("user15", 10, NewUser)
	cache.AddTable("user16", 10, NewUser)

	AddFunc("0 */1 * * * *", func() {
		cache.Save()
	})
}
