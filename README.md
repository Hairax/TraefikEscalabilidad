# Proyecto API Gateway con Traefik y Docker

Este proyecto implementa un **API Gateway** utilizando **Traefik** como **reverse proxy**, gestionando múltiples microservicios en un entorno Docker.

## 🚀 Tecnologías utilizadas

- **Traefik v2.9** como API Gateway y reverse proxy.
- **Docker Compose** para la gestión de contenedores.
- **Go** para los microservicios.
- **Red de Docker** para comunicación entre servicios.

---

## 📌 Arquitectura del Proyecto

El sistema está compuesto por:

1. **Traefik (API Gateway)**: Encargado de recibir las solicitudes HTTP y enrutar a los microservicios correctos.
2. **Microservicios**: Implementados en **Go**, cada uno escuchando en diferentes rutas.
3. **Red Docker `traefik`**: Permite la comunicación entre los servicios de manera segura y eficiente.

### 🔄 Flujo de una solicitud

1. Un cliente hace una petición a `http://api-gateway.local/service1`.
2. Traefik recibe la petición y la redirige al microservicio correspondiente.
3. El microservicio responde con los datos solicitados.

---

## 📂 Estructura del proyecto

```
📁 traefik-api-gateway
│── 📁 service1
│── 📁 service2
│── 📁 service3
│── 📄 docker-compose.yml
│── 📄 README.md
```

---

## 🛠️ Configuración de `docker-compose.yml`

```yaml
services:
  reverse-proxy:
    image: traefik:v2.9
    command: --api.insecure=true --providers.docker
    networks:
      - traefik
    ports:
      - "80:80"
      - "8080:8080"
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock

networks:
  traefik:
    external: true
```

### 🔍 Explicación de la configuración

- **`image: traefik:v2.9`** → Version de traefik a utilizar.
- **`command: --api.insecure=true --providers.docker`** → Activa la API web y la integración con Docker.
- **`networks: traefik`** → Conecta el servicio a la red Docker `traefik`.
- **`ports`**:
  - `80:80` → Para recibir tráfico HTTP.
  - `8080:8080` → Para acceder al dashboard de Traefik.
- **`volumes`**:
  - `- /var/run/docker.sock:/var/run/docker.sock` → Permite que Traefik detecte cambios en los contenedores.

---

## 🔧 Instalación y ejecución

### 1️⃣ Clonar el repositorio

```sh
git clone https://github.com/Hairax/TraefikEscalabilidad
cd TraefikEscalabilidad
```

### 2️⃣ Ejecutar Docker Compose

```sh
docker-compose up -d --build
```

Esto levantará Traefik, ahora realizalo en cada carpeta de los servicios.

### 3️⃣ Acceder al Dashboard de Traefik

Abre en tu navegador:

```
http://localhost:8080
```

---

## 🛑 Detener y eliminar los contenedores

Para detener los servicios:

```sh
docker-compose down
```

Para eliminar todos los contenedores, redes y volúmenes:

```sh
docker-compose down -v
```

---

## 📝 Notas adicionales

- Si estás usando **Windows/Mac con Docker Desktop**, asegurate de que la red `traefik` existe:
  ```sh
  docker network create traefik
  ```
- Si querés usar HTTPS con Let's Encrypt, se recomienda configurar Traefik con certificados TLS.
