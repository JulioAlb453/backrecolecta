package main

// Archivo generado por scripts/generate_swagger_stubs.py — no editar a mano.

// doc_get_health godoc
// @Summary Health check
// @Tags Sistema
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /health [get]
func doc_get_health() {}

// doc_post_api_tipo_camion godoc
// @Summary Crear tipo de camión
// @Tags TipoCamion
// @Accept json
// @Produce json
// @Param body body CreateTipoCamionRequest true "Payload JSON"
// @Success 200 {object} map[string]interface{}
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/tipo-camion/ [post]
func doc_post_api_tipo_camion() {}

// doc_get_api_tipo_camion godoc
// @Summary Listar tipos de camión
// @Tags TipoCamion
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/tipo-camion/ [get]
func doc_get_api_tipo_camion() {}

// doc_get_api_tipo_camion_nombre_nombre godoc
// @Summary Obtener tipo de camión por nombre
// @Tags TipoCamion
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/tipo-camion/nombre/{nombre} [get]
func doc_get_api_tipo_camion_nombre_nombre() {}

// doc_delete_api_tipo_camion_id godoc
// @Summary Eliminar tipo de camión
// @Tags TipoCamion
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/tipo-camion/{id} [delete]
func doc_delete_api_tipo_camion_id() {}

// doc_post_api_camion godoc
// @Summary Crear camión
// @Tags Camion
// @Accept json
// @Produce json
// @Param body body CreateCamionRequest true "Payload JSON"
// @Success 200 {object} map[string]interface{}
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/camion/ [post]
func doc_post_api_camion() {}

// doc_get_api_camion godoc
// @Summary Listar camiones
// @Tags Camion
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/camion/ [get]
func doc_get_api_camion() {}

// doc_get_api_camion_id godoc
// @Summary Obtener camión por ID
// @Tags Camion
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/camion/{id} [get]
func doc_get_api_camion_id() {}

// doc_put_api_camion_id godoc
// @Summary Actualizar camión
// @Tags Camion
// @Accept json
// @Produce json
// @Param body body CreateCamionRequest true "Payload JSON"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/camion/{id} [put]
func doc_put_api_camion_id() {}

// doc_delete_api_camion_id godoc
// @Summary Eliminar camión
// @Tags Camion
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/camion/{id} [delete]
func doc_delete_api_camion_id() {}

// doc_get_api_camion_placa_placa godoc
// @Summary Obtener camión por placa
// @Tags Camion
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/camion/placa/{placa} [get]
func doc_get_api_camion_placa_placa() {}

// doc_get_api_camion_modelo godoc
// @Summary Buscar camión por modelo
// @Tags Camion
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/camion/modelo [get]
func doc_get_api_camion_modelo() {}

// doc_post_api_estado_camion godoc
// @Summary Crear estado de camión
// @Tags EstadoCamion
// @Accept json
// @Produce json
// @Param body body CreateEstadoCamionRequest true "Payload JSON"
// @Success 200 {object} map[string]interface{}
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/estado-camion/ [post]
func doc_post_api_estado_camion() {}

// doc_get_api_estado_camion godoc
// @Summary Listar estados de camión
// @Tags EstadoCamion
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/estado-camion/ [get]
func doc_get_api_estado_camion() {}

// doc_get_api_estado_camion_camion_id godoc
// @Summary Obtener estado por camión
// @Tags EstadoCamion
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/estado-camion/camion/{id} [get]
func doc_get_api_estado_camion_camion_id() {}

// doc_put_api_estado_camion_id godoc
// @Summary Actualizar estado de camión
// @Tags EstadoCamion
// @Accept json
// @Produce json
// @Param body body CreateEstadoCamionRequest true "Payload JSON"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/estado-camion/{id} [put]
func doc_put_api_estado_camion_id() {}

// doc_delete_api_estado_camion_id godoc
// @Summary Eliminar estado de camión
// @Tags EstadoCamion
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/estado-camion/{id} [delete]
func doc_delete_api_estado_camion_id() {}

// doc_post_api_historial_asignacion godoc
// @Summary Crear historial de asignación
// @Tags HistorialAsignacion
// @Accept json
// @Produce json
// @Param body body CreateHistorialAsignacionRequest true "Payload JSON"
// @Success 200 {object} map[string]interface{}
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/historial-asignacion/ [post]
func doc_post_api_historial_asignacion() {}

// doc_get_api_historial_asignacion godoc
// @Summary Listar historial de asignación
// @Tags HistorialAsignacion
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/historial-asignacion/ [get]
func doc_get_api_historial_asignacion() {}

// doc_get_api_historial_asignacion_id godoc
// @Summary Obtener historial por ID
// @Tags HistorialAsignacion
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/historial-asignacion/{id} [get]
func doc_get_api_historial_asignacion_id() {}

// doc_put_api_historial_asignacion_id godoc
// @Summary Actualizar historial de asignación
// @Tags HistorialAsignacion
// @Accept json
// @Produce json
// @Param body body CreateHistorialAsignacionRequest true "Payload JSON"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/historial-asignacion/{id} [put]
func doc_put_api_historial_asignacion_id() {}

// doc_delete_api_historial_asignacion_id godoc
// @Summary Eliminar historial de asignación
// @Tags HistorialAsignacion
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/historial-asignacion/{id} [delete]
func doc_delete_api_historial_asignacion_id() {}

// doc_get_api_historial_asignacion_camion_camionId godoc
// @Summary Historial por camión
// @Tags HistorialAsignacion
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/historial-asignacion/camion/{camionId} [get]
func doc_get_api_historial_asignacion_camion_camionId() {}

// doc_get_api_historial_asignacion_chofer_choferId godoc
// @Summary Historial por chofer
// @Tags HistorialAsignacion
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/historial-asignacion/chofer/{choferId} [get]
func doc_get_api_historial_asignacion_chofer_choferId() {}

// doc_get_api_historial_asignacion_activo_camion_camionId godoc
// @Summary Asignación activa por camión
// @Tags HistorialAsignacion
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/historial-asignacion/activo/camion/{camionId} [get]
func doc_get_api_historial_asignacion_activo_camion_camionId() {}

