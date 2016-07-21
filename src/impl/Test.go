package impl

import (
	"fmt"
)

type Test struct {
	Id string
}

func (test *Test) Show(name string, age int, li interface{}) {
	fmt.Println("Test->Show")
}
