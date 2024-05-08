package lista

const PANICO = "La lista esta vacia"
const PANICOITER = "El iterador termino de iterar"

type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}
type nodoLista[T any] struct {
	dato      T
	siguiente *nodoLista[T]
}
type iterListaEnlazada[T any] struct {
	lista    *listaEnlazada[T]
	anterior *nodoLista[T]
	actual   *nodoLista[T]
}

func CrearListaEnlazada[T any]() Lista[T] {
	lista := new(listaEnlazada[T])
	return lista
}

func crearNodo[T any](dato T) *nodoLista[T] {
	nodo := new(nodoLista[T])
	nodo.dato = dato
	nodo.siguiente = nil
	return nodo
}

func (lista *listaEnlazada[T]) EstaVacia() bool {
	return lista.primero == nil && lista.ultimo == nil
}

func (lista *listaEnlazada[T]) InsertarPrimero(dato T) {
	nodo := crearNodo(dato)
	if lista.EstaVacia() {
		lista.ultimo = nodo
	}
	nodo.siguiente = lista.primero
	lista.primero = nodo
	lista.largo++
}

func (lista *listaEnlazada[T]) InsertarUltimo(dato T) {
	nodo := crearNodo(dato)
	if lista.EstaVacia() {
		lista.primero = nodo
	} else {
		lista.ultimo.siguiente = nodo
	}
	lista.ultimo = nodo
	lista.largo++
}

func (lista *listaEnlazada[T]) BorrarPrimero() T {
	if lista.EstaVacia() {
		panic(PANICO)
	}
	elem := lista.primero.dato
	lista.primero = lista.primero.siguiente
	if lista.primero == nil {
		lista.ultimo = nil
	}
	lista.largo--
	return elem
}

func (lista *listaEnlazada[T]) VerPrimero() T {
	if lista.EstaVacia() {
		panic(PANICO)
	}
	return lista.primero.dato
}

func (lista *listaEnlazada[T]) VerUltimo() T {
	if lista.EstaVacia() {
		panic(PANICO)
	}
	return lista.ultimo.dato
}

func (lista *listaEnlazada[T]) Largo() int {
	return lista.largo
}

func (lista *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	visito := lista.primero
	for visito != nil && visitar(visito.dato) {
		visito = visito.siguiente
	}
}

//PRIMITIVAS ITERADOR

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iterListaEnlazada[T]{lista, nil, lista.primero}
}

func (iterador *iterListaEnlazada[T]) VerActual() T {
	if !iterador.HaySiguiente() {
		panic(PANICOITER)
	}
	return iterador.actual.dato
}

func (iterador *iterListaEnlazada[T]) HaySiguiente() bool {
	return iterador.actual != nil
}

func (iterador *iterListaEnlazada[T]) Siguiente() {
	if !iterador.HaySiguiente() {
		panic(PANICOITER)
	}
	iterador.anterior = iterador.actual
	iterador.actual = iterador.actual.siguiente
}

func (iterador *iterListaEnlazada[T]) Insertar(dato T) {
	nodo := crearNodo(dato)
	if !iterador.HaySiguiente() {
		iterador.lista.ultimo = nodo
	}
	if iterador.anterior == nil {
		iterador.lista.primero = nodo
	} else {
		iterador.anterior.siguiente = nodo
	}

	nodo.siguiente = iterador.actual
	iterador.actual = nodo
	iterador.lista.largo++

}

func (iterador *iterListaEnlazada[T]) Borrar() T {
	if !iterador.HaySiguiente() {
		panic(PANICOITER)
	}
	elem := iterador.actual.dato
	if iterador.actual.siguiente == nil {
		iterador.lista.ultimo = iterador.anterior
	}
	if iterador.anterior == nil {
		iterador.lista.primero = iterador.actual.siguiente
	} else {
		iterador.anterior.siguiente = iterador.actual.siguiente
	}
	iterador.actual = iterador.actual.siguiente
	iterador.lista.largo--
	return elem

}
