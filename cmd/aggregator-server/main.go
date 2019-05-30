package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/jinzhu/gorm"

	"github.com/codecure/news-aggregator/database"
)

var db *gorm.DB

const port string = "8000"

var templates = template.Must(template.New("index").Parse(`
<html>
	<body>
		<form method="GET">
			<input type="search" name="s" placeholder="Type part of title to search" value="{{.SearchQuery}}">
			<input type="submit" value="Search">
			<input type="button" onclick="location.href='/';" value="Clear">
		</form>

		{{range .NewsItems}}
			<p><a href="{{.Link}}" target="_blank">{{.Title}}</a></p>
		{{end}}
	</body>
</html>
`))

type pageData struct {
	NewsItems   []database.NewsItem
	SearchQuery string
}

func renderTemplate(w http.ResponseWriter, tmpl string, item *database.NewsItem) {
	err := templates.ExecuteTemplate(w, tmpl, item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		query := ""
		if queryArr, ok := r.Form["s"]; ok {
			query = queryArr[0]
		}

		allNewsItems := []database.NewsItem{}
		db.Where("title LIKE ?", fmt.Sprintf("%%%s%%", query)).Find(&allNewsItems)

		err := templates.ExecuteTemplate(w, "index", pageData{NewsItems: allNewsItems, SearchQuery: query})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	dbDirPtr := flag.String("db", "", "Directory to store DB file")
	flag.Parse()

	db = database.GetDB(*dbDirPtr)
	defer db.Close()

	db.AutoMigrate(&database.NewsItem{})

	fmt.Println("Server started at: http://127.0.0.1:" + port)
	fmt.Println("Press CTRL+C to interrupt")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
