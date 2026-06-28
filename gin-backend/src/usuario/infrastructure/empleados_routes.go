package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/vicpoo/API_recolecta/src/core"
)

// RegisterEmpleadoRoutes expone /api/empleados como alias de usuarios internos
// (tabla usuario con roles distintos de ciudadano).
func RegisterEmpleadoRoutes(r *gin.Engine, deps *UsuarioDependencies) {
	empleados := r.Group("/api/empleados")
	empleados.POST("/login", deps.Login.Handle)

	protected := empleados.Group("")
	protected.Use(core.JWTAuthMiddleware(), core.RequireRole(core.ADMIN))
	{
		protected.POST("/", deps.Create.Handle)
		protected.GET("/", deps.List.Handle)
		protected.GET("/:id", deps.Get.Handle)
		protected.DELETE("/:id", deps.Delete.Handle)
	}
}
