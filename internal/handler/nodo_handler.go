package handler

// TODO: handlers para agregar/eliminar/renombrar nodo y agregar arista
import (
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"dijkstra-go/internal/model"
)

// AgregarNodo atiende POST /agregar-nodo
func (h *GrafoHandler) AgregarNodo(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	nombre := strings.TrimSpace(r.FormValue("nombre"))

	if nombre == "" {
		nombre = generarNombreAleatorio()
	} else {
		nombre = strings.ToUpper(nombre)
	}

	nodo := model.NuevoNodo(nombre)
	if !h.Grafo.ExisteNodo(nodo) {
		h.Grafo.AgregarNodo(nodo)
		redirigirConMensaje(w, r, "Nodo agregado: "+nodo.Nombre)
		return
	}
	redirigirConMensaje(w, r, "El nodo ya existe")
}

// EliminarNodo atiende POST /eliminar-nodo
func (h *GrafoHandler) EliminarNodo(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	nombre := strings.ToUpper(strings.TrimSpace(r.FormValue("nombre")))

	nodo := model.NuevoNodo(nombre)
	if !h.Grafo.ExisteNodo(nodo) {
		redirigirConMensaje(w, r, "Nodo no encontrado")
		return
	}

	h.Grafo.EliminarNodo(nodo)
	redirigirConMensaje(w, r, "Nodo eliminado: "+nombre)
}

// RenombrarNodo atiende POST /renombrar-nodo
func (h *GrafoHandler) RenombrarNodo(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	actual := strings.ToUpper(strings.TrimSpace(r.FormValue("nombreActual")))
	nuevo := strings.TrimSpace(r.FormValue("nuevoNombre"))

	if nuevo == "" {
		redirigirConMensaje(w, r, "El nombre no puede estar vacío")
		return
	}
	nuevo = strings.ToUpper(nuevo)

	nodoActual := model.NuevoNodo(actual)
	if !h.Grafo.ExisteNodo(nodoActual) {
		redirigirConMensaje(w, r, "Nodo no encontrado")
		return
	}

	nodoNuevo := model.NuevoNodo(nuevo)
	if h.Grafo.ExisteNodo(nodoNuevo) && nuevo != actual {
		redirigirConMensaje(w, r, "El nombre ya existe")
		return
	}

	h.Grafo.RenombrarNodo(nodoActual, nodoNuevo)
	redirigirConMensaje(w, r, "Nodo renombrado a: "+nuevo)
}

// AgregarArista atiende POST /agregar-arista
func (h *GrafoHandler) AgregarArista(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	origen := strings.TrimSpace(r.FormValue("origen"))
	destino := strings.TrimSpace(r.FormValue("destino"))
	peso, err := strconv.Atoi(r.FormValue("peso"))

	if origen == "" || destino == "" {
		redirigirConMensaje(w, r, "Debes indicar origen y destino")
		return
	}
	if err != nil || peso <= 0 {
		redirigirConMensaje(w, r, "El peso debe ser mayor que cero")
		return
	}

	nodoOrigen := model.NuevoNodo(strings.ToUpper(origen))
	nodoDestino := model.NuevoNodo(strings.ToUpper(destino))

	if !h.Grafo.ExisteNodo(nodoOrigen) || !h.Grafo.ExisteNodo(nodoDestino) {
		redirigirConMensaje(w, r, "Ambos nodos deben existir")
		return
	}

	h.Grafo.AgregarArista(nodoOrigen, nodoDestino, peso)
	redirigirConMensaje(w, r, "Arista agregada: "+origen+" -> "+destino+" ("+strconv.Itoa(peso)+")")
}

// --- Auxiliares ---

func generarNombreAleatorio() string {
	letras := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return string(letras[rand.Intn(len(letras))])
}

// redirigirConMensaje redirige a "/" pasando el mensaje como query param.
// Spring usa RedirectAttributes (flash attributes) que Go no tiene de
// fábrica; esta es la forma más simple de lograr el mismo efecto.
func redirigirConMensaje(w http.ResponseWriter, r *http.Request, mensaje string) {
	http.Redirect(w, r, "/?mensaje="+url.QueryEscape(mensaje), http.StatusSeeOther)
}
