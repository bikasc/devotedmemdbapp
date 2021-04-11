package cli

import (
	"bufio"
	"devotedbapp/data"
	"devotedbapp/database"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/hashicorp/go-memdb"
)

// used to hold our transactions
var transactions []*memdb.Txn = []*memdb.Txn{}

func processCommands(command, name, value string) {

	switch command {
	case "BEGIN_TRANSACTION":
		var txn *memdb.Txn
		if len(transactions) > 0 {
			//log.Println("BEGIN has been called in second time")
			txn = transactions[len(transactions)-1]
			snapshot := txn.Snapshot()
			//SWAP the transaction point
			transactions[len(transactions)-1] = snapshot
			transactions = append(transactions, txn) //Needs to be the end
		} else {
			txn = database.DbConn.Txn(true)
			transactions = append(transactions, txn)
		}

	//	log.Printf("Current Transaction no %d\n", len(transactions))

	case "ROLLBACK_TRANSACTION":
		if len(transactions) > 0 {
			txn := transactions[len(transactions)-1]
			txn.Abort()
			transactions = transactions[:len(transactions)-1]
		}
	case "COMMIT_TRANSACTION":
		if len(transactions) > 0 {
			txn := transactions[len(transactions)-1]
			txn.Commit()
			transactions = transactions[:len(transactions)-1]
		}

	case "GET":
		var txn *memdb.Txn
		if len(transactions) > 0 {
			txn = transactions[len(transactions)-1]
		}
		value := data.GetData(name, txn)
		fmt.Println(value)
	case "SET":
		var txn *memdb.Txn
		if len(transactions) > 0 {
			txn = transactions[len(transactions)-1]
		}
		sampledata := &data.Data{ID: name, Value: value}
		data.SetData(sampledata, txn)

	case "COUNT":
		var txn *memdb.Txn
		if len(transactions) > 0 {
			txn = transactions[len(transactions)-1]
		}
		value1 := name
		count := data.GetCount(value1, txn)
		fmt.Println("Count= ", count)
	case "DELETE":
		var txn *memdb.Txn
		if len(transactions) > 0 {
			txn = transactions[len(transactions)-1]
		}
		_ = data.DeleteData(name, txn)
		//fmt.Println("Delete Returned ", err)
	case "ENDCONNECTION":
		//fmt.Printf("Entered Command %s", command)
		data.EndDatabase()

	default:
		fmt.Printf("Unsupported Command %s", command)
		//data.EndDatabase()
	}
}
func processCLICommands() {
	argLength := len(os.Args[1:])
	//fmt.Println("Arg length is %d\n", argLength)

	if argLength < 2 {
		return
	}
	var method string
	var name string
	var value string
	for i, a := range os.Args[1:] {
		if i == 0 {
			method = a
		} else if i == 1 {
			name = a
		} else if i == 2 {
			value = a
		}
		fmt.Printf("Arg %d is %s\n", i+1, a)
	}
	//fmt.Printf("method = %s, name = %s, value=%s\n", method, name, value)
	processCommands(method, name, value)
}

func ProcessInterective() {
	argLength := len(os.Args[1:])
	//fmt.Println("Arg length is %d\n", argLength)

	if argLength > 1 {
		processCLICommands()
		return
	}
	fmt.Println("Welcome to this console app")

	isItExit := false

	for {
		fmt.Println("Please select your following options")
		fmt.Println("END for Close database and Exit, or enter command as SET [name] [value]")
		fmt.Print("Enter your Options: ")

		reader := bufio.NewReader(os.Stdin)
		stroption, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Did not read from input", err)
		}
		split := strings.Split(stroption, " ")
		//fmt.Println(split)
		var command, name, value string
		if len(split) > 0 {
			for i, v := range split {
				if i == 0 {
					command = v
				} else if i == 1 {
					name = v
				} else if i == 2 {
					value = v
				}
			}
		}
		if strings.Contains(stroption, "END") {
			isItExit = true
			command = "ENDCONNECTION"
		}
		if strings.Contains(stroption, "BEGIN") {
			command = "BEGIN_TRANSACTION"
		}
		if strings.Contains(stroption, "ROLLBACK") {
			command = "ROLLBACK_TRANSACTION"
		}
		if strings.Contains(stroption, "COMMIT") {
			command = "COMMIT_TRANSACTION"
		}

		processCommands(command, name, value)
		if isItExit {
			break
		}
	}
}
