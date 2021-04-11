package database

import (
	"log"

	"github.com/hashicorp/go-memdb"
)

//var DbConn *sql.DB
var DbConn *memdb.MemDB

//Setup In Memory Database
func SetupDatabase() {

	//Create a DB schema
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"data": &memdb.TableSchema{
				Name: "data",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
					"value": &memdb.IndexSchema{
						Name:    "value",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Value"},
					},
				},
			},
		},
	}
	err := schema.Validate()
	if err != nil {
		log.Fatalf("memdb.NewMemDB : Error : %s\n", err)
	}

	//Create a new data base
	DbConn, err = memdb.NewMemDB(schema)
	if err != nil {
		log.Fatalf("memdb.NewMemDB : Error : %s\n", err)
	}

}

func CloseDatabase() {

}
