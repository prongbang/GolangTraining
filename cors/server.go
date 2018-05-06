package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type Person struct {
    Name string `json:"name"`
}

type People []Person

func GetPeople(w http.ResponseWriter, r *http.Request) {

    if origin := r.Header.Get("Origin"); origin != "" {
        // w.Header().Set("Access-Control-Allow-Origin", origin)
        fmt.Println(origin)
    }
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")

    people := People{
        Person{Name: "John"},
        Person{Name: "Mike"},
        Person{Name: "Peter"},
    }

    if err := json.NewEncoder(w).Encode(people); err != nil {
        panic(err)
    }
}

func main() {
    http.HandleFunc("/api/register", GetPeople)
    http.ListenAndServe(":8080", nil)
}
