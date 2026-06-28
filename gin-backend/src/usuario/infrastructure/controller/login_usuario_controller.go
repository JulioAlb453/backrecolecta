package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/core"
	"github.com/vicpoo/API_recolecta/src/usuario/application"
)

type LoginUsersController struct {
	uc *application.LoginUser
}

func NewLoginUsersController(uc *application.LoginUser) *LoginUsersController {
	return &LoginUsersController{uc: uc}
}

type loginRequest struct {
	EmailOrAlias string `json:"email_or_alias"`
	Email        string `json:"email"`
	Password     string `json:"password" binding:"required"`
}

type loginUsuarioResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Token   string      `json:"token"`
	Data    interface{} `json:"data"`
	Code    int         `json:"code"`
}

// LoginUsuario godoc
// @Summary      Login de usuario/empleado
// @Tags         Usuario
// @Accept       json
// @Produce      json
// @Param        body body main.LoginUsuarioRequest true "Credenciales"
// @Success      200 {object} map[string]interface{}
// @Failure      401 {object} map[string]interface{}
// @Router       /api/usuarios/login [post]
func (c *LoginUsersController) Handle(ctx *gin.Context) {
	var body loginRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		core.RespondBadRequest(ctx, "json inválido", map[string]string{"detail": err.Error()})
		return
	}

	credential := strings.TrimSpace(body.EmailOrAlias)
	if credential == "" {
		credential = strings.TrimSpace(body.Email)
	}
	if credential == "" {
		core.RespondBadRequest(ctx, "email_or_alias es requerido", nil)
		return
	}

	usuario, valid, err := c.uc.Execute(ctx, application.LoginInput{
		EmailOrAlias: credential,
		Password:     body.Password,
	})
	if err != nil {
		core.RespondInternalServerError(ctx, "error al iniciar sesión", err)
		return
	}
	if !valid || usuario == nil {
		core.RespondError(ctx, http.StatusUnauthorized, core.ErrCodeUnauthorized, "credenciales inválidas", nil)
		return
	}

	token, err := core.GenerateToken(usuario.ID, usuario.RolID)
	if err != nil {
		core.RespondInternalServerError(ctx, "error al generar token", err)
		return
	}

	core.RespondOK(ctx, loginUsuarioResponse{
		Success: true,
		Message: "login correcto",
		Token:   token,
		Data:    usuario,
		Code:    http.StatusOK,
	})
}
