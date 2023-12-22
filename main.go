package main

import (
	"database/sql"
	"fmt"
	"github.com/emrekursatt/JuniorProject/api"
	"github.com/emrekursatt/JuniorProject/configs"
	_ "github.com/emrekursatt/JuniorProject/docs"
	"github.com/emrekursatt/JuniorProject/repository"
	"github.com/emrekursatt/JuniorProject/services"
	_ "github.com/go-sql-driver/mysql"
	"github.com/prometheus/client_golang/prometheus"
	_ "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	_ "github.com/prometheus/client_golang/prometheus/promhttp"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
	"os"
	"sync"
)

var mt sync.Mutex
var tableCreated = make(chan bool)

var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Number of HTTP requests",
		},
		[]string{"path"},
	)
)

func init() {
	prometheus.MustRegister(httpRequestsTotal)
}

func prometheusMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		httpRequestsTotal.With(prometheus.Labels{"path": r.URL.Path}).Inc()
		next.ServeHTTP(w, r)
	})
}

// @title Swagger
// @version v1
// @description This is a sample server Task server.
// @host localhost:8080
// @BasePath /api
func main() {

	var wg sync.WaitGroup

	wg.Add(2)

	var cfg = configs.DefaultConfig()

	var database, _ = configs.NewConnection(cfg)

	go func() {
		defer wg.Done()
		createTaskTable(database, configs.TASK_TABLE, wg)
	}()

	go func() {
		defer wg.Done()
		createUserTable(database, configs.USER_TABLE, wg)
	}()

	wg.Wait()

	http.Handle("/swagger/", httpSwagger.WrapHandler)

	TaskRepositoryDB := repository.NewTaskRepository(database)
	ts := api.TaskHandler{Service: services.NewTaskService(&TaskRepositoryDB)}
	http.Handle("/api/task/create", prometheusMiddleware(http.HandlerFunc(ts.Insert)))
	http.Handle("/api/task/delete", prometheusMiddleware(http.HandlerFunc(ts.Delete)))
	http.Handle("/api/task/update", prometheusMiddleware(http.HandlerFunc(ts.Update)))
	http.Handle("/api/task/getAll", prometheusMiddleware(http.HandlerFunc(ts.GetAllTasks)))

	UserRepositoryDB := repository.NewUserRepository(database)
	us := api.UserHandler{Service: services.NewUserService(&UserRepositoryDB)}
	http.Handle("/api/user/register", prometheusMiddleware(http.HandlerFunc(us.Insert)))
	http.Handle("/api/user/delete", prometheusMiddleware(http.HandlerFunc(us.Delete)))
	http.Handle("/api/user/update", prometheusMiddleware(http.HandlerFunc(us.Update)))
	http.Handle("/api/user/getAll", prometheusMiddleware(http.HandlerFunc(us.GetAllUsers)))

	http.Handle("/metrics", promhttp.Handler())

	http.Handle("/", prometheusMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome Junior Project :) \nYou can use /metrics for prometheus and /swagger/index.html for swagger ui \n  "))
	})))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)

	}

	log.Printf("Listening on port %s", port)
	log.Printf("Open http://localhost:%s in the browser", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))

}

func createTaskTable(database *sql.DB, tableName string, wg sync.WaitGroup) (*sql.DB, error) {
	mt.Lock()
	defer mt.Unlock()
	var createTaskTable, err = configs.CreateTaskTable(database, tableName)
	if createTaskTable == nil {
		log.Fatal("Database crate table failed : ", tableName)
		return nil, err
	}
	return createTaskTable, nil
}

func createUserTable(database *sql.DB, tableName string, wg sync.WaitGroup) (*sql.DB, error) {
	mt.Lock()
	defer mt.Unlock()
	var createUserTable, err = configs.CreateUserTable(database, tableName)
	if createUserTable == nil {
		log.Fatal("Database crate table failed : ", tableName)
		return nil, err
	}

	return createUserTable, nil
}
