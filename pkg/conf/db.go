package conf

type db struct {
	Username string
	Password string
	Host     string
	Port     int
	DbName   string
	Timeout  string
}

var DB = db{
	Username: "root",
	Password: "1",
	Host:     "127.0.0.1",
	Port:     3306,
	DbName:   "Dousheng",
	Timeout:  "10s",
}
