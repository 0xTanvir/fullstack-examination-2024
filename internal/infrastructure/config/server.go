// Package config provides the config for the application.
package config

// App is the configuration for the application.
type App struct {
	UI            UI
	APIServer     Server
	SwaggerServer Server
	SQLite        SQLite
}

// UI is the configuration for the UI.
type UI struct {
	URL string `validate:"required"`
}

// Server is the configuration for the server.
type Server struct {
	Enable bool
	Port   int
}

// SQLite is the configuration for the SQLite database.
type SQLite struct {
	DBFilename string `validate:"required"`
}