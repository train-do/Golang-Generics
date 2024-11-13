// package main

// import "fmt"

// // Fungsi dengan variadic parameter
// func printData(data ...any) {
// 	// Iterasi dan cetak data
// 	for _, d := range data {
// 		fmt.Println(d)
// 	}
// }

// func main() {
// 	// Slice dengan data yang dinamis
// 	data := []any{"Alice", "Bob", "Charlie", "John", 2}
// 	data = append(data, struct{ Id int }{Id: 2})
// 	// Memanggil printData dan "menyebarkan" elemen slice ke variadic parameter
// 	printData(data...) // Output: Alice\nBob\nCharlie\nJohn
// }

package main

import (
	"log"
	"net/http"

	"github.com/train-do/Golang-Generics/router"
)

func main() {
	router := router.RouterAPI()

	log.Println("Starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
