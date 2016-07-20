package impl

import (
	"fmt"
)

type Test struct {
	Id string
}

func (test *Test) Show() {
	fmt.Println("Test->Show")
}
