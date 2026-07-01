package grafo

// Arista representa una conexión hacia otro nodo.
type Arista struct {
	Destino *Nodo
	Peso    int
}

// Constructor
func NuevaArista(destino *Nodo, peso int) Arista {
	return Arista{
		Destino: destino,
		Peso:    peso,
	}
}	