// doc_get_api_historial_asignacion_activo_chofer_choferId godoc
// @Summary Asignación activa por chofer
// @Tags HistorialAsignacion
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/historial-asignacion/activo/chofer/{choferId} [get]
func doc_get_api_historial_asignacion_activo_chofer_choferId() {}

// doc_put_api_historial_asignacion_baja_id godoc
// @Summary Dar de baja historial
// @Tags HistorialAsignacion
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/historial-asignacion/baja/{id} [put]
func doc_put_api_historial_asignacion_baja_id() {}

// doc_put_api_historial_asignacion_cerrar_camion_camionId godoc
// @Summary Cerrar asignación por camión
// @Tags HistorialAsignacion
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/historial-asignacion/cerrar/camion/{camionId} [put]
func doc_put_api_historial_asignacion_cerrar_camion_camionId() {}

// doc_put_api_historial_asignacion_cerrar_chofer_choferId godoc
// @Summary Cerrar asignación por chofer
// @Tags HistorialAsignacion
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/historial-asignacion/cerrar/chofer/{choferId} [put]
func doc_put_api_historial_asignacion_cerrar_chofer_choferId() {}

// doc_post_api_rutas godoc
// @Summary Crear ruta
// @Tags Ruta
// @Accept json
// @Produce json
// @Param body body CreateRutaRequest true "Payload JSON"
// @Success 200 {object} map[string]interface{}
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/rutas/ [post]
func doc_post_api_rutas() {}

// doc_get_api_rutas godoc
// @Summary Listar rutas
// @Tags Ruta
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/rutas/ [get]
func doc_get_api_rutas() {}

// doc_get_api_rutas_id godoc
// @Summary Obtener ruta por ID
// @Tags Ruta
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/rutas/{id} [get]
func doc_get_api_rutas_id() {}

// doc_put_api_rutas_id godoc
// @Summary Actualizar ruta
// @Tags Ruta
// @Accept json
// @Produce json
// @Param body body CreateRutaRequest true "Payload JSON"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/rutas/{id} [put]
func doc_put_api_rutas_id() {}

// doc_delete_api_rutas_id godoc
// @Summary Eliminar ruta
// @Tags Ruta
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/rutas/{id} [delete]
func doc_delete_api_rutas_id() {}

// doc_get_api_rutas_activas godoc
// @Summary Listar rutas activas
// @Tags Ruta
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/rutas/activas [get]
func doc_get_api_rutas_activas() {}

// doc_post_api_puntos_recoleccion godoc
// @Summary Crear punto de recolección
// @Tags PuntoRecoleccion
// @Accept json
// @Produce json
// @Param body body CreatePuntoRecoleccionRequest true "Payload JSON"
// @Success 200 {object} map[string]interface{}
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/puntos-recoleccion/ [post]
func doc_post_api_puntos_recoleccion() {}

// doc_get_api_puntos_recoleccion godoc
// @Summary Listar puntos de recolección
// @Tags PuntoRecoleccion
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/puntos-recoleccion/ [get]
func doc_get_api_puntos_recoleccion() {}

// doc_get_api_puntos_recoleccion_id godoc
// @Summary Obtener punto por ID
// @Tags PuntoRecoleccion
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/puntos-recoleccion/{id} [get]
func doc_get_api_puntos_recoleccion_id() {}

// doc_get_api_puntos_recoleccion_ruta_rutaId godoc
// @Summary Puntos por ruta
// @Tags PuntoRecoleccion
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/puntos-recoleccion/ruta/{rutaId} [get]
func doc_get_api_puntos_recoleccion_ruta_rutaId() {}

// doc_put_api_puntos_recoleccion_id godoc
// @Summary Actualizar punto de recolección
// @Tags PuntoRecoleccion
// @Accept json
// @Produce json
// @Param body body CreatePuntoRecoleccionRequest true "Payload JSON"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/puntos-recoleccion/{id} [put]
func doc_put_api_puntos_recoleccion_id() {}

// doc_delete_api_puntos_recoleccion_id godoc
// @Summary Eliminar punto de recolección
// @Tags PuntoRecoleccion
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/puntos-recoleccion/{id} [delete]
func doc_delete_api_puntos_recoleccion_id() {}

// doc_post_api_relleno_sanitario godoc
// @Summary Crear relleno sanitario
// @Tags RellenoSanitario
// @Accept json
// @Produce json
// @Param body body CreateRellenoSanitarioRequest true "Payload JSON"
// @Success 200 {object} map[string]interface{}
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/relleno-sanitario/ [post]
func doc_post_api_relleno_sanitario() {}

// doc_get_api_relleno_sanitario godoc
// @Summary Listar rellenos sanitarios
// @Tags RellenoSanitario
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/relleno-sanitario/ [get]
func doc_get_api_relleno_sanitario() {}

// doc_get_api_relleno_sanitario_id godoc
// @Summary Obtener relleno por ID
// @Tags RellenoSanitario
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/relleno-sanitario/{id} [get]
func doc_get_api_relleno_sanitario_id() {}

// doc_put_api_relleno_sanitario_id godoc
// @Summary Actualizar relleno sanitario
// @Tags RellenoSanitario
// @Accept json
// @Produce json
// @Param body body CreateRellenoSanitarioRequest true "Payload JSON"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/relleno-sanitario/{id} [put]
func doc_put_api_relleno_sanitario_id() {}

// doc_delete_api_relleno_sanitario_id godoc
// @Summary Eliminar relleno sanitario
// @Tags RellenoSanitario
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/relleno-sanitario/{id} [delete]
func doc_delete_api_relleno_sanitario_id() {}

// doc_get_api_relleno_sanitario_buscar godoc
// @Summary Buscar relleno sanitario
// @Tags RellenoSanitario
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/relleno-sanitario/buscar [get]
func doc_get_api_relleno_sanitario_buscar() {}

// doc_get_api_relleno_sanitario_exists_id godoc
// @Summary Verificar existencia de relleno
// @Tags RellenoSanitario
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/relleno-sanitario/exists/{id} [get]
func doc_get_api_relleno_sanitario_exists_id() {}

