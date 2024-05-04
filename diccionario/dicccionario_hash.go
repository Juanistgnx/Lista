package diccionario

import (
	"fmt"
)

type estadoNodo int

const (
	VACIO = estadoNodo(iota)
	OCUPADO
	BORRADO
)
const LARGOINICIAL int = 13
const PANICO = "La clave no pertenece al diccionario"
const PANICOITER = "El iterador termino de iterar"

type nodoHash[K comparable, V any] struct {
	clave  K
	dato   V
	estado estadoNodo
}

type hashCerrado[K comparable, V any] struct {
	contenido []nodoHash[K, V]
	tamaño    int
	ocupados  int
	borrados  int
}

type iterDiccionarioHash[K comparable, V any] struct {
	hash     *hashCerrado[K, V]
	posicion int
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	nuevo := new(hashCerrado[K, V])
	nuevo.contenido = make([]nodoHash[K, V], LARGOINICIAL)
	nuevo.tamaño = LARGOINICIAL
	nuevo.ocupados, nuevo.borrados = 0, 0
	return nuevo
}

// Primitivas del Hash
func (hash *hashCerrado[K, V]) Guardar(clave K, dato V) {
	busqueda := buscar(hash, clave)
	if busqueda != nil {
		busqueda.dato = dato
		return
	}
	indice0 := indiceHash(clave, hash.tamaño)
	for i := 0; i < hash.tamaño; i++ {
		indice := (indice0 + i) % (hash.tamaño)
		if hash.contenido[indice].estado == VACIO {
			hash.contenido[indice].clave = clave
			hash.contenido[indice].dato = dato
			hash.contenido[indice].estado = OCUPADO
			break
		}
	}
	hash.ocupados++
}

func (hash *hashCerrado[K, V]) Pertenece(clave K) bool {
	nodo := buscar(hash, clave)
	return (nodo != nil)
}

func (hash *hashCerrado[K, V]) Obtener(clave K) V {
	nodo := buscar(hash, clave)
	if nodo == nil {
		panic("La clave no pertenece al diccionario")
	} else {
		return nodo.dato
	}
}

func (hash *hashCerrado[K, V]) Borrar(clave K) V {
	dato := hash.Obtener(clave) //Obtener hace saltar el panic en caso tal que la clave no haga parte del hash
	nodo := buscar(hash, clave)
	nodo.estado = BORRADO
	hash.borrados++
	return dato
}

func (hash *hashCerrado[K, V]) Cantidad() int {
	return (hash.ocupados - hash.borrados)
}

func (hash *hashCerrado[K, V]) Iterar(f func(clave K, dato V) bool) {
	for _, nodo := range hash.contenido {
		if nodo.estado == OCUPADO && !f(nodo.clave, nodo.dato) {
			break
		}
	}
}

func (hash *hashCerrado[K, V]) Iterador() IterDiccionario[K, V] {
	iterador := new(iterDiccionarioHash[K, V])
	iterador.hash = hash
	for i := 0; i < hash.tamaño; i++ {
		if hash.contenido[i].estado == OCUPADO {
			iterador.posicion = i
			break
		}
	}
	return iterador
}

// Primitivas del iterador
func (iter *iterDiccionarioHash[K, V]) HaySiguiente() bool {
	return iter.posicion < iter.hash.tamaño && iter.hash.contenido[iter.posicion].estado == OCUPADO

}

func (iter *iterDiccionarioHash[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic(PANICOITER)
	}
	elem := iter.hash.contenido[iter.posicion]
	return elem.clave, elem.dato
}

func (iter *iterDiccionarioHash[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic(PANICOITER)
	}
	tabla := iter.hash.contenido
	iter.posicion++
	for iter.posicion < iter.hash.tamaño {
		if tabla[iter.posicion].estado == OCUPADO {
			break
		}
		iter.posicion++
	}
}

//funciones auxiliares

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

func indiceHash[K comparable](clave K, largo_tabla int) int { //usando la funcion hash me da el indice en que inicialmente le corresponde a mi valor
	convertirABytes(clave)
	//hacer el hash
	//hacemos return hash % largo_tabla
}

// Recibe una cantidad para redimensionar
func redimensionar[K comparable, V any](hash *hashCerrado[K, V], cantidad_nueva int) {
	nuevo_contenido := make([]nodoHash[K, V], cantidad_nueva)
	for _, e := range hash.contenido {
		if e.estado == OCUPADO {
			ind_ini := indiceHash(e.clave, cantidad_nueva)
			indice := ind_ini
			for ; indice < cantidad_nueva && nuevo_contenido[indice].estado == OCUPADO; indice++ {
			}
			if indice == cantidad_nueva {
				indice = 0
				for ; indice < ind_ini && nuevo_contenido[indice].estado == OCUPADO; indice++ {
				}
			}
			nuevo_contenido[indice].clave, nuevo_contenido[indice].dato = e.clave, e.dato
		}
	}
	hash.contenido = nuevo_contenido
	hash.tamaño = cantidad_nueva
	hash.borrados = 0
}

func buscar[K comparable, V any](hash *hashCerrado[K, V], clave K) *nodoHash[K, V] {
	indice0 := indiceHash(clave, hash.tamaño)
	for i := 0; i < hash.tamaño; i++ {
		indice := (indice0 + i) % (hash.tamaño)
		if hash.contenido[indice].estado == BORRADO {
			continue
		} else if hash.contenido[indice].estado == VACIO {
			break
		} else if hash.contenido[indice].clave == clave {
			return &(hash.contenido[indice])
		}
	}
	return nil
}
