package core

// IDs alineados con docker/postgresql/seeds/seed.sql (tabla rol).
const (
	ADMIN       = 1
	COORDINADOR = 2
	OPERADOR    = 3
	CONDUCTOR   = 4
	CIUDADANO   = 5

	// Alias legacy usado en alerta_usuario.
	SUPERVISOR = OPERADOR
)
