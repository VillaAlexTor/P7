@echo off
setlocal
cd /d "%~dp0"

if not defined P7_API_PORT set "P7_API_PORT=8080"
if not defined P7_WEB_PORT set "P7_WEB_PORT=3000"
if not defined NEXT_PUBLIC_API_URL set "NEXT_PUBLIC_API_URL=http://localhost:%P7_API_PORT%"
if not defined P7_ALLOWED_ORIGINS set "P7_ALLOWED_ORIGINS=http://localhost:%P7_WEB_PORT%"

echo.
echo ========================================
echo   p7 - Inicio completo con Docker
echo ========================================
echo.

where docker >nul 2>&1
if errorlevel 1 (
  echo [ERROR] Docker no esta instalado o no esta en PATH.
  echo Instala Docker Desktop y vuelve a ejecutar este archivo.
  pause
  exit /b 1
)

docker info >nul 2>&1
if errorlevel 1 (
  echo [ERROR] Docker no esta iniciado.
  echo Abre Docker Desktop, espera a que termine de arrancar y vuelve a intentarlo.
  pause
  exit /b 1
)

docker compose version >nul 2>&1
if errorlevel 1 (
  echo [ERROR] Docker Compose no esta disponible.
  echo Actualiza Docker Desktop y vuelve a intentarlo.
  pause
  exit /b 1
)

echo [1/2] Construyendo e iniciando web, API y Ollama...
echo La primera ejecucion puede tardar porque descargara el modelo de IA.
docker compose up --build --wait
if errorlevel 1 goto :error

echo.
echo [2/2] p7 esta listo en http://localhost:%P7_WEB_PORT%
echo Para detenerlo ejecuta: docker compose down
start "" "http://localhost:%P7_WEB_PORT%"
exit /b 0

:error
echo.
echo [ERROR] No se pudo iniciar p7. Revisa el detalle anterior.
echo Puedes consultar los registros con: docker compose logs
pause
exit /b 1
