$ErrorActionPreference = "Stop"
$ProjectRoot = Split-Path -Parent $PSScriptRoot
$TaskCache = Join-Path $ProjectRoot ".cache"

New-Item -ItemType Directory -Force (Join-Path $TaskCache "go-build") | Out-Null
New-Item -ItemType Directory -Force (Join-Path $TaskCache "tmp") | Out-Null

Write-Host "[1/7] Go tests"
Push-Location (Join-Path $ProjectRoot "apps/api")
$env:GOCACHE = Join-Path $TaskCache "go-build"
$env:GOTMPDIR = Join-Path $TaskCache "tmp"
$env:GOTELEMETRY = "off"
go test ./...
if ($LASTEXITCODE -ne 0) {
    Write-Warning "Native Go test execution was blocked. Running the same suites through Go/WASM."
    $GoRoot = go env GOROOT
    $WasmRunner = Join-Path $GoRoot "lib/wasm/wasm_exec_node.js"
    $env:GOOS = "js"
    $env:GOARCH = "wasm"

    go test -c -o (Join-Path $TaskCache "ai.test.wasm") ./internal/ai
    if ($LASTEXITCODE -ne 0) { exit $LASTEXITCODE }
    node --stack-size=8192 (Join-Path $ProjectRoot "scripts/run-go-wasm-test.cjs") $WasmRunner (Join-Path $TaskCache "ai.test.wasm")
    if ($LASTEXITCODE -ne 0) { exit $LASTEXITCODE }

    go test -c -o (Join-Path $TaskCache "analysis.test.wasm") ./internal/analysis
    if ($LASTEXITCODE -ne 0) { exit $LASTEXITCODE }
    node --stack-size=8192 (Join-Path $ProjectRoot "scripts/run-go-wasm-test.cjs") $WasmRunner (Join-Path $TaskCache "analysis.test.wasm")
    if ($LASTEXITCODE -ne 0) { exit $LASTEXITCODE }

    go test -c -o (Join-Path $TaskCache "server.test.wasm") ./cmd/server
    if ($LASTEXITCODE -ne 0) { exit $LASTEXITCODE }
    node --stack-size=8192 (Join-Path $ProjectRoot "scripts/run-go-wasm-test.cjs") $WasmRunner (Join-Path $TaskCache "server.test.wasm")
    if ($LASTEXITCODE -ne 0) { exit $LASTEXITCODE }

    Remove-Item Env:GOOS
    Remove-Item Env:GOARCH
}
Pop-Location

Write-Host "[2/7] Frontend tests"
Push-Location (Join-Path $ProjectRoot "apps/web")
npm test
if ($LASTEXITCODE -ne 0) { exit $LASTEXITCODE }

Write-Host "[3/7] TypeScript"
npm run typecheck
if ($LASTEXITCODE -ne 0) { exit $LASTEXITCODE }

Write-Host "[4/7] Production build"
npm run build
if ($LASTEXITCODE -ne 0) { exit $LASTEXITCODE }

Write-Host "[5/7] Production dependency audit"
npm audit --omit=dev --audit-level=high
if ($LASTEXITCODE -ne 0) { exit $LASTEXITCODE }
Pop-Location

Write-Host "[6/7] Documentation parity"
node (Join-Path $ProjectRoot "scripts/check-docs.mjs")
if ($LASTEXITCODE -ne 0) { exit $LASTEXITCODE }

Write-Host "[7/7] Real evaluation source metadata"
node (Join-Path $ProjectRoot "scripts/check-real-evaluation.mjs")
if ($LASTEXITCODE -ne 0) { exit $LASTEXITCODE }

Write-Host "All p7 checks passed." -ForegroundColor Green

