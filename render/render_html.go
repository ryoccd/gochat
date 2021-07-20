package render

import (
	"fmt"
	"html/template"
	"net/http"
)

func ParseTemplateFile(filenames ...string) (t *template.Template) {
	var files []string
	t = template.New("layout")

	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	t = template.Must(t.ParseFiles(files...))
	return
}

func RenderHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string

	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(writer, "layout", data)
}
