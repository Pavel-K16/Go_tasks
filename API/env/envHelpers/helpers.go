package envHelpers

import "os"

func Dsn() string {
	host := os.Getenv("HOST")
	user := os.Getenv("USER")
	pass := os.Getenv("PASS")
	dbname := os.Getenv("DBNAME")
	port := os.Getenv("PORT")
	ssl := os.Getenv("SSL")
	dsn := "host=" + host + " user=" + user + " password=" + pass + " dbname=" + dbname + " port=" + port + " sslmode=" + ssl
	return dsn
}
