package route

import (
	"net/http"
	"strings"
)

// Convenience function to redirect to the error message page
func Error_message(writer http.ResponseWriter, request *http.Request, msg string) {
	url := []string{"/err?msg=", msg}
	http.Redirect(writer, request, strings.Join(url, ""), 302)
}
