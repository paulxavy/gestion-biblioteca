package config

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConectarBD() (*sql.DB, error) {
	godotenv.Load()
	host := os.Getenv("DB_HOST")
	puerto := os.Getenv("DB_PORT")
	usuario := os.Getenv("DB_USER")
	contraseña := os.Getenv("DB_PASSWORD")
	nombreBD := os.Getenv("DB_NAME")

	cadenaConexion := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, puerto, usuario, contraseña, nombreBD,
	)

	db, err := sql.Open("postgres", cadenaConexion)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Conexión exitosa a PostgreSQL")
	return db, nil
}
