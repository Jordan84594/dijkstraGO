package algoritmo

// TODO: implementar el algoritmo de Dijkstra (equivalente a Dijkstra.java)

import (
	"container/heap"
	"math"

	"dijkstra-go/internal/model"
)

// Infinito representa una distancia no alcanzable (equivalente a Integer.MAX_VALUE en Java)
const Infinito = math.MaxInt

// item es un par (nodo, distancia) que vive dentro de la cola de prioridad
type item struct {
	nodo      model.Nodo
	distancia int
}

// colaPrioridad implementa heap.Interface para que container/heap
// pueda ordenar los nodos por distancia (menor distancia primero)
type colaPrioridad []item

func (cp colaPrioridad) Len() int            { return len(cp) }
func (cp colaPrioridad) Less(i, j int) bool  { return cp[i].distancia < cp[j].distancia }
func (cp colaPrioridad) Swap(i, j int)       { cp[i], cp[j] = cp[j], cp[i] }
func (cp *colaPrioridad) Push(x any)         { *cp = append(*cp, x.(item)) }
func (cp *colaPrioridad) Pop() any {
	old := *cp
	n := len(old)
	ultimo := old[n-1]
	*cp = old[:n-1]
	return ultimo
}

// Dijkstra encapsula el estado del algoritmo (equivalente a Dijkstra.java)
type Dijkstra struct {
	grafo      *model.Grafo
	distancias map[model.Nodo]int
	anteriores map[model.Nodo]*model.Nodo
}

// NuevoDijkstra crea una nueva instancia (equivalente al constructor Dijkstra(Grafo))
func NuevoDijkstra(grafo *model.Grafo) *Dijkstra {
	return &Dijkstra{grafo: grafo}
}

// Calcular ejecuta el algoritmo de Dijkstra desde un nodo origen
func (d *Dijkstra) Calcular(origen model.Nodo) {
	d.distancias = make(map[model.Nodo]int)
	d.anteriores = make(map[model.Nodo]*model.Nodo)

	// Inicializar distancias
	for nodo := range d.grafo.ListaAdyacencia {
		d.distancias[nodo] = Infinito
		d.anteriores[nodo] = nil
	}
	d.distancias[origen] = 0

	cola := &colaPrioridad{{nodo: origen, distancia: 0}}
	heap.Init(cola)

	for cola.Len() > 0 {
		actual := heap.Pop(cola).(item)

		// Entrada obsoleta: este nodo ya se procesó con una distancia mejor.
		// (Java resolvía esto con cola.remove(vecino) + cola.add(vecino);
		// en Go es más idiomático simplemente ignorar la entrada vieja).
		if actual.distancia > d.distancias[actual.nodo] {
			continue
		}

		for _, arista := range d.grafo.ObtenerVecinos(actual.nodo) {
			vecino := arista.Destino
			nuevaDistancia := d.distancias[actual.nodo] + arista.Peso

			if nuevaDistancia < d.distancias[vecino] {
				d.distancias[vecino] = nuevaDistancia

				origenCopia := actual.nodo
				d.anteriores[vecino] = &origenCopia

				heap.Push(cola, item{nodo: vecino, distancia: nuevaDistancia})
			}
		}
	}
}

// GetDistancia devuelve la distancia mínima hasta un nodo
func (d *Dijkstra) GetDistancia(destino model.Nodo) int {
	if dist, existe := d.distancias[destino]; existe {
		return dist
	}
	return Infinito
}

// GetDistancias devuelve todas las distancias calculadas
func (d *Dijkstra) GetDistancias() map[model.Nodo]int {
	return d.distancias
}

// ObtenerCamino devuelve el camino mínimo desde el origen hasta el destino
func (d *Dijkstra) ObtenerCamino(destino model.Nodo) []model.Nodo {
	camino := []model.Nodo{}

	dist, existe := d.distancias[destino]
	if !existe || dist == Infinito {
		return camino
	}

	for actual := &destino; actual != nil; actual = d.anteriores[*actual] {
		camino = append(camino, *actual)
	}

	// Invertir el slice (Collections.reverse en Java)
	for i, j := 0, len(camino)-1; i < j; i, j = i+1, j-1 {
		camino[i], camino[j] = camino[j], camino[i]
	}

	return camino
}

// GetAnteriores devuelve el nodo anterior de cada vértice
func (d *Dijkstra) GetAnteriores() map[model.Nodo]*model.Nodo {
	return d.anteriores
}