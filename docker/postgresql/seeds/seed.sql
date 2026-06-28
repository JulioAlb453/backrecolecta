-- =============================================================================
-- SEED alineado con gin-backend/db_script.sql
-- Compatible con PostgreSQL (Neon) y schema public de proyecto_recolecta
-- =============================================================================

BEGIN;

-- -----------------------------------------------------------------------------
-- 1. ROLES
-- -----------------------------------------------------------------------------
INSERT INTO rol (role_id, nombre, eliminado) VALUES
  (1, 'Administrador', FALSE),
  (2, 'Coordinador', FALSE),
  (3, 'Operador', FALSE),
  (4, 'Conductor', FALSE),
  (5, 'Ciudadano', FALSE)
ON CONFLICT (role_id) DO NOTHING;

-- -----------------------------------------------------------------------------
-- 2. COLONIAS
-- -----------------------------------------------------------------------------
INSERT INTO colonia (colonia_id, nombre, zona, created_at) VALUES
  (1, 'Centro Histórico', 'Centro', '2024-01-15 08:00:00'),
  (2, 'Colonia Industrial', 'Norte', '2024-01-15 08:00:00'),
  (3, 'Las Palmas', 'Norte', '2024-01-15 08:00:00'),
  (4, 'Vista Hermosa', 'Sur', '2024-01-15 08:00:00'),
  (5, 'Jardines del Valle', 'Sur', '2024-01-15 08:00:00'),
  (6, 'El Mirador', 'Centro', '2024-01-15 08:00:00'),
  (7, 'Residencial San Miguel', 'Norte', '2024-01-15 08:00:00'),
  (8, 'Fraccionamiento Los Pinos', 'Sur', '2024-01-15 08:00:00')
ON CONFLICT (colonia_id) DO NOTHING;

-- -----------------------------------------------------------------------------
-- 3. USUARIOS STAFF (IDs 1-12)
-- -----------------------------------------------------------------------------
INSERT INTO usuario (user_id, nombre, alias, email, password, role_id, eliminado, created_at, updated_at)
VALUES
  (1, 'Roberto García Méndez', 'rgarcia', 'roberto.garcia@recolecta.mx', '$2a$10$EixZaYVK1fsbw1ZfbX3OXe', 1, FALSE, '2024-01-10 09:00:00', '2024-01-10 09:00:00'),
  (2, 'María Elena Torres', 'mtorres', 'maria.torres@recolecta.mx', '$2a$10$EixZaYVK1fsbw1ZfbX3OXe', 2, FALSE, '2024-01-12 10:00:00', '2024-01-12 10:00:00'),
  (3, 'Carlos Ramírez López', 'cramirez', 'carlos.ramirez@recolecta.mx', '$2a$10$EixZaYVK1fsbw1ZfbX3OXe', 2, FALSE, '2024-01-12 10:30:00', '2024-01-12 10:30:00'),
  (4, 'Ana Patricia Morales', 'amorales', 'ana.morales@recolecta.mx', '$2a$10$EixZaYVK1fsbw1ZfbX3OXe', 3, FALSE, '2024-01-15 08:30:00', '2024-01-15 08:30:00'),
  (5, 'Jorge Luis Sánchez', 'jsanchez', 'jorge.sanchez@recolecta.mx', '$2a$10$EixZaYVK1fsbw1ZfbX3OXe', 3, FALSE, '2024-01-15 08:45:00', '2024-01-15 08:45:00'),
  (6, 'Patricia Hernández Cruz', 'phernandez', 'patricia.hernandez@recolecta.mx', '$2a$10$EixZaYVK1fsbw1ZfbX3OXe', 3, FALSE, '2024-01-15 09:00:00', '2024-01-15 09:00:00'),
  (7, 'Juan Manuel Flores', 'jflores', 'juan.flores@recolecta.mx', '$2a$10$EixZaYVK1fsbw1ZfbX3OXe', 4, FALSE, '2024-02-01 07:00:00', '2024-02-01 07:00:00'),
  (8, 'Pedro Ávila Gómez', 'pavila', 'pedro.avila@recolecta.mx', '$2a$10$EixZaYVK1fsbw1ZfbX3OXe', 4, FALSE, '2024-02-01 07:15:00', '2024-02-01 07:15:00'),
  (9, 'Luis Alberto Vargas', 'lvargas', 'luis.vargas@recolecta.mx', '$2a$10$EixZaYVK1fsbw1ZfbX3OXe', 4, FALSE, '2024-02-01 07:30:00', '2024-02-01 07:30:00'),
  (10, 'Miguel Ángel Medina', 'mmedina', 'miguel.medina@recolecta.mx', '$2a$10$EixZaYVK1fsbw1ZfbX3OXe', 4, FALSE, '2024-02-05 07:00:00', '2024-02-05 07:00:00'),
  (11, 'José Antonio Ruiz', 'jruiz', 'jose.ruiz@recolecta.mx', '$2a$10$EixZaYVK1fsbw1ZfbX3OXe', 4, FALSE, '2024-02-05 07:15:00', '2024-02-05 07:15:00'),
  (12, 'Francisco Javier Castro', 'fcastro', 'francisco.castro@recolecta.mx', '$2a$10$EixZaYVK1fsbw1ZfbX3OXe', 4, FALSE, '2024-02-05 07:30:00', '2024-02-05 07:30:00')
