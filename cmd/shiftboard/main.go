package main

import (
	"context"
	"fmt"
	"os"

	"shiftboard/internal/model"
	"shiftboard/internal/service"
)

func main() {
	board := service.NewBoard()
	orders := []model.WorkOrder{
		{ID: "WO-101", Title: "Inspect loading dock", Zone: "north", Tags: []string{"safety", "daily"}},
		{ID: "WO-102", Title: "Restock first-aid kits", Zone: "south", Tags: []string{"inventory"}},
	}
	if err := board.Import(context.Background(), orders); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Print(board.Report())
}
