package model

// TODO: definir struct Arista (equivalente a Arista.java)
// Arista representa una conexión ponderada hacia un nodo destino
// (equivalente a Arista.java).
//
// Nota: aquí Destino es un Nodo por VALOR, no un puntero. Como Nodo
// ya es comparable por valor (su Nombre), no necesitamos punteros
// para identificarlo — simplifica el código frente al Java original.
type Arista struct {
	Destino Nodo
	Peso    int
}

func NuevaArista(destino Nodo, peso int) Arista {
	return Arista{Destino: destino, Peso: peso}
}

func (a Arista) String() string {
	return a.Destino.Nombre
}
