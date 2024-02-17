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
	returnID, answer := Evaluate(id, expr)
	returnAnswer := strconv.FormatFloat(answer, 'g', -1, 64)
	if err := PostAnswer(returnID + ":" + returnAnswer); err != nil {
		log.Fatalf("Error sending Post: %v", err)
		return err
	}
	return nil
}

func Evaluate(id, expr string) (string, float64) {
	var stack Stack
	tokens := strings.Split(expr, " ")

	for _, token := range tokens {
		if token == "+" || token == "-" || token == "*" || token == "/" {
			// если токен - оператор, то забираем 2 последних элемента со стека
			op1 := stack.Pop()
			op2 := stack.Pop()
			ans := Calculate(op1, op2, token)
			stack.Push(ans)
		} else {
			// если токен не оператор - то операнд(число), пушим в стек :)
			op, _ := strconv.ParseFloat(token, 64)

			stack.Push(op)
		}
	}
	// в этом случае, последний элемент стека и есть наш ответ
	return id, stack.Pop()
}

// Calculate - вычисляет
func Calculate(op1, op2 float64, operator string) float64 {
	switch operator {
	case "+":
		return op2 + op1
	case "-":
		return op2 - op1
	case "*":
		return op2 * op1
	case "/":
		if op1 == 0 {
			return 0
		}
		return op2 / op1
	}
	return 0
}

func PostAnswer(sendingData string) error {
	fmt.Println("PostAnswer", sendingData)
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
