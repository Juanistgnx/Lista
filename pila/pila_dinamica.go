package pila

/* Definición del struct pila proporcionado por la cátedra. */

// Funciones son solo las que me guardo para mi,los metodos son los que estan en la interfaz

const largo_minimo int = 5                 //Establezco un largo minimo del arreglo,tal como se ve en los ejemplos
const panico string = "La pila esta vacia" //Hago como cte el mensaje de panico en caso de que la pila este vacia

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func CrearPilaDinamica[T any]() Pila[T] {
	pila := new(pilaDinamica[T])
	pila.datos = make([]T, largo_minimo)
	return pila
}

func (pila *pilaDinamica[T]) Apilar(elem T) {
	if pila.cantidad == len(pila.datos) {
		redimensionar[T](pila, pila.cantidad*2)
	}
	pila.datos[pila.cantidad] = elem
	pila.cantidad++
}
func (pila *pilaDinamica[T]) EstaVacia() bool {
	return pila.cantidad == 0
}

func (pila *pilaDinamica[T]) VerTope() T {
	if pila.EstaVacia() {
		panic(panico)
	}
	return pila.datos[pila.cantidad-1]
}

func (pila *pilaDinamica[T]) Desapilar() T {
	if pila.EstaVacia() {
		panic(panico)
	}
	var basura T = pila.datos[pila.cantidad-1]
	pila.cantidad--
	if len(pila.datos) >= 4*pila.cantidad && len(pila.datos) > largo_minimo {
		redimensionar[T](pila, len(pila.datos)/2)
	}
	return basura
}

func redimensionar[T any](pila *pilaDinamica[T], cant int) {
	aux := make([]T, cant)
	copy(aux, pila.datos)
	pila.datos = aux
}

// En las funciones pongo el [T any] y en losm metodos no porque estos ultimos ya tienen el tipo T definido en la interfaz,
//en cambio para nuestras funciones le tenemos que pasar con que tipos vamos a trabajar(con any)