ON CONFLICT (user_id) DO NOTHING;

-- -----------------------------------------------------------------------------
-- 4. CIUDADANOS (IDs 100-299, alineado con seed Redis)
-- -----------------------------------------------------------------------------
INSERT INTO usuario (user_id, nombre, alias, email, password, role_id, eliminado, created_at, updated_at)
SELECT
  100 + g,
  'Usuario ' || (100 + g),
  'user' || (100 + g),
  'user' || (100 + g) || '@example.com',
  '$2a$10$EixZaYVK1fsbw1ZfbX3OXe',
  5,
  FALSE,
  now(),
  now()
FROM generate_series(0, 199) AS g
ON CONFLICT (user_id) DO NOTHING;

-- -----------------------------------------------------------------------------
-- 5. DOMICILIOS CIUDADANOS
-- -----------------------------------------------------------------------------
INSERT INTO domicilio (domicilio_id, usuario_id, alias, direccion, colonia_id, eliminado, created_at, updated_at)
SELECT
  100 + g,
  100 + g,
  'Domicilio Principal',
  (
    (ARRAY['Calle Olmo','Calle Lirio','Calle Roble','Avenida Reforma','Calle Cedro','Calle Laurel','Calle Magnolia','Calle Nogal','Calle Pino','Calle Sauce'])[ (g % 10) + 1 ]
    || ' ' || ((100 + g) % 200 + 10)::text
    || ', Colonia ' || (ARRAY['Centro Histórico','Colonia Industrial','Las Palmas','Vista Hermosa','Jardines del Valle','El Mirador','Residencial San Miguel','Fraccionamiento Los Pinos'])[ (g % 8) + 1 ]
    || ', Ciudad Recolecta, CP ' || (ARRAY['38000','38100','38110','38200','38210','38010','38120','38220'])[ (g % 8) + 1 ]
  ),
  (g % 8) + 1,
  FALSE,
  now(),
  now()
FROM generate_series(0, 199) AS g
ON CONFLICT (domicilio_id) DO NOTHING;

-- -----------------------------------------------------------------------------
-- 6. TIPOS DE CAMIÓN
-- -----------------------------------------------------------------------------
INSERT INTO tipo_camion (tipo_camion_id, nombre, descripcion, created_at) VALUES
  (1, 'Compactador 12m³', 'Camión compactador estándar capacidad 12 metros cúbicos', '2024-01-20 08:00:00'),
  (2, 'Compactador 15m³', 'Camión compactador gran capacidad 15 metros cúbicos', '2024-01-20 08:00:00'),
  (3, 'Camión de Volteo', 'Camión de volteo para escombros y residuos voluminosos', '2024-01-20 08:00:00')
ON CONFLICT (tipo_camion_id) DO NOTHING;

