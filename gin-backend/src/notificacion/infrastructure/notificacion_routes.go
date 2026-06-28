// notificacion_routes.go
package infrastructure

import (
	"github.com/gin-gonic/gin"
)

type NotificacionRouter struct {
	engine *gin.Engine
}

func NewNotificacionRouter(engine *gin.Engine) *NotificacionRouter {
	return &NotificacionRouter{
		engine: engine,
	}
}

func (router *NotificacionRouter) Run() {
	countActivasByUsuarioIDController,
		countByCamionIDController,
		countByTipoController,
		countByUsuarioIDController,
		crearEmergenciaController,
		crearFallaController,
		crearMantenimientoController,
		createNotificacionController,
		deleteNotificacionController,
		getAllNotificacionesController,
		getNotificacionByIdController,
		updateNotificacionController,
		getActivasByUsuarioIDController,
		getActivasController,
		getInactivasController,
		getByCamionIDController,
		getByCamionYTipoController,
		getByUsuarioIDController,
		getByUsuarioYTipoController,
		getByCreadoPorController,
		getByFallaIDController,
		getByMantenimientoIDController,
		getByTipoController,
		getByFechaRangeController,
		getGlobalesController,
		marcarComoActivaController,
		marcarComoLeidaController,
		marcarTodasComoLeidasController,
		notificarUsuarioController,
		obtenerNoLeidasController,
		notificarMultiplesController,
		notificarTodosController := InitNotificacionDependencies()

	notificacionGroup := router.engine.Group("/api/notificaciones")
	{
		notificacionGroup.POST("/", createNotificacionController.Run)
		notificacionGroup.GET("/", getAllNotificacionesController.Run)
		notificacionGroup.GET("/:id", getNotificacionByIdController.Run)
		notificacionGroup.PUT("/:id", updateNotificacionController.Run)
		notificacionGroup.DELETE("/:id", deleteNotificacionController.Run)

		notificacionGroup.GET("/count/usuario/:usuario_id", countByUsuarioIDController.Run)
		notificacionGroup.GET("/count/activas/usuario/:usuario_id", countActivasByUsuarioIDController.Run)
		notificacionGroup.GET("/count/tipo/:tipo", countByTipoController.Run)
		notificacionGroup.GET("/count/camion/:camion_id", countByCamionIDController.Run)

		notificacionGroup.POST("/emergencia", crearEmergenciaController.Run)
		notificacionGroup.POST("/falla", crearFallaController.Run)
		notificacionGroup.POST("/mantenimiento", crearMantenimientoController.Run)
		notificacionGroup.POST("/notificar", notificarUsuarioController.Run)
		notificacionGroup.POST("/enviar-multiples", notificarMultiplesController.Run)
		notificacionGroup.POST("/enviar-todos", notificarTodosController.Run)

		notificacionGroup.PATCH("/:id/marcar-leida", marcarComoLeidaController.Run)
		notificacionGroup.PATCH("/:id/reactivar", marcarComoActivaController.Run)
		notificacionGroup.PATCH("/usuario/:usuario_id/marcar-todas-leidas", marcarTodasComoLeidasController.Run)

		notificacionGroup.GET("/usuario/:usuario_id", getByUsuarioIDController.Run)
		notificacionGroup.GET("/activas/usuario/:usuario_id", getActivasByUsuarioIDController.Run)
		notificacionGroup.GET("/usuario/:usuario_id/tipo/:tipo", getByUsuarioYTipoController.Run)

		notificacionGroup.GET("/camion/:camion_id", getByCamionIDController.Run)
		notificacionGroup.GET("/camion/:camion_id/tipo/:tipo", getByCamionYTipoController.Run)
		notificacionGroup.GET("/tipo/:tipo", getByTipoController.Run)

		notificacionGroup.GET("/creado-por/:creado_por", getByCreadoPorController.Run)
		notificacionGroup.GET("/falla/:falla_id", getByFallaIDController.Run)
		notificacionGroup.GET("/mantenimiento/:mantenimiento_id", getByMantenimientoIDController.Run)

		notificacionGroup.GET("/activas", getActivasController.Run)
		notificacionGroup.GET("/inactivas", getInactivasController.Run)
		notificacionGroup.GET("/globales", getGlobalesController.Run)
		notificacionGroup.GET("/rango-fecha", getByFechaRangeController.Run)
		notificacionGroup.GET("/no-leidas/usuario/:usuario_id", obtenerNoLeidasController.Run)
	}
}
