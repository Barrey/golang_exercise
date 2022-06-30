package main

import "fmt"

func queryDatabase() {
	fmt.Println("SELECT * from users where active=1")
}

func configDatabase() (string, string, string, string) {
	var mysqlHost string = "localhost"
	var mysqlUser string = "root"
	var mysqlPassword string = "root"
	var mysqlDatabase string = "test"

	fmt.Println("Connect database")
	return mysqlHost, mysqlUser, mysqlPassword, mysqlDatabase //return multiple value
}

func closeDatabase() {
	fmt.Println("Close database")
}

func main() {
	defer closeDatabase() //defer always run at the end of main function
	var mysqlHost, mysqlUser, mysqlPassword, mysqlDatabase string
	mysqlHost, mysqlUser, mysqlPassword, mysqlDatabase = configDatabase()

	fmt.Println("MySQL Host:", mysqlHost)
	fmt.Println("MySQL User:", mysqlUser)
	fmt.Println("MySQL Password:", mysqlPassword)
	fmt.Println("MySQL Database:", mysqlDatabase)

	queryDatabase()
}
