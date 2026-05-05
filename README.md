# 🍽️ Restaurante API

¡Bienvenido a la API del restaurante! Esta API está construida en Go con Gin y sirve datos de comidas, proveedores y cartas. Aquí tienes todo lo necesario para levantarla, consumirla y extenderla con Swagger.

---

## 🚀 Qué hace esta API

- `GET /` → Mensaje de bienvenida
- `GET /comidas` → Devuelve la lista de comidas
- `GET /proveedores` → Devuelve la lista de proveedores
- `GET /cartas` → Devuelve las cartas del restaurante

---

## 🧱 Tecnologías usadas

- 🟦 **Go** (Go 1.26)
- 🌐 **Gin** para el servidor HTTP
- 📦 Estructura limpia con paquetes:
  - `Domain/` → modelos de datos
  - `Repository/` → datos de ejemplo (repositorios)
  - `Services/` → lógica de servicio
- 🛠️ `go.mod` para manejar dependencias

---

## ▶️ Cómo levantar la API

1. Abre la terminal en la carpeta del proyecto:
   ```bash
   cd restaurante-api
   ```

2. Instala las dependencias si es necesario:
   ```bash
   go mod tidy
   ```

3. Compilar el proyecto:
``` bash
    go build
    ```

4. Ejecuta el servidor:
   ```bash
   go run main.go
   ```

5. El servidor arrancará en:
   ```text
   http://localhost:8080
   ```

---

## 📡 Cómo consumir los endpoints

Puedes usar Postman, Insomnia, curl o el navegador.

### Ejemplos con `curl`

- Obtener mensaje de bienvenida:
  ```bash
  curl http://localhost:8080/
  ```

- Obtener comidas:
  ```bash
  curl http://localhost:8080/comidas
  ```

- Obtener proveedores:
  ```bash
  curl http://localhost:8080/proveedores
  ```

- Obtener cartas:
  ```bash
  curl http://localhost:8080/cartas
  ```

---

## 🧾 Ejemplo de respuesta JSON

### `/comidas`

```json
[
  {
    "id": 1,
    "nombre": "Hamburguesa",
    "descripcion": "Hamburguesa con queso, lechuga y tomate",
    "precio": 8.99,
    "ingredientes": ["Carne de res", "Queso", "Lechuga", "Tomate", "Pan"],
    "promocion": "VERANO"
  }
]
```

---

## 🛠️ Servicios que ofrece

Actualmente la API incluye:

- 🍔 `comidas` → Modelo `Comida`
- 🚚 `proveedores` → Modelo `Proveedor`
- 🗂️ `cartas` → Modelo `Carta`

Cada servicio retorna datos definidos en `Repository/`.

---

## 🌱 Qué se puede agregar después

- ✅ Endpoint `POST /reservas`
- ✅ CRUD completo para `comidas`, `proveedores`, `cartas`
- ✅ Conexión a base de datos (PostgreSQL, MySQL, SQLite)
- ✅ Autenticación con JWT
- ✅ Validaciones de payload
- ✅ Swagger / OpenAPI

---

## 🧾 Cómo agregar Swagger

Si quieres documentar la API con Swagger, puedes usar `swaggo/gin-swagger`.

### 1. Instala dependencias

```bash
cd /home/piero/Descargas/restaurante-api
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files
go get -u github.com/swaggo/swag/cmd/swag
```

### 2. Añade comentarios Swagger en `main.go`

```go
// @title Restaurante API
// @version 1.0
// @description API de ejemplo para restaurante
// @host localhost:8080
// @BasePath /

package main

import (
    "restaurante-api/Services"
    "github.com/gin-gonic/gin"
    ginSwagger "github.com/swaggo/gin-swagger"
    swaggerFiles "github.com/swaggo/files"
)
```

### 3. Genera los archivos Swagger

```bash
swag init
```

### 4. Agrega la ruta Swagger

```go
router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
```

### 5. Accede a Swagger UI

Visita:

```text
http://localhost:8080/swagger/index.html
```

---

## 💡 Nota final
Este proyecto ya tiene la estructura básica para separar:

- dominio (`Domain/`)
- repositorios (`Repository/`)
- servicios (`Services/`)
- servidor HTTP (`main.go`)

Si quieres, también puedo agregar ya mismo:

- ✅ ruta `POST /reservas`
- ✅ Swagger configurado con ejemplos
- ✅ documento de endpoints más detallado
