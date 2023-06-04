#!/bin/bash

# Variables
API_BASE_URL="http://localhost:8080"

# Función para realizar una solicitud HTTP
http_request() {
  local method=$1
  local url=$2
  local body=$3
  local response

  response=$(curl -X "$method" -s -H "Content-Type: application/json" -d "$body" "$url")
  echo "$response"
}

# Función para autenticar como un administrador
authenticate_admin() {
  local username=$1
  local password=$2
  local response
  response=$(http_request "POST" "$API_BASE_URL/login" "{\"usuario\": \"$username\", \"password\": \"$password\"}")

  if [ -z "$response" ]; then
    echo "Error: No se pudo realizar la solicitud HTTP"
    exit 1
  else
    echo "Se inició sesión correctamente como admin $username con la contraseña $password y el token $response"
  fi
}

# Función para autenticar como un conductor
authenticate_conductor() {
  local username=$1
  local password=$2
  local response
  response=$(http_request "POST" "$API_BASE_URL/loginconductor" "{\"usuario\": \"$username\", \"password\": \"$password\"}")

  if [ -z "$response" ]; then
    echo "Error: No se pudo realizar la solicitud HTTP"
    exit 1
  else
    echo "Se inició sesión correctamente como conductor $username con la contraseña $password y el token $response"
  fi
}

# Ejemplo de uso
authenticate_admin "admin" "PiritaAdmin"
echo ""
authenticate_conductor "conductor" "conductordemo"
