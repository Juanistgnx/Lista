package diccionario

//TDAPila "tdas/pila"

const PANICO = "La clave no pertenece al diccionario"
const PANICOITER = "El iterador termino de iterar"

type nodoAbb[K comparable, V any] struct {
	izquierdo *nodoAbb[K, V]
	derecho   *nodoAbb[K, V]
	clave     K
	dato      V
}
type abb[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	cantidad int
	cmp      funcCmp[K]
}
type iterAbb_test[K comparable, V any] struct { //A modo de prueba tomo esto,despues hay que ver como queda

}

type funcCmp[K comparable] func(K, K) int

// debe devoler 0 si son iguales
// debe devolver -1 si el primero es menor al segundo
// debe devolver 1 si es primero es mayor al segundo

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{nil, 0, funcion_cmp}
}

func crearnodo[K comparable, V any](clave K, dato V) *nodoAbb[K, V] {
	return &nodoAbb[K, V]{nil, nil, clave, dato}
}

func (abb *abb[K, V]) Guardar(clave K, dato V) {
	ubi := buscar(&abb.raiz, clave, abb.cmp)
	nodo := *ubi
	if nodo == nil {
		nodo = crearnodo(clave, dato)
	} else {
		nodo.dato = dato
	}
	abb.cantidad++
	*ubi = nodo
}

func (abb *abb[K, V]) Pertenece(clave K) bool {
	ubi := buscar(&abb.raiz, clave, abb.cmp)
	return *ubi != nil
}

func (abb *abb[K, V]) Obtener(clave K) V {
	ubi := buscar(&abb.raiz, clave, abb.cmp)
	nodo := *ubi
	if nodo == nil {
		panic(PANICO)
	}
	return nodo.dato
}

func (abb *abb[K, V]) Borrar(clave K) V {
	ubi := buscar(&abb.raiz, clave, abb.cmp)
	if *ubi == nil {
		panic(PANICO)
	}
	nodo := *ubi
	basura := nodo.dato
	if nodo.derecho == nil && nodo.izquierdo == nil {
		*ubi = nil
	} else if nodo.derecho == nil && nodo.izquierdo != nil {
		*ubi = nodo.izquierdo
	} else if nodo.derecho != nil && nodo.izquierdo == nil {
		*ubi = nodo.derecho
	} else {
		//puedo buscar el mas derecho de la izquierda o el mas izquierdo de la derecha,pero seguro esta despues de este nodo
		//decidi el mas derecho de izquierda
		ubi_r := buscar_mas_grande(&nodo.izquierdo) //Busco desde su nodo izquierdo por que no puede haber mas grandes que este nodo a borrar desde este lado.Por ende,puedo acortar la busqueda mas facil
		remplazo := *ubi_r
		//No importa si justo el hijo izquierdo era el mas grande,luego recupero esa posicion de izquierdp con la linea 77,por que ahora ,como actualize a donde apunta ubi_r ya no apunta a ese nodo sino a su hijo izquierdo
		*ubi_r = remplazo.izquierdo

		remplazo.izquierdo = nodo.izquierdo
		remplazo.derecho = nodo.derecho
		*ubi = remplazo
	}
	abb.cantidad--
	return basura
}

func (abb *abb[K, V]) Cantidad() int {
	return abb.cantidad
}

func (abb *abb[K, V]) Iterar(f func(clave K, dato V) bool) {
	iterar(abb.raiz, f)
}

func (abb *abb[K, V]) Iterador() IterDiccionario[K, V] {
	return &iterAbb_test[K, V]{abb}
}

func (abb *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	//...

}

func (abb *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	return &iterAbb_test[K, V]{abb}
}

func (iter *iterAbb_test[K, V]) HaySiguiente() bool {
	//A modo de prueba tomo esto
	return true
}

func (iter *iterAbb_test[K, V]) VerActual() (K, V) {
	//A modo de prueba se toma esto
	return iter.arbol.raiz.clave, iter.arbol.raiz.dato
}

func (iter *iterAbb_test[K, V]) Siguiente() {
	//A modo de prueba tomo esto
	//...
}

func buscar[K comparable, V any](raiz **nodoAbb[K, V], clave K, funcion_cmp func(K, K) int) **nodoAbb[K, V] { //Tuve que usar doble puntero para el nodo que recibe sino no me cambia a donde apunta mi cajita ,solamente me cambiaba la cajita a la que apuntaba con lo cual si esgaba en la raiz se rompia porqe apuntaba a oyrs caja que no esra la del abb
	nodo := *raiz
	if nodo == nil || funcion_cmp(nodo.clave, clave) == 0 {
		return raiz
	}
	if funcion_cmp(nodo.clave, clave) > 0 {
		return buscar(&nodo.izquierdo, clave, funcion_cmp)
	}
	return buscar(&nodo.derecho, clave, funcion_cmp)
	//EN teoria saque la fomra mas dificil,me devuelve la cajita del puntero (termino cajita es lo que usaba santisi creo)
}
func iterar[K comparable, V any](nodo *nodoAbb[K, V], visitar func(K, V) bool) { //Esta función está en preorder
	if nodo == nil || !visitar(nodo.clave, nodo.dato) {
		return
	}
	iterar(nodo.izquierdo, visitar)
	iterar(nodo.derecho, visitar)
}

func buscar_mas_grande[K comparable, V any](raiz **nodoAbb[K, V]) **nodoAbb[K, V] {
	nodo := *raiz
	if nodo.derecho == nil { //si donde apunta la caja que tengo ,no tiene a donde apunta a su derecha,entoces a donde apunta mi caja es el mas grande
		return raiz
	}
	return buscar_mas_grande(&nodo.derecho)
}
