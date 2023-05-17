# 🚀 Pirita Backend 🚀

¡Bienvenidos al proyecto "Pirita"! Este es el servidor backend escrito en Go que utiliza una base de datos SQLite para gestionar datos de conductores, contratos, pagos, vehículos y viajes.

> Nota: Te recomendamos instalar [podman](https://podman.io) y [podman desktop](https://podman-desktop.io/) en lugar de Docker para que trabajes más tranquil@ :)

## 📂 Estructura del proyecto

El proyecto se organiza de la siguiente manera:

```
.
├── db
│   └── db.sqlite
├── Dockerfile
├── go.mod
├── go.sum
├── main.go
├── Makefile
├── models
│   ├── conductor.go
│   ├── contrato.go
│   ├── pago.go
│   ├── vehiculo.go
│   └── viaje.go
└── routes
    ├── conductor.go
    ├── contrato.go
    ├── observability.go
    ├── pago.go
    ├── vehiculo.go
    └── viaje.go
```

## 🔧 Compilación y ejecución

### Local 🖥

Para compilar un binario productivo y distribuible, podemos hacerlo con el comando:

```bash
make build
```

Para probar el programa sin producir un binario productivo:

```bash
make debug
```

### Docker 🐋

Aunque en este proyecto hacemos referencia a Docker y proporcionamos un Dockerfile y comandos make para trabajar con Docker, recomendamos usar Podman y Podman Desktop para manejar los contenedores, ya que consideramos que es una opción más segura y flexible.

Para construir la imagen Docker (o Podman) del proyecto, puedes usar el comando:

```bash
make docker-build
```

Esto creará una imagen Docker con el nombre `pirita_backend`.

Para ejecutar la imagen en un contenedor, puedes usar el comando:

```bash
make docker-run
```

Esto iniciará un contenedor Docker en segundo plano con el nombre pirita_backend y expondrá el puerto 3000.

Para detener el contenedor, puedes usar el comando:

```bash
make docker-stop
```

Y para eliminar el contenedor, puedes usar el comando:

```bash
make docker-rm
```
