package data

import (
	"devotedbapp/database"
	"log"

	"github.com/hashicorp/go-memdb"
)

//Create a sample struct for data
type Data struct {
	ID    string
	Value string
}

func SetData(data *Data, txn *memdb.Txn) {
	//log.Println(data)
	needCommit := false
	if txn == nil {
		txn = database.DbConn.Txn(true)
		needCommit = true
	}

	if err := txn.Insert("data", data); err != nil {
		log.Println(err.Error())
	}
	if needCommit {
		txn.Commit()
	}

}

func GetData(name string, txn *memdb.Txn) string {
	//log.Printf("Entered GetData function with name=%s\n", name)
	if txn == nil {
		txn = database.DbConn.Txn(false)
		defer txn.Abort()
	}

	raw, err := txn.First("data", "id", name)
	if err != nil {
		log.Println(err.Error())
	}

	value := "NULL"
	if raw != nil {
		value = raw.(*Data).Value
	} else {
		//log.Println("else block in 1st")
		it, err := txn.Get("data", "id")
		if err != nil {
			log.Println(err.Error())
		}
		for obj := it.Next(); obj != nil; obj = it.Next() {
			d := obj.(*Data)
			log.Printf("Name =%s, Value = %s\n", d.ID, d.Value)
			if name == d.ID {
				//log.Println("They are equal")
				value = d.Value
				break
			}
		}

	}

	return value

}

func DeleteData(name string, txn *memdb.Txn) error {

	needCommit := false
	if txn == nil {
		txn = database.DbConn.Txn(true)
		needCommit = true
	}

	it, err := txn.Get("data", "id")
	if err == nil {
		for obj := it.Next(); obj != nil; obj = it.Next() {
			d := obj.(*Data)
			name1 := d.ID

			if name == name1 {
				//log.Println("They are equal")
				deleteObject := &Data{ID: d.ID, Value: d.Value}
				err = txn.Delete("data", deleteObject)
				break
			}
		}
	} else {
		log.Println(err.Error())
	}

	if needCommit {
		txn.Commit()
	}

	return err

}

func GetCount(value string, txn *memdb.Txn) int {
	if txn == nil {
		txn = database.DbConn.Txn(false)
		defer txn.Abort()
	}

	count := 0
	it, err := txn.Get("data", "value", value)
	if err != nil {
		log.Printf("get error: %v\n", err)
	}
	if it != nil {
		for obj := it.Next(); obj != nil; obj = it.Next() {
			count++
		}
	}
	return count

}

func EndDatabase() {
	database.CloseDatabase()
}
