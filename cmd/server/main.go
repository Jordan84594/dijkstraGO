package main

import (
	"html/template"
	"log"
	"net/http"

	"dijkstra-go/internal/handler"
)

func main() {
	tmpl := template.Must(template.ParseGlob("web/templates/*.html"))
	h := handler.NuevoGrafoHandler(tmpl)

	mux := http.NewServeMux()

	// Sintaxis "MÉTODO /ruta" disponible desde Go 1.22 (ver go.mod)
	mux.HandleFunc("GET /{$}", h.MostrarFormulario)
	mux.HandleFunc("POST /calcular", h.Calcular)
	mux.HandleFunc("POST /agregar-nodo", h.AgregarNodo)
	mux.HandleFunc("POST /eliminar-nodo", h.EliminarNodo)
	mux.HandleFunc("POST /renombrar-nodo", h.RenombrarNodo)
	mux.HandleFunc("POST /agregar-arista", h.AgregarArista)

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	log.Println("Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
