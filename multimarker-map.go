package main

import (
    "bufio"
    "fmt"
    "html/template"
    "net/http"
    "os"
)

type Calle struct {
    Nombre string
}

func handler(w http.ResponseWriter, r *http.Request) {
    file, err := os.Open("data.txt")
    if err != nil {
        fmt.Fprintf(w, "Error al leer los datos")
        return
    }
    defer file.Close()

    data := make([]Calle, 1)

    s := bufio.NewScanner(file)
    s.Split(bufio.ScanLines)
    for s.Scan() {
        data = append(data, Calle{s.Text()})
    }

    t := template.New("map")
    if err != nil {
        fmt.Fprintf(w, "Error al cargar los datos")
        return
    }
    t, err = template.ParseFiles("map.html")
    if err != nil {
        fmt.Fprintf(w, "Error al cargar el mapa")
        return
    }
    t.Execute(w, data)
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
