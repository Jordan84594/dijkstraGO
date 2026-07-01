package grafo

// Grafo utiliza una lista de adyacencia.
type Grafo struct {
	ListaAdyacencia map[*Nodo][]Arista
}

// Constructor
func NuevoGrafo() *Grafo {

	return &Grafo{
		ListaAdyacencia: make(map[*Nodo][]Arista),
	}

}

// Agrega un nodo al grafo.
func (g *Grafo) AgregarNodo(n *Nodo) {

	if _, existe := g.ListaAdyacencia[n]; !existe {
		g.ListaAdyacencia[n] = []Arista{}
	}

}

// Agrega una arista dirigida.
func (g *Grafo) AgregarArista(origen, destino *Nodo, peso int) {

	g.AgregarNodo(origen)
	g.AgregarNodo(destino)

	g.ListaAdyacencia[origen] = append(
		g.ListaAdyacencia[origen],
		NuevaArista(destino, peso),
	)

}

// Devuelve los vecinos de un nodo.
func (g *Grafo) ObtenerVecinos(n *Nodo) []Arista {
	return g.ListaAdyacencia[n]
}

// Devuelve todos los nodos del grafo.
func (g *Grafo) ObtenerNodos() []*Nodo {

	nodos := []*Nodo{}

	for nodo := range g.ListaAdyacencia {
		nodos = append(nodos, nodo)
	}

	return nodos
}