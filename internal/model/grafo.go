package model

// Grafo representa un grafo no dirigido mediante lista de adyacencia
// (equivalente a Grafo.java)
type Grafo struct {
	ListaAdyacencia map[Nodo][]Arista
}

// NuevoGrafo crea un grafo vacío (equivalente al constructor Grafo())
func NuevoGrafo() *Grafo {
	return &Grafo{ListaAdyacencia: make(map[Nodo][]Arista)}
}

// AgregarNodo agrega un nodo al grafo
func (g *Grafo) AgregarNodo(nodo Nodo) {
	if _, existe := g.ListaAdyacencia[nodo]; !existe {
		g.ListaAdyacencia[nodo] = []Arista{}
	}
}

// AgregarArista agrega una arista entre dos nodos en ambos sentidos
func (g *Grafo) AgregarArista(origen, destino Nodo, peso int) {
	g.AgregarNodo(origen)
	g.AgregarNodo(destino)

	g.ListaAdyacencia[origen] = append(g.ListaAdyacencia[origen], NuevaArista(destino, peso))
	g.ListaAdyacencia[destino] = append(g.ListaAdyacencia[destino], NuevaArista(origen, peso))
}

// ObtenerVecinos devuelve los vecinos de un nodo
func (g *Grafo) ObtenerVecinos(nodo Nodo) []Arista {
	return g.ListaAdyacencia[nodo]
}

// ExisteNodo verifica si existe un nodo
func (g *Grafo) ExisteNodo(nodo Nodo) bool {
	_, existe := g.ListaAdyacencia[nodo]
	return existe
}

// EliminarNodo quita un nodo del grafo (usado por eliminar-nodo)
func (g *Grafo) EliminarNodo(nodo Nodo) {
	delete(g.ListaAdyacencia, nodo)
}

// RenombrarNodo cambia el nombre de un nodo y actualiza todas las
// aristas que apuntaban a él (Java lo resolvía "gratis" porque el
// objeto Nodo mantenía su identidad; en Go, como Nodo es un valor,
// hay que reescribir las referencias a mano)
func (g *Grafo) RenombrarNodo(actual, nuevo Nodo) {
	aristas := g.ListaAdyacencia[actual]
	delete(g.ListaAdyacencia, actual)
	g.ListaAdyacencia[nuevo] = aristas

	for nodo, lista := range g.ListaAdyacencia {
		for i, a := range lista {
			if a.Destino == actual {
				lista[i].Destino = nuevo
			}
		}
		g.ListaAdyacencia[nodo] = lista
	}
}