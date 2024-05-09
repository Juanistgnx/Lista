package diccionario

import (
	TDAPila "tdas/pila"
)

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
type iterAbb[K comparable, V any] struct { //A modo de prueba tomo esto,despues hay que ver como queda
	datos TDAPila.Pila[*nodoAbb[K, V]]
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
	continuar := true
	iterar(abb.raiz, f, &continuar)
}

func (abb *abb[K, V]) Iterador() IterDiccionario[K, V] {
	iterador := &iterAbb[K, V]{TDAPila.CrearPilaDinamica[*nodoAbb[K, V]]()}
	continuar := true
	iterar(abb.raiz, func(clave K, dato V) bool {
		iterador.datos.Apilar(&nodoAbb[K, V]{nil, nil, clave, dato})
		return true
	}, &continuar)
	iterador.datos = invertir_pilas(iterador.datos)
	return iterador
}

func (abb *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	continuar := true
	iterar_rango(abb.raiz, visitar, &continuar, abb.cmp, *desde, *hasta)
}

func (abb *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	iterador := &iterAbb[K, V]{TDAPila.CrearPilaDinamica[*nodoAbb[K, V]]()}
	continuar := true
	iterar_rango(abb.raiz, func(clave K, dato V) bool {
		iterador.datos.Apilar(&nodoAbb[K, V]{nil, nil, clave, dato})
		return true
	}, &continuar, abb.cmp, *desde, *hasta)
	iterador.datos = invertir_pilas(iterador.datos)
	return iterador
}

func (iter *iterAbb[K, V]) HaySiguiente() bool {
	return !iter.datos.EstaVacia()
}

func (iter *iterAbb[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic(PANICOITER)
	}
	actual := iter.datos.VerTope()
	return actual.clave, actual.dato
}

func (iter *iterAbb[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic(PANICOITER)
	}
	iter.datos.Desapilar()
}

// FUNCIONES AUXILIARES

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
func iterar[K comparable, V any](nodo *nodoAbb[K, V], visitar func(K, V) bool, flag *bool) {
	//Esta función está en pre-order
	/*if nodo == nil || !visitar(nodo.clave, nodo.dato) {
		return
	}
	iterar(nodo.izquierdo, visitar)
	iterar(nodo.derecho, visitar)*/

	//Esta función itera en in-order
	if nodo != nil && *flag {
		iterar(nodo.izquierdo, visitar, flag)
		if *flag && !visitar(nodo.clave, nodo.dato) {
			*flag = false
			return
		}
		iterar(nodo.derecho, visitar, flag)
	}
}

func iterar_rango[K comparable, V any](nodo *nodoAbb[K, V], visitar func(K, V) bool, flag *bool, cmp funcCmp[K], desde K, hasta K) {
	if nodo != nil && *flag {
		pertenencia_alta := cmp(nodo.clave, hasta)
		pertenencia_baja := cmp(nodo.clave, desde)
		if pertenencia_alta >= 0 {
			iterar_rango(nodo.izquierdo, visitar, flag, cmp, desde, hasta)
			if pertenencia_alta == 0 {
				if *flag && !visitar(nodo.clave, nodo.dato) {
					*flag = false
					return
				}
			}
		} else if pertenencia_baja <= 0 {
			if pertenencia_baja == 0 {
				if *flag && !visitar(nodo.clave, nodo.dato) {
					*flag = false
					return
				}
			}
			iterar_rango(nodo.derecho, visitar, flag, cmp, desde, hasta)
		} else {
			iterar_rango(nodo.izquierdo, visitar, flag, cmp, desde, hasta)
			if *flag && !visitar(nodo.clave, nodo.dato) {
				*flag = false
				return
			}
			iterar_rango(nodo.derecho, visitar, flag, cmp, desde, hasta)
		}
	}
}

func buscar_mas_grande[K comparable, V any](raiz **nodoAbb[K, V]) **nodoAbb[K, V] {
	nodo := *raiz
	if nodo.derecho == nil { //si donde apunta la caja que tengo ,no tiene a donde apunta a su derecha,entoces a donde apunta mi caja es el mas grande
		return raiz
	}
	return buscar_mas_grande(&nodo.derecho)
}

func invertir_pilas[K comparable, V any](pila TDAPila.Pila[*nodoAbb[K, V]]) TDAPila.Pila[*nodoAbb[K, V]] {
	pila_auxiliar := TDAPila.CrearPilaDinamica[*nodoAbb[K, V]]()
	for !pila.EstaVacia() {
		pila_auxiliar.Apilar(pila.Desapilar())
	}
	return pila_auxiliar
}
