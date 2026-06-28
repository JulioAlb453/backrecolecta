package main

// Modelos de request documentados en Swagger (alineados con los handlers reales).

// CreateCamionRequest POST/PUT /api/camion/
type CreateCamionRequest struct {
	Placa            string `json:"placa" example:"TSU-001"`
	Modelo           string `json:"modelo" example:"International 2020"`
	TipoCamionID     int32  `json:"tipo_camion_id" example:"1"`
	EsRentado        bool   `json:"es_rentado" example:"false"`
	DisponibilidadID int32  `json:"disponibilidad_id" example:"1"`
}

// CreateEstadoCamionRequest POST/PUT /api/estado-camion/
type CreateEstadoCamionRequest struct {
	CamionID      int32  `json:"camion_id" example:"1"`
	Estado        string `json:"estado" example:"en_ruta"`
	Timestamp     string `json:"timestamp" example:"2024-06-25T08:00:00Z"`
	Observaciones string `json:"observaciones" example:"Salida del depósito"`
}

// CreateTipoCamionRequest POST /api/tipo-camion/
type CreateTipoCamionRequest struct {
	Nombre      string `json:"nombre" example:"Compactador"`
	Descripcion string `json:"descripcion" example:"Camión compactador de 15 m³"`
}

// CreateRutaRequest POST/PUT /api/rutas/
type CreateRutaRequest struct {
	Nombre      string `json:"nombre" example:"Ruta Centro Norte"`
	Descripcion string `json:"descripcion" example:"Recolección zona centro y norte"`
	JsonRuta    string `json:"json_ruta" example:"{\"puntos\":[{\"lat\":16.6253,\"lng\":-93.1021}]}"`
	CreatedAt   string `json:"created_at,omitempty" example:"2024-06-25T08:00:00Z"`
}

// CreatePuntoRecoleccionRequest POST/PUT /api/puntos-recoleccion/
type CreatePuntoRecoleccionRequest struct {
	RutaID int32  `json:"ruta_id" example:"1"`
	CP     string `json:"cp" example:"29150"`
}

// CreateRellenoSanitarioRequest POST/PUT /api/relleno-sanitario/
type CreateRellenoSanitarioRequest struct {
	Nombre             string  `json:"nombre" example:"Relleno Suchiapa"`
	Direccion          string  `json:"direccion" example:"Carretera Suchiapa Km 3"`
	EsRentado          bool    `json:"es_rentado" example:"false"`
	CapacidadToneladas float64 `json:"capacidad_toneladas" example:"5000"`
}

// CreateRutaCamionRequest POST/PUT /api/ruta-camion/
type CreateRutaCamionRequest struct {
	RutaID   int32  `json:"ruta_id" example:"1"`
	CamionID int32  `json:"camion_id" example:"1"`
	Fecha    string `json:"fecha" example:"2024-06-25T08:00:00Z"`
}

// CreateHistorialAsignacionRequest POST/PUT /api/historial-asignacion/
type CreateHistorialAsignacionRequest struct {
	IDChofer        *int   `json:"id_chofer" example:"7"`
	IDCamion        *int   `json:"id_camion" example:"1"`
	FechaAsignacion string `json:"fecha_asignacion" example:"2024-06-25T07:00:00Z"`
	FechaBaja       string `json:"fecha_baja,omitempty" example:""`
	Eliminado       bool   `json:"eliminado" example:"false"`
}

// CreateRegistroVaciadoRequest POST /api/registro-vaciado/
type CreateRegistroVaciadoRequest struct {
	RellenoID    int32  `json:"relleno_id" example:"1"`
	RutaCamionID int32  `json:"ruta_camion_id" example:"1"`
	Hora         string `json:"hora" example:"2024-06-25T14:30:00Z"`
}

// CreateUsuarioRequest POST /api/usuarios y /api/empleados/
type CreateUsuarioRequest struct {
	Nombre   string `json:"nombre" example:"Juan Conductor"`
	Email    string `json:"email" example:"conductor@recolecta.mx"`
	Password string `json:"password" example:"Conductor123"`
	RolID    int    `json:"rol_id" example:"4"`
}

// LoginUsuarioRequest POST /api/usuarios/login y /api/empleados/login
type LoginUsuarioRequest struct {
	Email    string `json:"email" example:"conductor@recolecta.mx"`
	Password string `json:"password" example:"Conductor123"`
}

