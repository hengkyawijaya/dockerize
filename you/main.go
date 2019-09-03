package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

var (
	DbHost     = os.Getenv("DB_HOST")
	DbPort     = os.Getenv("DB_PORT")
	DbUser     = os.Getenv("DB_USER")
	DbPassword = os.Getenv("DB_PASSWORD")
	DbName     = os.Getenv("DB_NAME")
	SslMode    = os.Getenv("SSL_MODE")
)

type User struct {
	Name  string `db:"name" json:"name"`
	Email string `db:"email" json:"email"`
}

type Response struct {
	Data   interface{} `json:"data"`
	Status Status      `json:"status"`
}

type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Succeed bool   `json:"succeed"`
}

func main() {
	log.Println("server you started")
	postgresqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		DbHost,
		DbPort,
		DbUser,
		DbPassword,
		DbName,
		SslMode,
	)

	DB, err := sql.Open("postgres", postgresqlInfo)

	if err != nil {
		log.Println("Error ", err)
	}

	http.HandleFunc("/you", func(w http.ResponseWriter, r *http.Request) {

		var user User
		err = DB.QueryRow("SELECT name, email FROM public.users LIMIT 1").Scan(&user.Name, &user.Email)
		if err != nil {
			log.Println("Error ", err)
		}

		respJSON, err := json.Marshal(Response{
			Data: user,
			Status: Status{
				Code:    200,
				Message: "Operation handly correctly",
				Succeed: true,
			},
		})

		if err != nil {
			log.Println("Error ", err)
		}

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		w.Write(respJSON)
	})

	http.ListenAndServe(":8001", nil)
}
