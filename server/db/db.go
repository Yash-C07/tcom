package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

var Conn *pgx.Conn

func Connect() {

	conn, err := pgx.Connect(context.Background(), "postgres://postgres:postgres@localhost:5432/tcom-db")
	if err != nil {
		fmt.Println(err)
		return
	}
	Conn = conn

}
func GetConn() *pgx.Conn {
	return Conn

}
func Disconnect() {
	defer Conn.Close(context.Background())
}