-- -----------------------------------------------------------------------------
-- 7. CAMIONES
-- -----------------------------------------------------------------------------
INSERT INTO camion (
  camion_id, placa, modelo, tipo_camion_id, es_rentado, eliminado,
  disponibilidad_id, nombre_disponibilidad, color_disponibilidad, created_at, updated_at
) VALUES
  (1, 'ABC-123-MX', 'Freightliner M2 106 2022', 1, FALSE, FALSE, 1, 'Disponible', 'green', '2024-01-20 08:00:00', '2024-01-20 08:00:00'),
  (2, 'DEF-456-MX', 'International DuraStar 2021', 2, FALSE, FALSE, 1, 'Disponible', 'green', '2024-01-20 08:15:00', '2024-01-20 08:15:00'),
  (3, 'GHI-789-MX', 'Kenworth T370 2023', 1, FALSE, FALSE, 1, 'Disponible', 'green', '2024-01-20 08:30:00', '2024-01-20 08:30:00'),
  (4, 'JKL-012-MX', 'Volvo VHD 2020', 2, TRUE, FALSE, 1, 'Disponible', 'green', '2024-02-01 09:00:00', '2024-02-01 09:00:00'),
  (5, 'MNO-345-MX', 'Peterbilt 337 2021', 1, TRUE, FALSE, 1, 'Disponible', 'green', '2024-02-01 09:15:00', '2024-02-01 09:15:00'),
  (6, 'PQR-678-MX', 'Mack LR 2019', 3, TRUE, FALSE, 1, 'Disponible', 'green', '2024-02-01 09:30:00', '2024-02-01 09:30:00')
ON CONFLICT (camion_id) DO NOTHING;

-- -----------------------------------------------------------------------------
-- 8. HISTORIAL ASIGNACIÓN CONDUCTOR-CAMIÓN
-- -----------------------------------------------------------------------------
INSERT INTO historial_asignacion_camion (
  id_historial, id_chofer, id_camion, fecha_asignacion, fecha_baja, eliminado, created_at, updated_at
) VALUES
  (1, 7, 1, '2024-02-10 06:00:00', NULL, FALSE, '2024-02-10 06:00:00', '2024-02-10 06:00:00'),
  (2, 8, 2, '2024-02-10 06:00:00', NULL, FALSE, '2024-02-10 06:00:00', '2024-02-10 06:00:00'),
  (3, 9, 3, '2024-02-10 06:00:00', NULL, FALSE, '2024-02-10 06:00:00', '2024-02-10 06:00:00'),
  (4, 10, 4, '2024-02-15 06:00:00', NULL, FALSE, '2024-02-15 06:00:00', '2024-02-15 06:00:00'),
  (5, 11, 5, '2024-02-15 06:00:00', NULL, FALSE, '2024-02-15 06:00:00', '2024-02-15 06:00:00'),
  (6, 12, 6, '2024-02-15 06:00:00', NULL, FALSE, '2024-02-15 06:00:00', '2024-02-15 06:00:00')
ON CONFLICT (id_historial) DO NOTHING;

