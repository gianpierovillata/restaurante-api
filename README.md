# 🍽️ Restaurante API

¡Bienvenido a la API del restaurante! Esta API está construida en Go con Gin y sirve datos de comidas, proveedores y cartas.

---

## 🚀 Qué hace esta API

- `GET /` → Mensaje de bienvenida
- `GET /comidas` → Lista de comidas
- `GET /proveedores` → Lista de proveedores
- `GET /cartas` → Lista de cartas con menús

---

## 🧱 Tecnologías usadas

- 🟦 **Go**
- 🌐 **Gin** para el servidor HTTP
- 📦 Estructura con paquetes:
  - `Api/` → controladores HTTP
  - `Domain/` → modelos de datos
  - `Repository/` → datos de ejemplo
  - `Services/` → lógica de negocio
- 🛠️ `go.mod` para manejar dependencias

---

## ▶️ Cómo levantar la API

1. Abre la terminal en la carpeta del proyecto:
   ```bash
   cd /home/piero/Descargas/restaurante-api
   ```

2. Instala las dependencias si es necesario:
   ```bash
   go mod tidy
   ```

3. Ejecuta el servidor:
   ```bash
   go run main.go
   ```

4. El servidor arrancará en:
   ```text
   http://localhost:8081
   ```

---

## 📡 Cómo consumir los endpoints

Puedes usar Postman, Insomnia, curl o el navegador.

### Ejemplos con `curl`

- Obtener mensaje de bienvenida:
  ```bash
  curl http://localhost:8081/
  ```

- Obtener comidas:
  ```bash
  curl http://localhost:8081/comidas
  ```

- Obtener proveedores:
  ```bash
  curl http://localhost:8081/proveedores
  ```

- Obtener cartas:
  ```bash
  curl http://localhost:8081/cartas
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

### `/cartas`

```json
[
  {
    "id": 1,
    "nombre": "Carta de Verano",
    "fecha": "2024-06-01",
    "menu": [
      {
        "id": 1,
        "nombre": "Ensalada César",
        "descripcion": "Lechuga, pollo, crutones y aderezo César",
        "precio": 8.99,
        "ingredientes": ["Lechuga", "Pollo", "Crutones", "Aderezo César"],
        "promocion": "VERANO"
      }
    ]
  }
]
```

---

## 🛠️ Qué retorna cada endpoint

- `GET /comidas` → arreglo de `Comida`
- `GET /proveedores` → arreglo de `Proveedor`
- `GET /cartas` → arreglo de `Carta` con un campo `menu` que contiene comidas

---

## 📁 Estructura actual del proyecto

- `main.go` → servidor y rutas
- `Api/` → controladores HTTP
- `Domain/` → definiciones de los modelos
- `Repository/` → datos de ejemplo estáticos
- `Services/` → funciones que llaman a los repositorios

---

## 🌱 Qué se puede agregar después

- ✅ CRUD completo para `comidas`, `proveedores` y `cartas`
- ✅ Conexión a base de datos (PostgreSQL, MySQL, SQLite)
- ✅ Autenticación con JWT
- ✅ Validaciones de payload
- ✅ Swagger / OpenAPI

---

## 💡 Nota final
Esta versión no incluye Swagger automáticamente, pero la estructura ya está preparada para agregarlo más adelante.
