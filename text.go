package resputil

import (
	"fmt"
	"io"
	"net/http"
)

func Text(w http.ResponseWriter, format string, a ...interface{}) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	if len(a) > 0 {
		fmt.Fprintf(w, format, a...)
	} else {
		io.WriteString(w, format)
	}
	w.WriteHeader(200)
}
