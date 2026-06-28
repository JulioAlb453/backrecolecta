# Carga el seed Redis contra Upstash (o cualquier Redis remoto con TLS).
# Uso desde PowerShell en la raíz del repo:
#   .\docker\redis\init-scripts\load-redis-upstash.ps1 `
#     -RedisHost "popular-lamb-129664.upstash.io" `
#     -Password "TU_PASSWORD"

param(
    [Parameter(Mandatory = $true)]
    [string]$RedisHost,

    [Parameter(Mandatory = $true)]
    [string]$Password,

    [int]$Port = 6379,
    [int]$Db = 0
)

$ErrorActionPreference = "Stop"

$repoRoot = Resolve-Path (Join-Path $PSScriptRoot "..\..\..")
$redisDir = Join-Path $repoRoot "docker\redis"
$seedLink = Join-Path $redisDir "seeds\redis-seed-latest.txt"

if (-not (Test-Path $seedLink)) {
    Write-Error "Seed no encontrado. Ejecuta primero: bash docker/redis/init-scripts/generate-seed-data.sh"
}

$hostClean = $RedisHost -replace '^https?://', '' -replace '/$', ''

Write-Host "[INFO] Cargando seed contra ${hostClean}:${Port} (TLS)..." -ForegroundColor Cyan

docker run --rm `
    -e "REDISCLI_AUTH=$Password" `
    -v "${redisDir}:/redis" `
    -w /redis/init-scripts `
    redis:7.2-alpine `
    sh load-redis.sh --tls $hostClean $Port $Password $Db

if ($LASTEXITCODE -ne 0) {
    exit $LASTEXITCODE
}

Write-Host "[OK] Seed cargado. Validando contrato..." -ForegroundColor Green

docker run --rm `
    -e "REDISCLI_AUTH=$Password" `
    -v "${redisDir}:/redis" `
    -w /redis/init-scripts `
    redis:7.2-alpine `
    sh verify-redis.sh --tls $hostClean $Port $Password $Db

exit $LASTEXITCODE
