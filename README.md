
<p align="center"><h1 align="center">BACKEND-REVIEW-APP</h1></p>
<p align="center">

</p>
<p align="center">
	<img src="https://img.shields.io/github/license/NetKBs/backend-review-app?style=default&logo=opensourceinitiative&logoColor=white&color=0080ff" alt="license">
	<img src="https://img.shields.io/github/last-commit/NetKBs/backend-review-app?style=default&logo=git&logoColor=white&color=0080ff" alt="last-commit">
	<img src="https://img.shields.io/github/languages/top/NetKBs/backend-review-app?style=default&color=0080ff" alt="repo-top-language">
	<img src="https://img.shields.io/github/languages/count/NetKBs/backend-review-app?style=default&color=0080ff" alt="repo-language-count">
</p>
<p align="center"><!-- default option, no dependency badges. -->
</p>
<p align="center">
	<!-- default option, no dependency badges. -->
</p>
<br>


## 🚀 Tecnologías

* Go
* Gin
* GORM
* SQLite

## 🔧 Variables De Entorno

Para ejecutar este proyecto, necesitarás agregar las siguientes variables de entorno en tu `.env` en la raíz del proyecto:

* `PORT`: Puerto en el que se ejecutará el servidor.
* `TURSO_DATABASE_URL` *(opcional)*: URL de la base de datos de TURSO.
* `TURSO_AUTH_TOKEN` *(opcional)*: Token de autenticación para la base de datos de TURSO.
* `GEOAPIFY_KEY`: Key del mapa
* `SECRET_KEY`: Key usada para generar las firmas del jwt
---

## 🗄️ Base de Datos Local

Si no se proporcionan las variables de entorno `TURSO_DATABASE_URL` y `TURSO_AUTH_TOKEN`, el proyecto utilizará automáticamente una base de datos local SQLite.

---

### 🛠️ Makefile

El Makefile incluye comandos útiles para gestionar el proyecto. Por favor asegúrate de tener `make` disponible en tu PATH


* `seed`: realiza un seeding (llenado de datos), en la base de datos
	

Modo de Uso
```sh
❯ make comando
```
---

## 🖥️ Problemas Para Usuarios Windows

Si estás utilizando **Windows**, es posible que encuentres problemas con dependencias que usen **cgo** (por ejemplo, go-sqlite3), debido a la falta de compiladores adecuados. A continuación, se describen los pasos que debes seguir para resolver este problema:

1. **Instala MinGW:** Descarga e instala MinGW desde su sitio web oficial. Asegúrate de seleccionar la opción de instalar los compiladores necesarios.

2. **Agrega los compiladores a la variable PATH:** Una vez instalado MinGW, debes agregar los compiladores a la variable PATH de tu sistema. Puedes hacer esto agregando la ruta de instalación de MinGW a la variable PATH.

---

### ⚙️ Instalación


1. Clonar el repositorio:
```sh
❯ git clone https://github.com/NetKBs/backend-review-app/
```

2. Navegar al directorio del proyecto:
```sh
❯ cd backend-review-app
```

3. Instalar las dependencias del proyecto:
```sh
  ❯ go mod tidy
```

### 🤖 Uso &nbsp; [<img align="center" src="https://img.shields.io/badge/Go-00ADD8.svg?style={badge_style}&logo=go&logoColor=white" />](https://golang.org/)

```sh
❯ go run main.go
```
