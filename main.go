package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

// main func
func main() {
	connStr := "postgres://bfajpwvj:Drqhw3J301xOmN-ULwUbfs2bUNvnRTdF@isilo.db.elephantsql.com/bfajpwvj?sslmode=verify-full"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	conn := &Conn{db: db}

	r := gin.Default()
	r.Use(cors.Default())

	r.Static("/css", "./client/dist/css")
	r.Static("/js", "./client/dist/js")

	r.NoRoute(func(c *gin.Context) {
		c.File("./client/dist/index.html")
	})

	r.GET("/api", conn.Index)
	r.POST("/api", conn.Store)
	r.PUT("/api", conn.Edit)
	r.DELETE("/api/:id", conn.Delete)
	r.POST("/api/find", conn.Find)
	r.POST("/api/search", conn.Search)

	port := os.Getenv("PORT")
	if port == "" {
		r.Run(":4000")
	} else {
		r.Run(":" + port)
	}
}

// Conn struct
type Conn struct {
	db *sql.DB
}

// Menu struct
type Menu struct {
	MenuID    string
	MenuName  string
	MenuType  int
	MenuPrice NullInt64
}

// NullInt64 is an alias for sql.NullInt64 data type
type NullInt64 struct {
	sql.NullInt64
}

// MarshalJSON for NullInt64
func (ni *NullInt64) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ni.Int64)
}

// NewNullString func
func NewNullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}
