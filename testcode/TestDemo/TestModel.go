package testdemo

// User ...
type User struct {
	Name string
}

//Person ...
type Person struct {
	Uid  int32
	Name string //名字
	Age  int32  //年龄
}

// Student ...
type Student struct {
	Name  string
	Score float64
}

// NewStudent ...
func NewStudent(n string, s float64) *Student {
	return &Student{
		Name:  n,
		Score: s,
	}
}
