package grafo

// Nodo representa un vértice del grafo.
type Nodo struct {
	Nombre string
}

// Constructor
func NuevoNodo(nombre string) *Nodo {
	return &Nodo{
		Nombre: nombre,
	}
}