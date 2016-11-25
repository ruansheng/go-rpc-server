package impl

import (
	//"encoding/json"
	"fmt"
	//"utils/goconfig"
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

/*
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

func (test *Test) GetProfile2(uid string) User {
	c, err := LoadConfigFile("conf.ini")
	if err == nil {
		fmt.Println(c)
		port, _ := c.GetValue("", "port")
		fmt.Println(port)
		vInt, _ := c.Int("", "port")
		fmt.Println(vInt)
		//invoke(vInt)
	}
	userinfo := User{"ruansheng", 25, "M"}
	fmt.Println(userinfo)
	return userinfo
}

func getCmd(id string, action string, m string, args interface{}) string {
	params := make(map[string]interface{})
	params["m"] = m
	params["args"] = args

	cmd := make(map[string]interface{})
	cmd["id"] = id
	cmd["action"] = action
	cmd["params"] = params

	j, err := json.Marshal(cmd)
	if err != nil {
		panic(err)
	}
	js := string(j)
	return js
}
*/
