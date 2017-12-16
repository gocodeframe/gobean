package example

import (
	"fmt"
)

type Person struct {
	Phone Phone
}

func (this *Person) PhoneType() {
	fmt.Println(this.Phone.Type())
}
