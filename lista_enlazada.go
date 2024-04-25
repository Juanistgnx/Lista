ckage lista

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
	anterior *nodoLista[T]
	actual   *nodoLista[T]
}

func CrearListaEnlazada[T any]() Lista[T] {
	lista := new(listaEnlazada[T])
	lista.primero, lista.ultimo = nil, nil
	lista.largo = 0
	return lista
}

func crearNodo[T any](dato T) *nodoLista[T] {
	nodo := new(nodoLista[T])
	nodo.dato = dato
	nodo.siguiente = nil
	return nodo
}

//Crear funcion para crear iterador

func (lista *listaEnlazada[T]) EstaVacia() bool {
	return lista.primero == nil && lista.ultimo == nil
}

func (lista *listaEnlazada[T]) InsertarPrimero(dato T) {
	nodo := crearNodo(dato)
	if lista.EstaVacia() {
		lista.ultimo = nodo
	} else {
		nodo.siguiente = lista.primero
	}
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
		panic("La lista esta vacia")
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
		panic("La lista esta vacia")
	}
	return lista.primero.dato
}

func (lista *listaEnlazada[T]) VerUltimo() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	return lista.ultimo.dato
}

func (lista *listaEnlazada[T]) Largo() int {
	return lista.largo
}

// DISCUTIR
func (lista *listaEnlazada[T]) Iterar(visitar func(T) bool) {

}

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] { //Se supone que crea un iterador de las lista pero interfaz

}
