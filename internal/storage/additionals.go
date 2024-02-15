package storage

import (
	"encoding/json"
	"io/ioutil"
)

func saveToFile(fileName string, val interface{}) error {
	data, err := json.Marshal(val)
	if err != nil {
		logger.Printf("Failed to marshal: %v", err)
		return err
	}

	if err := ioutil.WriteFile(fileName, data, 0644); err != nil {
		logger.Printf("Failed to write to file: %v", err)
		return err
	}
	logger.Printf("Successfully wrote to file: %s", fileName)

	return nil
}

func ClearExpressionsList() error {
	defaultState := NewExpressionsList()

	jsonDefault, err := json.Marshal(defaultState)
	if err != nil {
		logger.Printf("Failed to marshal default state: %v", err)
		return err
	}

	err = ioutil.WriteFile(expressionsFile, jsonDefault, 0644)
	if err != nil {
		logger.Printf("Failed to write default state to file: %v", err)
		return err
	}
	logger.Printf("Expressions list cleared")
	return nil
}
