package configs

const TASK_TABLE = "task_table"
const USER_TABLE = "user_table"

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
}

func DefaultConfig() Config {
	return Config{
		Host:     "localhost",
		Port:     "3306",
		User:     "root",
		Password: "12345",
		Dbname:   "task",
	}
}
