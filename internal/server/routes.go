package server

import (
	"context"
	"net/http"
	"time"

	"PersonalWebsite2026/cmd/web"

	"github.com/a-h/templ"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", s.homeHandler)
	mux.Handle("GET /projects", templ.Handler(web.ProjectsPage()))
	mux.HandleFunc("GET /blog", web.BlogListHandler)
	mux.HandleFunc("GET /blog/{slug}", web.BlogPostHandler)

	fileServer := http.FileServer(http.FS(web.Files))
	mux.Handle("GET /assets/", fileServer)

	return s.corsMiddleware(mux)
}

func (s *Server) homeHandler(w http.ResponseWriter, r *http.Request) {
	go wakeService("https://sous-chef.koletackney.dev/")
	go wakeService("https://short.koletackney.dev/")
	templ.Handler(web.HomePage()).ServeHTTP(w, r)
}

func wakeService(url string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return
	}
	http.DefaultClient.Do(req) //nolint:errcheck
}

func (s *Server) corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, X-CSRF-Token")
		w.Header().Set("Access-Control-Allow-Credentials", "false")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
