package models

import (
	"database/sql"

	aah "aahframework.org/aah.v0"
	"github.com/go-pg/pg"
)

var (
	db *pg.DB
)

// Transaction is to handle Db transactions
type Transaction struct {
	Transaction *pg.Tx
}

func initDb(_ *aah.Event) {
	db = pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "postgres",
		Database: "test",
	})

	// err := createSchema(db)
	// if err != nil {
	// 	panic(err)
	// }
}

func closeDb(_ *aah.Event) {
	if db != nil {
		_ = db.Close()
	}
}

// GetTx returns the DB transaction.
func GetTx() *Transaction {
	tx, err := db.Begin()
	if err != nil {
		panic(err)
		return nil
	}
	return &Transaction{Transaction: tx}
}

// CommitOrRollback commits or rollback if error.
func (t *Transaction) CommitOrRollback() {
	if t.Transaction != nil {
		if err := t.Transaction.Commit(); err != nil && err != sql.ErrTxDone {
			if err = t.Transaction.Rollback(); err != nil && err != sql.ErrTxDone {
				panic(err)
			}
		}
		t.Transaction = nil
	}
}

func createSchema(db *pg.DB) error {
	for _, model := range []interface{}{&Post{}} {
		err := db.CreateTable(model, nil)
		if err != nil {
			return err
		}
	}
	return nil
}

func init() {
	aah.OnStart(initDb)
	aah.OnShutdown(closeDb)
}
