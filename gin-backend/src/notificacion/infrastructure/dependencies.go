package infrastructure

import (
	"github.com/vicpoo/API_recolecta/src/notificacion/application"
)

func InitNotificacionDependencies() (
	*CountNotificacionesActivasByUsuarioIDController,
	*CountNotificacionesByCamionIDController,
	*CountNotificacionesByTipoController,
	*CountNotificacionesByUsuarioIDController,
	*CrearNotificacionEmergenciaController,
	*CrearNotificacionFallaController,
	*CrearNotificacionMantenimientoController,
	*CreateNotificacionController,
	*DeleteNotificacionController,
	*GetAllNotificacionesController,
	*GetNotificacionByIdController,
	*UpdateNotificacionController,
	*GetNotificacionesActivasByUsuarioIDController,
	*GetNotificacionesActivasController,
	*GetNotificacionesInactivasController,
	*GetNotificacionesByCamionIDController,
	*GetNotificacionesByCamionYTipoController,
	*GetNotificacionesByUsuarioIDController,
	*GetNotificacionesByUsuarioYTipoController,
	*GetNotificacionesByCreadoPorController,
	*GetNotificacionesByFallaIDController,
	*GetNotificacionesByMantenimientoIDController,
	*GetNotificacionesByTipoController,
	*GetNotificacionesByFechaRangeController,
	*GetNotificacionesGlobalesController,
	*MarcarNotificacionComoActivaController,
	*MarcarNotificacionComoLeidaController,
	*MarcarTodasNotificacionesComoLeidasController,
	*NotificarUsuarioController,
	*ObtenerNumeroNotificacionesNoLeidasController,
	*NotificarMultiplesUsuariosController,
	*NotificarTodosUsuariosController,
) {
	repo := NewPostgresNotificacionRepository()

	countActivasByUsuarioIDUseCase := application.NewCountNotificacionesActivasByUsuarioIDUseCase(repo)
	countByCamionIDUseCase := application.NewCountNotificacionesByCamionIDUseCase(repo)
	countByTipoUseCase := application.NewCountNotificacionesByTipoUseCase(repo)
	countByUsuarioIDUseCase := application.NewCountNotificacionesByUsuarioIDUseCase(repo)

	crearEmergenciaUseCase := application.NewCrearNotificacionEmergenciaUseCase(repo)
	crearFallaUseCase := application.NewCrearNotificacionFallaUseCase(repo)
	crearMantenimientoUseCase := application.NewCrearNotificacionMantenimientoUseCase(repo)
	createNotificacionUseCase := application.NewCreateNotificacionUseCase(repo)

	deleteNotificacionUseCase := application.NewDeleteNotificacionUseCase(repo)
	getAllNotificacionesUseCase := application.NewGetAllNotificacionesUseCase(repo)
	getNotificacionByIdUseCase := application.NewGetNotificacionByIdUseCase(repo)
	updateNotificacionUseCase := application.NewUpdateNotificacionUseCase(repo)

	getActivasByUsuarioIDUseCase := application.NewGetNotificacionesActivasByUsuarioIDUseCase(repo)
	getActivasUseCase := application.NewGetNotificacionesActivasUseCase(repo)
	getInactivasUseCase := application.NewGetNotificacionesInactivasUseCase(repo)

	getByCamionIDUseCase := application.NewGetNotificacionesByCamionIDUseCase(repo)
	getByCamionYTipoUseCase := application.NewGetNotificacionesByCamionYTipoUseCase(repo)

	getByUsuarioIDUseCase := application.NewGetNotificacionesByUsuarioIDUseCase(repo)
	getByUsuarioYTipoUseCase := application.NewGetNotificacionesByUsuarioYTipoUseCase(repo)

	getByCreadoPorUseCase := application.NewGetNotificacionesByCreadoPorUseCase(repo)
	getByFallaIDUseCase := application.NewGetNotificacionesByFallaIDUseCase(repo)
	getByMantenimientoIDUseCase := application.NewGetNotificacionesByMantenimientoIDUseCase(repo)

	getByTipoUseCase := application.NewGetNotificacionesByTipoUseCase(repo)
	getByFechaRangeUseCase := application.NewGetNotificacionesByFechaRangeUseCase(repo)

	getGlobalesUseCase := application.NewGetNotificacionesGlobalesUseCase(repo)

	marcarComoActivaUseCase := application.NewMarcarNotificacionComoActivaUseCase(repo)
	marcarComoLeidaUseCase := application.NewMarcarNotificacionComoLeidaUseCase(repo)
	marcarTodasComoLeidasUseCase := application.NewMarcarTodasNotificacionesComoLeidasUseCase(repo)

	notificarUsuarioUseCase := application.NewNotificarUsuarioUseCase(repo)
	obtenerNoLeidasUseCase := application.NewObtenerNumeroNotificacionesNoLeidasUseCase(repo)

	notificarMultiplesUseCase := application.NewNotificarMultiplesUsuariosUseCase(repo)
	notificarTodosUseCase := application.NewNotificarTodosUsuariosUseCase(repo)

	countActivasByUsuarioIDController := NewCountNotificacionesActivasByUsuarioIDController(countActivasByUsuarioIDUseCase)
	countByCamionIDController := NewCountNotificacionesByCamionIDController(countByCamionIDUseCase)
	countByTipoController := NewCountNotificacionesByTipoController(countByTipoUseCase)
	countByUsuarioIDController := NewCountNotificacionesByUsuarioIDController(countByUsuarioIDUseCase)

	crearEmergenciaController := NewCrearNotificacionEmergenciaController(crearEmergenciaUseCase)
	crearFallaController := NewCrearNotificacionFallaController(crearFallaUseCase)
	crearMantenimientoController := NewCrearNotificacionMantenimientoController(crearMantenimientoUseCase)
	createNotificacionController := NewCreateNotificacionController(createNotificacionUseCase)

	deleteNotificacionController := NewDeleteNotificacionController(deleteNotificacionUseCase)
	getAllNotificacionesController := NewGetAllNotificacionesController(getAllNotificacionesUseCase)
	getNotificacionByIdController := NewGetNotificacionByIdController(getNotificacionByIdUseCase)
	updateNotificacionController := NewUpdateNotificacionController(updateNotificacionUseCase)

	getActivasByUsuarioIDController := NewGetNotificacionesActivasByUsuarioIDController(getActivasByUsuarioIDUseCase)
	getActivasController := NewGetNotificacionesActivasController(getActivasUseCase)
	getInactivasController := NewGetNotificacionesInactivasController(getInactivasUseCase)

	getByCamionIDController := NewGetNotificacionesByCamionIDController(getByCamionIDUseCase)
	getByCamionYTipoController := NewGetNotificacionesByCamionYTipoController(getByCamionYTipoUseCase)

	getByUsuarioIDController := NewGetNotificacionesByUsuarioIDController(getByUsuarioIDUseCase)
	getByUsuarioYTipoController := NewGetNotificacionesByUsuarioYTipoController(getByUsuarioYTipoUseCase)

	getByCreadoPorController := NewGetNotificacionesByCreadoPorController(getByCreadoPorUseCase)
	getByFallaIDController := NewGetNotificacionesByFallaIDController(getByFallaIDUseCase)
	getByMantenimientoIDController := NewGetNotificacionesByMantenimientoIDController(getByMantenimientoIDUseCase)

	getByTipoController := NewGetNotificacionesByTipoController(getByTipoUseCase)
	getByFechaRangeController := NewGetNotificacionesByFechaRangeController(getByFechaRangeUseCase)

	getGlobalesController := NewGetNotificacionesGlobalesController(getGlobalesUseCase)

	marcarComoActivaController := NewMarcarNotificacionComoActivaController(marcarComoActivaUseCase)
	marcarComoLeidaController := NewMarcarNotificacionComoLeidaController(marcarComoLeidaUseCase)
	marcarTodasComoLeidasController := NewMarcarTodasNotificacionesComoLeidasController(marcarTodasComoLeidasUseCase)

	notificarUsuarioController := NewNotificarUsuarioController(notificarUsuarioUseCase)
	obtenerNoLeidasController := NewObtenerNumeroNotificacionesNoLeidasController(obtenerNoLeidasUseCase)

	notificarMultiplesController := NewNotificarMultiplesUsuariosController(notificarMultiplesUseCase)
	notificarTodosController := NewNotificarTodosUsuariosController(notificarTodosUseCase)

	return countActivasByUsuarioIDController,
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
		notificarTodosController
}
