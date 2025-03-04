package main

import (
    "os"
	"fmt"
	"html/template"
	"net/http"

    "github.com/gorilla/mux"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

// mdToHTML
// Takes in a byte stream e.g. file I/O and converts it to HTML based on Markdown rules.
// Parameters:
//   - md ([]byte): The byte stream to be converted to HTML
func mdToHTML(md []byte) []byte {
	// Create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	// Create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer)
}

func renderPage(html string, args any) func(http.ResponseWriter, *http.Request) {
    return func(writer http.ResponseWriter, request *http.Request) {
        path := fmt.Sprintf("pages/%s.html", html)
        vars := mux.Vars(request)
        blogPost := vars["post"]

        if blogPost != "" {
            data, err := os.ReadFile(fmt.Sprintf("static/md/blog/%s.md", blogPost))
            if err != nil {
                http.Error(writer, err.Error(), http.StatusNotFound)
            } else {
                writer.Write(mdToHTML(data))
            }
        } else {
            // Parse HTML file
            pageTemplate, err := template.ParseFiles(path)
            if err != nil {
                http.Error(writer, err.Error(), http.StatusInternalServerError)
                return
            }

            // Render the page
            err = pageTemplate.Execute(writer, args)
            if err != nil {
                http.Error(writer, err.Error(), http.StatusInternalServerError)
                return
            }
        }
    }
}

func main() {
    var router *mux.Router = mux.NewRouter()
    httpHandler := http.FileServer(http.Dir("static"))

    router.Handle("/static/{filename:[a-zA-Z0-9._/-]+}", http.StripPrefix("/static/", httpHandler))
    router.HandleFunc("/", renderPage("index", nil))
    router.HandleFunc("/home", renderPage("home", nil))
    router.HandleFunc("/about", renderPage("about", nil))
    router.HandleFunc("/projects", renderPage("projects", nil))
    router.HandleFunc("/blog", renderPage("blog", nil))
    router.HandleFunc("/blog/{post}", renderPage("blog", nil))
    // Fallback route
    router.HandleFunc("/{anything:.*}", renderPage("index", nil))

    fmt.Print("Starting server at localhost:8080...")
    http.ListenAndServe(":8080", router)
}
