package main

import (
    "fmt"
    "html/template"
    "net/http"
)

type Link struct {
    Name string
    Href string
}

type HomeData struct {
    Links []Link
}

func main() {
    // Serve static content
    fs := http.FileServer(http.Dir("static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
        home_data := HomeData{
            Links: []Link{
                { Name: "Home",     Href: "" },
                { Name: "About",    Href: "" },
                { Name: "Projects", Href: "" },
                { Name: "Blog",     Href: "" },
            },
        }

        // Parse the home page
        home, err := template.ParseFiles("pages/index.html")
        if err != nil {
            http.Error(writer, err.Error(), http.StatusInternalServerError)
            return
        }

        // Load the home page
        err = home.Execute(writer, home_data)
        if err != nil {
            http.Error(writer, err.Error(), http.StatusInternalServerError)
            return
        }
    })

    fmt.Print("Starting server at localhost:8080...")
    http.ListenAndServe(":8080", nil)
}
