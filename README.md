# go-web-service-starter

This project is basis for my golang based web services. It is not the best and there are certainly some improvements 
needs to be made.

Only 3 pages exists: Home, Login and Register.

Login and Register pages has required components, however Home page only has some placeholder text. Please see how it looks on `Screenshots` section.

## Usage

Build and run `src/cmd/main.go`

Use `--help` flag to print flags

```
  -db-host string
        postgres database host (default "localhost")
  -db-name string
        postgres database name (default "postgres")
  -db-password string
        postgres database password
  -db-port string
        postgres database port (default "5432")
  -db-username string
        postgres database username (default "postgres")
  -session-secret string
        session secret (default "secret")
```

`-db-password` must be provided to run!

## Screenshots

![Login Screen](https://github.com/oguzhand95/go-web-service-starter/raw/master/screenshots/login.PNG)
![Register Screen](https://github.com/oguzhand95/go-web-service-starter/raw/master/screenshots/register.PNG)
