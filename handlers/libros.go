package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"gestion-biblioteca/interfaces"
	"gestion-biblioteca/modelos"
)

type ManejadorLibros struct {
	Servicio interfaces.Biblioteca
}

func (m *ManejadorLibros) ManejarLibros(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		m.obtenerLibros(w, r)
	case http.MethodPost:
		m.agregarLibro(w, r)
	case http.MethodPut:
		m.actualizarLibro(w, r)
	case http.MethodDelete:
		m.eliminarLibro(w, r)
	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

func (m *ManejadorLibros) obtenerLibros(w http.ResponseWriter, r *http.Request) {
	libros, err := m.Servicio.ObtenerLibros()
	if err != nil {
		http.Error(w, "Error al obtener los libros", http.StatusInternalServerError)
		log.Println("Error en SELECT:", err)
		return
	}

	if libros == nil {
		libros = []modelos.Libro{}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(libros)
}

func (m *ManejadorLibros) agregarLibro(w http.ResponseWriter, r *http.Request) {
	var libro modelos.Libro

	err := json.NewDecoder(r.Body).Decode(&libro)
	if err != nil {
		http.Error(w, "Error: el JSON enviado no es válido", http.StatusBadRequest)
		return
	}

	if libro.Titulo == "" || libro.Autor == "" {
		http.Error(w, "Error: el título y el autor son obligatorios", http.StatusBadRequest)
		return
	}

	err = m.Servicio.AgregarLibro(libro)
	if err != nil {
		http.Error(w, "Error al guardar el libro en la base de datos", http.StatusInternalServerError)
		log.Println("Error en INSERT:", err)
		return
	}

	canal := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		canal <- fmt.Sprintf("Reporte generado para el libro: '%s' de %s", libro.Titulo, libro.Autor)
	}()

	go func() {
		mensaje := <-canal
		fmt.Println("Reporte:", mensaje)
	}()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"mensaje": "Libro agregado exitosamente",
		"titulo":  libro.Titulo,
		"autor":   libro.Autor,
	})
}

func (m *ManejadorLibros) actualizarLibro(w http.ResponseWriter, r *http.Request) {
	var libro modelos.Libro

	err := json.NewDecoder(r.Body).Decode(&libro)
	if err != nil {
		http.Error(w, "Error: el JSON enviado no es válido", http.StatusBadRequest)
		return
	}

	if libro.ID <= 0 {
		http.Error(w, "Error: el ID del libro es obligatorio", http.StatusBadRequest)
		return
	}

	if libro.Titulo == "" || libro.Autor == "" {
		http.Error(w, "Error: el título y el autor son obligatorios", http.StatusBadRequest)
		return
	}

	err = m.Servicio.ActualizarLibro(libro)
	if err != nil {
		http.Error(w, "Error al actualizar el libro: "+err.Error(), http.StatusInternalServerError)
		log.Println("Error en UPDATE:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"mensaje": "Libro actualizado exitosamente",
	})
}

func (m *ManejadorLibros) eliminarLibro(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Error: el parámetro 'id' es obligatorio", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Error: el ID debe ser un número válido", http.StatusBadRequest)
		return
	}

	err = m.Servicio.EliminarLibro(id)
	if err != nil {
		http.Error(w, "Error al eliminar el libro: "+err.Error(), http.StatusInternalServerError)
		log.Println("Error en DELETE:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"mensaje": "Libro eliminado exitosamente",
	})
}
