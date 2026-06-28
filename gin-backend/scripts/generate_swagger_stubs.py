#!/usr/bin/env python3
"""Genera swagger_stubs.go con anotaciones @Router para todas las rutas activas."""

from __future__ import annotations

import re
from pathlib import Path

ROOT = Path(__file__).resolve().parents[1]
OUTPUT = ROOT / "swagger_stubs.go"

# (method, path, tag, summary, secured)
ROUTES: list[tuple[str, str, str, str, bool]] = [
    ("get", "/health", "Sistema", "Health check", False),

    ("post", "/api/tipo-camion/", "TipoCamion", "Crear tipo de camión", False),
    ("get", "/api/tipo-camion/", "TipoCamion", "Listar tipos de camión", False),
    ("get", "/api/tipo-camion/nombre/{nombre}", "TipoCamion", "Obtener tipo de camión por nombre", False),
    ("delete", "/api/tipo-camion/{id}", "TipoCamion", "Eliminar tipo de camión", False),

    ("post", "/api/camion/", "Camion", "Crear camión", False),
    ("get", "/api/camion/", "Camion", "Listar camiones", False),
    ("get", "/api/camion/{id}", "Camion", "Obtener camión por ID", False),
    ("put", "/api/camion/{id}", "Camion", "Actualizar camión", False),
    ("delete", "/api/camion/{id}", "Camion", "Eliminar camión", False),
    ("get", "/api/camion/placa/{placa}", "Camion", "Obtener camión por placa", False),
    ("get", "/api/camion/modelo", "Camion", "Buscar camión por modelo", False),

    ("post", "/api/estado-camion/", "EstadoCamion", "Crear estado de camión", False),
    ("get", "/api/estado-camion/", "EstadoCamion", "Listar estados de camión", False),
    ("get", "/api/estado-camion/camion/{id}", "EstadoCamion", "Obtener estado por camión", False),
    ("put", "/api/estado-camion/{id}", "EstadoCamion", "Actualizar estado de camión", False),
    ("delete", "/api/estado-camion/{id}", "EstadoCamion", "Eliminar estado de camión", False),

    ("post", "/api/historial-asignacion/", "HistorialAsignacion", "Crear historial de asignación", False),
    ("get", "/api/historial-asignacion/", "HistorialAsignacion", "Listar historial de asignación", False),
    ("get", "/api/historial-asignacion/{id}", "HistorialAsignacion", "Obtener historial por ID", False),
    ("put", "/api/historial-asignacion/{id}", "HistorialAsignacion", "Actualizar historial de asignación", False),
    ("delete", "/api/historial-asignacion/{id}", "HistorialAsignacion", "Eliminar historial de asignación", False),
    ("get", "/api/historial-asignacion/camion/{camionId}", "HistorialAsignacion", "Historial por camión", False),
    ("get", "/api/historial-asignacion/chofer/{choferId}", "HistorialAsignacion", "Historial por chofer", False),
    ("get", "/api/historial-asignacion/activo/camion/{camionId}", "HistorialAsignacion", "Asignación activa por camión", False),
    ("get", "/api/historial-asignacion/activo/chofer/{choferId}", "HistorialAsignacion", "Asignación activa por chofer", False),
    ("put", "/api/historial-asignacion/baja/{id}", "HistorialAsignacion", "Dar de baja historial", False),
    ("put", "/api/historial-asignacion/cerrar/camion/{camionId}", "HistorialAsignacion", "Cerrar asignación por camión", False),
    ("put", "/api/historial-asignacion/cerrar/chofer/{choferId}", "HistorialAsignacion", "Cerrar asignación por chofer", False),

    ("post", "/api/rutas/", "Ruta", "Crear ruta", False),
    ("get", "/api/rutas/", "Ruta", "Listar rutas", False),
    ("get", "/api/rutas/{id}", "Ruta", "Obtener ruta por ID", False),
    ("put", "/api/rutas/{id}", "Ruta", "Actualizar ruta", False),
    ("delete", "/api/rutas/{id}", "Ruta", "Eliminar ruta", False),
    ("get", "/api/rutas/activas", "Ruta", "Listar rutas activas", False),

    ("post", "/api/puntos-recoleccion/", "PuntoRecoleccion", "Crear punto de recolección", False),
    ("get", "/api/puntos-recoleccion/", "PuntoRecoleccion", "Listar puntos de recolección", False),
    ("get", "/api/puntos-recoleccion/{id}", "PuntoRecoleccion", "Obtener punto por ID", False),
    ("get", "/api/puntos-recoleccion/ruta/{rutaId}", "PuntoRecoleccion", "Puntos por ruta", False),
    ("put", "/api/puntos-recoleccion/{id}", "PuntoRecoleccion", "Actualizar punto de recolección", False),
    ("delete", "/api/puntos-recoleccion/{id}", "PuntoRecoleccion", "Eliminar punto de recolección", False),

    ("post", "/api/relleno-sanitario/", "RellenoSanitario", "Crear relleno sanitario", False),
    ("get", "/api/relleno-sanitario/", "RellenoSanitario", "Listar rellenos sanitarios", False),
    ("get", "/api/relleno-sanitario/{id}", "RellenoSanitario", "Obtener relleno por ID", False),
    ("put", "/api/relleno-sanitario/{id}", "RellenoSanitario", "Actualizar relleno sanitario", False),
    ("delete", "/api/relleno-sanitario/{id}", "RellenoSanitario", "Eliminar relleno sanitario", False),
    ("get", "/api/relleno-sanitario/buscar", "RellenoSanitario", "Buscar relleno sanitario", False),
    ("get", "/api/relleno-sanitario/exists/{id}", "RellenoSanitario", "Verificar existencia de relleno", False),

    ("post", "/api/ruta-camion/", "RutaCamion", "Crear ruta-camión", False),
    ("get", "/api/ruta-camion/", "RutaCamion", "Listar rutas-camión", False),
    ("get", "/api/ruta-camion/{id}", "RutaCamion", "Obtener ruta-camión por ID", False),
    ("get", "/api/ruta-camion/camion/{camion_id}", "RutaCamion", "Rutas por camión", False),
    ("get", "/api/ruta-camion/ruta/{ruta_id}", "RutaCamion", "Camiones por ruta", False),
    ("get", "/api/ruta-camion/exists/{id}", "RutaCamion", "Verificar existencia ruta-camión", False),
    ("put", "/api/ruta-camion/{id}", "RutaCamion", "Actualizar ruta-camión", False),
    ("delete", "/api/ruta-camion/{id}", "RutaCamion", "Eliminar ruta-camión", False),

    ("post", "/api/registro-vaciado/", "RegistroVaciado", "Crear registro de vaciado", False),
    ("get", "/api/registro-vaciado/", "RegistroVaciado", "Listar registros de vaciado", False),
    ("get", "/api/registro-vaciado/{id}", "RegistroVaciado", "Obtener registro de vaciado por ID", False),
    ("get", "/api/registro-vaciado/relleno/{relleno_id}", "RegistroVaciado", "Registros por relleno", False),
    ("get", "/api/registro-vaciado/ruta-camion/{ruta_camion_id}", "RegistroVaciado", "Registros por ruta-camión", False),
    ("get", "/api/registro-vaciado/exists/{id}", "RegistroVaciado", "Verificar existencia registro vaciado", False),
    ("delete", "/api/registro-vaciado/{id}", "RegistroVaciado", "Eliminar registro de vaciado", False),

    ("post", "/api/notificaciones/", "Notificacion", "Crear notificación", False),
    ("get", "/api/notificaciones/", "Notificacion", "Listar notificaciones", False),
    ("get", "/api/notificaciones/{id}", "Notificacion", "Obtener notificación por ID", False),
    ("put", "/api/notificaciones/{id}", "Notificacion", "Actualizar notificación", False),
    ("delete", "/api/notificaciones/{id}", "Notificacion", "Eliminar notificación", False),
    ("get", "/api/notificaciones/count/usuario/{usuario_id}", "Notificacion", "Contar notificaciones por usuario", False),
    ("get", "/api/notificaciones/count/activas/usuario/{usuario_id}", "Notificacion", "Contar activas por usuario", False),
    ("get", "/api/notificaciones/count/tipo/{tipo}", "Notificacion", "Contar por tipo", False),
    ("get", "/api/notificaciones/count/camion/{camion_id}", "Notificacion", "Contar por camión", False),
    ("post", "/api/notificaciones/emergencia", "Notificacion", "Crear notificación de emergencia", False),
    ("post", "/api/notificaciones/falla", "Notificacion", "Crear notificación de falla", False),
    ("post", "/api/notificaciones/mantenimiento", "Notificacion", "Crear notificación de mantenimiento", False),
    ("post", "/api/notificaciones/notificar", "Notificacion", "Enviar notificación", False),
    ("post", "/api/notificaciones/enviar-multiples", "Notificacion", "Enviar múltiples notificaciones", False),
    ("post", "/api/notificaciones/enviar-todos", "Notificacion", "Enviar a todos", False),
    ("patch", "/api/notificaciones/{id}/marcar-leida", "Notificacion", "Marcar notificación leída", False),
    ("patch", "/api/notificaciones/{id}/reactivar", "Notificacion", "Reactivar notificación", False),
    ("patch", "/api/notificaciones/usuario/{usuario_id}/marcar-todas-leidas", "Notificacion", "Marcar todas leídas", False),
    ("get", "/api/notificaciones/usuario/{usuario_id}", "Notificacion", "Notificaciones por usuario", False),
    ("get", "/api/notificaciones/activas/usuario/{usuario_id}", "Notificacion", "Activas por usuario", False),
    ("get", "/api/notificaciones/usuario/{usuario_id}/tipo/{tipo}", "Notificacion", "Por usuario y tipo", False),
    ("get", "/api/notificaciones/camion/{camion_id}", "Notificacion", "Notificaciones por camión", False),
    ("get", "/api/notificaciones/camion/{camion_id}/tipo/{tipo}", "Notificacion", "Por camión y tipo", False),
    ("get", "/api/notificaciones/tipo/{tipo}", "Notificacion", "Notificaciones por tipo", False),
    ("get", "/api/notificaciones/creado-por/{creado_por}", "Notificacion", "Notificaciones por creador", False),
    ("get", "/api/notificaciones/falla/{falla_id}", "Notificacion", "Notificaciones por falla", False),
    ("get", "/api/notificaciones/mantenimiento/{mantenimiento_id}", "Notificacion", "Notificaciones por mantenimiento", False),
    ("get", "/api/notificaciones/activas", "Notificacion", "Listar notificaciones activas", False),
    ("get", "/api/notificaciones/inactivas", "Notificacion", "Listar notificaciones inactivas", False),
    ("get", "/api/notificaciones/globales", "Notificacion", "Listar notificaciones globales", False),
    ("get", "/api/notificaciones/rango-fecha", "Notificacion", "Notificaciones por rango de fecha", False),
    ("get", "/api/notificaciones/no-leidas/usuario/{usuario_id}", "Notificacion", "No leídas por usuario", False),

    ("post", "/api/alertas", "AlertaUsuario", "Crear alerta de usuario", True),
    ("get", "/api/alertas", "AlertaUsuario", "Listar mis alertas", True),
    ("put", "/api/alertas/{id}/leida", "AlertaUsuario", "Marcar alerta como leída", True),

    ("get", "/api/usuarios", "Usuario", "Listar usuarios", False),
    ("get", "/api/usuarios/{id}", "Usuario", "Obtener usuario por ID", False),
    ("delete", "/api/usuarios/{id}", "Usuario", "Eliminar usuario", False),

    ("post", "/api/empleados/login", "Empleado", "Login de empleado", False),
    ("post", "/api/empleados/", "Empleado", "Crear empleado", True),
    ("get", "/api/empleados/", "Empleado", "Listar empleados", True),
    ("get", "/api/empleados/{id}", "Empleado", "Obtener empleado por ID", True),
    ("delete", "/api/empleados/{id}", "Empleado", "Eliminar empleado", True),

    ("post", "/api/roles", "Rol", "Crear rol", True),
    ("get", "/api/roles", "Rol", "Listar roles", True),
    ("put", "/api/roles/{id}", "Rol", "Actualizar rol", True),
    ("delete", "/api/roles/{id}", "Rol", "Eliminar rol", True),

    ("post", "/api/anomalias/", "Anomalia", "Crear anomalía", False),
    ("get", "/api/anomalias/", "Anomalia", "Listar anomalías", False),
    ("get", "/api/anomalias/{id}", "Anomalia", "Obtener anomalía por ID", False),
    ("put", "/api/anomalias/{id}", "Anomalia", "Actualizar anomalía", False),
    ("delete", "/api/anomalias/{id}", "Anomalia", "Eliminar anomalía", False),
    ("get", "/api/anomalias/punto/{puntoId}", "Anomalia", "Anomalías por punto", False),
    ("get", "/api/anomalias/chofer/{choferId}", "Anomalia", "Anomalías por chofer", False),
    ("get", "/api/anomalias/estado", "Anomalia", "Anomalías por estado", False),
    ("get", "/api/anomalias/tipo", "Anomalia", "Anomalías por tipo", False),
    ("get", "/api/anomalias/por-fecha", "Anomalia", "Anomalías por fecha", False),

    ("post", "/api/incidencias/", "Incidencia", "Crear incidencia", False),
    ("get", "/api/incidencias/", "Incidencia", "Listar incidencias", False),
    ("get", "/api/incidencias/{id}", "Incidencia", "Obtener incidencia por ID", False),
    ("put", "/api/incidencias/{id}", "Incidencia", "Actualizar incidencia", False),
    ("delete", "/api/incidencias/{id}", "Incidencia", "Eliminar incidencia", False),
    ("get", "/api/incidencias/conductor/{conductor_id}", "Incidencia", "Incidencias por conductor", False),
    ("get", "/api/incidencias/punto/{punto_recoleccion_id}", "Incidencia", "Incidencias por punto", False),
    ("get", "/api/incidencias/fecha", "Incidencia", "Incidencias por fecha", False),

    ("post", "/api/reportes-conductor/", "ReporteConductor", "Crear reporte de conductor", False),
    ("get", "/api/reportes-conductor/", "ReporteConductor", "Listar reportes de conductor", False),
    ("get", "/api/reportes-conductor/{id}", "ReporteConductor", "Obtener reporte por ID", False),
    ("put", "/api/reportes-conductor/{id}", "ReporteConductor", "Actualizar reporte de conductor", False),
    ("delete", "/api/reportes-conductor/{id}", "ReporteConductor", "Eliminar reporte de conductor", False),
    ("get", "/api/reportes-conductor/camion/{camion_id}", "ReporteConductor", "Reportes por camión", False),
    ("get", "/api/reportes-conductor/conductor/{conductor_id}", "ReporteConductor", "Reportes por conductor", False),
    ("get", "/api/reportes-conductor/ruta/{ruta_id}", "ReporteConductor", "Reportes por ruta", False),
    ("get", "/api/reportes-conductor/fecha", "ReporteConductor", "Reportes por fecha", False),

    ("post", "/api/registros-mantenimiento/", "RegistroMantenimiento", "Crear registro de mantenimiento", False),
    ("get", "/api/registros-mantenimiento/", "RegistroMantenimiento", "Listar registros de mantenimiento", False),
    ("get", "/api/registros-mantenimiento/{id}", "RegistroMantenimiento", "Obtener registro por ID", False),
    ("put", "/api/registros-mantenimiento/{id}", "RegistroMantenimiento", "Actualizar registro de mantenimiento", False),
    ("delete", "/api/registros-mantenimiento/{id}", "RegistroMantenimiento", "Eliminar registro de mantenimiento", False),
    ("get", "/api/registros-mantenimiento/alerta/{alerta_id}", "RegistroMantenimiento", "Registros por alerta", False),
    ("get", "/api/registros-mantenimiento/camion/{camion_id}", "RegistroMantenimiento", "Registros por camión", False),
    ("get", "/api/registros-mantenimiento/coordinador/{coordinador_id}", "RegistroMantenimiento", "Registros por coordinador", False),
    ("get", "/api/registros-mantenimiento/fecha", "RegistroMantenimiento", "Registros por fecha", False),

    ("post", "/api/reportes-falla-critica/", "ReporteFallaCritica", "Crear reporte de falla crítica", False),
    ("get", "/api/reportes-falla-critica/", "ReporteFallaCritica", "Listar reportes de falla crítica", False),
    ("get", "/api/reportes-falla-critica/{id}", "ReporteFallaCritica", "Obtener reporte por ID", False),
    ("put", "/api/reportes-falla-critica/{id}", "ReporteFallaCritica", "Actualizar reporte de falla crítica", False),
    ("delete", "/api/reportes-falla-critica/{id}", "ReporteFallaCritica", "Eliminar reporte de falla crítica", False),
    ("get", "/api/reportes-falla-critica/camion/{camionId}", "ReporteFallaCritica", "Reportes por camión", False),
    ("get", "/api/reportes-falla-critica/conductor/{conductorId}", "ReporteFallaCritica", "Reportes por conductor", False),
    ("get", "/api/reportes-falla-critica/por-fecha", "ReporteFallaCritica", "Reportes por fecha", False),

    ("post", "/api/reportes-mantenimiento-generado/", "ReporteMantenimientoGenerado", "Crear reporte de mantenimiento generado", False),
    ("get", "/api/reportes-mantenimiento-generado/", "ReporteMantenimientoGenerado", "Listar reportes generados", False),
    ("get", "/api/reportes-mantenimiento-generado/{id}", "ReporteMantenimientoGenerado", "Obtener reporte generado por ID", False),
    ("put", "/api/reportes-mantenimiento-generado/{id}", "ReporteMantenimientoGenerado", "Actualizar reporte generado", False),
    ("delete", "/api/reportes-mantenimiento-generado/{id}", "ReporteMantenimientoGenerado", "Eliminar reporte generado", False),
    ("get", "/api/reportes-mantenimiento-generado/coordinador/{coordinador_id}", "ReporteMantenimientoGenerado", "Reportes por coordinador", False),
    ("get", "/api/reportes-mantenimiento-generado/fecha", "ReporteMantenimientoGenerado", "Reportes por fecha", False),
    ("get", "/api/reportes-mantenimiento-generado/fecha-generacion", "ReporteMantenimientoGenerado", "Reportes por fecha de generación", False),

    ("post", "/api/seguimientos-falla-critica/", "SeguimientoFallaCritica", "Crear seguimiento de falla crítica", False),
    ("get", "/api/seguimientos-falla-critica/", "SeguimientoFallaCritica", "Listar seguimientos", False),
    ("get", "/api/seguimientos-falla-critica/{id}", "SeguimientoFallaCritica", "Obtener seguimiento por ID", False),
    ("put", "/api/seguimientos-falla-critica/{id}", "SeguimientoFallaCritica", "Actualizar seguimiento", False),
    ("delete", "/api/seguimientos-falla-critica/{id}", "SeguimientoFallaCritica", "Eliminar seguimiento", False),
    ("get", "/api/seguimientos-falla-critica/falla/{fallaId}", "SeguimientoFallaCritica", "Seguimientos por falla", False),
    ("get", "/api/seguimientos-falla-critica/por-fecha", "SeguimientoFallaCritica", "Seguimientos por fecha", False),

    ("post", "/api/tipos-mantenimiento/", "TipoMantenimiento", "Crear tipo de mantenimiento", False),
    ("get", "/api/tipos-mantenimiento/", "TipoMantenimiento", "Listar tipos de mantenimiento", False),
    ("get", "/api/tipos-mantenimiento/{id}", "TipoMantenimiento", "Obtener tipo por ID", False),
    ("put", "/api/tipos-mantenimiento/{id}", "TipoMantenimiento", "Actualizar tipo de mantenimiento", False),
    ("delete", "/api/tipos-mantenimiento/{id}", "TipoMantenimiento", "Eliminar tipo de mantenimiento", False),

    ("post", "/domicilios", "Domicilio", "Crear domicilio (compat)", True),
    ("get", "/domicilios", "Domicilio", "Listar domicilios (compat)", True),
    ("get", "/domicilios/{id}", "Domicilio", "Obtener domicilio (compat)", True),
    ("put", "/domicilios/{id}", "Domicilio", "Actualizar domicilio (compat)", True),
    ("delete", "/domicilios/{id}", "Domicilio", "Eliminar domicilio (compat)", True),

    ("get", "/colonias", "Colonia", "Listar colonias (compat)", False),
    ("get", "/colonias/{id}", "Colonia", "Obtener colonia (compat)", False),
    ("post", "/colonias", "Colonia", "Crear colonia (compat)", True),
    ("put", "/colonias/{id}", "Colonia", "Actualizar colonia (compat)", True),
    ("delete", "/colonias/{id}", "Colonia", "Eliminar colonia (compat)", True),
]


