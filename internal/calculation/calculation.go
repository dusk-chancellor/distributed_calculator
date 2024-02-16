package calculation

import "fmt"

type Expression struct {
	expression string
}

func ReceiveExpression(expr string) error {
	fmt.Println(expr)
	return nil
}
