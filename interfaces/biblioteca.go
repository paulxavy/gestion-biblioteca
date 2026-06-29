package interfaces

import "gestion-biblioteca/modelos"

type Biblioteca interface {
	AgregarLibro(libro modelos.Libro) error
	ObtenerLibros() ([]modelos.Libro, error)
	ActualizarLibro(libro modelos.Libro) error
	EliminarLibro(id int) error
}
