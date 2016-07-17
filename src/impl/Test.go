package impl

import (
	"fmt"
)

type Test struct {
}

func (test *Test) Show() {
	fmt.Println("Test->Show")
}