-- -----------------------------------------------------------------------------
-- 9. RUTAS  (json_ruta incluye waypoints reales en Suchiapas, Chiapas)
-- -----------------------------------------------------------------------------
INSERT INTO ruta (ruta_id, nombre, descripcion, json_ruta, eliminado, created_at) VALUES
  (1, 'Ruta Norte A', 'Cobertura Colonia Industrial y Las Palmas',
   '{"zona":"Norte","turno":"matutino","puntos":[{"id":"PR-NA-001","orden":1,"lat":16.6420,"lng":-93.1180,"nombre":"Esq. Calle Industrial / Av. Las Palmas"},{"id":"PR-NA-002","orden":2,"lat":16.6395,"lng":-93.1210,"nombre":"Esq. Calle 3a Nte. / Calle Industrial"},{"id":"PR-NA-003","orden":3,"lat":16.6370,"lng":-93.1155,"nombre":"Esq. Calle 2a Nte. / Pte. Las Palmas"},{"id":"PR-NA-004","orden":4,"lat":16.6345,"lng":-93.1130,"nombre":"Esq. Calle 1a Nte. / Pte. Las Palmas"},{"id":"PR-NA-005","orden":5,"lat":16.6325,"lng":-93.1095,"nombre":"Esq. Av. Las Palmas / Central"}]}',
   FALSE, '2024-02-01 08:00:00'),
  (2, 'Ruta Norte B', 'Cobertura Residencial San Miguel',
   '{"zona":"Norte","turno":"vespertino","puntos":[{"id":"PR-NB-001","orden":1,"lat":16.6510,"lng":-93.1080,"nombre":"Esq. Residencial San Miguel / 4a Nte."},{"id":"PR-NB-002","orden":2,"lat":16.6480,"lng":-93.1055,"nombre":"Esq. 3a Nte. / Calle San Miguel"},{"id":"PR-NB-003","orden":3,"lat":16.6455,"lng":-93.1030,"nombre":"Esq. 2a Nte. / Calle San Miguel"},{"id":"PR-NB-004","orden":4,"lat":16.6430,"lng":-93.1010,"nombre":"Esq. 1a Nte. / Calle San Miguel"},{"id":"PR-NB-005","orden":5,"lat":16.6405,"lng":-93.0985,"nombre":"Esq. San Miguel / Oriente"}]}',
   FALSE, '2024-02-01 08:15:00'),
  (3, 'Ruta Centro', 'Cobertura Centro Histórico y El Mirador',
   '{"zona":"Centro","turno":"matutino","puntos":[{"id":"PR-CE-001","orden":1,"lat":16.6278,"lng":-93.1045,"nombre":"Esq. Calle 1a Nte. / Av. Central"},{"id":"PR-CE-002","orden":2,"lat":16.6261,"lng":-93.1038,"nombre":"Esq. Av. Central / Calle 1a Sur"},{"id":"PR-CE-003","orden":3,"lat":16.6241,"lng":-93.1025,"nombre":"Esq. Calle 1a Sur / 1a Ote."},{"id":"PR-CE-004","orden":4,"lat":16.6235,"lng":-93.0998,"nombre":"Esq. 1a Ote. / Calle 2a Sur"},{"id":"PR-CE-005","orden":5,"lat":16.6270,"lng":-93.1070,"nombre":"Esq. 1a Nte. / Av. 1a Pte."}]}',
   FALSE, '2024-02-01 08:30:00'),
  (4, 'Ruta Sur A', 'Cobertura Vista Hermosa y Jardines del Valle',
   '{"zona":"Sur","turno":"matutino","puntos":[{"id":"PR-SA-001","orden":1,"lat":16.6185,"lng":-93.0950,"nombre":"Esq. Vista Hermosa / Av. Sur"},{"id":"PR-SA-002","orden":2,"lat":16.6162,"lng":-93.0925,"nombre":"Esq. Jardines del Valle / 1a Sur"},{"id":"PR-SA-003","orden":3,"lat":16.6140,"lng":-93.0910,"nombre":"Esq. 2a Sur / Calle Valle"},{"id":"PR-SA-004","orden":4,"lat":16.6120,"lng":-93.0935,"nombre":"Esq. Calle Valle / 3a Sur"},{"id":"PR-SA-005","orden":5,"lat":16.6145,"lng":-93.0960,"nombre":"Esq. 3a Sur / Av. Vista"}]}',
   FALSE, '2024-02-01 08:45:00'),
  (5, 'Ruta Sur B', 'Cobertura Fraccionamiento Los Pinos',
   '{"zona":"Sur","turno":"vespertino","puntos":[{"id":"PR-SB-001","orden":1,"lat":16.6115,"lng":-93.1050,"nombre":"Esq. Los Pinos / Calle Sur"},{"id":"PR-SB-002","orden":2,"lat":16.6090,"lng":-93.1080,"nombre":"Esq. Fracc. Los Pinos / 2a Nte."},{"id":"PR-SB-003","orden":3,"lat":16.6068,"lng":-93.1110,"nombre":"Esq. 2a Nte. / Calle Pinos"},{"id":"PR-SB-004","orden":4,"lat":16.6085,"lng":-93.1140,"nombre":"Esq. Calle Pinos / Calle Sur"},{"id":"PR-SB-005","orden":5,"lat":16.6110,"lng":-93.1120,"nombre":"Esq. Calle Sur / Av. Los Pinos"}]}',
   FALSE, '2024-02-01 09:00:00')
