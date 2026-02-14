package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ProArash/todo-manager-cli/internal/db"
	"github.com/ProArash/todo-manager-cli/internal/todo"
)

func main() {
	dsn := "host=localhost user=postgres password=password dbname=task port=5432 sslmode=disable TimeZone=Asia/Tehran"
	db, err := db.Init(dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.AutoMigrate(&todo.TaskModel{}); err != nil {
		log.Fatal(err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)

	defer stop()

	todoService := todo.ServiceInstance(db)

	task, err := todoService.CreateTask(ctx, "My task 2")
	if err != nil {
		log.Fatalf("error while creating new task=> %v", err)
	}
	fmt.Println(task)
}
