// package playground

// import (
// 	"database/sql"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"os"

// 	// _ "github.com/go-sql-driver/mysql"
// 	"github.com/joncarr/trebuchet/app/configuration"
// 	_ "github.com/lib/pq"
// )

// type animal struct {
// 	id         int
// 	animalType string
// 	nickname   string
// 	zone       int
// 	age        int
// }

// func main() {

// 	//*****************************************************************************************************************************************************
// 	//* Open file holding configuration settings and assign to a variable
// 	file, err := os.Open("configuration/config.json")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()

// 	config := new(configuration.Configuration)
// 	json.NewDecoder(file).Decode(config)

// 	//*****************************************************************************************************************************************************
// 	//* Construct the DB connection string 'connString'

// 	var connString string

// 	switch config.DBType {
// 	case "mysql":
// 		connString = fmt.Sprintf("%s:%s@/%s", config.Dev.DBUserName, config.Dev.DBPassword, config.Dev.DBName)
// 	case "postgres":
// 		connString = fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", config.Dev.DBUserName, config.Dev.DBPassword, config.Dev.DBName)
// 	case "sqlite":
// 		connString = fmt.Sprintf("%s.db", config.AppName)

// 	}

// 	//*****************************************************************************************************************************************************
// 	//* Make the connection to the DB
// 	db, err := sql.Open(config.DBType, connString)
// 	if err != nil {
// 		fmt.Println("DB Connection Error: Connection was not made to the database")
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	//*****************************************************************************************************************************************************
// 	//* TEST QUERIES
// 	//*****************************************************************************************************************************************************
// 	query := fmt.Sprintf("SELECT * FROM animals WHERE %s > %d", "age", 10)
// 	fmt.Println(query)
// 	rows, err := db.Query(query)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer rows.Close()

// 	animals := []animal{}

// 	for rows.Next() {
// 		a := animal{}
// 		err := rows.Scan(&a.id, &a.animalType, &a.nickname, &a.zone, &a.age)
// 		if err != nil {
// 			log.Println(err)
// 			continue
// 		}
// 		animals = append(animals, a)
// 	}

// 	if rows.Err(); err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println(animals)

// 	//***************************************** Execute single query with args*************************************************************
// 	query = fmt.Sprintf("SELECT * FROM animals WHERE age > %d", 10)
// 	fmt.Println(query)
// 	row := db.QueryRow(query)
// 	al := animal{}
// 	err = row.Scan(&al.id, &al.animalType, &al.nickname, &al.zone, &al.age)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(al)

// 	//********************************************** Execute DB record insert****************************************************************
// 	// anml := animal{
// 	// 	animalType: "Steggosaurus",
// 	// 	nickname:   "Dozer",
// 	// 	zone:       3,
// 	// 	age:        60,
// 	// }

// 	// query = fmt.Sprintf("INSERT INTO animals (animal_type, nickname, zone, age) VALUES ('%s','%s',%d,%d)", anml.animalType, anml.nickname, anml.zone, anml.age)
// 	// fmt.Println(query)

// 	// result, err := db.Exec(query)
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// lastID, _ := result.LastInsertId()
// 	// affRows, _ := result.RowsAffected()
// 	// fmt.Printf("ID of last inserted record: %d\n", lastID)
// 	// fmt.Printf("Number of affected rows: %d\n", affRows)

// 	//********************************************** Execute DB record update****************************************************************
// 	// query = fmt.Sprintf("UPDATE animals SET age = %d where id = %d", 33, 6)
// 	// result, err := db.Exec(query)
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// id, _ := result.LastInsertId()
// 	// aff, _ := result.RowsAffected()
// 	// fmt.Printf("Last Insert ID: %d\n", id)
// 	// fmt.Printf("Rows Affected: %d\n", aff)

// 	//**********************************************Postgres Only, Get Back ID****************************************************************
// 	// var rtnID int
// 	// query = fmt.Sprintf("UPDATE animals SET age = %d WHERE id = %d returning id", 33, 7)
// 	// db.QueryRow(query).Scan(&rtnID)
// 	// fmt.Printf("Returned ID after updating record: %d\n", rtnID)

// 	//********************************************** Execute DB prepared STMT*****************************************************************
// 	fmt.Println("Statements...")
// 	fmt.Println("")
// 	stmt, err := db.Prepare("SELECT * FROM animals WHERE age > $1")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer stmt.Close()

// 	// The prepared statement allows us to "kind of" cache a query to be run.
// 	// The query is the query used inside the Prepare() method above which returns a pointer to STMT struct.
// 	// When we run the 'stmt' we defined, all we need to pass it is any of the variables which are listed
// 	// inside the original query.  In this particular example, all we need to pass is an int to represent the
// 	// age we are querying against.
// 	rows, err = stmt.Query(20)
// 	results := handleRows(rows, err)
// 	fmt.Println("Executed via Prepared Statment:", results)

// 	fmt.Println("")
// 	fmt.Println("A COMPLETELY DIFFERENT PREPARED STATEMENT BELOW HERE")
// 	fmt.Println("*****************************************************")
// 	fmt.Println("")

// 	rows, err = stmt.Query(30)
// 	results = handleRows(rows, err)
// 	fmt.Println("Executed via Prepared Statment:", results)

// 	fmt.Println("")
// 	fmt.Println("A COMPLETELY DIFFERENT PREPARED STATEMENT BELOW HERE")
// 	fmt.Println("*****************************************************")
// 	fmt.Println("")

// 	testTransaction(db)

// }

// func handleRows(rows *sql.Rows, err error) []animal {
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	animals := []animal{}

// 	for rows.Next() {
// 		a := animal{}
// 		if err := rows.Scan(&a.id, &a.animalType, &a.nickname, &a.zone, &a.age); err != nil {
// 			log.Fatal(err)
// 		}
// 		animals = append(animals, a)
// 	}
// 	if err := rows.Err(); err != nil {
// 		fmt.Println(err)
// 	}
// 	return animals
// }

// func testTransaction(db *sql.DB) {
// 	fmt.Println("")
// 	fmt.Println("Transactions...")
// 	fmt.Println("")

// 	tx, err := db.Begin()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer tx.Rollback()

// 	// Probably bad examples here since we are just grabbing data.
// 	// Would probably be better to demonstrate this when inserting
// 	// or updating a record.

// 	stmt, err := tx.Prepare("SELECT * FROM animals WHERE age > $1")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer stmt.Close()

// 	rows, err := stmt.Query(15)
// 	results := handleRows(rows, err)
// 	fmt.Println(results)

// 	rows, err = stmt.Query(30)
// 	results2 := handleRows(rows, err)
// 	fmt.Println(results2)

// 	err = tx.Commit()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// }
