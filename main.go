package main

import (
	"fmt"
	"github.com/emrekursatt/JuniorProject/api"
	"github.com/emrekursatt/JuniorProject/configs"
	_ "github.com/emrekursatt/JuniorProject/docs"
	"github.com/emrekursatt/JuniorProject/repository"
	"github.com/emrekursatt/JuniorProject/services"
	_ "github.com/go-sql-driver/mysql"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
	"os"
)

// @title Swagger Example API
// @version v1
// @description This is a sample server Petstore server.
// @host localhost:8080
// @BasePath /api
func main() {
	var cfg = configs.DefaultConfig()

	var database, _ = configs.NewConnection(cfg)

	createTaskTable, _ := configs.CreateTaskTable(database, configs.TABLE_NAME)
	if createTaskTable == nil {
		log.Fatal("Database crate table failed : ", configs.TABLE_NAME)
	}

	createUserTable, _ := configs.CreateUserTable(database, configs.USER)
	if createUserTable == nil {
		log.Fatal("Database crate table failed : ", configs.USER)
	}

	TaskRepositoryDB := repository.NewTaskRepository(database)

	ts := api.TaskHandler{Service: services.NewTaskService(&TaskRepositoryDB)}
	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)
	http.HandleFunc("/api/task/create", ts.Insert)
	http.HandleFunc("/api/task/delete", ts.Delete)
	http.HandleFunc("/api/task/update", ts.Update)
	http.HandleFunc("/api/task/getAll", ts.GetAllTasks)

	UserRepositoryDB := repository.NewUserRepository(database)
	us := api.UserHandler{Service: services.NewUserService(&UserRepositoryDB)}
	http.HandleFunc("/api/user/register", us.Insert)
	http.HandleFunc("/api/user/delete", us.Delete)
	http.HandleFunc("/api/user/update", us.UpdatePassword)
	http.HandleFunc("/api/user/getAll", us.GetAllUsers)
	http.HandleFunc("/api/user/login", us.Login)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Printf("Open http://localhost:%s in the browser", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))

}
