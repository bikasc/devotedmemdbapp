package data

import (
	"devotedbapp/data"
	"devotedbapp/database"
	"testing"
)

func TestDataSet(t *testing.T) {

	database.SetupDatabase()
	obj := &Data{ID: "a", Value: "foo"}
	data.SetData(obj, nil)

	txn := database.DbConn.Txn(false)
	defer txn.Abort()

	raw, err := txn.First("data", "id", obj.ID)
	if err != nil {
		t.Errorf("err: %v", err)
	}
	if raw != obj {
		t.Errorf("bad: %#v %#v", raw, obj)
	}

}

func TestDataGet(t *testing.T) {
	database.SetupDatabase()
	obj := &Data{ID: "b", Value: "bar"}

	txn := database.DbConn.Txn(true)
	err := txn.Insert("data", obj)
	if err != nil {
		txn.Commit()
		txn := database.DbConn.Txn(false)
		defer txn.Abort()
		value := data.GetData(obj.ID, txn)
		if value != obj.Value {
			t.Errorf("The original value %s does not match with retrieved value %s", obj.Value, value)
		}
	}
}

func TestDataCount(t *testing.T) {

	database.SetupDatabase()
	obj1 := &Data{ID: "c", Value: "baz"}
	obj2 := &Data{ID: "d", Value: "baz"}
	data.SetData(obj1, nil)
	data.SetData(obj2, nil)

	count := data.GetCount(obj1.Value)

	if count != 2 {
		t.Errorf("Expected 2, but actual received %d", count)
	}

}

func TestDataDelete(t *testing.T) {
	database.SetupDatabase()
	obj := &Data{ID: "ef", Value: "ghbaz"}
	data.SetData(obj, nil)

	data.DeleteData(obj.ID, nil)

	count := data.GetCount(obj.Value)

	if count != 0 {
		t.Errorf("Expected 0, but actual received %d", count)
	}
}
