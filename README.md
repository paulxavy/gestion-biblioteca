# Sistema de Gestión de Libros Electrónicos

## Objetivo del Programa

Este proyecto es un **Sistema de Gestión de Libros Electrónicos** desarrollado en el lenguaje de programación **Go (Golang)**, utilizando **PostgreSQL** como sistema de gestión de base de datos relacional. Su propósito principal es permitir el registro y consulta de libros a través de una interfaz web sencilla, aplicando conceptos fundamentales de la programación orientada a objetos, bases de datos relacionales y concurrencia.

## Funcionalidades Principales

| Funcionalidad | Descripción |
|---|---|
| **Registro de libros** | Permite agregar nuevos libros (título y autor) a la base de datos mediante un formulario web |
| **Consulta de inventario** | Muestra todos los libros registrados en una tabla dinámica que se actualiza en tiempo real |
| **API REST** | Endpoints GET /libros y POST /libros para la comunicación entre el frontend y el backend |
| **Concurrencia** | Uso de Goroutines y Channels para simular la generación asíncrona de reportes |
| **Interfaz gráfica** | Página HTML con formulario y tabla de datos, conectada al backend vía fetch |

## Tecnologías Utilizadas

- **Go 1.21+** — Lenguaje de programación del backend
- **PostgreSQL** — Base de datos relacional
- **HTML/CSS/JavaScript** — Interfaz gráfica del usuario
- **net/http** — Servidor web nativo de Go
- **database/sql + pq** — Conexión nativa a PostgreSQL
- **net/http/httptest** — Framework de testing para los endpoints

## Requisitos Previos

- Go instalado (versión 1.21 o superior)
- PostgreSQL instalado y corriendo
- Git (opcional, para clonar el repositorio)

## Instrucciones de Instalación y Ejecución

### 1. Clonar el repositorio

```bash
git clone https://github.com/tu-usuario/gestion-biblioteca.git
cd gestion-biblioteca
```

### 2. Configurar PostgreSQL

Primero, asegúrate de tener PostgreSQL corriendo. Luego, crea la base de datos:

```sql
CREATE DATABASE biblioteca;
```

Conéctate a la base de datos biblioteca y ejecuta el siguiente script para crear la tabla:

```sql
CREATE TABLE IF NOT EXISTS libros (
    id SERIAL PRIMARY KEY,
    titulo VARCHAR(255) NOT NULL,
    autor VARCHAR(255) NOT NULL,
    disponible BOOLEAN DEFAULT true
);
```

### 3. Configurar las variables de entorno

Crea un archivo .env en la raíz del proyecto con tus credenciales:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=tu_contraseña
DB_NAME=biblioteca
```

### 4. Descargar las dependencias

```bash
go mod tidy
```

### 5. Ejecutar el programa

```bash
go run main.go
```

El servidor se iniciará en http://localhost:8080. Abre esa URL en tu navegador para acceder a la interfaz gráfica.

### 6. Ejecutar las pruebas

```bash
go test -v
```

## Estructura del Proyecto

```
gestion-biblioteca/
├── .env
├── .gitignore
├── go.mod
├── go.sum
├── main.go
├── main_test.go
├── schema.sql
├── README.md
├── config/
│   └── database.go
├── modelos/
│   └── libro.go
├── interfaces/
│   └── biblioteca.go
├── servicios/
│   └── gestor.go
├── handlers/
│   └── libros.go
└── static/
    └── index.html
```

## Justificación Técnica

### Programación Orientada a Objetos (POO)

Aunque Go no es un lenguaje orientado a objetos en el sentido tradicional, permite aplicar sus principios fundamentales:

- **Encapsulación**: El struct GestorBiblioteca tiene un atributo privado db, lo que impide el acceso directo desde otros paquetes. Solo los métodos del gestor pueden interactuar con la base de datos.
- **Abstracción**: La interfaz Biblioteca define un contrato con los métodos AgregarLibro y ObtenerLibros, permitiendo que cualquier struct que implemente estos métodos pueda ser utilizado de forma intercambiable.
- **Polimorfismo**: Gracias a las interfaces de Go, es posible sustituir la implementación concreta sin modificar el código que depende de la interfaz.

### PostgreSQL como Base de Datos

Se eligió PostgreSQL por ser un sistema de gestión de bases de datos relacional robusto, de código abierto y ampliamente utilizado en la industria. La conexión se realiza mediante el paquete nativo database/sql de Go junto con el driver lib/pq, sin necesidad de ORMs externos, lo que permite un control directo sobre las consultas SQL.

### Concurrencia

Go tiene soporte nativo para la concurrencia a través de Goroutines y Channels. En este proyecto, cuando se agrega un libro exitosamente, se lanza una goroutine que simula la generación de un reporte asíncrono. Se utiliza un canal (channel) para la comunicación y sincronización entre goroutines, demostrando el modelo de concurrencia de Go.

---

**Autor:** Estudiante Universitario
**Materia:** Programación Avanzada
**Año:** 2026