// doc_post_api_ruta_camion godoc
// @Summary Crear ruta-camión
// @Tags RutaCamion
// @Accept json
// @Produce json
// @Param body body CreateRutaCamionRequest true "Payload JSON"
// @Success 200 {object} map[string]interface{}
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/ruta-camion/ [post]
func doc_post_api_ruta_camion() {}

// doc_get_api_ruta_camion godoc
// @Summary Listar rutas-camión
// @Tags RutaCamion
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/ruta-camion/ [get]
func doc_get_api_ruta_camion() {}

// doc_get_api_ruta_camion_id godoc
// @Summary Obtener ruta-camión por ID
// @Tags RutaCamion
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/ruta-camion/{id} [get]
func doc_get_api_ruta_camion_id() {}

// doc_get_api_ruta_camion_camion_camion_id godoc
// @Summary Rutas por camión
// @Tags RutaCamion
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/ruta-camion/camion/{camion_id} [get]
func doc_get_api_ruta_camion_camion_camion_id() {}

// doc_get_api_ruta_camion_ruta_ruta_id godoc
// @Summary Camiones por ruta
// @Tags RutaCamion
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/ruta-camion/ruta/{ruta_id} [get]
func doc_get_api_ruta_camion_ruta_ruta_id() {}

// doc_get_api_ruta_camion_exists_id godoc
// @Summary Verificar existencia ruta-camión
// @Tags RutaCamion
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/ruta-camion/exists/{id} [get]
func doc_get_api_ruta_camion_exists_id() {}

// doc_put_api_ruta_camion_id godoc
// @Summary Actualizar ruta-camión
// @Tags RutaCamion
// @Accept json
// @Produce json
// @Param body body CreateRutaCamionRequest true "Payload JSON"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/ruta-camion/{id} [put]
func doc_put_api_ruta_camion_id() {}

// doc_delete_api_ruta_camion_id godoc
// @Summary Eliminar ruta-camión
// @Tags RutaCamion
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/ruta-camion/{id} [delete]
func doc_delete_api_ruta_camion_id() {}

// doc_post_api_registro_vaciado godoc
// @Summary Crear registro de vaciado
// @Tags RegistroVaciado
// @Accept json
// @Produce json
// @Param body body CreateRegistroVaciadoRequest true "Payload JSON"
// @Success 200 {object} map[string]interface{}
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/registro-vaciado/ [post]
func doc_post_api_registro_vaciado() {}

// doc_get_api_registro_vaciado godoc
// @Summary Listar registros de vaciado
// @Tags RegistroVaciado
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/registro-vaciado/ [get]
func doc_get_api_registro_vaciado() {}

// doc_get_api_registro_vaciado_id godoc
// @Summary Obtener registro de vaciado por ID
// @Tags RegistroVaciado
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/registro-vaciado/{id} [get]
func doc_get_api_registro_vaciado_id() {}

// doc_get_api_registro_vaciado_relleno_relleno_id godoc
// @Summary Registros por relleno
// @Tags RegistroVaciado
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/registro-vaciado/relleno/{relleno_id} [get]
func doc_get_api_registro_vaciado_relleno_relleno_id() {}

// doc_get_api_registro_vaciado_ruta_camion_ruta_camion_id godoc
// @Summary Registros por ruta-camión
// @Tags RegistroVaciado
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/registro-vaciado/ruta-camion/{ruta_camion_id} [get]
func doc_get_api_registro_vaciado_ruta_camion_ruta_camion_id() {}

// doc_get_api_registro_vaciado_exists_id godoc
// @Summary Verificar existencia registro vaciado
// @Tags RegistroVaciado
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/registro-vaciado/exists/{id} [get]
func doc_get_api_registro_vaciado_exists_id() {}

// doc_delete_api_registro_vaciado_id godoc
// @Summary Eliminar registro de vaciado
// @Tags RegistroVaciado
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/registro-vaciado/{id} [delete]
func doc_delete_api_registro_vaciado_id() {}

// doc_post_api_notificaciones godoc
// @Summary Crear notificación
// @Tags Notificacion
// @Accept json
// @Produce json
// @Param body body CreateNotificacionRequest true "Payload JSON"
// @Success 200 {object} map[string]interface{}
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/notificaciones/ [post]
func doc_post_api_notificaciones() {}

// doc_get_api_notificaciones godoc
// @Summary Listar notificaciones
// @Tags Notificacion
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/notificaciones/ [get]
func doc_get_api_notificaciones() {}

// doc_get_api_notificaciones_id godoc
// @Summary Obtener notificación por ID
// @Tags Notificacion
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/notificaciones/{id} [get]
func doc_get_api_notificaciones_id() {}

// doc_put_api_notificaciones_id godoc
// @Summary Actualizar notificación
// @Tags Notificacion
// @Accept json
// @Produce json
// @Param body body UpdateNotificacionRequest true "Payload JSON"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/notificaciones/{id} [put]
func doc_put_api_notificaciones_id() {}

// doc_delete_api_notificaciones_id godoc
// @Summary Eliminar notificación
// @Tags Notificacion
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/notificaciones/{id} [delete]
func doc_delete_api_notificaciones_id() {}

// doc_get_api_notificaciones_count_usuario_usuario_id godoc
// @Summary Contar notificaciones por usuario
// @Tags Notificacion
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/notificaciones/count/usuario/{usuario_id} [get]
func doc_get_api_notificaciones_count_usuario_usuario_id() {}

// doc_get_api_notificaciones_count_activas_usuario_usuario_id godoc
// @Summary Contar activas por usuario
// @Tags Notificacion
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/notificaciones/count/activas/usuario/{usuario_id} [get]
func doc_get_api_notificaciones_count_activas_usuario_usuario_id() {}

// doc_get_api_notificaciones_count_tipo_tipo godoc
// @Summary Contar por tipo
// @Tags Notificacion
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/notificaciones/count/tipo/{tipo} [get]
func doc_get_api_notificaciones_count_tipo_tipo() {}

// doc_get_api_notificaciones_count_camion_camion_id godoc
// @Summary Contar por camión
// @Tags Notificacion
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/notificaciones/count/camion/{camion_id} [get]
func doc_get_api_notificaciones_count_camion_camion_id() {}