ON CONFLICT (ruta_id) DO UPDATE SET
  json_ruta   = EXCLUDED.json_ruta,
  nombre      = EXCLUDED.nombre,
  descripcion = EXCLUDED.descripcion;

-- -----------------------------------------------------------------------------
-- 10. PUNTOS DE RECOLECCIÓN
-- -----------------------------------------------------------------------------
INSERT INTO punto_recoleccion (punto_id, ruta_id, cp, eliminado) VALUES
  (1, 1, 'PR-NA-001', FALSE),
  (2, 1, 'PR-NA-002', FALSE),
  (3, 1, 'PR-NA-003', FALSE),
  (4, 1, 'PR-NA-004', FALSE),
  (5, 1, 'PR-NA-005', FALSE),
  (6, 2, 'PR-NB-001', FALSE),
  (7, 2, 'PR-NB-002', FALSE),
  (8, 2, 'PR-NB-003', FALSE),
  (9, 2, 'PR-NB-004', FALSE),
  (10, 2, 'PR-NB-005', FALSE),
  (11, 3, 'PR-CE-001', FALSE),
  (12, 3, 'PR-CE-002', FALSE),
  (13, 3, 'PR-CE-003', FALSE),
  (14, 3, 'PR-CE-004', FALSE),
  (15, 3, 'PR-CE-005', FALSE),
  (16, 4, 'PR-SA-001', FALSE),
  (17, 4, 'PR-SA-002', FALSE),
  (18, 4, 'PR-SA-003', FALSE),
  (19, 4, 'PR-SA-004', FALSE),
  (20, 4, 'PR-SA-005', FALSE),
  (21, 5, 'PR-SB-001', FALSE),
  (22, 5, 'PR-SB-002', FALSE),
  (23, 5, 'PR-SB-003', FALSE),
  (24, 5, 'PR-SB-004', FALSE),
  (25, 5, 'PR-SB-005', FALSE)
ON CONFLICT (punto_id) DO NOTHING;

-- -----------------------------------------------------------------------------
-- 11. ASIGNACIÓN RUTA-CAMIÓN
-- -----------------------------------------------------------------------------
INSERT INTO ruta_camion (ruta_camion_id, ruta_id, camion_id, fecha, created_at, eliminado) VALUES
  (1, 1, 1, CURRENT_DATE, now(), FALSE),
  (2, 2, 5, CURRENT_DATE, now(), FALSE),
  (3, 3, 2, CURRENT_DATE, now(), FALSE),
  (4, 4, 3, CURRENT_DATE, now(), FALSE),
  (5, 5, 4, CURRENT_DATE, now(), FALSE)
ON CONFLICT (ruta_camion_id) DO NOTHING;

-- -----------------------------------------------------------------------------
-- 12. TIPOS DE MANTENIMIENTO
-- -----------------------------------------------------------------------------
INSERT INTO tipo_mantenimiento (tipo_mantenimiento_id, nombre, categoria, eliminado) VALUES
  (1, 'Cambio de Aceite', 'preventivo', FALSE),
  (2, 'Revisión de Frenos', 'preventivo', FALSE),
  (3, 'Alineación y Balanceo', 'preventivo', FALSE),
  (4, 'Cambio de Filtros', 'preventivo', FALSE),
  (5, 'Reparación Motor', 'correctivo', FALSE),
  (6, 'Reparación Transmisión', 'correctivo', FALSE),
  (7, 'Reparación Sistema Hidráulico', 'correctivo', FALSE),
  (8, 'Reemplazo Neumáticos', 'correctivo', FALSE)
ON CONFLICT (tipo_mantenimiento_id) DO NOTHING;

-- -----------------------------------------------------------------------------
-- 13. ALERTAS Y REGISTROS DE MANTENIMIENTO
-- -----------------------------------------------------------------------------
INSERT INTO alerta_mantenimiento (alerta_id, camion_id, tipo_mantenimiento_id, descripcion, observaciones, created_at, atendido) VALUES
  (1, 2, 1, 'Servicio por kilometraje', 'Unidad alcanzó kilometraje para servicio', '2026-01-20 10:00:00', TRUE),
  (2, 5, 2, 'Ruido al frenar', 'Conductor reportó ruido al frenar', '2026-01-22 14:30:00', TRUE),
  (3, 6, 4, 'Mantenimiento trimestral', 'Mantenimiento trimestral programado', '2026-01-25 09:00:00', TRUE)
