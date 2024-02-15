package storage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Timeouts struct {
	Plus     string `json:"plus"`
	Minus    string `json:"minus"`
	Multiply string `json:"multiply"`
	Divide   string `json:"divide"`
}

const timeoutsFile = "timeouts.json"

func SetNewTimeout(timeout string) error {
	tmt := ""
	var oper string
	for _, ch := range timeout {
		char := string(ch)
		switch char {
		case "+":
			oper = "Plus"
		case "-":
			oper = "Minus"
		case "*":
			oper = "Multiply"
		case "/":
			oper = "Divide"
		default:
			tmt += char
		}
	}

	var newTimeout Timeouts
	if oper == "Plus" {
		newTimeout.Plus = tmt
	} else if oper == "Minus" {
		newTimeout.Minus = tmt
	} else if oper == "Multiply" {
		newTimeout.Multiply = tmt
	} else if oper == "Divide" {
		newTimeout.Divide = tmt
	} else {
		return nil
	}

	return saveToFile(timeoutsFile, newTimeout)
}

func GetTimeouts() (*Timeouts, error) {
	return getFromTimeoutsFile(timeoutsFile)
}

func getFromTimeoutsFile(fileName string) (*Timeouts, error) {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		logger.Printf("File does not exist: %s", err)
		return &Timeouts{}, nil
	}

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		logger.Printf("Cannot read the file: %s", err)
		return nil, fmt.Errorf("failed to read file: %v", err)
	}

	if len(data) == 0 {
		logger.Println("File is empty, returning new Expressions object")
		return &Timeouts{}, nil
	}

	var timeout Timeouts
	if err := json.Unmarshal(data, &timeout); err != nil {
		logger.Printf("Cannot unmarshal data: %s", err)
		return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}
	logger.Printf("List retrieved from %s", expressionsFile)
	return &timeout, nil
}
