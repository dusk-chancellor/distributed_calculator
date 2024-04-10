package main

import (
	"calculator_yandex/internal/storage"
	"context"
	"fmt"
)

func main() {
	ctx := context.TODO()

	db, err := storage.New("./database/storage.db")
	if err != nil {
		panic(err)
	}
}