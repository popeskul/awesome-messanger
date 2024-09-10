package swagger

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"path/filepath"
	"strings"

	"github.com/popeskul/awesome-messanger/services/message/internal/config"
)

type Server struct {
	address    string
	apiBaseURL string
	server     *http.Server
}

type SwaggerAddress string
type APIBaseURL string

func ProvideSwaggerAddress(cfg *config.Config) SwaggerAddress {
	return SwaggerAddress(cfg.Server.SwaggerAddress)
}

func ProvideAPIBaseURL(cfg *config.Config) APIBaseURL {
	return APIBaseURL(fmt.Sprintf("http://%s", cfg.Server.GatewayAddress))
}

func NewSwaggerServer(address SwaggerAddress, apiBaseURL APIBaseURL) *Server {
	log.Printf("NewSwaggerServer: address=%s, apiBaseURL=%s", address, apiBaseURL)
	return &Server{
		address:    string(address),
		apiBaseURL: string(apiBaseURL),
	}
}

func (s *Server) Run() error {
	tmpl, err := template.ParseFiles("swagger/index.html")
	if err != nil {
		return fmt.Errorf("error parsing index.html template: %v", err)
	}

	apiURL, err := url.Parse(s.apiBaseURL)
	if err != nil {
		return fmt.Errorf("invalid API base URL: %v", err)
	}
	proxy := httputil.NewSingleHostReverseProxy(apiURL)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
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

	s.server = &http.Server{
		Addr:    s.address,
		Handler: mux,
	}

	log.Printf("Swagger UI listening on %s with API base URL %s", s.address, s.apiBaseURL)
	return s.server.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	if s.server != nil {
		log.Println("Shutting down Swagger server...")
		return s.server.Shutdown(ctx)
	}
	return nil
}
