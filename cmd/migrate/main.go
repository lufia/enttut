package main

import (
	"context"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/lufia/enttut/ent"
	"github.com/lufia/enttut/ent/migrate"
)

func main() {
	c, err := ent.Open("postgres", "user=lufia host=localhost port=5432 dbname=enttut sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	ctx := context.Background()
	err = c.Schema.WriteTo(ctx, os.Stdout, migrate.WithDropIndex(true), migrate.WithDropColumn(true))
	if err != nil {
		log.Fatal(err)
	}
}
