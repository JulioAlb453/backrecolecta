package http

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/core"
	"github.com/vicpoo/API_recolecta/src/domicilio/application"
	"github.com/vicpoo/API_recolecta/src/domicilio/domain"
)

type DomicilioController struct {
	create *application.CreateDomicilio
	get    *application.GetDomicilio
	list   *application.ListDomicilios
	update *application.UpdateDomicilio
	delete *application.DeleteDomicilio
}

type createDomicilioRequest = domain.CreateDomicilioRequest

func NewDomicilioController(
	create *application.CreateDomicilio,
	get *application.GetDomicilio,
	list *application.ListDomicilios,
	update *application.UpdateDomicilio,
	delete *application.DeleteDomicilio,
) *DomicilioController {
	return &DomicilioController{create, get, list, update, delete}
}

func (c *DomicilioController) RegisterRoutes(r *gin.Engine) {
	registerDomicilioRoutes(r, c, "/api/domicilios")
	registerDomicilioRoutes(r, c, "/domicilios") // compatibilidad
}

func registerDomicilioRoutes(r *gin.Engine, c *DomicilioController, basePath string) {
	protected := r.Group(basePath, core.JWTAuthMiddleware())
	{
		protected.POST("", c.Create)
		protected.GET("", c.List)
		protected.GET("/:id", c.GetByID)
		protected.PUT("/:id", c.Update)
		protected.DELETE("/:id", c.Delete)
	}
}

func buildDireccion(req createDomicilioRequest) string {
	if strings.TrimSpace(req.Direccion) != "" {
		return strings.TrimSpace(req.Direccion)
	}

	parts := make([]string, 0, 3)
	if strings.TrimSpace(req.Calle) != "" {
		parts = append(parts, strings.TrimSpace(req.Calle))
	}
	if strings.TrimSpace(req.Numero) != "" {
		parts = append(parts, strings.TrimSpace(req.Numero))
	}
	if strings.TrimSpace(req.Referencia) != "" {
		parts = append(parts, strings.TrimSpace(req.Referencia))
	}
	return strings.Join(parts, " ")
}

// @Summary      Crear domicilio
// @Description  Requiere colonia_id, alias y direccion (o calle). usuario_id es opcional si hay JWT.
// @Tags         Domicilio
// @Accept       json
// @Produce      json
// @Param        body body domain.CreateDomicilioRequest true "Datos del domicilio"
// @Success      201 {object} map[string]interface{}
// @Failure      400 {object} core.ErrorResponse
// @Failure      401 {object} core.ErrorResponse
// @Failure      500 {object} core.ErrorResponse
// @Security     BearerAuth
// @Router       /api/domicilios [post]
func (c *DomicilioController) Create(ctx *gin.Context) {
	var req createDomicilioRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		core.RespondValidationError(ctx, "Datos de domicilio inválidos", map[string]string{"error": err.Error()})
		return
	}

	usuarioID := req.UsuarioID
	if usuarioID == 0 {
		usuarioID = ctx.GetInt("user_id")
	}

	direccion := buildDireccion(req)
	if req.ColoniaID == 0 || strings.TrimSpace(req.Alias) == "" || direccion == "" {
		core.RespondBadRequest(ctx, "colonia_id, alias y direccion (o calle) son requeridos", nil)
		return
	}

	body := domain.Domicilio{
		UsuarioID: usuarioID,
		Alias:     strings.TrimSpace(req.Alias),
		Direccion: direccion,
		ColoniaID: req.ColoniaID,
	}

	if err := c.create.Execute(&body); err != nil {
		core.RespondInternalServerError(ctx, "Error al crear domicilio", err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "domicilio creado correctamente",
		"code":    http.StatusCreated,
	})
}

// @Summary      Listar domicilios del usuario autenticado
// @Tags         Domicilio
// @Produce      json
// @Success      200 {object} map[string]interface{}
// @Failure      401 {object} core.ErrorResponse
// @Failure      500 {object} core.ErrorResponse
// @Security     BearerAuth
// @Router       /api/domicilios [get]
func (c *DomicilioController) List(ctx *gin.Context) {
	usuarioID := ctx.GetInt("user_id")
	domicilios, err := c.list.ExecuteByUsuarioID(usuarioID)
	if err != nil {
		core.RespondInternalServerError(ctx, "Error al listar domicilios", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "domicilios listados correctamente",
		"data":    domicilios,
		"code":    http.StatusOK,
	})
}

// @Summary      Obtener domicilio por ID
// @Tags         Domicilio
// @Produce      json
// @Param        id path int true "ID del domicilio"
// @Success      200 {object} map[string]interface{}
// @Failure      404 {object} core.ErrorResponse
// @Security     BearerAuth
// @Router       /api/domicilios/{id} [get]
func (c *DomicilioController) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		core.RespondInvalidInput(ctx, "ID de domicilio inválido")
		return
	}

	domicilio, err := c.get.Execute(id)
	if err != nil {
		core.RespondError(ctx, http.StatusNotFound, core.ErrCodeNotFound, "Domicilio no encontrado", nil)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "domicilio obtenido correctamente",
		"data":    domicilio,
		"code":    http.StatusOK,
	})
}

// @Summary      Actualizar domicilio
// @Tags         Domicilio
// @Accept       json
// @Produce      json
// @Param        id path int true "ID del domicilio"
// @Param        body body domain.CreateDomicilioRequest true "Datos del domicilio"
// @Success      200 {object} map[string]interface{}
// @Security     BearerAuth
// @Router       /api/domicilios/{id} [put]
func (c *DomicilioController) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		core.RespondInvalidInput(ctx, "ID de domicilio inválido")
		return
	}

	var req createDomicilioRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		core.RespondValidationError(ctx, "Datos de domicilio inválidos", map[string]string{"error": err.Error()})
		return
	}

	body := domain.Domicilio{
		DomicilioID: id,
		Alias:       strings.TrimSpace(req.Alias),
		Direccion:   buildDireccion(req),
		ColoniaID:   req.ColoniaID,
	}

	if err := c.update.Execute(&body); err != nil {
		core.RespondInternalServerError(ctx, "Error al actualizar domicilio", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "domicilio actualizado correctamente",
		"code":    http.StatusOK,
	})
}

// @Summary      Eliminar domicilio
// @Tags         Domicilio
// @Produce      json
// @Param        id path int true "ID del domicilio"
// @Success      200 {object} map[string]interface{}
// @Security     BearerAuth
// @Router       /api/domicilios/{id} [delete]
func (c *DomicilioController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		core.RespondInvalidInput(ctx, "ID de domicilio inválido")
		return
	}

	usuarioID := ctx.GetInt("user_id")
	if err := c.delete.Execute(id, usuarioID); err != nil {
		core.RespondError(ctx, http.StatusForbidden, core.ErrCodeForbidden, "No autorizado para eliminar este domicilio", map[string]string{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": fmt.Sprintf("domicilio %d eliminado correctamente", id),
		"code":    http.StatusOK,
	})
}
