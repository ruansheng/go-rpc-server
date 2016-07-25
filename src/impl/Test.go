package impl

import (
	"fmt"
)

type Test struct {
	Id string
}

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  string `json:"sex"`
}

func (test *Test) Show(name string, age float64, li []interface{}, price float64) {
	fmt.Println("Test->Show:", name, age, li, price)
}

func (test *Test) GetProfile(uid string) map[string]interface{} {
	userinfo := make(map[string]interface{})
	userinfo["name"] = "ruansheng"
	userinfo["age"] = 25
	userinfo["sex"] = "M"
	fmt.Println(userinfo)
	return userinfo
}

func (test *Test) GetProfile1(uid string) User {
	userinfo := User{"ruansheng", 25, "M"}
	fmt.Println(userinfo)
	return userinfo
}
