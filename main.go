package main

import (
	"database/sql"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"sync"
	"time"

	"log"
	"math/rand"

	_ "github.com/go-sql-driver/mysql"
)

func generate_random_user_data(db *sql.DB) {
	names := [6]string{"Mike_", "Eric_", "Michael_", "William_", "Jeremy_", "Morpheous_"}
	offsets := []int{12, -15, 11, 124, -124, 458, 495, 127, -30, -202, -4938, -22, -32, 2933}
	future := []int{124, 242, 24142, 4534, 67, 876, 9967, 76, 434, 643, 34634, 3454, 43, 214, 12412, 24}
	enums := []string{"active", "inactive", "deleted"}
	r := rand.Intn(6)
	o := rand.Intn(14)
	f := rand.Intn(16)
	e := rand.Intn(3)
	nameID := rand.Intn(10000000)

	username := names[r] + strconv.Itoa(nameID)
	status := enums[e]
	currentTime := time.Now().UTC().AddDate(0, 0, offsets[o])
	updatedTime := time.Now().UTC().AddDate(0, 0, future[f])
	_, err := db.Exec("INSERT INTO users(username, status, created_on, updated_on) VALUES(?, ?, ?, ?)", username, status, currentTime, updatedTime)
	if err != nil {
		log.Fatal(err)
	}
}

func update_status_for_customers(db *sql.DB) {
	start := time.Now()
	var wg sync.WaitGroup
	var count int
	err := db.QueryRow("SELECT count(*) FROM users").Scan(&count)
	if err != nil {
		panic(err)
	}
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			var status string
			err := db.QueryRow("SELECT status FROM users where id = ?", i+1).Scan(&status)
			if err != nil {
				log.Fatal(err)
			}
			if status == "inactive" {
				db.Exec("UPDATE users set status = 'deleted' where id = ?", i+1)
			}
		}()
		wg.Wait()
	}
	fmt.Println("Finished in: ")
	fmt.Println(time.Now().Sub(start))
}

func main() {
	args := os.Args[1:]
	fmt.Println(reflect.TypeOf(args), args)
	flow := args[0]
	var db *sql.DB

	db, err := sql.Open("mysql", "root:yes@tcp(gomysql)/gotest1")
	if err != nil {
		log.Fatal(err)
	}
	if flow == "update" {
		update_status_for_customers(db)
	} else {
		for i := 0; i < 75000; i++ {
			generate_random_user_data(db)
		}
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err.Error)
	}
}