# (method, path) -> tipo Swagger del body JSON
BODY_MODELS: dict[tuple[str, str], str] = {
    ("post", "/api/camion/"): "CreateCamionRequest",
    ("put", "/api/camion/{id}"): "CreateCamionRequest",
    ("post", "/api/estado-camion/"): "CreateEstadoCamionRequest",
    ("put", "/api/estado-camion/{id}"): "CreateEstadoCamionRequest",
    ("post", "/api/tipo-camion/"): "CreateTipoCamionRequest",
    ("post", "/api/rutas/"): "CreateRutaRequest",
    ("put", "/api/rutas/{id}"): "CreateRutaRequest",
    ("post", "/api/puntos-recoleccion/"): "CreatePuntoRecoleccionRequest",
    ("put", "/api/puntos-recoleccion/{id}"): "CreatePuntoRecoleccionRequest",
    ("post", "/api/relleno-sanitario/"): "CreateRellenoSanitarioRequest",
    ("put", "/api/relleno-sanitario/{id}"): "CreateRellenoSanitarioRequest",
    ("post", "/api/ruta-camion/"): "CreateRutaCamionRequest",
    ("put", "/api/ruta-camion/{id}"): "CreateRutaCamionRequest",
    ("post", "/api/historial-asignacion/"): "CreateHistorialAsignacionRequest",
    ("put", "/api/historial-asignacion/{id}"): "CreateHistorialAsignacionRequest",
    ("post", "/api/registro-vaciado/"): "CreateRegistroVaciadoRequest",
    ("post", "/api/empleados/login"): "LoginUsuarioRequest",
    ("post", "/api/empleados/"): "CreateUsuarioRequest",
    ("post", "/api/anomalias/"): "CreateAnomaliaRequest",
    ("put", "/api/anomalias/{id}"): "UpdateAnomaliaRequest",
    ("post", "/api/incidencias/"): "CreateIncidenciaRequest",
    ("put", "/api/incidencias/{id}"): "CreateIncidenciaRequest",
    ("post", "/api/reportes-conductor/"): "CreateReporteConductorRequest",
    ("put", "/api/reportes-conductor/{id}"): "CreateReporteConductorRequest",
    ("post", "/api/registros-mantenimiento/"): "CreateRegistroMantenimientoRequest",
    ("put", "/api/registros-mantenimiento/{id}"): "CreateRegistroMantenimientoRequest",
    ("post", "/api/reportes-falla-critica/"): "CreateReporteFallaCriticaRequest",
    ("put", "/api/reportes-falla-critica/{id}"): "CreateReporteFallaCriticaRequest",
    ("post", "/api/tipos-mantenimiento/"): "CreateTipoMantenimientoRequest",
    ("put", "/api/tipos-mantenimiento/{id}"): "CreateTipoMantenimientoRequest",
    ("post", "/api/notificaciones/"): "CreateNotificacionRequest",
    ("put", "/api/notificaciones/{id}"): "UpdateNotificacionRequest",
    ("post", "/api/notificaciones/emergencia"): "NotificacionEmergenciaRequest",
    ("post", "/api/notificaciones/falla"): "NotificacionFallaRequest",
    ("post", "/api/notificaciones/mantenimiento"): "NotificacionMantenimientoRequest",
    ("post", "/api/notificaciones/notificar"): "NotificarUsuarioRequest",
    ("post", "/api/notificaciones/enviar-multiples"): "NotificarMultiplesRequest",
    ("post", "/api/notificaciones/enviar-todos"): "NotificarTodosRequest",
    ("post", "/api/alertas"): "CreateAlertaRequest",
    ("post", "/domicilios"): "domain.CreateDomicilioRequest",
    ("put", "/domicilios/{id}"): "domain.CreateDomicilioRequest",
    ("post", "/colonias"): "domain.CreateColoniaRequest",
    ("put", "/colonias/{id}"): "domain.UpdateColoniaRequest",
}

