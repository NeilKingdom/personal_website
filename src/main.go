package main

import (
	// Core
    "os"
	"fmt"
	"html/template"
	"net/http"
	// External
    "github.com/gorilla/mux"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

/**
 * Takes in a byte stream (e.g., file contents) and converts it to HTML based on Markdown rules.
 * @param stream The byte stream to be converted to HTML
 * @returns A byte stream representing the HTML output
 */
func mdToHTML(stream []byte) []byte {
	// Markdown parser
	mdOpts := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	mdParser := parser.NewWithExtensions(mdOpts)
	mdAST := mdParser.Parse(stream)

	// HTML renderer
	htmlOpts := html.RendererOptions{ Flags: html.CommonFlags | html.HrefTargetBlank }
	htmlRenderer := html.NewRenderer(htmlOpts)

	return markdown.Render(mdAST, htmlRenderer)
}

/**
 * Returns a callback function that will render the page specified by __htmlPage__.
 * @param htmlPage Name of the page to render (.html extension ought to be omitted)
 * @param args Arguments to pass to the page
 * @returns A callback function intended to be utilized by the router
 */
func renderPage(htmlPage string, args any) func(http.ResponseWriter, *http.Request) {
    return func(writer http.ResponseWriter, httpRequest *http.Request) {
        path := fmt.Sprintf("pages/%s.html", htmlPage)
        params := mux.Vars(httpRequest)

        if params["post"] != "" {
            data, err := os.ReadFile(fmt.Sprintf("md/blog/%s.md", params["post"]))
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
    pageServer := http.FileServer(http.Dir("pages"))
	assetsServer := http.FileServer(http.Dir("../assets"))

	// Map path to the appropriate static file server
	router.PathPrefix("/pages/").Handler(http.StripPrefix("/pages/", pageServer))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", assetsServer))

    router.HandleFunc("/", renderPage("index", nil))
    router.HandleFunc("/about", renderPage("about", nil))
    router.HandleFunc("/resume", renderPage("resume", nil))
    router.HandleFunc("/projects", renderPage("projects", nil))
    router.HandleFunc("/blog", renderPage("blog", nil))
    router.HandleFunc("/blog/{post}", renderPage("blog", nil))
    router.HandleFunc("/{anything:.*}", renderPage("index", nil))

    fmt.Print("Starting server at localhost:8080...")
    http.ListenAndServe(":8080", router)
}
