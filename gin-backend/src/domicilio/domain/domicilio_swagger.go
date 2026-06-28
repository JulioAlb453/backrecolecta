package domain

// CreateDomicilioRequest cuerpo para POST /api/domicilios y /domicilios.
type CreateDomicilioRequest struct {
	UsuarioID  int    `json:"usuario_id" example:"100"`
	ColoniaID  int    `json:"colonia_id" example:"1"`
	Alias      string `json:"alias" example:"Casa"`
	Direccion  string `json:"direccion" example:"Av. Segunda Sur Ote. 448, Centro, Suchiapa"`
	Calle      string `json:"calle" example:"Av. Segunda Sur Ote."`
	Numero     string `json:"numero" example:"448"`
	Referencia string `json:"referencia" example:"Frente al parque"`
}
