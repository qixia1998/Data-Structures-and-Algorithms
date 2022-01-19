//main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing fmt,database/sql, net/http, text/template package
import (
	//    "fmt"
	"database/sql"
	//  "log"
	//  "net/http"
	//  "text/template"
	//  "errors"
	_ "github.com/go-sql-driver/mysql"
)

type Customer struct {
	CustomerId   int
	CustomerName string
	SSN          string
}
// GetConnection method which returns sql.DB
func GetConnection() (database *sql.DB) {
	databaseDriver := "mysql"
	databaseUser := "root"
	databasePass := "123456"
	databaseName := "crm"
	database, error := sql.Open(databaseDriver, databaseUser + ":" + databasePass + "@/" + databaseName)
	if error != nil {
		panic(error.Error())
	}
	return database	
}
// GetCustomerById with parameter customerId returns Customer
func GetCustomerById(CustomerId int) Customer {
	var database *sql.DB
	database = GetConnection()

	var error error
	var rows *sql.Rows
	rows, error = database.Query("SELECT * FROM Customer WHERE CustomerId=?", CustomerId)
	if error != nil {
		panic(error.Error())
	}
	//fmt.Println(rows)
	var customer Customer
	customer = Customer{}

	for rows.Next() {
		var customerId int
		var customerName string
		var SSN string
		error = rows.Scan(&customerId, &customerName, &SSN)
		if error != nil {
			panic(error.Error())
		}
		customer.CustomerId = customerId
		customer.CustomerName = customerName
		customer.SSN = SSN
	}
	defer database.Close()
	return customer
}
// GetCustomers method returns Customer Array
func GetCustomers() []Customer {
	var database *sql.DB
	database = GetConnection()

	var error error
	var rows *sql.Rows
	rows, error = database.Query("SELECT * FROM Customer ORDER BY CustomerId DESC")
	if error != nil {
		panic(error.Error())
	}
	var customer Customer
	customer = Customer{}

	var customers []Customer
	customers = []Customer{}
	for rows.Next() {
		var customerId int
		var customerName string
		var SSN string
		error = rows.Scan(&customerId, &customerName, &SSN)
		if error != nil {
			panic(error.Error())
		}
		customer.CustomerId = customerId
		customer.CustomerName = customerName
		customer.SSN = SSN
		customers = append(customers, customer)
	}

	defer database.Close()

	return customers
}
// InsertCustomer method with parameter customer
func InsertCustomer(customer Customer) {
	var database *sql.DB
	database = GetConnection()

	var error error
	var insert *sql.Stmt
	insert, error = database.Prepare("INSERT INTO CUSTOMER(CustomerName,SSN) VALUES(?,?)")
	if error != nil {
		panic(error.Error())
	}
	insert.Exec(customer.CustomerName, customer.SSN)
	//log.Println("INSERT: Customer Name: " + customer.name + " | SSN: " + customer.ssn)

	defer database.Close()

	//return Customer{}
}
// Update Customer method with parameter customer
func UpdateCustomer(customer Customer) {
	var database *sql.DB
	database = GetConnection()

	var error error
	var update *sql.Stmt
	update, error = database.Prepare("UPDATE CUSTOMER SET CustomerName=?, SSN=? WHERE CustomerId=?")
	if error != nil {
		panic(error.Error())
	}
	update.Exec(customer.CustomerName, customer.SSN, customer.CustomerId)
	//log.Println("INSERT: Customer Name: " + customer.name + " | SSN: " + customer.ssn)

	defer database.Close()

	//return Customer{}
}
// Delete Customer method with parameter customer
func DeleteCustomer(customer Customer) {
	var database *sql.DB
	database = GetConnection()

	var error error
	var delete *sql.Stmt
	delete, error = database.Prepare("DELETE FROM Customer WHERE CustomerId=?")
	if error != nil {
		panic(error.Error())
	}
	delete.Exec(customer.CustomerId)
	//log.Println("INSERT: Customer Name: " + customer.name + " | SSN: " + customer.ssn)

	defer database.Close()

	//return Customer{}
}

/*func main() {
     var customers []Customer
    customers = GetCustomers()
    fmt.Println(customers)
  //  var customer Customer
//    customer.CustomerName = "Thomas Smith"
  //  customer.SSN = "2323343"
  //  InsertCustomer(customer)
  //var customer Customer
  //  customer.CustomerName = "George Thompson"
  //  customer.SSN = "23233432"
  //  customer.CustomerId = 2
  //  UpdateCustomer(customer)
var customer Customer
  //customer.CustomerName = "George Thompson"
  //customer.SSN = "23233432"
 customer.CustomerId = 2
    deleteCustomer(customer)
    customers = GetCustomers()
    fmt.Println(customers)
}
*/