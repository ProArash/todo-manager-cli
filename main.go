package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/ProArash/todo-manager-cli/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("task.db"), &gorm.Config{})
	ctx := context.Background()
	db.AutoMigrate(&model.Task{})

	tasks, err := gorm.G[model.Task](db).Find(ctx)

	fmt.Println("Current tasks:")
	for _, task := range tasks {
		fmt.Printf("%s\n", task.Title)
	}

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("%s\n", "=> TODO manager CLI by @ProArash")
	for range 40 {
		fmt.Print("-")
	}
	fmt.Print("\n[Enter the task name]: ")
	scanner.Scan()
	taskInput := scanner.Text()
	fmt.Println(taskInput)

	gorm.G[model.Task](db).Create(ctx, &model.Task{Title: taskInput})
	fmt.Printf("Task %s created!", taskInput)
}
