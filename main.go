package main

import (
	"fmt"
	"invoicer/main/config"
	"invoicer/main/controller"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func init() {
	_ = godotenv.Load(".env")

	err := config.InitializeConnectionToMySQLDatabase()

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Connection to SQL Database Successful")
}

func main() {
	http.HandleFunc("/getPhoneCodeList", controller.GetPhoneCodeList)

	http.ListenAndServe(":8080", nil)
}
