package servicios

import (
	"database/sql"
	"fmt"

	"gestion-biblioteca/modelos"
)

type GestorBiblioteca struct {
	db *sql.DB
}

func NuevoGestor(db *sql.DB) *GestorBiblioteca {
	return &GestorBiblioteca{db: db}
}

func (g *GestorBiblioteca) AgregarLibro(libro modelos.Libro) error {
	consulta := "INSERT INTO libros (titulo, autor, disponible) VALUES ($1, $2, $3)"
	_, err := g.db.Exec(consulta, libro.Titulo, libro.Autor, true)
	return err
}

func (g *GestorBiblioteca) ObtenerLibros() ([]modelos.Libro, error) {
	consulta := "SELECT id, titulo, autor, disponible FROM libros"
	filas, err := g.db.Query(consulta)
	if err != nil {
		return nil, err
	}
	defer filas.Close()

	var libros []modelos.Libro
	for filas.Next() {
		var libro modelos.Libro
		err := filas.Scan(&libro.ID, &libro.Titulo, &libro.Autor, &libro.Disponible)
		if err != nil {
			return nil, err
		}
		libros = append(libros, libro)
	}

	return libros, nil
}

func (g *GestorBiblioteca) ActualizarLibro(libro modelos.Libro) error {
	consulta := "UPDATE libros SET titulo = $1, autor = $2, disponible = $3 WHERE id = $4"
	resultado, err := g.db.Exec(consulta, libro.Titulo, libro.Autor, libro.Disponible, libro.ID)
	if err != nil {
		return err
	}

	filasAfectadas, err := resultado.RowsAffected()
	if err != nil {
		return err
	}
	if filasAfectadas == 0 {
		return fmt.Errorf("no se encontró el libro con ID %d", libro.ID)
	}

	return nil
}

func (g *GestorBiblioteca) EliminarLibro(id int) error {
	consulta := "DELETE FROM libros WHERE id = $1"
	resultado, err := g.db.Exec(consulta, id)
	if err != nil {
		return err
	}

	filasAfectadas, err := resultado.RowsAffected()
	if err != nil {
		return err
	}
	if filasAfectadas == 0 {
		return fmt.Errorf("no se encontró el libro con ID %d", id)
	}

	return nil
}