// doc_post_api_notificaciones_emergencia godoc
// @Summary Crear notificación de emergencia
// @Tags Notificacion
// @Accept json
// @Produce json
// @Param body body NotificacionEmergenciaRequest true "Payload JSON"
// @Success 200 {object} map[string]interface{}
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/notificaciones/emergencia [post]
func doc_post_api_notificaciones_emergencia() {}

// doc_post_api_notificaciones_falla godoc
// @Summary Crear notificación de falla
// @Tags Notificacion
// @Accept json
// @Produce json
// @Param body body NotificacionFallaRequest true "Payload JSON"
// @Success 200 {object} map[string]interface{}
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/notificaciones/falla [post]
func doc_post_api_notificaciones_falla() {}

// doc_post_api_notificaciones_mantenimiento godoc
// @Summary Crear notificación de mantenimiento
// @Tags Notificacion
// @Accept json
// @Produce json
// @Param body body NotificacionMantenimientoRequest true "Payload JSON"
// @Success 200 {object} map[string]interface{}
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/notificaciones/mantenimiento [post]
func doc_post_api_notificaciones_mantenimiento() {}

// doc_post_api_notificaciones_notificar godoc
// @Summary Enviar notificación
// @Tags Notificacion
// @Accept json
// @Produce json
// @Param body body NotificarUsuarioRequest true "Payload JSON"
// @Success 200 {object} map[string]interface{}
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/notificaciones/notificar [post]
func doc_post_api_notificaciones_notificar() {}

// doc_post_api_notificaciones_enviar_multiples godoc
// @Summary Enviar múltiples notificaciones
// @Tags Notificacion
// @Accept json
// @Produce json
// @Param body body NotificarMultiplesRequest true "Payload JSON"
// @Success 200 {object} map[string]interface{}
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/notificaciones/enviar-multiples [post]
func doc_post_api_notificaciones_enviar_multiples() {}

// doc_post_api_notificaciones_enviar_todos godoc
// @Summary Enviar a todos
// @Tags Notificacion
// @Accept json
// @Produce json
// @Param body body NotificarTodosRequest true "Payload JSON"
// @Success 200 {object} map[string]interface{}
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/notificaciones/enviar-todos [post]
func doc_post_api_notificaciones_enviar_todos() {}

// doc_patch_api_notificaciones_id_marcar_leida godoc
// @Summary Marcar notificación leída
// @Tags Notificacion
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/notificaciones/{id}/marcar-leida [patch]
func doc_patch_api_notificaciones_id_marcar_leida() {}

// doc_patch_api_notificaciones_id_reactivar godoc
// @Summary Reactivar notificación
// @Tags Notificacion
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/notificaciones/{id}/reactivar [patch]
func doc_patch_api_notificaciones_id_reactivar() {}

// doc_patch_api_notificaciones_usuario_usuario_id_marcar_todas_leidas godoc
// @Summary Marcar todas leídas
// @Tags Notificacion
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/notificaciones/usuario/{usuario_id}/marcar-todas-leidas [patch]
func doc_patch_api_notificaciones_usuario_usuario_id_marcar_todas_leidas() {}

// doc_get_api_notificaciones_usuario_usuario_id godoc
// @Summary Notificaciones por usuario
// @Tags Notificacion
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/notificaciones/usuario/{usuario_id} [get]
func doc_get_api_notificaciones_usuario_usuario_id() {}

// doc_get_api_notificaciones_activas_usuario_usuario_id godoc
// @Summary Activas por usuario
// @Tags Notificacion
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/notificaciones/activas/usuario/{usuario_id} [get]
func doc_get_api_notificaciones_activas_usuario_usuario_id() {}

// doc_get_api_notificaciones_usuario_usuario_id_tipo_tipo godoc
// @Summary Por usuario y tipo
// @Tags Notificacion
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/notificaciones/usuario/{usuario_id}/tipo/{tipo} [get]
func doc_get_api_notificaciones_usuario_usuario_id_tipo_tipo() {}

// doc_get_api_notificaciones_camion_camion_id godoc
// @Summary Notificaciones por camión
// @Tags Notificacion
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/notificaciones/camion/{camion_id} [get]
func doc_get_api_notificaciones_camion_camion_id() {}

// doc_get_api_notificaciones_camion_camion_id_tipo_tipo godoc
// @Summary Por camión y tipo
// @Tags Notificacion
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/notificaciones/camion/{camion_id}/tipo/{tipo} [get]
func doc_get_api_notificaciones_camion_camion_id_tipo_tipo() {}

// doc_get_api_notificaciones_tipo_tipo godoc
// @Summary Notificaciones por tipo
// @Tags Notificacion
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/notificaciones/tipo/{tipo} [get]
func doc_get_api_notificaciones_tipo_tipo() {}

// doc_get_api_notificaciones_creado_por_creado_por godoc
// @Summary Notificaciones por creador
// @Tags Notificacion
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/notificaciones/creado-por/{creado_por} [get]
func doc_get_api_notificaciones_creado_por_creado_por() {}

// doc_get_api_notificaciones_falla_falla_id godoc
// @Summary Notificaciones por falla
// @Tags Notificacion
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/notificaciones/falla/{falla_id} [get]
func doc_get_api_notificaciones_falla_falla_id() {}

// doc_get_api_notificaciones_mantenimiento_mantenimiento_id godoc
// @Summary Notificaciones por mantenimiento
// @Tags Notificacion
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/notificaciones/mantenimiento/{mantenimiento_id} [get]
func doc_get_api_notificaciones_mantenimiento_mantenimiento_id() {}

// doc_get_api_notificaciones_activas godoc
// @Summary Listar notificaciones activas
// @Tags Notificacion
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/notificaciones/activas [get]
func doc_get_api_notificaciones_activas() {}

// doc_get_api_notificaciones_inactivas godoc
// @Summary Listar notificaciones inactivas
// @Tags Notificacion
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/notificaciones/inactivas [get]
func doc_get_api_notificaciones_inactivas() {}

