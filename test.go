package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-pg/pg"
)

type Item struct {
	Id    int    `json:"id"`
	Lists []List `json:"lists" pg:",many2many:item_to_items,fk:Item,joinFK:List"`
}

type List struct {
	Id    int    `json:"id"`
	Items []Item `json:"items" pg:",many2many:item_to_items,fk:List,joinFK:Item"`
}

func main() {

	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "postgres",
		Database: "test",
	})

	db.OnQueryProcessed(func(event *pg.QueryProcessedEvent) {
		query, err := event.FormattedQuery()
		if err != nil {
			panic(err)
		}

		log.Printf("%s %s", time.Since(event.StartTime), query)
	})

	defer db.Close()

	qs := []string{
		"CREATE TEMP TABLE items (id int)",
		"CREATE TEMP TABLE lists (id int)",
		"CREATE TEMP TABLE item_to_items (id int, item_id int, list_id int)",
		"INSERT INTO items VALUES (1), (2), (3)",
		"INSERT INTO lists VALUES (1), (2), (3)",
		"INSERT INTO item_to_items VALUES (1, 1, 2), (2, 1, 3)",
	}
	for _, q := range qs {
		_, err := db.Exec(q)
		if err != nil {
			panic(err)
		}
	}

	// Select item and all subitems with following queries:
	//
	// SELECT "item".* FROM "items" AS "item" ORDER BY "item"."id" LIMIT 1
	//
	// SELECT * FROM "items" AS "item"
	// JOIN "item_to_items" ON ("item_to_items"."item_id") IN ((1))
	// WHERE ("item"."id" = "item_to_items"."sub_id")

	var list []List
	err := db.Model(&list).Column("list.*", "Items").Select()
	if err != nil {
		panic(err)
	}
	fmt.Println(list)

}
