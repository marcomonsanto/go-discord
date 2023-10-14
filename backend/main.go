package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/marcomonsanto/go-discord/db"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/guregu/null.v4"
)

type User struct {
	ID          int32       `json:"id"`
	DisplayName string      `json:"display_name"`
	AboutMe     null.String `json:"about_me"`
	Avatar      null.String `json:"avatar"`
}

func main() {
	env := os.Getenv("APP_ENV")

	fmt.Printf("what %s", env)

	// Determine the corresponding environment file based on the environment.
	envFile := ".env." + env

	// Load the environment variables from the specified file.
	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	dbString := os.Getenv("DB_STRING")

	ctx := context.Background()

	database, err := sql.Open("mysql", dbString)
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	queries := db.New(database)

	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		dbUsers, err := queries.ListUsers(ctx)
		var users = make([]User, 0)

		for _, s := range dbUsers {
			users = append(users, User{
				ID:          s.ID,
				DisplayName: s.DisplayName,
				AboutMe:     null.NewString(s.AboutMe.String, s.AboutMe.Valid),
				Avatar:      null.NewString(s.Avatar.String, s.Avatar.Valid),
			})
		}

		log.Println(users)
		if err != nil {
			json.NewEncoder(w).Encode(err)
			return
		}
		json.NewEncoder(w).Encode(users)
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	fmt.Println("Litening on port 8080")
	http.ListenAndServe(":8080", nil)
}