// doc_get_api_notificaciones_globales godoc
// @Summary Listar notificaciones globales
// @Tags Notificacion
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/notificaciones/globales [get]
func doc_get_api_notificaciones_globales() {}

// doc_get_api_notificaciones_rango_fecha godoc
// @Summary Notificaciones por rango de fecha
// @Tags Notificacion
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/notificaciones/rango-fecha [get]
func doc_get_api_notificaciones_rango_fecha() {}

// doc_get_api_notificaciones_no_leidas_usuario_usuario_id godoc
// @Summary No leídas por usuario
// @Tags Notificacion
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/notificaciones/no-leidas/usuario/{usuario_id} [get]
func doc_get_api_notificaciones_no_leidas_usuario_usuario_id() {}

// doc_post_api_alertas godoc
// @Summary Crear alerta de usuario
// @Tags AlertaUsuario
// @Accept json
// @Produce json
// @Param body body CreateAlertaRequest true "Payload JSON"
// @Success 200 {object} map[string]interface{}
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security BearerAuth
// @Router /api/alertas [post]
func doc_post_api_alertas() {}

// doc_get_api_alertas godoc
// @Summary Listar mis alertas
// @Tags AlertaUsuario
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security BearerAuth
// @Router /api/alertas [get]
func doc_get_api_alertas() {}

// doc_put_api_alertas_id_leida godoc
// @Summary Marcar alerta como leída
// @Tags AlertaUsuario
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security BearerAuth
// @Router /api/alertas/{id}/leida [put]
func doc_put_api_alertas_id_leida() {}

// doc_get_api_usuarios godoc
// @Summary Listar usuarios
// @Tags Usuario
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/usuarios [get]
func doc_get_api_usuarios() {}

// doc_get_api_usuarios_id godoc
// @Summary Obtener usuario por ID
// @Tags Usuario
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/usuarios/{id} [get]
func doc_get_api_usuarios_id() {}

// doc_delete_api_usuarios_id godoc
// @Summary Eliminar usuario
// @Tags Usuario
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/usuarios/{id} [delete]
func doc_delete_api_usuarios_id() {}

// doc_post_api_empleados_login godoc
// @Summary Login de empleado
// @Tags Empleado
// @Accept json
// @Produce json
// @Param body body LoginUsuarioRequest true "Payload JSON"
// @Success 200 {object} map[string]interface{}
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/empleados/login [post]
func doc_post_api_empleados_login() {}

// doc_post_api_empleados godoc
// @Summary Crear empleado
// @Tags Empleado
// @Accept json
// @Produce json
// @Param body body CreateUsuarioRequest true "Payload JSON"
// @Success 200 {object} map[string]interface{}
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security BearerAuth
// @Router /api/empleados/ [post]
func doc_post_api_empleados() {}

// doc_get_api_empleados godoc
// @Summary Listar empleados
// @Tags Empleado
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security BearerAuth
// @Router /api/empleados/ [get]
func doc_get_api_empleados() {}

// doc_get_api_empleados_id godoc
// @Summary Obtener empleado por ID
// @Tags Empleado
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security BearerAuth
// @Router /api/empleados/{id} [get]
func doc_get_api_empleados_id() {}

// doc_delete_api_empleados_id godoc
// @Summary Eliminar empleado
// @Tags Empleado
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security BearerAuth
// @Router /api/empleados/{id} [delete]
func doc_delete_api_empleados_id() {}

// doc_post_api_roles godoc
// @Summary Crear rol
// @Tags Rol
// @Accept json
// @Produce json
// @Param body body map[string]interface{} true "Payload JSON"
// @Success 200 {object} map[string]interface{}
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security BearerAuth
// @Router /api/roles [post]
func doc_post_api_roles() {}

// doc_get_api_roles godoc
// @Summary Listar roles
// @Tags Rol
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security BearerAuth
// @Router /api/roles [get]
func doc_get_api_roles() {}

// doc_put_api_roles_id godoc
// @Summary Actualizar rol
// @Tags Rol
// @Accept json
// @Produce json
// @Param body body map[string]interface{} true "Payload JSON"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security BearerAuth
// @Router /api/roles/{id} [put]
func doc_put_api_roles_id() {}

// doc_delete_api_roles_id godoc
// @Summary Eliminar rol
// @Tags Rol
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security BearerAuth
// @Router /api/roles/{id} [delete]
func doc_delete_api_roles_id() {}

// doc_post_api_anomalias godoc
// @Summary Crear anomalía
// @Tags Anomalia
// @Accept json
// @Produce json
// @Param body body CreateAnomaliaRequest true "Payload JSON"
// @Success 200 {object} map[string]interface{}
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/anomalias/ [post]
func doc_post_api_anomalias() {}

// doc_get_api_anomalias godoc
// @Summary Listar anomalías
// @Tags Anomalia
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/anomalias/ [get]
func doc_get_api_anomalias() {}

// doc_get_api_anomalias_id godoc
// @Summary Obtener anomalía por ID
// @Tags Anomalia
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/anomalias/{id} [get]
func doc_get_api_anomalias_id() {}

// doc_put_api_anomalias_id godoc
// @Summary Actualizar anomalía
// @Tags Anomalia
// @Accept json
// @Produce json
// @Param body body UpdateAnomaliaRequest true "Payload JSON"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/anomalias/{id} [put]
func doc_put_api_anomalias_id() {}

// doc_delete_api_anomalias_id godoc
// @Summary Eliminar anomalía
// @Tags Anomalia
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/anomalias/{id} [delete]
func doc_delete_api_anomalias_id() {}

// doc_get_api_anomalias_punto_puntoId godoc
// @Summary Anomalías por punto
// @Tags Anomalia
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/anomalias/punto/{puntoId} [get]
func doc_get_api_anomalias_punto_puntoId() {}

// doc_get_api_anomalias_chofer_choferId godoc
// @Summary Anomalías por chofer
// @Tags Anomalia
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/anomalias/chofer/{choferId} [get]
func doc_get_api_anomalias_chofer_choferId() {}

// doc_get_api_anomalias_estado godoc
// @Summary Anomalías por estado
// @Tags Anomalia
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/anomalias/estado [get]
func doc_get_api_anomalias_estado() {}

