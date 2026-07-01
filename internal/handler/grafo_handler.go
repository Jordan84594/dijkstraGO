package handler

// TODO: handlers para mostrar formulario y calcular camino mínimo
import (
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"dijkstra-go/internal/algoritmo"
	"dijkstra-go/internal/model"
)

// GrafoHandler agrupa las dependencias de los handlers HTTP
// (equivalente a DijkstraWebController.java)
type GrafoHandler struct {
	Grafo     *model.Grafo
	Templates *template.Template
}

// NuevoGrafoHandler crea el handler con un grafo de demostración precargado
func NuevoGrafoHandler(tmpl *template.Template) *GrafoHandler {
	return &GrafoHandler{
		Grafo:     construirGrafoDemo(),
		Templates: tmpl,
	}
}

// datosVista es lo que se envía a la plantilla (equivalente al Model de Spring)
type datosVista struct {
	Nodos                 []string
	Aristas               []map[string]any
	Inicio                string
	Fin                   string
	Camino                string
	Distancia             string
	MostrarModalResultado bool
	Mensaje               string
}

// MostrarFormulario atiende GET /
func (h *GrafoHandler) MostrarFormulario(w http.ResponseWriter, r *http.Request) {
	datos := datosVista{
		Nodos:   obtenerNombres(h.Grafo),
		Aristas: obtenerAristas(h.Grafo),
		Mensaje: r.URL.Query().Get("mensaje"),
	}
	h.Templates.ExecuteTemplate(w, "index.html", datos)
}

// Calcular atiende POST /calcular
func (h *GrafoHandler) Calcular(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	inicio := r.FormValue("inicio")
	fin := r.FormValue("fin")

	nodos := indexarNodos(h.Grafo)
	origen, origenExiste := nodos[strings.ToUpper(inicio)]
	destino, destinoExiste := nodos[strings.ToUpper(fin)]

	datos := datosVista{
		Nodos:                 obtenerNombres(h.Grafo),
		Aristas:               obtenerAristas(h.Grafo),
		Inicio:                inicio,
		Fin:                   fin,
		MostrarModalResultado: true,
	}

	if !origenExiste || !destinoExiste {
		datos.Camino = "Nodo inválido"
		datos.Distancia = "No disponible"
		h.Templates.ExecuteTemplate(w, "index.html", datos)
		return
	}

	dijkstra := algoritmo.NuevoDijkstra(h.Grafo)
	dijkstra.Calcular(origen)
	camino := dijkstra.ObtenerCamino(destino)

	nombres := make([]string, 0, len(camino))
	for _, nodo := range camino {
		nombres = append(nombres, nodo.Nombre)
	}
	datos.Camino = strings.Join(nombres, " -> ")

	distancia := dijkstra.GetDistancia(destino)
	if distancia == algoritmo.Infinito {
		datos.Distancia = "No hay camino"
	} else {
		datos.Distancia = strconv.Itoa(distancia)
	}

	h.Templates.ExecuteTemplate(w, "index.html", datos)
}

// --- Funciones auxiliares ---

func construirGrafoDemo() *model.Grafo {
	g := model.NuevoGrafo()
	nodos := map[string]model.Nodo{}
	for _, nombre := range []string{"A", "B", "C", "D", "E"} {
		n := model.NuevoNodo(nombre)
		g.AgregarNodo(n)
		nodos[nombre] = n
	}

	g.AgregarArista(nodos["A"], nodos["B"], 4)
	g.AgregarArista(nodos["A"], nodos["C"], 2)
	g.AgregarArista(nodos["B"], nodos["C"], 3)
	g.AgregarArista(nodos["B"], nodos["D"], 3)
	g.AgregarArista(nodos["C"], nodos["D"], 6)
	g.AgregarArista(nodos["D"], nodos["E"], 1)

	return g
}

func indexarNodos(g *model.Grafo) map[string]model.Nodo {
	mapa := make(map[string]model.Nodo)
	for nodo := range g.ListaAdyacencia {
		mapa[nodo.Nombre] = nodo
	}
	return mapa
}

func obtenerNombres(g *model.Grafo) []string {
	nombres := make([]string, 0, len(g.ListaAdyacencia))
	for nodo := range g.ListaAdyacencia {
		nombres = append(nombres, nodo.Nombre)
	}
	return nombres
}

func obtenerAristas(g *model.Grafo) []map[string]any {
	aristas := []map[string]any{}
	for origen, lista := range g.ListaAdyacencia {
		for _, arista := range lista {
			aristas = append(aristas, map[string]any{
				"origen":  origen.Nombre,
				"destino": arista.Destino.Nombre,
				"peso":    arista.Peso,
			})
		}
	}
	return aristas
}
