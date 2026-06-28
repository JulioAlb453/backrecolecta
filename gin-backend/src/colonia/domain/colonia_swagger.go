package domain

// swagger:model CreateColoniaRequest
type CreateColoniaRequest struct {
	Nombre string `json:"nombre" binding:"required" example:"Centro Histórico"`
	Zona   string `json:"zona" binding:"required" example:"Centro"`
}

// swagger:model UpdateColoniaRequest
type UpdateColoniaRequest struct {
	Nombre string `json:"nombre" example:"Centro Histórico"`
	Zona   string `json:"zona" example:"Centro"`
}

// swagger:model ColoniaResponse
type ColoniaResponse struct {
	Success bool    `json:"success"`
	Message string  `json:"message"`
	Data    Colonia `json:"data"`
	Code    int     `json:"code"`
}

// swagger:model ColoniaListResponse
type ColoniaListResponse struct {
	Success bool      `json:"success"`
	Message string    `json:"message"`
	Data    []Colonia `json:"data"`
	Code    int       `json:"code"`
}

// swagger:model ColoniaMessageResponse
type ColoniaMessageResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Code    int    `json:"code,omitempty"`
}