// doc_get_api_anomalias_tipo godoc
// @Summary Anomalías por tipo
// @Tags Anomalia
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/anomalias/tipo [get]
func doc_get_api_anomalias_tipo() {}

// doc_get_api_anomalias_por_fecha godoc
// @Summary Anomalías por fecha
// @Tags Anomalia
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/anomalias/por-fecha [get]
func doc_get_api_anomalias_por_fecha() {}

// doc_post_api_incidencias godoc
// @Summary Crear incidencia
// @Tags Incidencia
// @Accept json
// @Produce json
// @Param body body CreateIncidenciaRequest true "Payload JSON"
// @Success 200 {object} map[string]interface{}
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/incidencias/ [post]
func doc_post_api_incidencias() {}

// doc_get_api_incidencias godoc
// @Summary Listar incidencias
// @Tags Incidencia
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/incidencias/ [get]
func doc_get_api_incidencias() {}

// doc_get_api_incidencias_id godoc
// @Summary Obtener incidencia por ID
// @Tags Incidencia
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/incidencias/{id} [get]
func doc_get_api_incidencias_id() {}

// doc_put_api_incidencias_id godoc
// @Summary Actualizar incidencia
// @Tags Incidencia
// @Accept json
// @Produce json
// @Param body body CreateIncidenciaRequest true "Payload JSON"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/incidencias/{id} [put]
func doc_put_api_incidencias_id() {}

// doc_delete_api_incidencias_id godoc
// @Summary Eliminar incidencia
// @Tags Incidencia
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/incidencias/{id} [delete]
func doc_delete_api_incidencias_id() {}

// doc_get_api_incidencias_conductor_conductor_id godoc
// @Summary Incidencias por conductor
// @Tags Incidencia
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/incidencias/conductor/{conductor_id} [get]
func doc_get_api_incidencias_conductor_conductor_id() {}

// doc_get_api_incidencias_punto_punto_recoleccion_id godoc
// @Summary Incidencias por punto
// @Tags Incidencia
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/incidencias/punto/{punto_recoleccion_id} [get]
func doc_get_api_incidencias_punto_punto_recoleccion_id() {}

// doc_get_api_incidencias_fecha godoc
// @Summary Incidencias por fecha
// @Tags Incidencia
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/incidencias/fecha [get]
func doc_get_api_incidencias_fecha() {}

// doc_post_api_reportes_conductor godoc
// @Summary Crear reporte de conductor
// @Tags ReporteConductor
// @Accept json
// @Produce json
// @Param body body CreateReporteConductorRequest true "Payload JSON"
// @Success 200 {object} map[string]interface{}
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/reportes-conductor/ [post]
func doc_post_api_reportes_conductor() {}

// doc_get_api_reportes_conductor godoc
// @Summary Listar reportes de conductor
// @Tags ReporteConductor
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/reportes-conductor/ [get]
func doc_get_api_reportes_conductor() {}

// doc_get_api_reportes_conductor_id godoc
// @Summary Obtener reporte por ID
// @Tags ReporteConductor
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/reportes-conductor/{id} [get]
func doc_get_api_reportes_conductor_id() {}

// doc_put_api_reportes_conductor_id godoc
// @Summary Actualizar reporte de conductor
// @Tags ReporteConductor
// @Accept json
// @Produce json
// @Param body body CreateReporteConductorRequest true "Payload JSON"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/reportes-conductor/{id} [put]
func doc_put_api_reportes_conductor_id() {}

// doc_delete_api_reportes_conductor_id godoc
// @Summary Eliminar reporte de conductor
// @Tags ReporteConductor
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/reportes-conductor/{id} [delete]
func doc_delete_api_reportes_conductor_id() {}

// doc_get_api_reportes_conductor_camion_camion_id godoc
// @Summary Reportes por camión
// @Tags ReporteConductor
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/reportes-conductor/camion/{camion_id} [get]
func doc_get_api_reportes_conductor_camion_camion_id() {}

// doc_get_api_reportes_conductor_conductor_conductor_id godoc
// @Summary Reportes por conductor
// @Tags ReporteConductor
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/reportes-conductor/conductor/{conductor_id} [get]
func doc_get_api_reportes_conductor_conductor_conductor_id() {}

// doc_get_api_reportes_conductor_ruta_ruta_id godoc
// @Summary Reportes por ruta
// @Tags ReporteConductor
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/reportes-conductor/ruta/{ruta_id} [get]
func doc_get_api_reportes_conductor_ruta_ruta_id() {}

// doc_get_api_reportes_conductor_fecha godoc
// @Summary Reportes por fecha
// @Tags ReporteConductor
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/reportes-conductor/fecha [get]
func doc_get_api_reportes_conductor_fecha() {}

// doc_post_api_registros_mantenimiento godoc
// @Summary Crear registro de mantenimiento
// @Tags RegistroMantenimiento
// @Accept json
// @Produce json
// @Param body body CreateRegistroMantenimientoRequest true "Payload JSON"
// @Success 200 {object} map[string]interface{}
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/registros-mantenimiento/ [post]
func doc_post_api_registros_mantenimiento() {}

// doc_get_api_registros_mantenimiento godoc
// @Summary Listar registros de mantenimiento
// @Tags RegistroMantenimiento
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/registros-mantenimiento/ [get]
func doc_get_api_registros_mantenimiento() {}

// doc_get_api_registros_mantenimiento_id godoc
// @Summary Obtener registro por ID
// @Tags RegistroMantenimiento
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/registros-mantenimiento/{id} [get]
func doc_get_api_registros_mantenimiento_id() {}

// doc_put_api_registros_mantenimiento_id godoc
// @Summary Actualizar registro de mantenimiento
// @Tags RegistroMantenimiento
// @Accept json
// @Produce json
// @Param body body CreateRegistroMantenimientoRequest true "Payload JSON"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/registros-mantenimiento/{id} [put]
func doc_put_api_registros_mantenimiento_id() {}

// doc_delete_api_registros_mantenimiento_id godoc
// @Summary Eliminar registro de mantenimiento
// @Tags RegistroMantenimiento
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/registros-mantenimiento/{id} [delete]
func doc_delete_api_registros_mantenimiento_id() {}

