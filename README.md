
# Backend - Review App

...


## Tecnologías

* Go
* Gin
* GORM
* SQLite


## API Documentación

...

## Variables De Entorno

Para ejecutar este proyecto, necesitarás agregar las siguientes variables de entorno en tu `.env` en la raíz del proyecto

- `PORT` Puerto del localhost

Para trabajar con una base de datos remota usando Turso (opcional):
- `TURSO_DATABASE_URL`
- `TURSO_AUTH_TOKEN`

## Ejecución

Clona el proyecto

```bash
  git clone https://github.com/NetKBs/backend-review-app.git
```

Accede a la ruta del proyecto

```bash
  cd backend-review-app
```

Crea las variables de entorno en tu .env

```bash
  touch .env
```

Instala las dependencias

```bash
  go mod download
```

Inicia el servidor
```bash
  go run main.go
```



