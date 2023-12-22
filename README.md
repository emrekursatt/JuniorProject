# JuniorProject

JuniorProject is a web API written in Go. This project provides RESTful APIs for user and task management. It includes features such as metric collection with Prometheus, API documentation with Swagger, and security against SQL injections.

## Design Pattern


The project uses the Repository Pattern design pattern.  

The project primarily includes the following packages:

- `configs`: This package contains the database connection settings and table names.

- `models`: This package defines the data models used in the project.

- `repository`: This package defines the `UserRepository` and `TaskRepository` interfaces that perform database operations. It also includes the `UserRepositoryDB` and `TaskRepositoryDB` structs that implement these interfaces.

- `services`: This package contains the business logic. It defines the `UserService` and `TaskService` interfaces, which perform operations using the methods defined in the `repository` package interfaces.

- `api`: This package handles HTTP requests and uses the services defined in the `services` package.



## Installation

1. Clone the project:
    ```
    gh repo clone emrekursatt/JuniorProject
    ```
2. Run the project:
    ```
    go run main.go
    ```

## Loading Dependencies
    ```
    go mod tidy
    ```

Before running these commands, make sure that Go is installed on your system. If Go is not installed, you can download and install it from Go's official website.

## API Usage
This project provides the following APIs for user and task management:  
Create user: POST /api/user/register  
Delete user: DELETE /api/user/delete  
Update user: PUT /api/user/update  
Get all users: GET /api/user/getAll  
User login: POST /api/user/login  
Create task: POST /api/task/create  
Delete task: DELETE /api/task/delete  
Update task: PUT /api/task/update  
Get all tasks: GET /api/task/getAll


## Metric Collection with Prometheus
This project collects metrics of HTTP requests with Prometheus. The metrics are collected under the name http_requests_total and each one is labeled with the request path. You can access the metrics at http://localhost:8080/metrics.  
## API Documentation with Swagger
This project uses Swagger to document its APIs. You can access the API documents at http://localhost:8080/swagger/index.html.  
## Security
This project provides protection against SQL injections. During operations such as user login and task creation, incoming data is validated and cleaned against SQL injections.  
## Error Management and Validation
This project uses its own error management system to manage errors and validate incoming data. When an error occurs, the error information is logged and an appropriate error message is returned to the user.  
## Tests
The tests for this project can be found in the tests folder.

## Concurrent Processing
This project uses concurrent processing to create database tables. The createTaskTable and createUserTable functions are run concurrently using sync.WaitGroup.  
## Logging
This project uses Go's built-in log package for logging. Logs are written to standard output, and in case of errors, the error information is logged.