// doc_get_api_registros_mantenimiento_alerta_alerta_id godoc
// @Summary Registros por alerta
// @Tags RegistroMantenimiento
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/registros-mantenimiento/alerta/{alerta_id} [get]
func doc_get_api_registros_mantenimiento_alerta_alerta_id() {}

// doc_get_api_registros_mantenimiento_camion_camion_id godoc
// @Summary Registros por camión
// @Tags RegistroMantenimiento
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/registros-mantenimiento/camion/{camion_id} [get]
func doc_get_api_registros_mantenimiento_camion_camion_id() {}

// doc_get_api_registros_mantenimiento_coordinador_coordinador_id godoc
// @Summary Registros por coordinador
// @Tags RegistroMantenimiento
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/registros-mantenimiento/coordinador/{coordinador_id} [get]
func doc_get_api_registros_mantenimiento_coordinador_coordinador_id() {}

// doc_get_api_registros_mantenimiento_fecha godoc
// @Summary Registros por fecha
// @Tags RegistroMantenimiento
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/registros-mantenimiento/fecha [get]
func doc_get_api_registros_mantenimiento_fecha() {}

// doc_post_api_reportes_falla_critica godoc
// @Summary Crear reporte de falla crítica
// @Tags ReporteFallaCritica
// @Accept json
// @Produce json
// @Param body body CreateReporteFallaCriticaRequest true "Payload JSON"
// @Success 200 {object} map[string]interface{}
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/reportes-falla-critica/ [post]
func doc_post_api_reportes_falla_critica() {}

// doc_get_api_reportes_falla_critica godoc
// @Summary Listar reportes de falla crítica
// @Tags ReporteFallaCritica
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/reportes-falla-critica/ [get]
func doc_get_api_reportes_falla_critica() {}

// doc_get_api_reportes_falla_critica_id godoc
// @Summary Obtener reporte por ID
// @Tags ReporteFallaCritica
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/reportes-falla-critica/{id} [get]
func doc_get_api_reportes_falla_critica_id() {}

// doc_put_api_reportes_falla_critica_id godoc
// @Summary Actualizar reporte de falla crítica
// @Tags ReporteFallaCritica
// @Accept json
// @Produce json
// @Param body body CreateReporteFallaCriticaRequest true "Payload JSON"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/reportes-falla-critica/{id} [put]
func doc_put_api_reportes_falla_critica_id() {}

// doc_delete_api_reportes_falla_critica_id godoc
// @Summary Eliminar reporte de falla crítica
// @Tags ReporteFallaCritica
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/reportes-falla-critica/{id} [delete]
func doc_delete_api_reportes_falla_critica_id() {}

// doc_get_api_reportes_falla_critica_camion_camionId godoc
// @Summary Reportes por camión
// @Tags ReporteFallaCritica
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/reportes-falla-critica/camion/{camionId} [get]
func doc_get_api_reportes_falla_critica_camion_camionId() {}

// doc_get_api_reportes_falla_critica_conductor_conductorId godoc
// @Summary Reportes por conductor
// @Tags ReporteFallaCritica
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/reportes-falla-critica/conductor/{conductorId} [get]
func doc_get_api_reportes_falla_critica_conductor_conductorId() {}

// doc_get_api_reportes_falla_critica_por_fecha godoc
// @Summary Reportes por fecha
// @Tags ReporteFallaCritica
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/reportes-falla-critica/por-fecha [get]
func doc_get_api_reportes_falla_critica_por_fecha() {}

// doc_post_api_reportes_mantenimiento_generado godoc
// @Summary Crear reporte de mantenimiento generado
// @Tags ReporteMantenimientoGenerado
// @Accept json
// @Produce json
// @Param body body map[string]interface{} true "Payload JSON"
// @Success 200 {object} map[string]interface{}
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/reportes-mantenimiento-generado/ [post]
func doc_post_api_reportes_mantenimiento_generado() {}

// doc_get_api_reportes_mantenimiento_generado godoc
// @Summary Listar reportes generados
// @Tags ReporteMantenimientoGenerado
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/reportes-mantenimiento-generado/ [get]
func doc_get_api_reportes_mantenimiento_generado() {}

// doc_get_api_reportes_mantenimiento_generado_id godoc
// @Summary Obtener reporte generado por ID
// @Tags ReporteMantenimientoGenerado
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/reportes-mantenimiento-generado/{id} [get]
func doc_get_api_reportes_mantenimiento_generado_id() {}

// doc_put_api_reportes_mantenimiento_generado_id godoc
// @Summary Actualizar reporte generado
// @Tags ReporteMantenimientoGenerado
// @Accept json
// @Produce json
// @Param body body map[string]interface{} true "Payload JSON"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/reportes-mantenimiento-generado/{id} [put]
func doc_put_api_reportes_mantenimiento_generado_id() {}

// doc_delete_api_reportes_mantenimiento_generado_id godoc
// @Summary Eliminar reporte generado
// @Tags ReporteMantenimientoGenerado
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/reportes-mantenimiento-generado/{id} [delete]
func doc_delete_api_reportes_mantenimiento_generado_id() {}

// doc_get_api_reportes_mantenimiento_generado_coordinador_coordinador_id godoc
// @Summary Reportes por coordinador
// @Tags ReporteMantenimientoGenerado
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/reportes-mantenimiento-generado/coordinador/{coordinador_id} [get]
func doc_get_api_reportes_mantenimiento_generado_coordinador_coordinador_id() {}

// doc_get_api_reportes_mantenimiento_generado_fecha godoc
// @Summary Reportes por fecha
// @Tags ReporteMantenimientoGenerado
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/reportes-mantenimiento-generado/fecha [get]
func doc_get_api_reportes_mantenimiento_generado_fecha() {}

// doc_get_api_reportes_mantenimiento_generado_fecha_generacion godoc
// @Summary Reportes por fecha de generación
// @Tags ReporteMantenimientoGenerado
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/reportes-mantenimiento-generado/fecha-generacion [get]
func doc_get_api_reportes_mantenimiento_generado_fecha_generacion() {}