// CreateAnomaliaRequest POST /api/anomalias/
type CreateAnomaliaRequest struct {
	PuntoID      *int32 `json:"punto_id" example:"1"`
	TipoAnomalia string `json:"tipo_anomalia" example:"basura_fuera_horario"`
	Descripcion  string `json:"descripcion" example:"Basura acumulada en esquina"`
	FechaReporte string `json:"fecha_reporte" example:"2024-06-25 08:30:00"`
	Estado       string `json:"estado" example:"pendiente"`
	IDChoferID   int32  `json:"id_chofer_id" example:"7"`
}

// UpdateAnomaliaRequest PUT /api/anomalias/{id}
type UpdateAnomaliaRequest struct {
	PuntoID         *int32 `json:"punto_id" example:"1"`
	TipoAnomalia    string `json:"tipo_anomalia" example:"basura_fuera_horario"`
	Descripcion     string `json:"descripcion" example:"Basura acumulada en esquina"`
	FechaReporte    string `json:"fecha_reporte" example:"2024-06-25 08:30:00"`
	FechaResolucion string `json:"fecha_resolucion,omitempty" example:"2024-06-25 12:00:00"`
	Estado          string `json:"estado" example:"resuelta"`
	IDChoferID      int32  `json:"id_chofer_id" example:"7"`
}

// CreateIncidenciaRequest POST/PUT /api/incidencias/
type CreateIncidenciaRequest struct {
	PuntoRecoleccionID *int32 `json:"punto_recoleccion_id" example:"1"`
	ConductorID        int32  `json:"conductor_id" example:"7"`
	Descripcion        string `json:"descripcion" example:"Calle bloqueada por obra"`
	JsonRuta           string `json:"json_ruta" example:"{\"desvio\":true}"`
	FechaReporte       string `json:"fecha_reporte" example:"2024-06-25T08:00:00Z"`
}

// CreateReporteConductorRequest POST/PUT /api/reportes-conductor/
type CreateReporteConductorRequest struct {
	ConductorID int32  `json:"conductor_id" example:"7"`
	CamionID    int32  `json:"camion_id" example:"1"`
	RutaID      int32  `json:"ruta_id" example:"1"`
	Descripcion string `json:"descripcion" example:"Ruta completada con retraso de 20 min"`
}

// CreateRegistroMantenimientoRequest POST/PUT /api/registros-mantenimiento/
type CreateRegistroMantenimientoRequest struct {
	AlertaID                 *int32  `json:"alerta_id" example:"1"`
	CamionID                 int32   `json:"camion_id" example:"1"`
	CoordinadorID            int32   `json:"coordinador_id" example:"2"`
	MecanicoResponsable      string  `json:"mecanico_responsable" example:"Taller Municipal"`
	FechaRealizada           string  `json:"fecha_realizada" example:"2024-06-25"`
	KilometrajeMantenimiento float64 `json:"kilometraje_mantenimiento" example:"45000"`
	Observaciones            string  `json:"observaciones" example:"Cambio de aceite y filtros"`
}

// CreateReporteFallaCriticaRequest POST/PUT /api/reportes-falla-critica/
type CreateReporteFallaCriticaRequest struct {
	CamionID    int32  `json:"camion_id" example:"1"`
	ConductorID int32  `json:"conductor_id" example:"7"`
	Descripcion string `json:"descripcion" example:"Falla en sistema hidráulico"`
}

// CreateTipoMantenimientoRequest POST/PUT /api/tipos-mantenimiento/
type CreateTipoMantenimientoRequest struct {
	Nombre    string `json:"nombre" example:"Preventivo"`
	Categoria string `json:"categoria" example:"motor"`
}

// CreateNotificacionRequest POST /api/notificaciones/
type CreateNotificacionRequest struct {
	UsuarioID                  *int32 `json:"usuario_id" example:"100"`
	Tipo                       string `json:"tipo" example:"info"`
	Titulo                     string `json:"titulo" example:"Camión cerca"`
	Mensaje                    string `json:"mensaje" example:"El camión está a 2 cuadras"`
	IDCamionRelacionado        *int32 `json:"id_camion_relacionado,omitempty" example:"1"`
	IDFallaRelacionado         *int32 `json:"id_falla_relacionado,omitempty"`
	IDMantenimientoRelacionado *int32 `json:"id_mantenimiento_relacionado,omitempty"`
	CreadoPor                  *int32 `json:"creado_por" example:"1"`
}

