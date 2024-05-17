package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"os"
)

func main() {
	conn, err := pgx.Connect(context.Background(), "postgres://Anno:Anno@localhost:5432/recordings")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var title string
	var price float32

	artist := "John Coltrane"
	err = conn.QueryRow(context.Background(), "select title, price from album where artist = $1", artist).Scan(&title, &price)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("The title is: %s and the artist is %s", title, artist)
}
