package main

import (
	"fmt"
	"log"
	"net/http"
)

type dollar float64

func (d dollar) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollar

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) readPrice(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]

	if !ok {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "%s\n", price)

}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	_, ok := db[item]

	if !ok {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	delete(db, item)
	fmt.Fprintf(w, "Deleted %s\n", item)

}

func main() {
	db := database{"shoes": 55, "tuxedo": 567}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.readPrice)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
