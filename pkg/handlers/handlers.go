package handlers

import (
	"fmt"
	"net/http"

	"github.com/rakibullahsazib/hotel-booking-api/pkg/config"
	"github.com/rakibullahsazib/hotel-booking-api/pkg/models"
	"github.com/rakibullahsazib/hotel-booking-api/pkg/render"
)

type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (repo *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	repo.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	n, err := fmt.Fprintf(w, "Hello World")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Bytes written: %d", n)
}
func (repo *Repository) About(w http.ResponseWriter, r *http.Request) {
	// add template data
	stringMap := map[string]string{
		"test": "Hello",
	}
	remoteIP := repo.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
