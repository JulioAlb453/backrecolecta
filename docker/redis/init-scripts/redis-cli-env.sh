#!/bin/sh
# Helpers compartidos para scripts que invocan redis-cli (local o Upstash con TLS).

REDIS_TLS=""
REDIS_HOST=""
REDIS_PORT=""
REDIS_PASSWORD=""
REDIS_DB=""

redis_cli_trim() {
    printf '%s' "$1" | sed -e 's/^[[:space:]]*//' -e 's/[[:space:]]*$//'
}

redis_cli_resolve_docker() {
    if command -v docker >/dev/null 2>&1; then
        printf '%s' "docker"
        return 0
    fi

    if command -v docker.exe >/dev/null 2>&1; then
        printf '%s' "docker.exe"
        return 0
    fi

    for candidate in \
        "/mnt/c/Program Files/Docker/Docker/resources/bin/docker.exe" \
        "/c/Program Files/Docker/Docker/resources/bin/docker.exe"
    do
        if [ -x "$candidate" ]; then
            printf '%s' "$candidate"
            return 0
        fi
    done

    return 1
}

# Si redis-cli no está en el host (común en Windows/WSL), re-ejecutar dentro de redis:7.2-alpine.
redis_cli_bootstrap() {
    script_dir="$1"
    script_path="$2"
    shift 2

    if command -v redis-cli >/dev/null 2>&1; then
        return 0
    fi

    if [ "${REDIS_CLI_DOCKER_REEXEC:-}" = "1" ]; then
        echo "[ERROR] redis-cli no está disponible dentro del contenedor Docker." >&2
        exit 127
    fi

    docker_bin=$(redis_cli_resolve_docker || true)
    if [ -z "$docker_bin" ]; then
        echo "[ERROR] redis-cli no está instalado y Docker no está disponible en este entorno." >&2
        echo "[DETALLE] Opción A (PowerShell, recomendada en Windows):" >&2
        echo "[DETALLE]   .\\docker\\redis\\init-scripts\\load-redis-upstash.ps1 -RedisHost \"xxx.upstash.io\" -Password \"...\"" >&2
        echo "[DETALLE] Opción B: activa WSL Integration en Docker Desktop → Settings → Resources → WSL Integration" >&2
        echo "[DETALLE] Opción C (WSL): sudo apt update && sudo apt install -y redis-tools" >&2
        exit 127
    fi

    redis_root="$(cd "$script_dir/.." && pwd)"
    script_name=$(basename "$script_path")

    echo "[INFO] redis-cli no encontrado en el host; usando contenedor redis:7.2-alpine via $docker_bin..." >&2

    exec "$docker_bin" run --rm \
        -e REDIS_CLI_DOCKER_REEXEC=1 \
        -v "${redis_root}:/redis" \
        -w /redis/init-scripts \
        redis:7.2-alpine \
        sh "$script_name" "$@"
}

redis_cli_parse_args() {
    while [ $# -gt 0 ]; do
        case "$1" in
            --tls)
                REDIS_TLS="--tls"
                shift
                ;;
            *)
                break
                ;;
        esac
    done

    REDIS_HOST=$(redis_cli_trim "${1:-localhost}")
    REDIS_PORT=$(redis_cli_trim "${2:-6379}")
    REDIS_PASSWORD=$(redis_cli_trim "${3:-}")
    REDIS_DB=$(redis_cli_trim "${4:-0}")

    # Aceptar hosts copiados con esquema (ej. https://xxx.upstash.io)
    REDIS_HOST=$(printf '%s' "$REDIS_HOST" | sed -e 's|^https://||' -e 's|^http://||' -e 's|/$||')

    if [ -z "$REDIS_TLS" ] && [ "${REDIS_TLS_ENABLED:-}" = "true" ]; then
        REDIS_TLS="--tls"
    fi

    if [ -n "$REDIS_PASSWORD" ]; then
        export REDISCLI_AUTH="$REDIS_PASSWORD"
    fi
}

redis_cmd() {
    if [ -n "$REDIS_TLS" ]; then
        redis-cli --tls -h "$REDIS_HOST" -p "$REDIS_PORT" -n "$REDIS_DB" "$@"
    else
        redis-cli -h "$REDIS_HOST" -p "$REDIS_PORT" -n "$REDIS_DB" "$@"
    fi
}

redis_cmd_ping() {
    redis_cmd PING 2>&1
}
