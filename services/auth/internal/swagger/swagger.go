package swagger

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"path/filepath"
	"strings"
)

type SwaggerServer struct {
	address    string
	apiBaseURL string
}

func NewSwaggerServer(address string, apiBaseURL string) *SwaggerServer {
	return &SwaggerServer{
		address:    address,
		apiBaseURL: apiBaseURL,
	}
}

func (s *SwaggerServer) Run() error {
	tmpl, err := template.ParseFiles("swagger/index.html")
	if err != nil {
		return fmt.Errorf("error parsing index.html template: %v", err)
	}

	apiURL, err := url.Parse(s.apiBaseURL)
	if err != nil {
		return fmt.Errorf("invalid API base URL: %v", err)
	}
	proxy := httputil.NewSingleHostReverseProxy(apiURL)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" || r.URL.Path == "/swagger" || r.URL.Path == "/swagger/" {
			data := struct {
				APIBaseURL string
			}{
				APIBaseURL: s.apiBaseURL,
			}
			if err := tmpl.Execute(w, data); err != nil {
				log.Printf("Error executing template: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		} else if r.URL.Path == "/swagger/swagger.swagger.json" {
			http.ServeFile(w, r, filepath.Join("swagger", "swagger.swagger.json"))
		} else if strings.HasPrefix(r.URL.Path, "/swagger/") {
			http.ServeFile(w, r, r.URL.Path[1:])
		} else {
			log.Printf("Proxying request to API: %s %s", r.Method, r.URL.Path)
			proxy.ServeHTTP(w, r)
		}
	})

	log.Printf("Swagger UI listening on %s with API base URL %s", s.address, s.apiBaseURL)
	return http.ListenAndServe(s.address, nil)
}
