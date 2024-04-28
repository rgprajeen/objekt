package http

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func contentTypeMiddleware(n httprouter.Handle, allowList []string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		contentType := r.Header.Get(HeaderContentType)
		for _, value := range allowList {
			if contentType == value {
				n(w, r, p)
				return
			}
		}
		w.WriteHeader(http.StatusUnsupportedMediaType)
	}
}
