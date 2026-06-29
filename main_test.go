package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"gestion-biblioteca/config"
	"gestion-biblioteca/handlers"
	"gestion-biblioteca/servicios"
)

func crearManejadorDePrueba(t *testing.T) *handlers.ManejadorLibros {
	db, err := config.ConectarBD()
	if err != nil {
		t.Skip("No se pudo conectar a la base de datos, saltando prueba:", err)
	}
	t.Cleanup(func() { db.Close() })

	gestor := servicios.NuevoGestor(db)
	return &handlers.ManejadorLibros{Servicio: gestor}
}

func TestObtenerLibrosEndpoint(t *testing.T) {
	manejador := crearManejadorDePrueba(t)

	peticion, err := http.NewRequest("GET", "/libros", nil)
	if err != nil {
		t.Fatal("Error al crear la petición:", err)
	}

	grabadora := httptest.NewRecorder()

	handler := http.HandlerFunc(manejador.ManejarLibros)
	handler.ServeHTTP(grabadora, peticion)

	if grabadora.Code != http.StatusOK {
		t.Errorf("Se esperaba status 200, se obtuvo %d", grabadora.Code)
	}

	tipoContenido := grabadora.Header().Get("Content-Type")
	if tipoContenido != "application/json" {
		t.Errorf("Se esperaba Content-Type 'application/json', se obtuvo '%s'", tipoContenido)
	}

	t.Log("GET /libros respondió correctamente con status 200")
}

func TestAgregarLibroEndpoint(t *testing.T) {
	manejador := crearManejadorDePrueba(t)

	jsonLibro := `{"titulo": "Libro de Prueba", "autor": "Autor Test"}`
	cuerpo := strings.NewReader(jsonLibro)

	peticion, err := http.NewRequest("POST", "/libros", cuerpo)
	if err != nil {
		t.Fatal("Error al crear la petición:", err)
	}
	peticion.Header.Set("Content-Type", "application/json")

	grabadora := httptest.NewRecorder()

	handler := http.HandlerFunc(manejador.ManejarLibros)
	handler.ServeHTTP(grabadora, peticion)

	if grabadora.Code != http.StatusCreated {
		t.Errorf("Se esperaba status 201, se obtuvo %d", grabadora.Code)
	}

	t.Log("POST /libros respondió correctamente con status 201")
}

func TestJSONInvalido(t *testing.T) {
	manejador := crearManejadorDePrueba(t)

	jsonMal := `{esto no es json}`
	cuerpo := strings.NewReader(jsonMal)

	peticion, err := http.NewRequest("POST", "/libros", cuerpo)
	if err != nil {
		t.Fatal("Error al crear la petición:", err)
	}

	grabadora := httptest.NewRecorder()

	handler := http.HandlerFunc(manejador.ManejarLibros)
	handler.ServeHTTP(grabadora, peticion)

	if grabadora.Code != http.StatusBadRequest {
		t.Errorf("Se esperaba status 400, se obtuvo %d", grabadora.Code)
	}

	t.Log("POST /libros con JSON inválido respondió con status 400")
}
