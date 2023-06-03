#!/bin/bash

# Variables
API_BASE_URL="http://localhost:8080"

# Función para realizar una solicitud HTTP
http_request() {
  local method=$1
  local url=$2
  local body=$3
  local response=$(curl -X "$method" -s -H "Content-Type: application/json" -d "$body" "$url")
  echo "$response"
}

# Función para autenticar como un usuario
authenticate_user() {
  local username=$1
  local password=$2
  local response=$(http_request "POST" "$API_BASE_URL/login" "{\"usuario\": \"$username\", \"password\": \"$password\"}")

  if [ -z "$response" ]; then
    echo "Error: No se pudo realizar la solicitud HTTP"
    exit 1
  else
    echo "Se inició sesión correctamente como $username con la contraseña $password y el token $response"
  fi
}

# Ejemplo de uso
authenticate_user "admin" "PiritaAdmin"
authenticate_user "conductor" "conductordemo"
