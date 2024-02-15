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
	timeouts, err := getFromTimeoutsFile(timeoutsFile)
	if err != nil {
		logger.Printf("Cannot read getFromTimeoutsFile: %s", err)
		return err
	}
	n := ""
	var oper int
	for _, ch := range timeout {
		char := string(ch)
		switch char {
		case "+":
			oper = 1
		case "-":
			oper = 2
		case "*":
			oper = 3
		case "/":
			oper = 4
		default:
			n += char
		}
	}

	switch oper {
	case 1:
		timeouts.Plus = n
	case 2:
		timeouts.Minus = n
	case 3:
		timeouts.Multiply = n
	case 4:
		timeouts.Divide = n
	}

	return saveToFile(timeoutsFile, timeouts)
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
