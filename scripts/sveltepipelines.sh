#!/bin/sh

if [ $# -eq 0 ]; then
	echo "Uso: $0 [install | test | check | unitest | lint]"
	exit 1
fi

# Función para instalar las dependencias
install() {
	echo "Instalando dependencias..."
	npm install
}

# Función para ejecutar los tests
tests() {
	echo "Ejecutando tests..."
	npm run test
}

# Función para ejecutar los tests de cobertura
check() {
	echo "Ejecutando tests de cobertura..."
	npm run check
}

# Función para ejecutar los tests unitarios
unitest() {
	echo "Ejecutando tests unitarios..."
	npm run test:unit
}

# Función para ejecutar el linter
lint() {
	echo "Ejecutando linter..."
	npm run lint
}


# Ejecuta la función correspondiente al parámetro
case $1 in
	install)
		install
		;;
	test)
		tests
		;;
	check)
		check
		;;
	unitest)
		unitest
		;;
	lint)
		lint
		;;
	*)
		echo "Uso: $0 [install | test | check | unitest | lint]"
		exit 1
		;;
esac