# Endpoints mutantes sin body en el handler real
NO_BODY: set[tuple[str, str]] = {
    ("put", "/api/historial-asignacion/baja/{id}"),
    ("put", "/api/historial-asignacion/cerrar/camion/{camionId}"),
    ("put", "/api/historial-asignacion/cerrar/chofer/{choferId}"),
    ("put", "/api/alertas/{id}/leida"),
    ("patch", "/api/notificaciones/{id}/marcar-leida"),
    ("patch", "/api/notificaciones/{id}/reactivar"),
    ("patch", "/api/notificaciones/usuario/{usuario_id}/marcar-todas-leidas"),
}


def slug(method: str, path: str) -> str:
    clean = re.sub(r"[{}]", "", path).replace("/", "_").replace("-", "_").strip("_")
    return f"doc_{method}_{clean}"


def write_method_block(method: str, path: str) -> list[str]:
    lines: list[str] = []
    if method in {"get", "delete"}:
        lines.append("// @Produce json")
    elif method in {"post", "put", "patch"}:
        lines.append("// @Accept json")
        lines.append("// @Produce json")
        if (method, path) not in NO_BODY:
            model = BODY_MODELS.get((method, path), "map[string]interface{}")
            lines.append(f'// @Param body body {model} true "Payload JSON"')
    else:
        lines.append("// @Produce json")
    return lines


