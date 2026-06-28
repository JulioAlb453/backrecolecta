package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/colonia/application"
	"github.com/vicpoo/API_recolecta/src/colonia/domain"
	"github.com/vicpoo/API_recolecta/src/core"
)

type ColoniaController struct {
	create *application.CreateColonia
	get    *application.GetColonia
	list   *application.ListColonias
	update *application.UpdateColonia
	delete *application.DeleteColonia
}

func NewColoniaController(
	create *application.CreateColonia,
	get *application.GetColonia,
	list *application.ListColonias,
	update *application.UpdateColonia,
	delete *application.DeleteColonia,
) *ColoniaController {
	return &ColoniaController{create, get, list, update, delete}
}

func (c *ColoniaController) RegisterRoutes(r *gin.Engine) {
	registerColoniaRoutes(r, c, "/api/colonia")
	registerColoniaRoutes(r, c, "/colonias") // compatibilidad con clientes previos
}

func registerColoniaRoutes(r *gin.Engine, c *ColoniaController, basePath string) {
	public := r.Group(basePath)
	{
		public.GET("", c.List)
		public.GET("/:id", c.GetByID)
	}

	admin := r.Group(
		basePath,
		core.JWTAuthMiddleware(),
		core.RequireRole(core.ADMIN),
	)
	{
		admin.POST("", c.Create)
		admin.PUT("/:id", c.Update)
		admin.DELETE("/:id", c.Delete)
	}
}

// @Summary      Crear colonia
// @Description  Requiere rol Administrador. Body: nombre y zona.
// @Tags         Colonia
// @Accept       json
// @Produce      json
// @Param        body body domain.CreateColoniaRequest true "Datos de colonia"
// @Success      201 {object} domain.ColoniaResponse
// @Failure      400 {object} core.ErrorResponse
// @Failure      401 {object} core.ErrorResponse
// @Failure      403 {object} core.ErrorResponse
// @Security     BearerAuth
// @Router       /api/colonia [post]
func (c *ColoniaController) Create(ctx *gin.Context) {
	var request domain.CreateColoniaRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		core.RespondValidationError(ctx, "Datos de colonia inválidos", map[string]string{"error": err.Error()})
		return
	}

	body := domain.Colonia{
		Nombre: request.Nombre,
		Zona:   request.Zona,
	}

	if err := c.create.Execute(&body); err != nil {
		core.RespondInternalServerError(ctx, "Error al crear colonia", err)
		return
	}

	ctx.JSON(http.StatusCreated, domain.ColoniaResponse{
		Success: true,
		Message: "Colonia creada correctamente",
		Data:    body,
		Code:    http.StatusCreated,
	})
}

// @Summary      Obtener colonia por ID
// @Tags         Colonia
// @Produce      json
// @Param        id path int true "ID de colonia"
// @Success      200 {object} domain.ColoniaResponse
// @Failure      404 {object} core.ErrorResponse
// @Router       /api/colonia/{id} [get]
func (c *ColoniaController) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		core.RespondInvalidInput(ctx, "ID de colonia inválido")
		return
	}

	colonia, err := c.get.Execute(id)
	if err != nil {
		core.RespondError(ctx, http.StatusNotFound, core.ErrCodeNotFound, "Colonia no encontrada", nil)
		return
	}

	ctx.JSON(http.StatusOK, domain.ColoniaResponse{
		Success: true,
		Message: "Colonia obtenida correctamente",
		Data:    *colonia,
		Code:    http.StatusOK,
	})
}

// @Summary      Listar colonias
// @Tags         Colonia
// @Produce      json
// @Success      200 {object} domain.ColoniaListResponse
// @Failure      500 {object} core.ErrorResponse
// @Router       /api/colonia [get]
func (c *ColoniaController) List(ctx *gin.Context) {
	colonias, err := c.list.Execute()
	if err != nil {
		core.RespondInternalServerError(ctx, "Error al listar colonias", err)
		return
	}

	ctx.JSON(http.StatusOK, domain.ColoniaListResponse{
		Success: true,
		Message: "Colonias listadas correctamente",
		Data:    colonias,
		Code:    http.StatusOK,
	})
}

// @Summary      Actualizar colonia
// @Tags         Colonia
// @Accept       json
// @Produce      json
// @Param        id path int true "ID de colonia"
// @Param        body body domain.UpdateColoniaRequest true "Datos de colonia"
// @Success      200 {object} domain.ColoniaMessageResponse
// @Security     BearerAuth
// @Router       /api/colonia/{id} [put]
func (c *ColoniaController) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		core.RespondInvalidInput(ctx, "ID de colonia inválido")
		return
	}

	var request domain.UpdateColoniaRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		core.RespondValidationError(ctx, "Datos de colonia inválidos", map[string]string{"error": err.Error()})
		return
	}

	body := domain.Colonia{
		ColoniaID: id,
		Nombre:    request.Nombre,
		Zona:      request.Zona,
	}

	if err := c.update.Execute(&body); err != nil {
		core.RespondInternalServerError(ctx, "Error al actualizar colonia", err)
		return
	}

	ctx.JSON(http.StatusOK, domain.ColoniaMessageResponse{
		Success: true,
		Message: "Colonia actualizada correctamente",
		Code:    http.StatusOK,
	})
}

// @Summary      Eliminar colonia
// @Tags         Colonia
// @Param        id path int true "ID de colonia"
// @Success      204
// @Security     BearerAuth
// @Router       /api/colonia/{id} [delete]
func (c *ColoniaController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		core.RespondInvalidInput(ctx, "ID de colonia inválido")
		return
	}

	if err := c.delete.Execute(id); err != nil {
		core.RespondInternalServerError(ctx, "Error al eliminar colonia", err)
		return
	}

	ctx.Status(http.StatusNoContent)
}
