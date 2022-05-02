package config

import (
	"html/template"

	"github.com/alexedwards/scs/v2"
)

// AppConfig holds the application config
type AppConfig struct {
	InProduction  bool
	UseCache      bool
	TemplateCache map[string]*template.Template
	Session       *scs.SessionManager
}