def main() -> None:
    lines = [
        "package main",
        "",
        "// Archivo generado por scripts/generate_swagger_stubs.py — no editar a mano.",
        "",
    ]

    seen: set[tuple[str, str]] = set()
    for method, path, tag, summary, secured in ROUTES:
        key = (method, path)
        if key in seen:
            continue
        seen.add(key)

        fn = slug(method, path)
        lines.append(f"// {fn} godoc")
        lines.append(f"// @Summary {summary}")
        lines.append(f"// @Tags {tag}")
        lines.extend(write_method_block(method, path))
        lines.append('// @Success 200 {object} map[string]interface{}')
        if method == "post":
            lines.append('// @Success 201 {object} map[string]interface{}')
        lines.append('// @Failure 400 {object} map[string]interface{}')
        lines.append('// @Failure 401 {object} map[string]interface{}')
        lines.append('// @Failure 500 {object} map[string]interface{}')
        if secured:
            lines.append("// @Security BearerAuth")
        lines.append(f"// @Router {path} [{method}]")
        lines.append(f"func {fn}() {{}}")
        lines.append("")

    OUTPUT.write_text("\n".join(lines) + "\n", encoding="utf-8")
    print(f"Generado {OUTPUT} con {len(seen)} rutas")


if __name__ == "__main__":
    main()
