package zako

import "net/http"

type Zako struct {
	mux *http.ServeMux
}

func New() *Zako {
	return &Zako{
		mux: http.NewServeMux(),
	}
}

func (z *Zako) Get(path string, handler http.HandlerFunc) {
	z.mux.HandleFunc("GET "+path, handler)
}

func (z *Zako) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	z.mux.ServeHTTP(w, r)
}