// UpdateNotificacionRequest PUT /api/notificaciones/{id}
type UpdateNotificacionRequest struct {
	UsuarioID                  *int32 `json:"usuario_id" example:"100"`
	Tipo                       string `json:"tipo" example:"info"`
	Titulo                     string `json:"titulo" example:"Camión cerca"`
	Mensaje                    string `json:"mensaje" example:"El camión está a 2 cuadras"`
	Activa                     bool   `json:"activa" example:"true"`
	IDCamionRelacionado        *int32 `json:"id_camion_relacionado,omitempty"`
	IDFallaRelacionado         *int32 `json:"id_falla_relacionado,omitempty"`
	IDMantenimientoRelacionado *int32 `json:"id_mantenimiento_relacionado,omitempty"`
	CreadoPor                  *int32 `json:"creado_por"`
}

// NotificacionEmergenciaRequest POST /api/notificaciones/emergencia
type NotificacionEmergenciaRequest struct {
	UsuarioID  *int32 `json:"usuario_id" example:"100"`
	Titulo     string `json:"titulo" example:"Emergencia"`
	Mensaje    string `json:"mensaje" example:"Contingencia en ruta 3"`
	CamionID   *int32 `json:"camion_id" example:"1"`
	CreadoPor  *int32 `json:"creado_por" example:"1"`
}

// NotificacionFallaRequest POST /api/notificaciones/falla
type NotificacionFallaRequest struct {
	UsuarioID  *int32 `json:"usuario_id" example:"100"`
	Titulo     string `json:"titulo" example:"Falla reportada"`
	Mensaje    string `json:"mensaje" example:"Camión detenido"`
	CamionID   *int32 `json:"camion_id" example:"1"`
	FallaID    *int32 `json:"falla_id" example:"1"`
	CreadoPor  *int32 `json:"creado_por" example:"1"`
}

// NotificacionMantenimientoRequest POST /api/notificaciones/mantenimiento
type NotificacionMantenimientoRequest struct {
	UsuarioID        *int32 `json:"usuario_id" example:"100"`
	Titulo           string `json:"titulo" example:"Mantenimiento programado"`
	Mensaje          string `json:"mensaje" example:"Camión en taller"`
	CamionID         *int32 `json:"camion_id" example:"1"`
	MantenimientoID  *int32 `json:"mantenimiento_id" example:"1"`
	CreadoPor        *int32 `json:"creado_por" example:"1"`
}

// NotificarUsuarioRequest POST /api/notificaciones/notificar
type NotificarUsuarioRequest struct {
	CreadorID      int32  `json:"creador_id" example:"1"`
	DestinatarioID int32  `json:"destinatario_id" example:"100"`
	Tipo           string `json:"tipo" example:"info"`
	Titulo         string `json:"titulo" example:"Aviso"`
	Mensaje        string `json:"mensaje" example:"Recolección reprogramada"`
}

// NotificarMultiplesRequest POST /api/notificaciones/enviar-multiples
type NotificarMultiplesRequest struct {
	UsuarioIDs []int32 `json:"usuario_ids" example:"100,101,102"`
	Tipo       string  `json:"tipo" example:"info"`
	Titulo     string  `json:"titulo" example:"Aviso general"`
	Mensaje    string  `json:"mensaje" example:"Servicio suspendido temporalmente"`
	CreadoPor  *int32  `json:"creado_por" example:"1"`
	CamionID   *int32  `json:"camion_id,omitempty"`
	FallaID    *int32  `json:"falla_id,omitempty"`
}

// NotificarTodosRequest POST /api/notificaciones/enviar-todos
type NotificarTodosRequest struct {
	Tipo      string `json:"tipo" example:"info"`
	Titulo    string `json:"titulo" example:"Aviso municipal"`
	Mensaje   string `json:"mensaje" example:"Recolección especial este sábado"`
	CreadoPor *int32 `json:"creado_por" example:"1"`
}

// CreateAlertaRequest POST /api/alertas
type CreateAlertaRequest struct {
	Titulo    string `json:"titulo" example:"Retraso en ruta"`
	Mensaje   string `json:"mensaje" example:"El camión llegará 30 min tarde"`
	UsuarioID int    `json:"usuario_id" example:"100"`
}
