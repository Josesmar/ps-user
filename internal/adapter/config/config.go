package config

import (
	"fmt"
	"os"
)

// const (
// 	host     = "ec2-34-233-115-14.compute-1.amazonaws.com"
// 	port     = "5432"
// 	user     = "tozekvgcqqqwbz"
// 	password = "6948c75bb05afd43909011b3fd176e13506b31659adf5268411c460b5baebdbd"
// 	name     = "dd2b5jm5v5gpsg"
// )

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "admin"
	name     = "ps-user"
)

var (
	// StringConectionBanco is string conection with postgres
	StringConectionBanco = ""

	// Port
	Port = 9000

	// Secretkey is key user signer token
	SecretKey []byte

	IP = ""
)

// Load go initializer variables environment
func Load() {

	StringConectionBanco = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		DbConfig.Host, DbConfig.Port, DbConfig.Username, DbConfig.Password, DbConfig.Name)

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
	IP = os.Getenv("API_IP")
}