ON CONFLICT (alerta_id) DO NOTHING;

INSERT INTO registro_mantenimiento (
  registro_id, alerta_id, camion_id, coordinador_id, mecanico_responsable,
  fecha_realizada, kilometraje_mantenimiento, observaciones, created_at
) VALUES
  (1, 1, 2, 2, 'Taller Central', '2026-01-21 10:00:00', 15000, 'Cambio de aceite completado', '2026-01-21 12:00:00'),
  (2, 2, 5, 3, 'Taller Norte', '2026-01-23 09:00:00', 0, 'Pastillas de freno reemplazadas', '2026-01-23 11:00:00'),
  (3, 3, 6, 2, 'Taller Sur', '2026-01-26 08:00:00', 0, 'Filtros reemplazados', '2026-01-26 10:00:00')
ON CONFLICT (registro_id) DO NOTHING;

-- -----------------------------------------------------------------------------
-- 14. AVISO GENERAL
-- -----------------------------------------------------------------------------
INSERT INTO aviso_general (aviso_id, titulo, mensaje, activo, created_at) VALUES
  (1, 'Mantenimiento Programado', 'Se realizará mantenimiento al sistema el próximo domingo de 00:00 a 06:00 hrs.', TRUE, '2026-01-25 08:00:00')
ON CONFLICT (aviso_id) DO NOTHING;

-- -----------------------------------------------------------------------------
-- 15. Sincronizar secuencias SERIAL
-- -----------------------------------------------------------------------------
SELECT setval(pg_get_serial_sequence('rol', 'role_id'), COALESCE((SELECT MAX(role_id) FROM rol), 1));
SELECT setval(pg_get_serial_sequence('colonia', 'colonia_id'), COALESCE((SELECT MAX(colonia_id) FROM colonia), 1));
SELECT setval(pg_get_serial_sequence('usuario', 'user_id'), COALESCE((SELECT MAX(user_id) FROM usuario), 1));
SELECT setval(pg_get_serial_sequence('domicilio', 'domicilio_id'), COALESCE((SELECT MAX(domicilio_id) FROM domicilio), 1));
SELECT setval(pg_get_serial_sequence('tipo_camion', 'tipo_camion_id'), COALESCE((SELECT MAX(tipo_camion_id) FROM tipo_camion), 1));
SELECT setval(pg_get_serial_sequence('camion', 'camion_id'), COALESCE((SELECT MAX(camion_id) FROM camion), 1));
SELECT setval(pg_get_serial_sequence('historial_asignacion_camion', 'id_historial'), COALESCE((SELECT MAX(id_historial) FROM historial_asignacion_camion), 1));
SELECT setval(pg_get_serial_sequence('ruta', 'ruta_id'), COALESCE((SELECT MAX(ruta_id) FROM ruta), 1));
SELECT setval(pg_get_serial_sequence('punto_recoleccion', 'punto_id'), COALESCE((SELECT MAX(punto_id) FROM punto_recoleccion), 1));
SELECT setval(pg_get_serial_sequence('ruta_camion', 'ruta_camion_id'), COALESCE((SELECT MAX(ruta_camion_id) FROM ruta_camion), 1));
SELECT setval(pg_get_serial_sequence('tipo_mantenimiento', 'tipo_mantenimiento_id'), COALESCE((SELECT MAX(tipo_mantenimiento_id) FROM tipo_mantenimiento), 1));
SELECT setval(pg_get_serial_sequence('alerta_mantenimiento', 'alerta_id'), COALESCE((SELECT MAX(alerta_id) FROM alerta_mantenimiento), 1));
SELECT setval(pg_get_serial_sequence('registro_mantenimiento', 'registro_id'), COALESCE((SELECT MAX(registro_id) FROM registro_mantenimiento), 1));
SELECT setval(pg_get_serial_sequence('aviso_general', 'aviso_id'), COALESCE((SELECT MAX(aviso_id) FROM aviso_general), 1));

COMMIT;
