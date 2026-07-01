package model

// TODO: definir struct Nodo (equivalente a Nodo.java)

// Nodo representa un vértice del grafo (equivalente a Nodo.java).
// En Go no hace falta equals()/hashCode(): un struct con campos
// comparables (como string) ya funciona automáticamente como clave de map.
type Nodo struct {
	Nombre string
}

// NuevoNodo crea un nodo nuevo (equivalente al constructor Nodo(String))
func NuevoNodo(nombre string) Nodo {
	return Nodo{Nombre: nombre}
}

func (n Nodo) String() string {
	return n.Nombre
}
