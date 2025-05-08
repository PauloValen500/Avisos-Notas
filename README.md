# 📝 Tablero de Avisos y Notas

Aplicación desarrollada en **Go** con interfaz web en HTML y CSS, que permite agregar, editar y eliminar notas y avisos de forma rápida y sencilla.

---

## 🧠 Descripción general

Este proyecto forma parte del curso _Tópicos para el Despliegue de Aplicaciones_. Su objetivo es implementar una aplicación cliente-servidor funcional y ligera, contenedorizada con Docker, versionada con GitHub y desplegada automáticamente mediante CI/CD.

---

## 🔧 Tecnologías utilizadas

- **Go**: Lenguaje backend
- **HTML + CSS + JS**: Interfaz web
- **SQLite**: Base de datos ligera integrada
- **Docker**: Contenerización y despliegue
- **GitHub Actions**: Automatización CI/CD
- **Docker Hub**: Repositorio de imágenes de contenedor

---

## 🚀 Cómo ejecutar el proyecto en un contenedor Docker (Linux)

> ✅ Asegúrate de tener Docker instalado y corriendo (Docker Desktop en Windows/Mac o Docker Engine en Linux).

### 1️⃣ Construir la imagen Docker

Abre la terminal en la carpeta raíz del proyecto y ejecuta:

```bash
docker build -t notas-y-avisos .
```

### 2️⃣ Ejecutar el contenedor

```bash
docker run -d -p 8080:8080 --name tablero \
-v ${PWD}/notas-y-avisos.db:/app/notas-y-avisos.db \
notas-y-avisos
```

> El volumen garantiza que los datos de la base de datos no se pierdan entre reinicios del contenedor.

### 3️⃣ Abrir en el navegador

```text
http://localhost:8080
```

---

## ⚙️ CI/CD con GitHub Actions y Docker Hub

Este repositorio incluye un flujo de trabajo automático configurado en `.github/workflows/ci.yml`. Cada vez que se haga un **push a la rama `main`**, se ejecuta el siguiente proceso:

1. Clona el repositorio
2. Inicia sesión en Docker Hub con los secrets del repositorio (`DOCKER_USERNAME` y `DOCKER_PASSWORD`)
3. Construye la imagen del proyecto
4. Publica automáticamente la imagen en Docker Hub:

🔗 [https://hub.docker.com/repository/docker/angelisrael03/notas-y-avisos](https://hub.docker.com/repository/docker/angelisrael03/notas-y-avisos)

---

## 📂 Estructura del repositorio

```
.
├── .github/
│   └── workflows/
│       └── ci.yml         # Configuración CI/CD
├── Dockerfile             # Instrucciones para contenerizar la app
├── go.mod / go.sum        # Dependencias Go
├── main.go                # Lógica backend del servidor
├── index.html             # Interfaz de usuario
├── styles.css             # Estilos visuales
├── notas-y-avisos.db      # Base de datos SQLite (generada al ejecutar)
└── README.md              # Este archivo
```

---

## 👥 Colaboradores

- Paulo Valenzuela  
- Aldo299  
- MichCelis  
- AndreAguilera10  
- AngelIsrael03

---

## 🛠️ Recomendaciones

- Usa `docker ps -a` para ver contenedores existentes.
- Usa `docker stop tablero && docker rm tablero` si necesitas reiniciar la imagen.
- Si cambias el código, reconstruye la imagen con `docker build`.

---

## ✅ Estado del CI/CD

> Consulta el historial en la pestaña **Actions** del repositorio:  
> `https://github.com/PauloValen500/Avisos-Notas/actions`

---
