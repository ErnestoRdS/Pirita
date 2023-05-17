# 🌐 Pirita Frontend 🌐

¡Bienvenidos al proyecto "Pirita" frontend! Este es el servidor frontend que utiliza Nginx para servir nuestra aplicación.

## 🔧 Compilación y ejecución

### Local 🖥

Para compilar y ejecutar el proyecto localmente, utilizaremos make. Primero, vamos a compilar el proyecto con el comando:

```bash
make build
```

Esto instalará las dependencias necesarias y construirá la aplicación.

Para iniciar la aplicación en modo de desarrollo, puedes usar el comando:

```bash
make debug
```

Antes de hacer push a git, se recomienda ejecutar las pruebas y verificaciones necesarias con el comando:


```bash
make test_before_push
```

Y durante el desarrollo, puedes correr pruebas y verificaciones rápidas con el comando:

```bash
make test
```

## Docker 🐋

Aunque en este proyecto hacemos referencia a Docker y proporcionamos un Dockerfile y comandos make para trabajar con Docker, recomendamos usar Podman y Podman Desktop para manejar los contenedores, ya que consideramos que es una opción más segura y flexible.

Para construir la imagen Docker (o Podman) del proyecto, puedes usar el comando:

```bash
make docker-build
```

Esto creará una imagen Docker con el nombre `pirita_frontend`.

Para ejecutar la imagen en un contenedor, puedes usar el comando:

```bash
make docker-run
```
Esto iniciará un contenedor Docker en segundo plano con el nombre `pirita_frontend` y expondrá el puerto 80.

Para detener el contenedor, puedes usar el comando:

```bash
make docker-stop
```

Y para eliminar el contenedor, puedes usar el comando:

```bash
make docker-rm
```