// doc_post_api_seguimientos_falla_critica godoc
// @Summary Crear seguimiento de falla crítica
// @Tags SeguimientoFallaCritica
// @Accept json
// @Produce json
// @Param body body map[string]interface{} true "Payload JSON"
// @Success 200 {object} map[string]interface{}
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/seguimientos-falla-critica/ [post]
func doc_post_api_seguimientos_falla_critica() {}

// doc_get_api_seguimientos_falla_critica godoc
// @Summary Listar seguimientos
// @Tags SeguimientoFallaCritica
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/seguimientos-falla-critica/ [get]
func doc_get_api_seguimientos_falla_critica() {}

// doc_get_api_seguimientos_falla_critica_id godoc
// @Summary Obtener seguimiento por ID
// @Tags SeguimientoFallaCritica
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/seguimientos-falla-critica/{id} [get]
func doc_get_api_seguimientos_falla_critica_id() {}

// doc_put_api_seguimientos_falla_critica_id godoc
// @Summary Actualizar seguimiento
// @Tags SeguimientoFallaCritica
// @Accept json
// @Produce json
// @Param body body map[string]interface{} true "Payload JSON"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/seguimientos-falla-critica/{id} [put]
func doc_put_api_seguimientos_falla_critica_id() {}

// doc_delete_api_seguimientos_falla_critica_id godoc
// @Summary Eliminar seguimiento
// @Tags SeguimientoFallaCritica
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/seguimientos-falla-critica/{id} [delete]
func doc_delete_api_seguimientos_falla_critica_id() {}

// doc_get_api_seguimientos_falla_critica_falla_fallaId godoc
// @Summary Seguimientos por falla
// @Tags SeguimientoFallaCritica
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/seguimientos-falla-critica/falla/{fallaId} [get]
func doc_get_api_seguimientos_falla_critica_falla_fallaId() {}

// doc_get_api_seguimientos_falla_critica_por_fecha godoc
// @Summary Seguimientos por fecha
// @Tags SeguimientoFallaCritica
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/seguimientos-falla-critica/por-fecha [get]
func doc_get_api_seguimientos_falla_critica_por_fecha() {}

// doc_post_api_tipos_mantenimiento godoc
// @Summary Crear tipo de mantenimiento
// @Tags TipoMantenimiento
// @Accept json
// @Produce json
// @Param body body CreateTipoMantenimientoRequest true "Payload JSON"
// @Success 200 {object} map[string]interface{}
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/tipos-mantenimiento/ [post]
func doc_post_api_tipos_mantenimiento() {}

// doc_get_api_tipos_mantenimiento godoc
// @Summary Listar tipos de mantenimiento
// @Tags TipoMantenimiento
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/tipos-mantenimiento/ [get]
func doc_get_api_tipos_mantenimiento() {}

// doc_get_api_tipos_mantenimiento_id godoc
// @Summary Obtener tipo por ID
// @Tags TipoMantenimiento
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/tipos-mantenimiento/{id} [get]
func doc_get_api_tipos_mantenimiento_id() {}

// doc_put_api_tipos_mantenimiento_id godoc
// @Summary Actualizar tipo de mantenimiento
// @Tags TipoMantenimiento
// @Accept json
// @Produce json
// @Param body body CreateTipoMantenimientoRequest true "Payload JSON"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/tipos-mantenimiento/{id} [put]
func doc_put_api_tipos_mantenimiento_id() {}

// doc_delete_api_tipos_mantenimiento_id godoc
// @Summary Eliminar tipo de mantenimiento
// @Tags TipoMantenimiento
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/tipos-mantenimiento/{id} [delete]
func doc_delete_api_tipos_mantenimiento_id() {}

// doc_post_domicilios godoc
// @Summary Crear domicilio (compat)
// @Tags Domicilio
// @Accept json
// @Produce json
// @Param body body domain.CreateDomicilioRequest true "Payload JSON"
// @Success 200 {object} map[string]interface{}
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security BearerAuth
// @Router /domicilios [post]
func doc_post_domicilios() {}

// doc_get_domicilios godoc
// @Summary Listar domicilios (compat)
// @Tags Domicilio
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security BearerAuth
// @Router /domicilios [get]
func doc_get_domicilios() {}

// doc_get_domicilios_id godoc
// @Summary Obtener domicilio (compat)
// @Tags Domicilio
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security BearerAuth
// @Router /domicilios/{id} [get]
func doc_get_domicilios_id() {}

// doc_put_domicilios_id godoc
// @Summary Actualizar domicilio (compat)
// @Tags Domicilio
// @Accept json
// @Produce json
// @Param body body domain.CreateDomicilioRequest true "Payload JSON"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security BearerAuth
// @Router /domicilios/{id} [put]
func doc_put_domicilios_id() {}

// doc_delete_domicilios_id godoc
// @Summary Eliminar domicilio (compat)
// @Tags Domicilio
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security BearerAuth
// @Router /domicilios/{id} [delete]
func doc_delete_domicilios_id() {}

// doc_get_colonias godoc
// @Summary Listar colonias (compat)
// @Tags Colonia
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /colonias [get]
func doc_get_colonias() {}

// doc_get_colonias_id godoc
// @Summary Obtener colonia (compat)
// @Tags Colonia
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /colonias/{id} [get]
func doc_get_colonias_id() {}

// doc_post_colonias godoc
// @Summary Crear colonia (compat)
// @Tags Colonia
// @Accept json
// @Produce json
// @Param body body domain.CreateColoniaRequest true "Payload JSON"
// @Success 200 {object} map[string]interface{}
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security BearerAuth
// @Router /colonias [post]
func doc_post_colonias() {}

// doc_put_colonias_id godoc
// @Summary Actualizar colonia (compat)
// @Tags Colonia
// @Accept json
// @Produce json
// @Param body body domain.UpdateColoniaRequest true "Payload JSON"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security BearerAuth
// @Router /colonias/{id} [put]
func doc_put_colonias_id() {}

// doc_delete_colonias_id godoc
// @Summary Eliminar colonia (compat)
// @Tags Colonia
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security BearerAuth
// @Router /colonias/{id} [delete]
func doc_delete_colonias_id() {}

