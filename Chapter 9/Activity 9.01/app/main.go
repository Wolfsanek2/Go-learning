// consuming multiple modules
package main

import (
	"fmt"

	"github.com/google/uuid"
	"rsc.io/quote"
)

func main() {
	id := uuid.New()
	str := quote.Opt()
	fmt.Println(id)
	fmt.Println(str)
}
