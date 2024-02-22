package calculation

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func ReceiveAndPost(id, expr string) error {
	answer, err1 := Evaluate(expr)

	returnAnswer := strconv.FormatFloat(answer, 'g', -1, 64)
	if err2 := PostAnswer(id+":"+returnAnswer, err1); err2 != nil {
		log.Fatalf("Error sending Post: %v", err2)
		return err2
	}
	return nil
}

func Evaluate(expr string) (float64, error) {
	var stack Stack
	tokens := strings.Split(expr, " ")

	for _, token := range tokens {
		if token == "+" || token == "-" || token == "*" || token == "/" {
			// если токен - оператор, то забираем 2 последних элемента со стека
			op1 := stack.Pop()
			op2 := stack.Pop()
			ans, err := Calculate(op1, op2, token)
			if err != nil {
				return 0, err
			}
			stack.Push(ans)
		} else {
			// если токен не оператор - то операнд(число), пушим в стек :)
			op, _ := strconv.ParseFloat(token, 64)

			stack.Push(op)
		}
	}
	// в этом случае, последний элемент стека и есть наш ответ
	return stack.Pop(), nil
}

// Calculate - вычисляет
func Calculate(op1, op2 float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return op2 + op1, nil
	case "-":
		return op2 - op1, nil
	case "*":
		return op2 * op1, nil
	case "/":
		if op1 == 0 {
			return 0, fmt.Errorf("Division by zero")
		}
		return op2 / op1, nil
	default:
		return 0, fmt.Errorf("Unknown operation")
	}
}

func PostAnswer(sendingData string, err error) error {
	fmt.Println("PostAnswer", sendingData)
	if err != nil {
		sendingData += ":error"
	} else {
		sendingData += ":ok"
	}
	req, err := http.NewRequest(
		http.MethodPost,
		"http://localhost:8080/getanswer",
		bytes.NewBuffer([]byte(sendingData)),
	)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return nil
}
