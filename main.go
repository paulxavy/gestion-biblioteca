package main

import (
	"fmt"
	"log"
	"net/http"

	"gestion-biblioteca/config"
	"gestion-biblioteca/handlers"
	"gestion-biblioteca/interfaces"
	"gestion-biblioteca/servicios"
)

func main() {
	db, err := config.ConectarBD()
	if err != nil {
		log.Fatal("Error al conectar a la base de datos: ", err)
	}
	defer db.Close()

	gestor := servicios.NuevoGestor(db)

	var _ interfaces.Biblioteca = gestor

	manejador := &handlers.ManejadorLibros{Servicio: gestor}

	http.HandleFunc("/libros", manejador.ManejarLibros)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})

	fmt.Println("Servidor corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
