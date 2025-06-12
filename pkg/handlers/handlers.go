package handlers

import (
	"net/http"

	"github.com/ArcticOctopus/bookings/pkg/config"
	"github.com/ArcticOctopus/bookings/pkg/models"
	"github.com/ArcticOctopus/bookings/pkg/render"
)

var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// TemplateData holds data set from handlers to template

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets teh repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{})

}

// Couple is the about couples-suite handler
func (m *Repository) Couples(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "couples.page.tmpl", &models.TemplateData{})

}

// Mountain is the mountain-view page handler
func (m *Repository) Mountain(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "mountain.page.tmpl", &models.TemplateData{})

}

// Book is the booking page handler
func (m *Repository) Book(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "book.page.tmpl", &models.TemplateData{})

}

// Contact is the contact page handler
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "contact.page.tmpl", &models.TemplateData{})

}
